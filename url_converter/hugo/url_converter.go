package hugo

import (
	"strings"
)

//https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#get-a-repository-directory-readme

// convert a github link to a github API link
// https://github.com/CodeYourFuture/Module-Databases/tree/main/E-Commerce to
// https://api.github.com/repos/CodeYourFuture/Module-Databases/readme/E-Commerce

func ConvertToGithubAPIURL(src string) string {
	src = strings.Replace(src, "https://github.com/CodeYourFuture/", "https://api.github.com/repos/CodeYourFuture/", 1)
	src = strings.Replace(src, "/tree/main", "/readme", 1)
	return src
}


// https://cyf-pd.netlify.app/blocks/* to https://api.github.com/repos/CodeYourFuture/CYF-PD/readme/content/blocks/*
// this is so PD can add their links without having to worry about the github API

func ConvertPDToGithubAPIURL(src string) string {
	src = strings.Replace(src, "https://cyf-pd.netlify.app/", "https://api.github.com/repos/CodeYourFuture/CYF-PD/readme/content/", 1)
	return src
}




