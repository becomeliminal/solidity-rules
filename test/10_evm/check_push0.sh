#!/bin/bash
# 0.8.20's own default is Shanghai, which emits PUSH0 (opcode 5f) for `x := 0`.
# paris must not. A foundry.toml that only asks for paris must match the flag.
set -euo pipefail

bin_of() {
    local name="$1"
    local found
    # No pipe: pipefail plus head's SIGPIPE fails this on some find builds.
    found=$(find . -path "*${name}_out/Push0.bin" -print)
    found=${found%%$'\n'*}
    if [ -z "$found" ]; then
        echo "FAIL: ${name}_out/Push0.bin not found" >&2
        exit 1
    fi
    echo "$found"
}

default=$(bin_of push0_default)
paris=$(bin_of push0_paris)
toml=$(bin_of push0_toml)

if grep -q 5f "$default"; then
    echo "PASS: default 0.8.20 bytecode contains PUSH0"
else
    echo "FAIL: default 0.8.20 bytecode has no PUSH0" >&2
    exit 1
fi

if grep -q 5f "$paris"; then
    echo "FAIL: evm_version=paris still emitted PUSH0" >&2
    exit 1
fi
echo "PASS: evm_version=paris has no PUSH0"

if grep -q 5f "$toml"; then
    echo "FAIL: foundry.toml evm_version=paris still emitted PUSH0" >&2
    exit 1
fi
echo "PASS: foundry.toml evm_version=paris has no PUSH0"
