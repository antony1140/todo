package cmd

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"strings"
)

	func isComment(scanner *bufio.Scanner)(bool, string){
		isComment := false
		write := false
		var comment []rune
		for _, char := range(scanner.Text()){
			if write == true {
				comment = append(comment, char)
				continue
			}
			if unicode.IsSpace(char){
				continue
			}
			if char != 47 {
				isComment = false
				return false, ""
			}
			if char == 47 {
				if isComment == true{
					isComment = false
					write = true
					fmt.Print(string(comment))
				}
				isComment = true
			}

		}
		return isComment, string(comment)
	}

	func listComments(fileName string)(int){

		file, err := os.Open("./" + fileName)
		if err != nil {
			fmt.Print(err)
		}
		scanner := bufio.NewScanner(file)

		comments := 0
		todos := 0
		for scanner.Scan(){
			if isComment, comment := isComment(scanner); isComment {
				comments += 1
				if strings.Contains(comment, "TODO"){
					todos += 1
					fmt.Print(comment,"\n")
				}
			}

		}
		return comments
	}

