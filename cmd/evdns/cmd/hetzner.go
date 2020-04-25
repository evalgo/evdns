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
	"errors"
	"fmt"

	"evalgo.org/evdns"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	apiURL  string
	token   string
	err     error
)

// hetznerCmd represents the hetzner command
var hetznerCmd = &cobra.Command{
	Use:   "hetzner",
	Short: "hetzner specific dns api calls",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(cfgFile) > 0 {
			apiURL = viper.GetString("url")
			token = viper.GetString("token")
		} else {
			apiURL, err = cmd.Flags().GetString("api")
			if err != nil {
				return err
			}
			token, err = cmd.Flags().GetString("token")
			if err != nil {
				return err
			}
		}
		zones, err := cmd.Flags().GetBool("zones")
		if err != nil {
			return err
		}
		h := evdns.NewHetzner(apiURL, token)
		switch true {
		case zones:
			hZones, err := h.Zones()
			if err != nil {
				return err
			}
			mZones := hZones.(map[string]interface{})
			for _, zone := range mZones["zones"].([]interface{}) {
				fmt.Println(zone.(map[string]interface{})["id"], zone.(map[string]interface{})["name"])
			}
			return nil
		}
		return errors.New("")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(hetznerCmd)
	hetznerCmd.Flags().StringVar(&cfgFile, "config", "./evdns.json", "config file (default is ./evdns.json)")
	hetznerCmd.Flags().String("url", "https://dns.hetzner.com/api/v1", "url to be used for api calls")
	hetznerCmd.Flags().String("token", "", "token to be used for api authorization")
	hetznerCmd.Flags().BoolP("zones", "z", false, "display zones")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("json")
	if len(cfgFile) > 0 {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("evdns.json")
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
