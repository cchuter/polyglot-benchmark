package markdown

import (
	"fmt"
	"strings"
)

const (
	headingMarker  = '#'
	listItemMarker = '*'
)

// Render translates markdown to HTML.
func Render(markdown string) string {
	var itemList []string
	var html = strings.Builder{}
	for _, line := range strings.Split(markdown, "\n") {
		if line[0] == listItemMarker {
			itemList = append(itemList, fmt.Sprintf("<li>%s</li>", renderInlineHTML(line[2:])))
			continue
		} else if len(itemList) != 0 {
			html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
			itemList = []string{}
		}
		if line[0] == headingMarker {
			level := getHeadingLevel(line)
			if level != -1 {
				html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", level, line[level+1:], level))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s</p>", line))
			}
			continue
		}
		html.WriteString(fmt.Sprintf("<p>%s</p>", renderInlineHTML(line)))
	}
	if len(itemList) != 0 {
		html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
	}
	return html.String()
}

// getHeadingLevel returns the heading level (1-6) or -1 if invalid.
func getHeadingLevel(line string) int {
	for i := 0; i <= 6; i++ {
		if line[i] != headingMarker {
			return i
		}
	}
	return -1
}

// renderInlineHTML converts markdown bold and italic markers to HTML tags.
func renderInlineHTML(text string) string {
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
