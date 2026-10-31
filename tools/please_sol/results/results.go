// Package results turns `forge test --json` into the Go test text Please's
// gentest parser reads from test.results.
package results

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

type suite struct {
	TestResults map[string]testResult `json:"test_results"`
}

type testResult struct {
	Status         string          `json:"status"`
	Reason         *string         `json:"reason"`
	Counterexample json.RawMessage `json:"counterexample"`
	Duration       string          `json:"duration"`
}

// Convert reads forge's JSON and writes Go test lines. The process exit is
// not decided here: a failing test is a --- FAIL line, and the caller exits
// with forge's status.
func Convert(r io.Reader, w io.Writer) error {
	buf, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	buf = bytes.TrimSpace(buf)
	if i := bytes.IndexByte(buf, '{'); i > 0 {
		buf = buf[i:]
	}
	var suites map[string]suite
	if err := json.Unmarshal(buf, &suites); err != nil {
		return fmt.Errorf("forge json: %w", err)
	}
	names := make([]string, 0, len(suites))
	for name := range suites {
		names = append(names, name)
	}
	sort.Strings(names)

	var b strings.Builder
	saw := false
	failed := false
	for _, suiteName := range names {
		s := suites[suiteName]
		tests := make([]string, 0, len(s.TestResults))
		for name := range s.TestResults {
			tests = append(tests, name)
		}
		sort.Strings(tests)
		contract := suiteName
		if i := strings.LastIndex(suiteName, ":"); i >= 0 {
			contract = suiteName[i+1:]
		}
		contract = strings.ReplaceAll(contract, " ", "_")
		for _, testName := range tests {
			tr := s.TestResults[testName]
			saw = true
			label := contract + "." + testLabel(testName)
			dur, err := parseDuration(tr.Duration)
			if err != nil {
				return fmt.Errorf("%s: %w", label, err)
			}
			status, err := goStatus(tr.Status)
			if err != nil {
				return fmt.Errorf("%s: %w", label, err)
			}
			if status == "FAIL" {
				failed = true
			}
			fmt.Fprintf(&b, "=== RUN   %s\n", label)
			if msg := message(tr); msg != "" && status != "PASS" {
				for _, line := range strings.Split(msg, "\n") {
					fmt.Fprintf(&b, "    %s\n", line)
				}
			}
			fmt.Fprintf(&b, "--- %s: %s (%s)\n", status, label, formatSeconds(dur))
		}
	}
	if !saw {
		return fmt.Errorf("forge json contained no tests")
	}
	if failed {
		b.WriteString("FAIL\n")
	} else {
		b.WriteString("PASS\n")
	}
	_, err = io.WriteString(w, b.String())
	return err
}

func testLabel(name string) string {
	if i := strings.IndexByte(name, '('); i >= 0 {
		name = name[:i]
	}
	return name
}

func goStatus(status string) (string, error) {
	switch strings.ToLower(status) {
	case "success":
		return "PASS", nil
	case "failure":
		return "FAIL", nil
	case "skipped", "skip":
		return "SKIP", nil
	default:
		return "", fmt.Errorf("unknown forge status %q", status)
	}
}

func message(tr testResult) string {
	var parts []string
	if tr.Reason != nil && *tr.Reason != "" {
		parts = append(parts, *tr.Reason)
	}
	if len(tr.Counterexample) > 0 && string(tr.Counterexample) != "null" {
		parts = append(parts, "counterexample: "+string(tr.Counterexample))
	}
	return strings.Join(parts, "\n")
}

func formatSeconds(d time.Duration) string {
	return strconv.FormatFloat(d.Seconds(), 'f', 6, 64) + "s"
}

func parseDuration(s string) (time.Duration, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, nil
	}
	var total time.Duration
	for _, field := range strings.Fields(s) {
		part, err := parsePart(field)
		if err != nil {
			return 0, fmt.Errorf("duration %q: %w", s, err)
		}
		total += part
	}
	return total, nil
}

func parsePart(field string) (time.Duration, error) {
	i := 0
	for i < len(field) && field[i] >= '0' && field[i] <= '9' {
		i++
	}
	if i == 0 {
		return 0, fmt.Errorf("bad component %q", field)
	}
	n, err := strconv.ParseInt(field[:i], 10, 64)
	if err != nil {
		return 0, err
	}
	unit := field[i:]
	var base time.Duration
	switch unit {
	case "ns":
		base = time.Nanosecond
	case "us", "\u00b5s", "\u03bcs":
		base = time.Microsecond
	case "ms":
		base = time.Millisecond
	case "s":
		base = time.Second
	case "m":
		base = time.Minute
	case "h":
		base = time.Hour
	default:
		return 0, fmt.Errorf("bad unit %q", unit)
	}
	return time.Duration(n) * base, nil
}
