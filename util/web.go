package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetdefaultList() []string {

	resp, err := http.Get("https://raw.githubusercontent.com/Paique/DesClientSider-API/main/recommended_keywords.txt")

	// Error checking of the http.Get() request
	if err != nil {
		log.Fatal(err)
	}

	// Resource leak if response body isn't closed
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	// Error checking of the ioutil.ReadAll() request
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)

	return strings.Split(bodyString, "\n")

}
