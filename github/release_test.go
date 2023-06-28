package github

import "testing"

type findSssetForArchTest struct {
	arg1     []Assets
	arg2     string
	expected string
}

type findReleaseUrlTest struct {
	arg1     []Assets
	arg2     string
	arg3     string
	expected string
}

func Test_findAssetForArch(t *testing.T) {

	var findSssetForArchTests = []findSssetForArchTest{
		{
			[]Assets{
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd65.tar.gz"},
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd68.tar.gz"},
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd64.tar.gz"},
			},
			"darwin_amd64",
			"https://github.com/banaan-darwin_amd64.tar.gz",
		},
		{
			[]Assets{
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd64.tar.gz"},
			},
			"darwin_amd644",
			"",
		},
	}

	for _, test := range findSssetForArchTests {
		if output, _ := findAssetForArch(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func Test_findReleaseUrl(t *testing.T) {
	var findReleaseUrlTests = []findReleaseUrlTest{
		{
			[]Assets{
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd65.tar.gz"},
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd68.tar.gz"},
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd64.tar.gz"},
			},
			"darwin_amd64",
			"https://github.com/1.3.3-source-code.tar.gz",
			"https://github.com/banaan-darwin_amd64.tar.gz",
		},
		{
			[]Assets{
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd64.tar.gz"},
			},
			"darwin_amd644",
			"https://github.com/1.3.3-source-code.tar.gz",
			"https://github.com/1.3.3-source-code.tar.gz",
		},
		{
			[]Assets{
				{BrowserDownloadUrl: "https://github.com/banaan-darwin_amd64.tar.gz"},
			},
			"darwin_amd644",
			"",
			"",
		},
	}

	for _, test := range findReleaseUrlTests {
		if output, _ := findReleaseUrl(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
