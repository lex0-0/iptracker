/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace the IP",
	Long:  `Trace the IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("pleas provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)

}

type IP struct {
	IP       string `json:"iP"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	TimeZone string `json:"timeZone"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	responsebyte := getData(url)

	data := IP{}

	err := json.Unmarshal(responsebyte, &data)
	if err != nil {
		log.Println("unabel to unmarshal the response")
	}

	fmt.Println("Data Found: ")
	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOCATION :%s\nTIMEZONE:%s\nPOSTAL :%s\n", data.IP, data.City, data.Region, data.Country, data.Location, data.TimeZone, data.Postal)

}
func getData(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("unabel to get response")
	}

	responsebyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("unabel to read response")
	}
	return responsebyte
}
