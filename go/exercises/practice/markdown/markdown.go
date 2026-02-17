package markdown

import (
	"fmt"
	"strings"
)

const (
	headingMarker  = '#'
	listItemPrefix = "* "
)

// Render translates markdown to HTML.
func Render(markdown string) string {
	var html strings.Builder
	inList := false

	for _, line := range strings.Split(markdown, "\n") {
		if strings.HasPrefix(line, listItemPrefix) {
			if !inList {
				html.WriteString("<ul>")
				inList = true
			}
			html.WriteString(fmt.Sprintf("<li>%s</li>", renderInline(line[2:])))
			continue
		}

		if inList {
			html.WriteString("</ul>")
			inList = false
		}

		if level := headingLevel(line); level > 0 {
			html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", level, line[level+1:], level))
		} else {
			html.WriteString(fmt.Sprintf("<p>%s</p>", renderInline(line)))
		}
	}

	if inList {
		html.WriteString("</ul>")
	}

	return html.String()
}

// headingLevel returns the heading level (1-6) for a line starting with #,
// or 0 if the line is not a valid heading.
func headingLevel(line string) int {
	count := 0
	for count < len(line) && line[count] == headingMarker {
		count++
	}
	if count >= 1 && count <= 6 {
		return count
	}
	return 0
}

// renderInline converts inline markdown (bold and italic) to HTML.
// Bold (__) is processed before italic (_) to avoid conflicts.
func renderInline(text string) string {
	result := text
	for strings.Contains(result, "__") {
		result = strings.Replace(result, "__", "<strong>", 1)
		result = strings.Replace(result, "__", "</strong>", 1)
	}
	for strings.Contains(result, "_") {
		result = strings.Replace(result, "_", "<em>", 1)
		result = strings.Replace(result, "_", "</em>", 1)
	}
	return result
}
