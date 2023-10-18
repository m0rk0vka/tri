/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tri",
	Short: "Tri is a todo application",
	Long:  `Tri will help you get more done in less time. It's designed as simple as possible to help you accomplish your goals`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set dataFile using --datafile.")
	}
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tridos.json", "data file to store todos")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tri.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
