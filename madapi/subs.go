package madapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SubscriptionResponse struct {
	// apiProductList     string `json:"api_product_list"`
	SubscriptionId          string `json:"subscriptionId"`
	SubscriptionDescription string `json:"subscriptionDescription"`
	AmountCharged           string `json:"amountCharged"`
	SendSMSNotification     string `json:"sendSMSNotification"`
	AutoRenew               string `json:"autoRenew"`
	AmountBefore            string `json:"amountBefore"`
	AmountAfter             string `json:"amountAfter"`
	CorrelationId           string `json:"correlationId"`
	Cvmoffer                string `json:"cvmoffer"`
	StatusCode              string `json:"statusCode"`
	StatusMessage           string `json:"statusMessage"`
	TransactionId           string `json:"transactionId"`
}

type UnSubscriptionResponse struct {
	// apiProductList     string `json:"api_product_list"`
	SubscriptionId          string `json:"subscriptionId"`
	SubscriptionDescription string `json:"subscriptionDescription"`
	AmountCharged           string `json:"amountCharged"`
	SendSMSNotification     string `json:"sendSMSNotification"`
	AutoRenew               string `json:"autoRenew"`
	AmountBefore            string `json:"amountBefore"`
	AmountAfter             string `json:"amountAfter"`
	CorrelationId           string `json:"correlationId"`
	Cvmoffer                string `json:"cvmoffer"`
	StatusCode              string `json:"statusCode"`
	StatusMessage           string `json:"statusMessage"`
	TransactionId           string `json:"transactionId"`
}

func Subscription(clientId, subscriptionProviderId, nodeId, subscriptionId, subscriptionDescription, registrationChannel, customerId, transactionId string) (SubscriptionResponse, error) {

	url := "https://prod-nigeria.api.mtn.com/v2/customers/" + customerId + "/subscriptions/"
	payload := map[string]interface{}{
		"subscriptionProviderId":  subscriptionProviderId,
		"nodeId":                  nodeId,
		"subscriptionId":          subscriptionId,
		"subscriptionDescription": subscriptionDescription,
		"registrationChannel":     registrationChannel,
	}

	bytesRepresentation, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("transactionId", transactionId)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-API-key", clientId)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))
	fmt.Println("API Response as String:\n" + string(body))

	var subscriptionResponse SubscriptionResponse
	json.Unmarshal(body, &subscriptionResponse)
	// response := authorization.AccessToken

	return subscriptionResponse, nil
}

func UnSubscription(clientId, subscriptionProviderId, subscriptionId, subscriptionDescription, customerId, transactionId string) (SubscriptionResponse, error) {

	url := "https://preprod-nigeria.api.mtn.com/v2/customers/" + customerId + "/subscriptions/" + subscriptionId + "?subscriptionProviderId=" + subscriptionProviderId + "&description=" + subscriptionDescription

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("transactionId", transactionId)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-API-key", clientId)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))
	fmt.Println("API Response as String:\n" + string(body))

	var subscriptionResponse SubscriptionResponse
	json.Unmarshal(body, &subscriptionResponse)
	// response := authorization.AccessToken

	return subscriptionResponse, nil
}
