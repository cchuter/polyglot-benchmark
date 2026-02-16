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
			itemList = append(itemList, fmt.Sprintf("<li>%s</li>", renderInline(line[2:])))
			continue
		} else if len(itemList) != 0 {
			html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
			itemList = []string{}
		}
		if line[0] == headingMarker {
			level := headingLevel(line)
			if level != -1 {
				html.WriteString(fmt.Sprintf("<h%d>%s</h%d>", level, line[level+1:], level))
			} else {
				html.WriteString(fmt.Sprintf("<p>%s</p>", line))
			}
			continue
		}
		html.WriteString(fmt.Sprintf("<p>%s</p>", renderInline(line)))
	}
	if len(itemList) != 0 {
		html.WriteString(fmt.Sprintf("<ul>%s</ul>", strings.Join(itemList, "")))
	}
	return html.String()
}

// headingLevel counts leading '#' characters and returns the heading level (1-6).
// Returns -1 if there are more than 6 '#' characters (not a valid heading).
func headingLevel(line string) int {
	for i := 0; i <= 6; i++ {
		if line[i] != headingMarker {
			return i
		}
	}
	return -1
}

// renderInline converts markdown inline formatting to HTML.
// Handles bold (__text__) and italic (_text_) markers.
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
