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

	func listComments(fileName string)([]string){
		var comments []string
		file, err := os.Open("./" + fileName)
		if err != nil {
			fmt.Print(err)
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan(){
			if isComment, comment := isComment(scanner); isComment {
				comments = append(comments, comment)
				if strings.Contains(comment, "TODO"){
					fmt.Print(comment,"\n")
				}
			}

		}
		return comments
	}

	func getDirFiles()([]string, error){
		var files []string
		dir, err := os.ReadDir(".")
		if err != nil {
			fmt.Print(err)
			return files, err
		}
		for _, file := range dir {
			files = append(files, file.Name())
		}
		return files, nil
	}

