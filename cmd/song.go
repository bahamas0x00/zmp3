/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/bahamas0x00/zmp3/pkg"
	"github.com/bahamas0x00/zmp3/pkg/common"
	"github.com/bahamas0x00/zmp3/pkg/zingmp3"
	"github.com/spf13/cobra"
)

var config *pkg.Config

// songCmd represents the song command
var songCmd = &cobra.Command{
	Use:   "song",
	Short: "Download MP3 song from Zing mp3",
	Long:  `Download MP3 song from Zing mp3`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !pkg.IsConfigFileExist() {
			err := pkg.WriteDefaultConfig()
			if err != nil {
				return err
			}
		}

		cfg, err := pkg.ReadConfigFile()
		if err != nil {
			return err
		}

		err = cfg.IsValidConfig()
		if err != nil {
			return err
		}

		config = cfg
		config.CreateDownloadFolderIfNotExist()
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return downloadSong()
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

func downloadSong() error {
	url, err := common.PromptString("Url")
	if err != nil {
		return err
	}

	if !strings.Contains(url, zingmp3.Song) {
		return common.InvalidSongUrl
	}

	song, err := zingmp3.GetDownloadLinks(url)
	if err != nil {
		return err
	}

	downloadObj := &zingmp3.DownloadObject{
		Title:  song.Title,
		Artist: strings.ReplaceAll(song.Artist, " ", ""),
		Type:   "mp3",
	}

	switch config.Mp3Quality {
	case zingmp3.Normal:
		downloadObj.DownloadUrl = song.Source.Audio.Num128.Download

	case zingmp3.VIP:
		downloadObj.DownloadUrl = song.Source.Audio.Num320.Download
		if downloadObj.DownloadUrl == "" {
			fmt.Println("This song is not supported 320kbps quality, 128kbps will be downloaded instead")
			downloadObj.DownloadUrl = song.Source.Audio.Num128.Download
		}

	}

	err = zingmp3.Download(downloadObj, config.GetDownloadFolder())
	if err != nil {
		return err
	}

	return nil

}
