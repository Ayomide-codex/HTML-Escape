package main

func replaceChar(b byte) string {
	switch b {
	case '&':
		return "&amp;"
	case '<':
		return "&lt;"
	case '>':
		return "&gt;"
	case '"':
		return "&quot;"
	case '\'':
		return "&#39;"
	default:
		return string(b)
	}
}
