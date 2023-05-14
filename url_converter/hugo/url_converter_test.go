package hugo

import (
	"testing"
)

// convert a github link to a github API link
// https://github.com/CodeYourFuture/Module-Databases/tree/main/E-Commerce to
// https://api.github.com/repos/CodeYourFuture/Module-Databases/readme/E-Commerce

func TestConvertToGithubAPIURL(t *testing.T) {
	result := ConvertToGithubAPIURL("https://github.com/CodeYourFuture/Module-Databases/tree/main/E-Commerce")
	expected := "https://api.github.com/repos/CodeYourFuture/Module-Databases/readme/E-Commerce"
	if result != expected {
		t.Errorf("ConvertToGithubAPIURL was incorrect, got: %s, want: %s.", result, expected)
	}
}

