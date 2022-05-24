package madapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type USSDResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode     string `json:"statusCode"`
	StatusMessage  string `json:"statusMessage"`
	SupportMessage string `json:"supportMessage"`
	TransactionId  string `json:"transactionId"`
	Data           struct {
		OutboundResponse string `json:"outboundResponse"`
		SessionId        string `json:"sessionId"`
		Msisdn           string `json:"msisdn"`
	} `json:"data"`
}

func SendOutBoundUSSD(accessToken string, sessionId string, messageType string, msisdn string, serviceCode string, ussdString string) (USSDResponse, error) {

	var bearer = "Bearer " + accessToken
	url := "https://preprod.api.mtn.com/v1/messages/ussd/outbound"
	payload := map[string]interface{}{
		"sessionId":   sessionId,
		"messageType": messageType,
		"msisdn":      msisdn,
		"serviceCode": serviceCode,
		"ussdString":  ussdString,
	}

	bytesRepresentation, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))
	fmt.Println("API Response as String:\n" + string(body))

	var ussdResponse USSDResponse
	json.Unmarshal(body, &ussdResponse)
	// response := authorization.AccessToken

	return ussdResponse, nil
}
