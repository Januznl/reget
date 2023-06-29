package compare

import "testing"

type normalizeTest struct {
	arg1, expected string
}

type normalizeToVersionTest struct {
	arg1     string
	expected Version
}

type compareTest struct {
	arg1, arg2 string
	expected   bool
}

func Test_normalizeRelease(t *testing.T) {
	var normalizeTests = []normalizeTest{
		{"v3.1.1", "3.1.1"},
		{"v3.1", "3.1.0"},
		{"3.1", "3.1.0"},
		{"5", "5.0.0"},
	}

	for _, test := range normalizeTests {
		if output := normalizeRelease(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func Test_releaseToVersion(t *testing.T) {
	var normalizeToVersionTests = []normalizeToVersionTest{
		{"v3.1.1", Version{Major: 3, Minor: 1, Patch: 1}},
		{"v3.1", Version{Major: 3, Minor: 1, Patch: 0}},
		{"v3", Version{Major: 3, Minor: 0, Patch: 0}},
		{"3", Version{Major: 3, Minor: 0, Patch: 0}},
	}

	for _, test := range normalizeToVersionTests {
		if output := releaseToVersion(test.arg1); output != test.expected {
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
