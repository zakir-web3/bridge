#!/usr/bin/env bash
# Refresh integration-pack generated bits from the current checkout.
# Safe to run locally; does not push or open a PR.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PACK="$ROOT/integration-pack"
SHA="$(git -C "$ROOT" rev-parse HEAD)"
BRANCH="$(git -C "$ROOT" rev-parse --abbrev-ref HEAD)"

mkdir -p "$PACK/abi" "$PACK/idl" "$PACK/eip712" "$PACK/interfaces"

echo "==> stamping VERSION ($SHA on $BRANCH)"
ver_line="integration-pack@0.1.0-preview"
if [[ -f "$PACK/VERSION" ]]; then
  ver_line="$(grep -E '^integration-pack@' "$PACK/VERSION" | head -1 || true)"
  ver_line="${ver_line:-integration-pack@0.1.0-preview}"
fi
cat > "$PACK/VERSION" <<VERSION
$ver_line
source_repo: zakir-web3/bridge
SOURCE_COMMIT: $SHA
branch: $BRANCH
exported_at: $(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION

VEC_SRC="$ROOT/internal/contract/solana_bridge_test.go"
VEC_OUT="$PACK/eip712/vectors-solana-withdraw.json"
echo "==> checking Solana EIP-712 vectors"
if [[ -f "$VEC_SRC" ]]; then
  if [[ -f "$VEC_OUT" ]]; then
    echo "    $VEC_OUT present (maintained / previously exported)"
  else
    echo "    warn: $VEC_OUT missing — copy or regenerate before release"
  fi
  if [[ "${RUN_VECTOR_TESTS:-0}" == "1" ]] && command -v go >/dev/null 2>&1; then
    echo "    RUN_VECTOR_TESTS=1 — running Go vector tests (60s timeout)"
    if (cd "$ROOT" && go test ./internal/contract/ -run 'TestSolanaBridgeDomainSeparator|TestBridgeHubWithdrawSolanaTypedData|TestVerifyingContractFromProgramID' -count=1 -timeout 60s); then
      echo "    Go vector tests: ok"
    else
      echo "    Go vector tests: failed (JSON left unchanged)"
    fi
  else
    echo "    skip Go tests (set RUN_VECTOR_TESTS=1 to enable)"
  fi
else
  echo "    warn: $VEC_SRC not in this checkout"
fi

echo "==> optional: Solidity ABIs via forge"
if command -v forge >/dev/null 2>&1 && [[ -d "$ROOT/solidity" ]]; then
  (cd "$ROOT/solidity" && forge build --silent 2>/dev/null) || true
  if [[ -d "$ROOT/solidity/out" ]]; then
    for name in Bridge BridgeHub BridgeERC20; do
      found="$(find "$ROOT/solidity/out" -name "${name}.json" 2>/dev/null | head -1 || true)"
      if [[ -n "${found}" ]]; then
        cp "$found" "$PACK/abi/${name}.json"
        echo "    copied abi/${name}.json"
      fi
    done
  fi
else
  echo "    forge not found or no solidity/ — skip abi/"
fi

echo "==> optional: Anchor IDL"
IDL_SRC="$ROOT/solana/target/idl/bridge.json"
if [[ -f "$IDL_SRC" ]]; then
  cp "$IDL_SRC" "$PACK/idl/bridge.json"
  echo "    copied idl/bridge.json"
else
  echo "    no solana/target/idl/bridge.json — run anchor build to refresh"
fi

echo "==> done. Review integration-pack/ then commit when ready."
echo "    Pack root: $PACK"
