package main

import (
	"fmt"
	fs "fs/ascii"
	"os"
	"strings"
)

func finalPrint(text string , banner string) string {
	name := ""
		if banner == "thinkertoy" || banner == "standard" || banner == "shadow" {
			name = banner
		} else {
			fmt.Println("incorrect banner")
			return ""
		}
	file := fs.Read_file(name)
	if file == nil {
		return ""
	}
	line := text
	if !fs.Is_ascii(line) {
		fmt.Println("Non Ascii character found")
		return ""
	}
	if len(line) < 1 {
		return ""
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
		return ""
	}
	defer resultFile.Close()

	_, err = resultFile.WriteString(finalResult)
	if err != nil {
		fmt.Println("Error writing the result")
		return ""
	}
	return finalResult
}


func main() {
	v := finalPrint("hamza" , "shadow")
	//fmt.Println(v)
}