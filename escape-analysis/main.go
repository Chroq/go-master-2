package main

import "fmt"

func escapeAnalysis(notEscaped *int) *int {
	escaped := 1

	notEscaped = &escaped

	return notEscaped
}

func main() {
	notEscaped := 1

	escapeAnalysis(&notEscaped)

	fmt.Println(notEscaped)
}
