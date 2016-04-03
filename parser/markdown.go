package parser

import gh "github.com/shurcooL/github_flavored_markdown"

func Markdown(b []byte) string {
	safe := gh.Markdown(b)
	return string(safe)
}
