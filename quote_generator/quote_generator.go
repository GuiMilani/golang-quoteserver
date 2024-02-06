package quote_generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseKanye struct {
 Quote     string `json:"quote"`
}

type ResponseTaylor struct {
 Quote     string `json:"quote"`
 Song      string `json:"song"`
 Album     string `json:"album"`
}

func GetQuote(artist string) string {
	fmt.Println("Calling API...")
	
	client := &http.Client{}

	var api_adress string

	if artist == "kanye" {
		api_adress = "https://api.kanye.rest/"
	} else if artist == "taylor" {
		api_adress = "https://taylorswiftapi.herokuapp.com/get"
	} else {
		return "Error selecting artist."
	}
		
	req, err := http.NewRequest("GET", api_adress, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	if artist == "kanye" {
		var responseObject ResponseKanye
		json.Unmarshal(bodyBytes, &responseObject)
		return "\"" + responseObject.Quote + "\"\n-Kanye West"
	} else if artist == "taylor" {
		var responseObject ResponseTaylor
		json.Unmarshal(bodyBytes, &responseObject)
		return "\"" + responseObject.Quote + "\"\n-Taylor Swift"
	} else {
		return "Error selecting artist."
	}

}