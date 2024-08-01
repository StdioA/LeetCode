import (
	"fmt"
	"strings"
)

func leftJustify(line []string, maxWidth int) string {
	var builder strings.Builder
	builder.Grow(maxWidth)
	lengthLeft := maxWidth
	for i, word := range line {
		if i > 0 {
			builder.WriteString(" ")
			lengthLeft--
		}
		builder.WriteString(word)
		lengthLeft -= len(word)
	}
	builder.WriteString(strings.Repeat(" ", lengthLeft))
	return builder.String()
}

func fullJustify(words []string, maxWidth int) []string {
	var (
		lines             [][]string
		lengths           []int
		line              []string
		currentLineLength = -1
	)
	for _, word := range words {
		if currentLineLength+len(word)+1 > maxWidth {
			lines = append(lines, line)
			lengths = append(lengths, currentLineLength)
			line = nil
			currentLineLength = -1
		}
		line = append(line, word)
		currentLineLength += len(word) + 1
	}
	var (
		result  = make([]string, 0, len(lines)+1)
		builder strings.Builder
	)
	builder.Grow(maxWidth)
	for i, tempLine := range lines {
		if len(tempLine) == 1 {
			result = append(result, leftJustify(tempLine, maxWidth))
			continue
		}
		spacesLeft := maxWidth - lengths[i]
		div, mod := spacesLeft/(len(tempLine)-1), spacesLeft%(len(tempLine)-1)
		builder.Reset()

		for j, word := range tempLine {
			if j > 0 {
				// print space
				spaceCount := div + 1
				if mod > 0 {
					spaceCount++
					mod--
				}
				builder.WriteString(strings.Repeat(" ", spaceCount))
			}
			builder.WriteString(word)
		}
		result = append(result, builder.String())
	}

	// print the last line
	result = append(result, leftJustify(line, maxWidth))
	return result
}
