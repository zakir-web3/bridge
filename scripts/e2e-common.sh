#!/usr/bin/env bash
# Shared helpers for local e2e scripts.

e2e_port_in_use() {
  local port="$1"
  if command -v lsof >/dev/null 2>&1; then
    lsof -iTCP:"$port" -sTCP:LISTEN >/dev/null 2>&1
  else
    (echo >/dev/tcp/127.0.0.1/"$port") >/dev/null 2>&1
  fi
}

e2e_free_port() {
  local port="$1"
  if ! e2e_port_in_use "$port"; then
    return 0
  fi
  if ! command -v lsof >/dev/null 2>&1; then
    return 0
  fi
  local pids
  pids="$(lsof -tiTCP:"$port" -sTCP:LISTEN 2>/dev/null || true)"
  if [[ -z "$pids" ]]; then
    return 0
  fi
  # shellcheck disable=SC2086
  kill $pids 2>/dev/null || true
  sleep 1
  if e2e_port_in_use "$port"; then
    # shellcheck disable=SC2086
    kill -9 $pids 2>/dev/null || true
  fi
}

e2e_stop_pid() {
  local pid="$1"
  if [[ -z "$pid" ]]; then
    return 0
  fi
  if kill -0 "$pid" 2>/dev/null; then
    kill "$pid" 2>/dev/null || true
    wait "$pid" 2>/dev/null || true
  fi
  if command -v pgrep >/dev/null 2>&1; then
    local child
    for child in $(pgrep -P "$pid" 2>/dev/null || true); do
      e2e_stop_pid "$child"
    done
  fi
}
