package markdown

import "strings"

// Render translates markdown input to HTML.
func Render(input string) string {
	var result strings.Builder
	lines := strings.Split(input, "\n")
	inList := false

	for _, line := range lines {
		if strings.HasPrefix(line, "* ") {
			if !inList {
				result.WriteString("<ul>")
				inList = true
			}
			content := applyInline(line[2:])
			result.WriteString("<li>" + content + "</li>")
		} else {
			if inList {
				result.WriteString("</ul>")
				inList = false
			}
			if strings.HasPrefix(line, "#") {
				level := headerLevel(line)
				if level > 0 && level <= 6 {
					content := applyInline(line[level+1:])
					lvl := string(rune('0' + level))
					result.WriteString("<h" + lvl + ">" + content + "</h" + lvl + ">")
				} else {
					result.WriteString("<p>" + applyInline(line) + "</p>")
				}
			} else {
				result.WriteString("<p>" + applyInline(line) + "</p>")
			}
		}
	}

	if inList {
		result.WriteString("</ul>")
	}

	return result.String()
}

func headerLevel(line string) int {
	level := 0
	for _, c := range line {
		if c == '#' {
			level++
		} else {
			break
		}
	}
	if level > 0 && level < len(line) && line[level] == ' ' {
		return level
	}
	return 0
}

func applyInline(text string) string {
	text = replaceMarker(text, "__", "<strong>", "</strong>")
	text = replaceMarker(text, "_", "<em>", "</em>")
	return text
}

func replaceMarker(text, marker, open, close string) string {
	parts := strings.Split(text, marker)
	var result strings.Builder
	for i, part := range parts {
		if i > 0 {
			if i%2 == 1 {
				result.WriteString(open)
			} else {
				result.WriteString(close)
			}
		}
		result.WriteString(part)
	}
	return result.String()
}
