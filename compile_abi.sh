#!/usr/bin/env bash

set -eo pipefail

# cross-platform in-place sed (GNU/BSD)
sedi() {
  if sed --version >/dev/null 2>&1; then
    sed -i "$@"
  else
    sed -i '' "$@"
  fi
}

commands=(git jq abigen)
for cmd in "${commands[@]}"; do
  if ! command -v "$cmd" &>/dev/null; then
    echo "$cmd command not found, please install $cmd first" && exit 1
  fi
done

contract_dir="$(git rev-parse --show-toplevel)/solidity/"
output_dir="$(git rev-parse --show-toplevel)/internal/contract/"

[[ ! -d "$output_dir" ]] && mkdir -p "$output_dir"
rm -rf "$output_dir"/*.sol.go

if [ ! -d "$contract_dir/node_modules" ]; then
  echo "===> Installing node modules"
  (cd "$contract_dir" && npm run install)
fi

echo "===> Compiling contracts"
(cd "$contract_dir" && npm run compile && npm run format)

contracts=(Bridge BridgeHub BridgeERC20)

abi_dir="$contract_dir/abi"
[[ ! -d "$abi_dir" ]] && mkdir -p "$abi_dir"

for contract in "${contracts[@]}"; do
  echo "===> Ethereum ABI wrapper code generator: $contract"
  file_path=$(find "$contract_dir/artifacts" -name "${contract}.json" -type f | head -n 1)
  jq -c '.abi' "$file_path" >"${abi_dir}/${contract}.abi"
  jq -r '.bytecode' "$file_path" >"${abi_dir}/${contract}.bin"
  go_filename=$(echo "${contract}" | sed -r 's/([a-z])([A-Z])/\1_\2/g' | tr '[:upper:]' '[:lower:]')
  abigen --abi "${abi_dir}/${contract}.abi" \
    --bin "${abi_dir}/${contract}.bin" \
    --type "${contract}" --pkg contract \
    --out "${output_dir}/${go_filename}.sol.go"
done

#
# Post-process generated bindings to deduplicate shared structs
#
# In our setup, both Bridge and BridgeHub ABIs contain a `Signature` struct.
# Since both bindings are generated into the same Go package `contract`, we
# keep the definition from `bridge.sol.go` and remove the duplicate from
# `bridge_hub.sol.go` to avoid "Signature redeclared" compile errors.
bridge_hub_file="${output_dir}/bridge_hub.sol.go"
if [[ -f "$bridge_hub_file" ]]; then
  echo "===> Removing duplicate Signature from BridgeHub binding"
  sedi '/^\/\/ Signature is an auto generated low-level Go binding around an user-defined struct\./,/^}/d' "$bridge_hub_file"
fi

rm -rf "$abi_dir"
