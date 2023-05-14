package hugo

import (
	"testing"
	"reflect"
	"strings"

)

func TestBase64Decode(t *testing.T) {
	encoded := "LS0tCnRpdGxlOiBUZXN0IFBhZ2UKLS0tCgpUaGlzIGlzIGEgdGVzdCBwYWdlLgoKYGBgb2JqZWN0aXZlcwphYmMKY2RlCmBgYApTb21lIG90aGVyIGNvbnRlbnQuCg=="
	expected := "This is a test page.\n\n```objectives\nabc\ncde\n```\n\nSome other content."

	result, err := base64Decode(encoded)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if strings.TrimSpace(result) != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}

func TestExtractFence(t *testing.T) {
	markdownContent := "This is a test page.\n\n```objectives\nabc\ncde\n```\n\nSome other content."
	flag := "objectives"
	expected := []string{"```objectives\nabc\ncde\n```"}

	result, err := extractFence(markdownContent, flag)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if strings.Join(result, "\n") != strings.Join(expected, "\n") {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", strings.Join(expected, "\n"), strings.Join(result, "\n"))
	}
}
func TestDeblobWithoutFenceFlag(t *testing.T) {
	apiResponse := "LS0tCnRpdGxlOiBUZXN0IFBhZ2UKLS0tCgpUaGlzIGlzIGEgdGVzdCBwYWdlLgoKYGBgb2JqZWN0aXZlcwphYmMKY2RlCmBgYApTb21lIG90aGVyIGNvbnRlbnQuCg=="
	expected := "This is a test page.\n\n```objectives\nabc\ncde\n```\n\nSome other content."

	result, err := deblob(apiResponse, "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if strings.TrimSpace(result) != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}
func TestDeblobWithFenceFlag(t *testing.T) {
	apiResponse := "LS0tCnRpdGxlOiBUZXN0IFBhZ2UKLS0tCgpUaGlzIGlzIGEgdGVzdCBwYWdlLgoKYGBgb2JqZWN0aXZlcwphYmMKY2RlCmBgYApTb21lIG90aGVyIGNvbnRlbnQuCg=="
	fenceType := "objectives"
	expected := "```objectives\nabc\ncde\n```"

	result, err := deblob(apiResponse, fenceType)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if strings.TrimSpace(result) != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}

func TestSplitContent(t *testing.T) {
    for _, tc := range testCases {
        tc.decodedContent = strings.ReplaceAll(tc.decodedContent, "<backtick>", "`")
        tc.expectedFrontMatter = strings.ReplaceAll(tc.expectedFrontMatter, "<backtick>", "`")
        tc.expectedMarkdownContent = strings.ReplaceAll(tc.expectedMarkdownContent, "<backtick>", "`")

        t.Run(tc.desc, func(t *testing.T) {
            frontMatter, markdownContent, err := splitContent(tc.decodedContent)
            if err != nil {
                t.Fatalf("Unexpected error: %v", err)
            }

            if !reflect.DeepEqual(frontMatter, tc.expectedFrontMatter) {
                t.Errorf("Expected front matter:\n%s\n\nGot:\n%s", tc.expectedFrontMatter, frontMatter)
            }

            if !reflect.DeepEqual(markdownContent, tc.expectedMarkdownContent) {
                t.Errorf("Expected markdown content:\n%s\n\nGot:\n%s", tc.expectedMarkdownContent, markdownContent)
            }
        })
    }
}
