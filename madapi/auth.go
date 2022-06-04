package madapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TokenDetails struct {
	// apiProductList     string `json:"api_product_list"`
	ApiProductListJson string `json:"api_product_list"`
	OrganizationName   string `json:"organization_name"`
	DeveloperEmail     string `json:"developer.email"`
	TokenType          string `json:"token_type"`
	IssuedType         string `json:"issued_at"`
	ClientID           string `json:"client_id"`
	AccessToken        string `json:"access_token"`
	ApplicationName    string `json:"application_name"`
	// scope              string `json:"scope"`
	ExpiresIN string `json:"expires_in"`
	// refreshCount       string `json:"refresh_count"`
	Status string `json:"status"`
}

func Authorization(clientId, clientSecret string) (string, error) {
	url := "https://api.mtn.com/v1/oauth/access_token/accesstoken?grant_type=client_credentials"

	s := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", clientId, clientSecret)

	payload := strings.NewReader(s)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(string(body))

	fmt.Println("API Response as String:\n" + string(body))

	var authorization TokenDetails
	json.Unmarshal(body, &authorization)
	accessToken := authorization.AccessToken

	return accessToken, nil
}
