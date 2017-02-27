// Copyright © 2017 Jun Sato <jnst@outlook.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"encoding/binary"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "qb",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open("/Users/jnst/Desktop/castle_lv1.qb")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		b := make([]byte, 1)
		_, err = f.Read(b)
		if err != nil {
			log.Fatal(err)
		}

		// Version
		var major uint8
		err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, &major)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		var minor uint8
		err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, &minor)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		var release uint8
		err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, &release)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		var build uint8
		err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, &build)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%d.%d.%d.%d\n", major, minor, release, build)

		//_, err = f.Read(b)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//// Color Format
		//colorFormat := string(b)
		//log.Println(colorFormat)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-qb.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-qb") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
