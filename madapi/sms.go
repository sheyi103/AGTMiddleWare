package madapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SMSResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TransactionId string `json:"transactionId"`
	Data          struct {
		RequestId        string `json:"requestId"`
		ClientCorrelator string `json:"clientCorrelator"`
	} `json:"data"`
}

func SendOutBoundSMS(accessToken string, clientCorrelator string, message string, receiverAddress []string, senderAddress string) (SMSResponse, error) {

	var bearer = "Bearer " + accessToken
	url := "https://preprod.api.mtn.com/v2/messages/sms/outbound"
	payload := map[string]interface{}{
		"clientCorrelator": clientCorrelator,
		"message":          message,
		"receiverAddress":  receiverAddress,
		"senderAddress":    senderAddress,
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

	var smsResponse SMSResponse
	json.Unmarshal(body, &smsResponse)
	// response := authorization.AccessToken

	return smsResponse, nil
}
