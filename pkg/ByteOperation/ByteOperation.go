package ByteOperation

import "strings"

func ToByteArray[T interface{ String() string }](entries []T, fieldDelim, entryDelim string) []byte {
	var lines []string
	for _, entry := range entries {
		formatted := strings.ReplaceAll(entry.String(), ";", fieldDelim)
		lines = append(lines, formatted)
	}
	return []byte(strings.Join(lines, entryDelim))
}

func ToArray(data []byte) []string {
	// Convert bytes to string and split by newlines
	str := string(data)
	// Split the string by newlines and remove empty lines
	lines := make([]string, 0)
	for _, line := range strings.Split(str, "\n") {
		// Trim spaces and check if line is not empty
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			lines = append(lines, trimmed)
		}
	}
	return lines

}
