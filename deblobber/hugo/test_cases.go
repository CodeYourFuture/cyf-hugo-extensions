package hugo

var testCases = []struct {
    desc               string
    decodedContent     string
    expectedFrontMatter string
    expectedMarkdownContent string
}{
    {
        desc: "regular front matter",
        decodedContent: `---
title: Test Page
---

This is a test page.

<backtick><backtick><backtick>objectives
Objective 1
Objective 2
<backtick><backtick><backtick>
Some other content.`,
        expectedFrontMatter: `title: Test Page`,
        expectedMarkdownContent: `This is a test page.

<backtick><backtick><backtick>objectives
Objective 1
Objective 2
<backtick><backtick><backtick>
Some other content.`,
    },
    {
        desc: "comment-wrapped front matter",
        decodedContent: `<!--
---
title: Test Page
---
-->

This is a test page.

<backtick><backtick><backtick>objectives
Objective 1
Objective 2
<backtick><backtick><backtick>
Some other content.`,
        expectedFrontMatter: `title: Test Page`,
        expectedMarkdownContent: `This is a test page.

<backtick><backtick><backtick>objectives
Objective 1
Objective 2
<backtick><backtick><backtick>
Some other content.`,
    },
}
