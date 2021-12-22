package fp1

import (
	"io/ioutil"
	"log"
	"net/http"
)

const DefaultRestUrl string = "https://api.foxpass.com/v1/users/"

func fp(url, token string) {
	//url := "https://api.foxpass.com/v1/users/"

	// Create a Bearer string by appending string access token
	//var bearer = "Bearer " + <ACCESS TOKEN HERE>

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", "Token "+token)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}
