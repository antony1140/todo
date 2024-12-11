package cmd

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

import (

	"github.com/spf13/cobra"
)

var just string
var ident string

func init(){
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&just, "just", "j","", "specify one file to read" )
	listCmd.Flags().StringVarP(&ident, "ident", "i","", "specify type of comment eg. TODO, FIXME" )
}
//Dummy Comment
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all comments in a directory",

	Long: 	"prints all commands regardless of type. \n use just, j to specify one file to be used. Use ident, i \nto specify comment type eg. TODO, BUG, FIXME, or custom type",

	//TODO: this works 
	Run: func(cmd *cobra.Command, args []string){
		printCommentsInDir(just, ident)
	},
}






