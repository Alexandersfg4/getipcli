package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const URL = "http://ip-api.com/json/"

type Error string

func (e Error) Error() string { return string(e) }

const ErrorParsingModel = Error("Error happened while parsing model")

type IpModel struct {
	Status   string `json:"status"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Timezone string `json:"timezone"`
	Query    string `jsson:"query"`
}

func GetBody() []byte {
	response, err := http.Get(URL)
	if err != nil {
		log.Default().Fatalln("Error happened while making a request")
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Default().Fatalln("Error happened while reading a body")
	}
	return content
}

func PrepareIpModel(body []byte) (*IpModel, error) {
	var result IpModel
	if err := json.Unmarshal(body, &result); err != nil {
		log.Default().Fatalln("Can not inmarshal JSON")
	}
	if result.Status != "success" {
		return &result, ErrorParsingModel
	}
	return &result, nil
}

func PrintIpInfo(ipModel *IpModel) {
	fmt.Printf("Your IP is %s\nCountry: %s\nCity: %s\n", ipModel.Query, ipModel.Country, ipModel.City)
}

func main() {
	body := GetBody()
	ipModel, err := PrepareIpModel(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	PrintIpInfo(ipModel)
}
