package compare

import "testing"

type normalizeTest struct {
	arg1, expected string
}

type compareTest struct {
	arg1, arg2 string
	expected   bool
}

func Test_normalizeRelease(t *testing.T) {
	var normalizeTests = []normalizeTest{
		{"v3.1.1", "v3.1.1"},
		{"v3.1", "v3.1.0"},
		{"3.1", "v3.1.0"},
		{"5", "v5.0.0"},
	}

	for _, test := range normalizeTests {
		if output := NormalizeSemVer(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func Test_compareEqualReleases(t *testing.T) {
	var compareEqualTests = []compareTest{
		{"v3.1.1", "3.1.1", true},
		{"3.1", "3.1.0", true},
		{"3.1", "3.1.1", false},
	}

	for _, test := range compareEqualTests {
		if output := CompareEqualReleases(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}

}

func Test_compareReleases(t *testing.T) {
	var compareReleasesTests = []compareTest{
		{"v3.1.1", "3.1.1", true},
		{"3.1", "3.1.1", true},
		{"v3", "3.1.1", true},
		{"3", "3.5.56", true},
		{"3.1", "3.0.1", false},
	}

	for _, test := range compareReleasesTests {
		if output := CompareReleases(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}

}
