/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// songCmd represents the song command
var songCmd = &cobra.Command{
	Use:   "song",
	Short: "Download MP3 song from Zing mp3",
	Long:  `Download MP3 song from Zing mp3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("song called")
		// TODO : Implement Download song from ZING mp3
	},
}

func init() {
	rootCmd.AddCommand(songCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// songCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// songCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
