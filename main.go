package main

import (
	"encoding/json"
	"fmt"
	"grafana2ufo/config"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/hashicorp/go-retryablehttp"
)

// GrafanaWebhook is the message format for the Grafana Alert Webhook
type GrafanaWebhook struct {
	DashboardID int `json:"dashboardId"`
	EvalMatches []struct {
		Value  int    `json:"value"`
		Metric string `json:"metric"`
		Tags   struct {
		} `json:"tags"`
	} `json:"evalMatches"`
	ImageURL string `json:"imageUrl"`
	Message  string `json:"message"`
	OrgID    int    `json:"orgId"`
	PanelID  int    `json:"panelId"`
	RuleID   int    `json:"ruleId"`
	RuleName string `json:"ruleName"`
	RuleURL  string `json:"ruleUrl"`
	State    string `json:"state"`
	Tags     struct {
		TagName string `json:"tag name"`
	} `json:"tags"`
	Title string `json:"title"`
}

func main() {

	// Define config file
	viper.SetDefault("listenOn", ":5022")
	viper.SetDefault("ufoAddr", "10.0.20.26")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // look for config in the working directory

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	var configuration config.Configuration
	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Log service informations on startup
	log.Printf("Service is listening on: '%s'", configuration.ListenOn)
	log.Printf("UFO server defined is: '%s'", configuration.UFOAddress)
	var grafanaMessage GrafanaWebhook

	callUFO(configuration.UFOAddress, configuration.UFOReset)

	log.Fatal(http.ListenAndServe(configuration.ListenOn, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("MSG RECEIVED\n")
		grafanaJSON, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		log.Printf(string(grafanaJSON))

		err = json.Unmarshal(grafanaJSON, &grafanaMessage)
		if err != nil {
			log.Println("Error Unmarshalling")
		}

		log.Printf("DashboardID: %v", grafanaMessage.DashboardID)
		if grafanaMessage.DashboardID == 2 {
			log.Printf("RuleID: %v\tRuleName: %v\n", grafanaMessage.RuleID, grafanaMessage.RuleName)
			if grafanaMessage.State == "alerting" {
				callUFO(configuration.UFOAddress, configuration.CameraPollerDown)
			} else if grafanaMessage.State == "ok" {
				callUFO(configuration.UFOAddress, configuration.CameraPollerUp)
			}
		}

	})))
}

func callUFO(ufoAddress string, ufoParams string) int {

	var ufoEndpoint = ufoAddress + ufoParams
	log.Printf("ufoEndpoint: %s\n", ufoEndpoint)
	response, err := retryablehttp.Get(ufoEndpoint)
	if err != nil {
		log.Printf("FATAL ERROR\n")
		panic(err)
	}
	log.Printf("UFO Response: %v", response.StatusCode)
	return response.StatusCode

}
