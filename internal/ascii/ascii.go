package fs

import (
	"fmt"
	"os"
	"strings"
)

// this function reads the files and return its content as a slice of string
func Read_file(s string) []string {
	file, err := os.ReadFile("internal/art/" + s + ".txt")
	if err != nil {
		fmt.Println("Error in the file !")
		return nil
	}
	ret := strings.Split(string(file), "\n")
	for i := 0; i < len(ret) && s == "thinkertoy"; i++ {
		ret[i] = strings.ReplaceAll(ret[i], "\r", "")
	}
	if len(ret) != 856 {
		fmt.Println("Error in the file !")
		return nil
	}
	return ret
}

// Checks if there is an empty string in the middle of a slice, surrounded by non-empty strings.
func Middle(slice []string) bool {
	before := false
	after := false
	middle := true
	i := 0
	for ; i < len(slice); i++ {
		if slice[i] == "" {
			break
		}
	}
	for j := i - 1; j >= 0; j++ {
		before = false
		if slice[j] != "" {
			before = true
			break
		}
	}

	for k := i + 1; k < len(slice); k++ {
		before = false
		if slice[k] != "" {
			after = true
			break
		}
	}
	if before && after {
		middle = true
	} else {
		middle = false
	}
	return middle
}

// this function cleans the slice to return it to the main function to apply the logic of ascii art
func CleanSlice(slice []string) []string {
	check := true
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			check = false
		} else if slice[i] == "" {
			if Middle(slice) && !check {
				check = true
				slice = append(slice[:i], slice[i+1:]...)
			} else {
				i++
			}
		}
	}
	return slice
}

// this function apllies the main logic of ascii art
func PrintAscii(slice []string, file []string) string {
	result := ""
	holder := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			result += "\r\n"
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(slice[i]); k++ {
					holder = (int(slice[i][k])-32)*9 + j
					result += file[holder]
				}
				result += "\r\n"
			}
		}
	}
	if slice[0] == "" {
		result = "\r\n" + result
	}
	result = result[:len(result)-1]
	return result
}

// this function returns a string have only the ascii art character
func Is_ascii(s string) string {
	var result string
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] == 10 || slice[i] == 13 {
			result += string(slice[i])
		} else if slice[i] >= 32 && slice[i] <= 126 {
			result += string(slice[i])
		} else {
			return "Non ascii character !"
		}
	}
	return result
}

// this fucntion take the text and the banner and apllies the main logic
func FinalPrint(text string, banner string) string {
	name := ""
	if banner == "thinkertoy" || banner == "standard" || banner == "shadow" {
		name = banner
	} else {
		return "incorrect banner"
	}
	file := Read_file(name)
	if file == nil {
		return "error in the file"
	}
	ret := Is_ascii(text)
	if ret == "Non ascii character !" {
		return "ascii error !"
	}
	if len(text) < 1 {
		return ""
	}
	finalResult := ""
	splitted_line := strings.Split(ret, "\r\n")
	cleaned := CleanSlice(splitted_line)
	finalResult = PrintAscii(cleaned, file[1:])
	return finalResult
}
