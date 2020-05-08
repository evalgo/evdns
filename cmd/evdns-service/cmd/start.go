/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"evalgo.org/evdns"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	webRoot      string
	address      string
	hetznerURL   string
	hetznerToken string
	clientID     string
	clientSecret string
	err          error
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		webRoot, err = cmd.Flags().GetString("webroot")
		if len(cfgFile) == 0 {
			if err != nil {
				return err
			}
			address, err = cmd.Flags().GetString("address")
			if err != nil {
				return err
			}
			hetznerURL, err = cmd.Flags().GetString("url")
			if err != nil {
				return err
			}
			hetznerToken, err = cmd.Flags().GetString("token")
			if err != nil {
				return err
			}
			clientID, err = cmd.Flags().GetString("id")
			if err != nil {
				return err
			}
			clientSecret, err = cmd.Flags().GetString("secret")
			if err != nil {
				return err
			}

		} else {
			webRoot = viper.GetString("webroot")
			address = viper.GetString("address")
			hetznerURL = viper.GetString("url")
			hetznerToken = viper.GetString("token")
			clientID = viper.GetString("id")
			clientSecret = viper.GetString("secret")
		}
		h := evdns.NewHetzner(hetznerURL, hetznerToken)
		return h.WSStart(address, clientID, clientSecret, webRoot)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /opt/evalgo.org/evdns-service/conf/evdns.json)")
	startCmd.Flags().String("address", "0.0.0.0:8989", "service starting address contains {ip}:{port}")
	startCmd.Flags().String("url", "https://dns.hetzner.com/api/v1", "url to connect this service with")
	startCmd.Flags().String("token", "", "access token for the given url")
	startCmd.Flags().String("id", "evdns", "client id")
	startCmd.Flags().String("secret", "secret", "client secret")
	startCmd.Flags().String("webroot", "./webroot", "path to frontend webroot")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("json")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/opt/evalgo.org/evdns-service/conf")
		viper.SetConfigName("evdns.json")
	}
	viper.AutomaticEnv()
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
