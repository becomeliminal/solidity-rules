subinclude("//build_defs:solidity")

# These two targets are expected to fail. `plz test` does not build them.
# `plz build //test/09_failures:broken` and `:partial` must exit non-zero.
sol_contract(
    name = "broken",
    src = "Broken.sol",
    solc_version = "0.8.20",
    languages = [],
)

sol_contract(
    name = "partial",
    src = "Partial.sol",
    solc_version = "0.8.20",
    languages = [],
)

# A real forge failure, converted to Go test text. The gentest itself passes
# only when the revert reason and a real duration are in that text. Its own
# test.results is a pass, so Please does not treat the inner failure as this
# target's failure.
gentest(
    name = "failure_reason",
    data = {"SRC": ["reason.t.sol"]},
    test_tools = {
        "forge": CONFIG.SOLIDITY.FORGE_TOOL,
        "plzsol": CONFIG.SOLIDITY.PLEASE_SOL_TOOL,
        "solc": "//test:solc_0_8_20",
    },
    sandbox = CONFIG.SOLIDITY.SANDBOX,
    test_cmd = """
set -e
mkdir -p src test
cp "$DATA_SRC" test/reason.t.sol
printf 'pragma solidity 0.8.20; contract Empty {}\\n' > src/Empty.sol
printf '[profile.default]\\nsrc = "src"\\ntest = "test"\\n' > foundry.toml
set +e
$TOOLS_FORGE test --root . --use "$TOOLS_SOLC" --json --match-contract PlainTest > forge.json 2>forge.err
code=$?
set -e
if [ "$code" -eq 0 ]; then
  echo "forge test passed; wanted a failure" >&2
  cat forge.err >&2
  exit 1
fi
$TOOLS_PLZSOL forge-results --json forge.json --out converted.txt
grep -q "one is not two" converted.txt
grep -q -- "--- FAIL: PlainTest.test_bad" converted.txt
if grep -q "(0.00s)" converted.txt; then
  echo "duration was flattened to 0.00s" >&2
  cat converted.txt >&2
  exit 1
fi
printf '=== RUN   converter_keeps_reason\\n--- PASS: converter_keeps_reason (0.01s)\\nPASS\\n' > test.results
""",
)
