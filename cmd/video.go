/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "Download MP4 song from Zing mp3",
	Long:  `Download MP4 song from Zing mp3`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("video called")
		// TODO: Implement Download video from ZING mp3
	},
}

func init() {
	rootCmd.AddCommand(videoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
