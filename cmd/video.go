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

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "Download MP4 song from Zing mp3",
	Long:  `Download MP4 song from Zing mp3`,
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
		return downloadVideo()
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

func downloadVideo() error {
	url, err := common.PromptString("Url")
	if err != nil {
		return err
	}

	if !strings.Contains(url, zingmp3.VideoClip) {
		return common.InvalidVideoUrl
	}

	video, err := zingmp3.GetDownloadLinks(url)
	if err != nil {
		return err
	}

	downloadObj := &zingmp3.DownloadObject{
		Title:  video.Title,
		Artist: strings.ReplaceAll(video.Artist, " ", ""),
		Type:   "mp4",
	}

	switch config.Mp4Quality {
	case zingmp3.SD_360:
		downloadObj.DownloadUrl = video.Source.Video.Num360.Download
	case zingmp3.SD_480:
		downloadObj.DownloadUrl = video.Source.Video.Num480.Download
	case zingmp3.HD_720:
		downloadObj.DownloadUrl = video.Source.Video.Num720.Download
	case zingmp3.FULL_HD_1080:
		downloadObj.DownloadUrl = video.Source.Video.Num1080.Download
	}

	if downloadObj.DownloadUrl == "" {
		fmt.Printf("Video with quality %dp is not exist, default quality 360p will be downloaded instead", config.Mp4Quality)
		downloadObj.DownloadUrl = video.Source.Video.Num360.Download
	}

	err = zingmp3.Download(downloadObj, config.GetDownloadFolder())
	if err != nil {
		return err
	}

	return nil
}
