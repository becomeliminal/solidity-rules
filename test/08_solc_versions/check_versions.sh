#!/bin/bash
# Each contract's bytecode records the solc that compiled it, in its CBOR
# metadata: "solc" (64736f6c63) as 3 bytes (43), then major, minor, patch.
set -euo pipefail

compiled_by() { # bytecode file -> "x.y.z"
    local v
    v=$(grep -oE "64736f6c6343[0-9a-f]{6}" "$1" | tail -1 | cut -c13-)
    [ -n "$v" ] || { echo "no solc version in $1's metadata" >&2; return 1; }
    echo "$((16#${v:0:2})).$((16#${v:2:2})).$((16#${v:4:2}))"
}

fail=0
for pair in "0820:0.8.20" "0823:0.8.23"; do
    tag="${pair%%:*}"
    want="${pair#*:}"
    # GIVEN a contract whose rule says solc_version = "$want"
    bin=$(find . -path "*pinned_${tag}_out/Pinned${tag}.bin" | head -1)
    [ -n "$bin" ] || { echo "FAIL: Pinned${tag}.bin not found"; fail=1; continue; }
    # WHEN its bytecode's metadata is read
    got=$(compiled_by "$bin")
    # THEN that solc compiled it
    if [ "$got" = "$want" ]; then
        echo "PASS: Pinned${tag} compiled by solc $got"
    else
        echo "FAIL: Pinned${tag} compiled by solc $got, rule asked for $want"
        fail=1
    fi
done
exit $fail
