package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func simpleMarkdownToHTML(markdown string) string {
	lines := strings.Split(markdown, "\n")
	var htmlLines []string

	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			htmlLines = append(htmlLines, "<h1>"+strings.TrimPrefix(line, "# ")+"</h1>")
		} else if strings.HasPrefix(line, "## ") {
			htmlLines = append(htmlLines, "<h2>"+strings.TrimPrefix(line, "## ")+"</h2>")
		} else if strings.HasPrefix(line, "### ") {
			htmlLines = append(htmlLines, "<h3>"+strings.TrimPrefix(line, "### ")+"</h3>")
		} else if strings.HasPrefix(line, "- ") {
			htmlLines = append(htmlLines, "<li>"+strings.TrimPrefix(line, "- ")+"</li>")
		} else {
			htmlLines = append(htmlLines, "<p>"+line+"</p>")
		}
	}

	return strings.Join(htmlLines, "\n")
}

func main() {
	inputFile, err := os.Open("input.md")
	if err != nil {
		fmt.Println("Error opening input.md:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	htmlContent := simpleMarkdownToHTML(content)

	outputFile, err := os.Create("output.html")
	if err != nil {
		fmt.Println("Error creating output.html:", err)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(htmlContent)
	if err != nil {
		fmt.Println("Error writing to output.html:", err)
		return
	}

	fmt.Println("Conversion completed! Check output.html")
}
