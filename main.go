package main

import (
	"fmt"
	fs "fs/ascii"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) < 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	name := ""
	if len(os.Args) == 2 {
		name = "standard"
	} else {
		if os.Args[2] == "thinkertoy" || os.Args[2] == "standard" || os.Args[2] == "shadow" {
			name = os.Args[2]
		} else {
			fmt.Println("incorrect banner")
			return
		}
	}
	file := fs.Read_file(name)
	if file == nil {
		return
	}
	line := os.Args[1]
	if !fs.Is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return
	}
	if len(line) < 1 {
		return
	}
	finalResult := ""
	lines_count := fs.Count_next_line(line)
	splitted_line := strings.Split(line, "\\n")
	splitted_line, lines_count = fs.Cleaned_split(splitted_line, lines_count)
	finalResult = fs.Print_art(file[1:], splitted_line, lines_count)

	filename := "result.txt"
	resultFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating a file:", err)
		return
	}
	defer resultFile.Close()

	_, err = resultFile.WriteString(finalResult)
	if err != nil {
		fmt.Println("Error writing the result")
		return
	}
}
