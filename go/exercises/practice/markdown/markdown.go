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
			itemList = nil
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

// headingLevel returns the heading level (1-6) or -1 if the line has 7+ '#' characters.
func headingLevel(line string) int {
	for i := 0; i <= 6; i++ {
		if line[i] != headingMarker {
			return i
		}
	}
	return -1
}

// renderInline converts bold (__) and italic (_) markers to HTML tags.
// Bold must be processed before italic so that __ is not consumed as two separate _.
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
