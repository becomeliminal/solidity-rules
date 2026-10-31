package results

import (
	"strings"
	"testing"
)

func TestConvert_PassFailAndDuration(t *testing.T) {
	// Captured from forge 1.5.0 `forge test --json` on a two-test contract.
	in := `{
	  "test/Plain.t.sol:PlainTest": {
	    "duration": "2ms 286µs 248ns",
	    "test_results": {
	      "test_bad()": {
	        "status": "Failure",
	        "reason": "one is not two",
	        "counterexample": null,
	        "duration": "486µs 904ns"
	      },
	      "test_ok()": {
	        "status": "Success",
	        "reason": null,
	        "counterexample": null,
	        "duration": "535µs 21ns"
	      }
	    }
	  }
	}`
	var b strings.Builder
	if err := Convert(strings.NewReader(in), &b); err != nil {
		t.Fatal(err)
	}
	got := b.String()
	for _, want := range []string{
		"=== RUN   PlainTest.test_bad",
		"    one is not two",
		"--- FAIL: PlainTest.test_bad (0.000487s)",
		"=== RUN   PlainTest.test_ok",
		"--- PASS: PlainTest.test_ok (0.000535s)",
		"FAIL\n",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in:\n%s", want, got)
		}
	}
	if strings.Contains(got, "(0.00s)") {
		t.Errorf("duration was flattened to 0.00s:\n%s", got)
	}
}

func TestConvert_Counterexample(t *testing.T) {
	in := `{
	  "Fuzz.t.sol:FuzzTest": {
	    "test_results": {
	      "test_fuzz(uint256)": {
	        "status": "Failure",
	        "reason": "assertion failed",
	        "counterexample": {"x": "1"},
	        "duration": "1s 2ms"
	      }
	    }
	  }
	}`
	var b strings.Builder
	if err := Convert(strings.NewReader(in), &b); err != nil {
		t.Fatal(err)
	}
	got := b.String()
	if !strings.Contains(got, "--- FAIL: FuzzTest.test_fuzz (1.002000s)") {
		t.Errorf("duration or name wrong:\n%s", got)
	}
	if !strings.Contains(got, "assertion failed") || !strings.Contains(got, `counterexample: {"x": "1"}`) {
		t.Errorf("reason dropped:\n%s", got)
	}
}

func TestConvert_NoTests(t *testing.T) {
	err := Convert(strings.NewReader(`{"S":{"test_results":{}}}`), &strings.Builder{})
	if err == nil {
		t.Fatal("expected an error")
	}
}
