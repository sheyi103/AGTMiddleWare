package madapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SMSOutboundResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TransactionId string `json:"transactionId"`
	Data          struct {
		RequestId        string `json:"requestId"`
		ClientCorrelator string `json:"clientCorrelator"`
	} `json:"data"`
}

type SMSSubscriptionResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TransactionId string `json:"transactionId"`
	Data          struct {
		Id string `json:"id"`
	} `json:"data"`
}

type SMSDeleteSubscriptionResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TransactionId string `json:"transactionId"`
	Data          struct {
	} `json:"data"`
}

func SendOutBoundSMS(accessToken string, clientCorrelator string, message string, receiverAddress []string, senderAddress string) (SMSOutboundResponse, error) {

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

	var smsOutboundResponse SMSOutboundResponse
	json.Unmarshal(body, &smsOutboundResponse)
	// response := authorization.AccessToken

	return smsOutboundResponse, nil
}

func SMSSubscription(accessToken string, senderAddress string, notifyUrl string, targetSystem string) (SMSSubscriptionResponse, error) {

	var bearer = "Bearer " + accessToken
	url := "https://preprod.api.mtn.com/v2/messages/sms/outbound/" + senderAddress + "/subscription"
	payload := map[string]interface{}{
		"notifyUrl":    notifyUrl,
		"targetSystem": targetSystem,
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

	var smsSubscriptionResponse SMSSubscriptionResponse
	json.Unmarshal(body, &smsSubscriptionResponse)
	// response := authorization.AccessToken

	return smsSubscriptionResponse, nil
}

func SMSNotifyUrl(senderAddress string, receiverAddress string, message string, created int64, notifyUrl string) (int32, error) {

	url := notifyUrl
	payload := map[string]interface{}{
		"senderAddress":   senderAddress,
		"receiverAddress": receiverAddress,
		"message":         message,
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

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// body, _ := ioutil.ReadAll(response.Status)
	// fmt.Println("response Body:", string(body))
	// fmt.Println("API Response as String:\n" + string(body))

	// var smsSubscriptionResponse SMSSubscriptionResponse
	// json.Unmarshal(body, &smsSubscriptionResponse)
	// response := authorization.AccessToken

	return http.StatusOK, nil
}

func SMSDeleteSubscription(accessToken string, senderAddress string, subscriptionId string) (SMSDeleteSubscriptionResponse, error) {

	var bearer = "Bearer " + accessToken
	url := "https://preprod.api.mtn.com/v2/messages/sms/outbound/" + senderAddress + "/subscription/" + subscriptionId

	request, err := http.NewRequest("DELETE", url, nil)
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

	var smsDeleteSubscriptionResponse SMSDeleteSubscriptionResponse
	json.Unmarshal(body, &smsDeleteSubscriptionResponse)
	// response := authorization.AccessToken

	return smsDeleteSubscriptionResponse, nil
}
