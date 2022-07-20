package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	token, err := Token()
	if err != nil {
		// handle the error
		return
	}
	fmt.Println(token)
	return
}

func Token() (out string, err error) {
	tokenUrl := "https://hownd.auth0.com/oauth/token"
	// For example only, this is not secure.
	clientID := "your id"
	clientSecret := "your secret"

	p := `{
	 	"client_id":"` + clientID + `",
	 	"client_secret":"` + clientSecret + `",
	 	"audience":"https://partner-api.hownd.com",
		"grant_type":"client_credentials"
	}`

	payload := strings.NewReader(p)

	req, err := http.NewRequest("POST", tokenUrl, payload)
	if err != nil {
		// handle the error
		return
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle the error
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// handle the error
		return
	}

	var authResponse TokenResponse

	err = json.Unmarshal(body, &authResponse)
	if err != nil {
		// handle the error
		return
	}
	out = authResponse.AccessToken
	return
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
