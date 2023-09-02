package ascii_art

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(textFromOutside string, artstylepath string) string {
	// 3 textstyles in a folder
	fileLines := ReadStandardTxt(artstylepath)
	asciiTemplates := return2dASCIIArray(fileLines)
	str := printAllStringASCII(textFromOutside, asciiTemplates)
	return str
}

func ReadStandardTxt(artstyle string) []string {
	readFile, err := os.Open(artstyle)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

func return2dASCIIArray(fileLines []string) [][]string {
	var asciiTemplates [][]string
	counter := 0
	var tempAsciArray []string
	for _, line := range fileLines {
		counter++
		// fmt.Println(index, line)
		if counter != 1 {
			tempAsciArray = append(tempAsciArray, line)
		}
		if counter == 9 {
			// fmt.Println("add to list") // but dont include the first line because it is empty line
			asciiTemplates = append(asciiTemplates, tempAsciArray)
			counter = 0
			tempAsciArray = nil
		} else {
		}
	}
	return asciiTemplates
}

// problem '\n' logic when we have 2 '\n' or 1 '\n' is different
func printMultipleCharacter(s string, asciiTemplates [][]string) string {
	returnString := ""
	// for ex 'hello'
	// we gonna write all letters index 0 after $ after \n after index 1  after $ after \n
	tempIntArrLetter := returnAsciiCodeInt(s)
	for i := 0; i < 8; i++ {
		for _, v := range tempIntArrLetter {
			returnString = fmt.Sprint(returnString, asciiTemplates[v][i])
		}
		returnString = fmt.Sprintln(returnString, "")
	}
	return returnString
}

func returnAsciiCodeInt(s string) []int {
	var tempIntArrLetter []int
	for _, v := range s {
		tempIntArrLetter = append(tempIntArrLetter, (int(v) - 32))
	}
	return tempIntArrLetter
}

func printAllStringASCII(text string, asciiTemplates [][]string) string {
	returnString := ""
	/*
		if ends w \n it gonna print println $
		if you can see text after \n chec;
		before \n
		if yes  println $
		if no println
	*/

	/*
	   func to uses printMultipleCharacters print whole stringfrom outside
	*/
	// Split the input string into an array of strings
	// split the line into words if there is a "\r\n" symbol
	// substrings := returnstring2EndlineArray(text)
	substrings := strings.Split(text, "\r\n")
	fmt.Println(substrings)
	lenOfsubstrings := len(substrings)
	for index, v := range substrings {
		if v == "\r\n" {
			// If it is last one
			if index == lenOfsubstrings-1 {
				returnString = fmt.Sprintln(returnString, "")
			} else if index == 0 {
				returnString = fmt.Sprintln(returnString, "") // no idea CHECK IT POTENTIAL ERROR
			} else {
				if substrings[index-1] == "\r\n" {
					returnString = fmt.Sprintln(returnString, "")
				} else {
					// "Hello\nWorld"
				}
			}
		} else {
			returnString = fmt.Sprint(returnString, printMultipleCharacter(v, asciiTemplates))
		}
	}
	return returnString
}

func returnstring2EndlineArray(text string) []string {
	substrings := make([]string, 0)
	escapedN := "\\n"
	newline := "\n"

	for {
		idx := strings.Index(text, escapedN)
		if idx == -1 {
			substrings = append(substrings, text)
			break
		}

		substrings = append(substrings, text[:idx])

		if idx+len(escapedN) < len(text) && text[idx+len(escapedN)] == 'n' {
			substrings = append(substrings, newline)
			text = text[idx+len(escapedN)+1:]
		} else {
			substrings = append(substrings, escapedN)
			text = text[idx+len(escapedN):]
		}
	}
	// fmt.Printf("%#v\n", substrings)
	var mysubstring2 []string
	for _, mysub := range substrings {
		if mysub != "" {
			mysubstring2 = append(mysubstring2, mysub)
		}
	}
	// fmt.Printf("%#v\n", mysubstring2)
	return mysubstring2
}
