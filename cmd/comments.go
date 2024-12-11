package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "slices"
	// "strings"
	"unicode"
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

	func getComments(fileName string, ident string)([]string){
		var comments []string
		file, err := os.Open("./" + fileName)
		if err != nil {
			fmt.Print(err)
		}
		scanner := bufio.NewScanner(file)
		lastLineComment := false
		comment := ""
		for scanner.Scan(){
			if isComment, line := isComment(scanner); isComment {
				if lastLineComment {
					comment += "\n" + line
					lastLineComment = true

				} else {
					comment += line
					lastLineComment = true
				}
			} else {
				if lastLineComment {
					if ident != " "{
						if strings.Contains(strings.ToLower(comment), strings.ToLower(ident)){
							comments = append(comments, comment)
						}
					} else{
						comments = append(comments, comment)
					}
				}
				lastLineComment = false
				comment = ""
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

	func getFile(just string)([]string, error){
		var files []string
		file, err := os.Open(just)
		if err != nil {
			fmt.Print(err)
			return files, err
		}
		files = append(files, file.Name())

		return files, nil

	}

	func printCommentsInDir(just string, ident string){
		var files []string
		if just != ""{
			fileList, err := getFile(just)
			if err != nil {
				fmt.Print(err)
			}
			files = fileList
		}else{
			fileList, err := getDirFiles()
			if err != nil {
				fmt.Print("Something went wrong")
			}
			files = fileList

		}
		for _,file := range files {
			// if just == ""{
			// 	if file != just{
			// 		continue
			// }
			//
			// }
			if comments := getComments(file, ident); len(comments) == 0 {
			}else {
				fmt.Println("---------------------")
				fmt.Println("---------------------")
				fmt.Println("---------------------")
				fmt.Println("\n FILE: " + file)
				fmt.Println("---------------------")
				for _, comment := range comments {
					fmt.Println(comment)
					fmt.Println("---------------------")
				}
			}

		}

	}

	func printCommentsInFile(fileName string){
			if comments := getComments(fileName, ident); len(comments) == 0 {
			}else {
				fmt.Println("\n" + fileName + ": ")
				for _, comment := range comments {
					fmt.Println(comment)
					fmt.Println("---------------------")
				}
			}
		
	}


