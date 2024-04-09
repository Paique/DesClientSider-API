package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetdefaultKeys(contra bool) []string {

	var resp *http.Response
	var err error

	var listType = "keywords"

	if contra {
		listType = "contra"
	}

	url := fmt.Sprintf("https://raw.githubusercontent.com/Paique/DesClientSider-API/main/keywords/recommended_%s.txt", listType)

	resp, err = http.Get(url)

	// Error checking of the http.Get() request
	if err != nil {
		log.Fatal(err)
	}

	// Resource leak if response body isn't closed
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	// Error checking of the ioutil.ReadAll() request
	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)

	keysList := strings.Split(bodyString, "\n")

	for i, key := range keysList {
		if key == "" || key == " " {
			keysList = append(keysList[:i], keysList[i+1:]...)
		}
	}

	return keysList

}
