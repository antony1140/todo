package cmd

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	func getComments(fileName string)([]string){
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
				// if strings.Contains(comment, "TODO"){
				// 	fmt.Print(comment,"\n")
				// }
			} else {
				if lastLineComment {
					comments = append(comments, comment)
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

	func printCommentsInDir(exclude []string){
		files, err := getDirFiles()
		if err != nil {
			fmt.Print("Something went wrong")
		}
		for _,file := range files {
			if slices.Contains(exclude, file){
				continue
			}
			if comments := getComments(file); len(comments) == 0 {
			}else {
				fmt.Println("---------------------")
				fmt.Println("---------------------")
				fmt.Println("\n FILE: " + file)
				fmt.Println("---------------------")
				fmt.Println("---------------------")
				for _, comment := range comments {
					fmt.Println(comment)
					fmt.Println("---------------------")
				}
			}

		}

	}

	func printCommentsInFile(fileName string){
			if comments := getComments(fileName); len(comments) == 0 {
			}else {
				fmt.Println("\n" + fileName + ": ")
				for _, comment := range comments {
					fmt.Println(comment)
					fmt.Println("---------------------")
				}
			}
		
	}


