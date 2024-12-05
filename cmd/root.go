package main

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "A brief description of your application",
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rootCmd.AddCommand(CmdWeb, gormCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
	}
}
