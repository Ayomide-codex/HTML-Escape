package main

func escapeHTML(input string) string {
	result := ""
	for i := 0; i < len(input); i++ {
		result += replaceChar(input[i])
	}
	return result
}
