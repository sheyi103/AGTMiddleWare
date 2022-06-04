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

type USSDSubscriptionResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TransactionId string `json:"transactionId"`
	Data          struct {
		Id string `json:"subscriptionId"`
	} `json:"data"`
}
type USSDFlowResponse struct {
	StatusCode string `json:"statusCode"`

	Data struct {
		InboundResponse   string `json:"inboundResponse"`
		UserInputRequired bool   `json:"userInputRequired"`
		MessageType       int32  `json:"messageType"`
		ServiceCode       string `json:"serviceCode"`
		Msisdn            string `json:"msisdn"`
		SessionId         string `json:"sessionId"`
	} `json:"data"`

	StatusMessage string `json:"statusMessage"`

	Llink struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_link"`
}

type USSDDeleteSubscriptionResponse struct {
	// apiProductList     string `json:"api_product_list"`
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
	TransactionId string `json:"transactionId"`
	Data          struct {
		SubscriptionId string `json:"subscriptionId"`
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

func USSDSubscription(accessToken string, senderAddress string, notifyUrl string, targetSystem string) (USSDSubscriptionResponse, error) {

	var bearer = "Bearer " + accessToken
	url := "https://api.mtn.com/v1/messages/ussd/subscription"
	payload := map[string]interface{}{
		"serviceCode":  senderAddress,
		"callbackUrl":  notifyUrl,
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

	var ussdSubscriptionResponse USSDSubscriptionResponse
	json.Unmarshal(body, &ussdSubscriptionResponse)
	// response := authorization.AccessToken

	return ussdSubscriptionResponse, nil
}
func USSDNotifyUrl(sessionId, messageType, msisdn, serviceCode, ussdString, notifyUrl string) (USSDFlowResponse, error) {

	url := notifyUrl
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

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	fmt.Println("API Response as String:\n" + string(body))

	var ussdResponse USSDFlowResponse
	json.Unmarshal(body, &ussdResponse)

	return ussdResponse, nil
}

func USSDDeleteSubscription(accessToken string, subscriptionId string) (USSDDeleteSubscriptionResponse, error) {

	var bearer = "Bearer " + accessToken
	url := "https://prod.api.mtn.com/v1/messages/ussd/subscription/" + subscriptionId
	log.Println(url)

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

	var ussdDeleteSubscriptionResponse USSDDeleteSubscriptionResponse
	json.Unmarshal(body, &ussdDeleteSubscriptionResponse)
	// response := authorization.AccessToken

	return ussdDeleteSubscriptionResponse, nil
}
