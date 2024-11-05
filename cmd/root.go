/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/bahamas0x00/zmp3/pkg"
	"github.com/bahamas0x00/zmp3/pkg/common"
	"github.com/spf13/cobra"
)

var (
	version    bool
	setConfig  bool
	showConfig bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zmp3",
	Short: "A simple CLI for download Song/Video from Zing mp3",
	Long:  `A simple CLI for download Song/Video from Zing mp3`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !pkg.IsConfigFileExist() {
			err := pkg.WriteDefaultConfig()
			if err != nil {
				return err
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			return printVersion()
		}
		if showConfig {
			return showCurrentConfig()
		}
		if setConfig {
			return setNewConfig()
		}

		return cmd.Help()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zmp3.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize()
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show cli version")
	rootCmd.Flags().BoolVarP(&showConfig, "show config", "s", false, "show current configuration")
	rootCmd.Flags().BoolVarP(&setConfig, "set config", "c", false, "set new configuration")

}

func printVersion() error {
	fmt.Println("Version: ", pkg.Version)
	return nil
}

func showCurrentConfig() error {
	cfg, err := pkg.ReadConfigFile()
	if err != nil {
		return err
	}
	err = cfg.IsValidConfig()
	if err != nil {
		return err
	}
	fmt.Printf("MP3 Quality: %d\n"+"MP4 Quality: %d\n"+"Directory: %s",
		cfg.Mp3Quality, cfg.Mp4Quality, cfg.GetDownloadFolder(),
	)
	return nil
}

func setNewConfig() error {
	cfg, err := promptSetNewConfig()
	if err != nil {
		return err
	}

	err = pkg.WriteConfigFile(cfg)
	if err != nil {
		return err
	}
	fmt.Println("Override new configuration with details:")

	fmt.Printf("MP3 Quality: %d\n"+
		"MP4 Quality: %d\n"+
		"Directory: %s",
		cfg.Mp3Quality, cfg.Mp4Quality, cfg.GetDownloadFolder())
	return nil
}

func promptSetNewConfig() (*pkg.Config, error) {
	cfg := &pkg.Config{}
	mp3Quality, err := common.PromptInteger("MP3 Quality")
	if err != nil {
		return nil, err
	}

	if err := pkg.IsValidMP3Quality(mp3Quality); err != nil {
		return nil, err
	}
	cfg.Mp3Quality = mp3Quality

	mp4Quality, err := common.PromptInteger("MP4 Quality")
	if err != nil {
		return nil, err
	}

	if err := pkg.IsValidMP4Quality(mp4Quality); err != nil {
		return nil, err
	}
	cfg.Mp4Quality = mp4Quality

	directory, err := common.PromptString("Directory")
	if err != nil {
		return nil, err
	}
	cfg.Directory = directory
	return cfg, nil
}
