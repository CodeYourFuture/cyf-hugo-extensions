package hugo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// this package expects a github API response
// here's an example https://api.github.com/repos/CodeYourFuture/Module-Databases/readme/E-Commerce

func deblob(apiResponse, flag string) (string, error) {
	decodedContent, err := base64Decode(apiResponse)
	if err != nil {
		return "", err
	}

	_, markdownContent, err := splitContent(decodedContent)
	if err != nil {
		return "", err
	}
	if flag != "" {
		fences, err := extractFence(markdownContent, flag)
		if err != nil {
			return "", err
		}
		return strings.Join(fences, "\n"), nil
	}

	return markdownContent, nil
}

func base64Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func splitContent(decodedContent string) (string, string, error) {
	re := regexp.MustCompile(`(?s)(<!--)?\s*---\s*(.*?)\s*---\s*(?:-->[\r\n]*)?`)
	matches := re.FindStringSubmatch(decodedContent)
	if len(matches) < 3 {
		return "", "", errors.New("failed to split content")
	}

	frontMatter := strings.TrimSpace(matches[2])
	markdownContent := strings.TrimSpace(decodedContent[len(matches[0]):])

	return frontMatter, markdownContent, nil
}
func extractFence(markdownContent, flag string) ([]string, error) {
	fenceRegex := regexp.MustCompile(fmt.Sprintf("(?s)```%s(.*?)```", regexp.QuoteMeta(flag)))
	matches := fenceRegex.FindAllString(markdownContent, -1)

	if matches == nil {
		return []string{}, nil
	}

	return matches, nil
}
