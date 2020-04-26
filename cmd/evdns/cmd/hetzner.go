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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

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
		records, err := cmd.Flags().GetBool("records")
		if err != nil {
			return err
		}
		zone, err := cmd.Flags().GetBool("zone")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		ID, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		rID, err := cmd.Flags().GetString("rid")
		if err != nil {
			return err
		}
		create, err := cmd.Flags().GetBool("create")
		if err != nil {
			return err
		}
		del, err := cmd.Flags().GetBool("delete")
		if err != nil {
			return err
		}
		record, err := cmd.Flags().GetBool("record")
		if err != nil {
			return err
		}
		rType, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		rValue, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}
		rTTL, err := cmd.Flags().GetInt("ttl")
		if err != nil {
			return err
		}
		update, err := cmd.Flags().GetBool("update")
		if err != nil {
			return err
		}
		export, err := cmd.Flags().GetBool("export")
		if err != nil {
			return err
		}
		validate, err := cmd.Flags().GetBool("validate")
		if err != nil {
			return err
		}
		zImport, err := cmd.Flags().GetBool("import")
		if err != nil {
			return err
		}
		h := evdns.NewHetzner(apiURL, token)
		switch true {
		case zImport:
			zFile, err := ioutil.ReadFile(rValue)
			if err != nil {
				return err
			}
			imported, err := h.ImportZone(ID, zFile)
			if err != nil {
				return err
			}
			zDetails := imported.(map[string]interface{})["zone"].(map[string]interface{})
			displayZone(zDetails)
			return nil
		case validate:
			zFile, err := ioutil.ReadFile(rValue)
			if err != nil {
				return err
			}
			validated, err := h.ValidateZone(zFile)
			if err != nil {
				return err
			}
			vInfo := validated.(map[string]interface{})
			fmt.Println("parsed records: ", vInfo["paresd_records"])
			if pErr, ok := vInfo["error"]; ok {
				fmt.Println(pErr.(map[string]interface{})["code"], pErr.(map[string]interface{})["message"])
			} else {
				for _, vr := range vInfo["valid_records"].([]interface{}) {
					rInfo := vr.(map[string]interface{})
					fmt.Println(rInfo["type"], rInfo["name"], rInfo["value"])
				}
			}
			return nil
		case export:
			exportFile, err := h.ExportZone(map[string]interface{}{"id": ID})
			if err != nil {
				return err
			}
			fmt.Println("write zone file to", ID+".zone")
			return ioutil.WriteFile(ID+".zone", exportFile.([]byte), 0777)
		case update:
			switch true {
			case zone:
				updated, err := h.UpdateZone(map[string]interface{}{"name": name, "ttl": rTTL, "id": ID})
				if err != nil {
					return err
				}
				zDetails := updated.(map[string]interface{})["zone"].(map[string]interface{})
				displayZone(zDetails)
				return nil
			case record:
				updated, err := h.UpdateRecord(map[string]interface{}{
					"id":      rID,
					"zone_id": ID,
					"type":    rType,
					"name":    name,
					"value":   rValue,
					"ttl":     rTTL,
				})
				if err != nil {
					return err
				}
				rDetails := updated.(map[string]interface{})["record"].(map[string]interface{})
				displayRecord(rDetails)
				return nil
			case records:
				var records interface{}
				err := json.Unmarshal([]byte(rValue), &records)
				if err != nil {
					return err
				}
				updated, err := h.UpdateRecords(records)
				if err != nil {
					return err
				}
				rInfo := updated.(map[string]interface{})
				fmt.Println("parsed records: ", len(rInfo["records"].([]interface{})))
				if pErr, ok := rInfo["error"]; ok {
					fmt.Println(pErr.(map[string]interface{})["code"], pErr.(map[string]interface{})["message"])
				} else {
					for _, vr := range rInfo["records"].([]interface{}) {
						rrInfo := vr.(map[string]interface{})
						fmt.Println(rrInfo["type"], rrInfo["name"], rrInfo["value"])
					}
				}
				return nil
			}
		case create:
			switch true {
			case zone:
				created, err := h.NewZone(map[string]interface{}{"name": name})
				if err != nil {
					return err
				}
				zDetails := created.(map[string]interface{})["zone"].(map[string]interface{})
				displayZone(zDetails)
				return nil
			case record:
				created, err := h.NewRecord(map[string]interface{}{
					"zone_id": ID,
					"type":    rType,
					"name":    name,
					"value":   rValue,
					"ttl":     rTTL,
				})
				if err != nil {
					return err
				}
				displayRecord(created.(map[string]interface{})["record"].(map[string]interface{}))
				return nil
			case records:
				var records interface{}
				err := json.Unmarshal([]byte(rValue), &records)
				if err != nil {
					return err
				}
				created, err := h.NewRecords(records)
				if err != nil {
					return err
				}

				rInfo := created.(map[string]interface{})
				fmt.Println("parsed records: ", len(rInfo["records"].([]interface{})))
				if pErr, ok := rInfo["error"]; ok {
					fmt.Println(pErr.(map[string]interface{})["code"], pErr.(map[string]interface{})["message"])
				} else {
					for _, vr := range rInfo["records"].([]interface{}) {
						rrInfo := vr.(map[string]interface{})
						fmt.Println(rrInfo["type"], rrInfo["name"], rrInfo["value"])
					}
				}
				return nil
			}
		case del:
			switch true {
			case zone:
				deleted, err := h.DeleteZone(ID)
				if err != nil {
					return err
				}
				fmt.Println(deleted)
				return nil
			case record:
				deleted, err := h.DeleteRecord(ID)
				if err != nil {
					return err
				}
				fmt.Println(deleted)
				return nil
			}
		case zones:
			hZones, err := h.Zones()
			if err != nil {
				return err
			}
			mZones := hZones.(map[string]interface{})
			for _, zone := range mZones["zones"].([]interface{}) {
				fmt.Println(zone.(map[string]interface{})["id"], zone.(map[string]interface{})["name"])
				//fmt.Println(zone.(map[string]interface{})["id"], zone.(map[string]interface{})["project"], zone.(map[string]interface{})["status"], zone.(map[string]interface{})["ttl"], zone.(map[string]interface{})["is_secondary_dns"])
				//fmt.Println("")
			}
			return nil
		case zone:
			zone, err := h.Zone(ID)
			if err != nil {
				return err
			}
			zDetails := zone.(map[string]interface{})["zone"].(map[string]interface{})
			displayZone(zDetails)
			return nil
		case record:
			rec, err := h.Zone(ID)
			if err != nil {
				return err
			}
			zDetails := rec.(map[string]interface{})["record"].(map[string]interface{})
			displayZone(zDetails)
			return nil
		case records:
			hZRecords, err := h.Records(ID)
			if err != nil {
				return err
			}
			mRecords := hZRecords.(map[string]interface{})
			for _, rec := range mRecords["records"].([]interface{}) {
				displayRecord(rec.(map[string]interface{}))
			}
			return nil
		}
		return errors.New("")
	},
}

func displayRecord(rDetails map[string]interface{}) {
	fmt.Println(rDetails["zone_id"], rDetails["id"], rDetails["type"], rDetails["name"], rDetails["value"], rDetails["ttl"])
}

func displayZone(zDetails map[string]interface{}) {
	fmt.Println(zDetails["name"])
	fmt.Println("---------")
	fmt.Println("id:", zDetails["id"])
	fmt.Println("project:", zDetails["project"])
	fmt.Println("records_count:", zDetails["records_count"])
	fmt.Println("created:", zDetails["created"])
	fmt.Println("modified:", zDetails["modified"])
	fmt.Println("verified:", zDetails["verified"])
	fmt.Println("status:", zDetails["status"])
	fmt.Println("owner:", zDetails["owner"])
	fmt.Println("paused:", zDetails["paused"])
	fmt.Println("ttl:", zDetails["ttl"])
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(hetznerCmd)
	hetznerCmd.Flags().StringVar(&cfgFile, "config", "./evdns.json", "config file (default is ./evdns.json)")
	hetznerCmd.Flags().String("url", "https://dns.hetzner.com/api/v1", "url to be used for api calls")
	hetznerCmd.Flags().String("token", "", "token to be used for api authorization")
	hetznerCmd.Flags().String("id", "", "id to be used in zones and record commands")
	hetznerCmd.Flags().String("rid", "", "id to be used in the record update command")
	hetznerCmd.Flags().String("name", "", "name to be used in create commands")
	hetznerCmd.Flags().String("type", "A", "record type")
	hetznerCmd.Flags().String("value", "", "record value")
	hetznerCmd.Flags().Int("ttl", 86400, "record ttl")
	hetznerCmd.Flags().BoolP("create", "", false, "create a zone or record")
	hetznerCmd.Flags().BoolP("update", "", false, "update a zone or record")
	hetznerCmd.Flags().BoolP("zones", "z", false, "display zones")
	hetznerCmd.Flags().BoolP("records", "r", false, "display records")
	hetznerCmd.Flags().BoolP("zone", "", false, "zone")
	hetznerCmd.Flags().BoolP("record", "", false, "record")
	hetznerCmd.Flags().BoolP("delete", "", false, "delete a zone or record")
	hetznerCmd.Flags().BoolP("export", "", false, "export a zone to a file")
	hetznerCmd.Flags().BoolP("import", "", false, "import a zone file")
	hetznerCmd.Flags().BoolP("validate", "", false, "validate a zone file")
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
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
