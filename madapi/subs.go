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
	SubscriptionId          string  `json:"subscriptionId"`
	SubscriptionDescription string  `json:"subscriptionDescription"`
	AmountCharged           float32 `json:"amountCharged"`
	SendSMSNotification     bool    `json:"sendSMSNotification"`
	AutoRenew               bool    `json:"autoRenew"`
	AmountBefore            float32 `json:"amountBefore"`
	AmountAfter             float32 `json:"amountAfter"`
	CorrelationId           string  `json:"correlationId"`
	Cvmoffer                bool    `json:"cvmoffer"`
	StatusCode              string  `json:"statusCode"`
	StatusMessage           string  `json:"statusMessage"`
	TransactionId           string  `json:"transactionId"`
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

type DataSyncRequestParams struct {
	ServiceType   string `json:"serviceType"`
	ChargingMode  string `json:"chargingMode"`
	AppliedPlan   string `json:"appliedPlan"`
	ContentId     string `json:"contentId"`
	ResultCode    string `json:"resultCode"`
	RenFlag       string `json:"renFlag"`
	Result        string `json:"result"`
	ValidityType  string `json:"validityType"`
	SequenceNo    string `json:"sequenceNo"`
	CallingParty  string `json:"callingParty"`
	BearerId      string `json:"bearerId"`
	OperationId   string `json:"operationId"`
	RequestedPlan string `json:"requestedPlan"`
	ChargeAmount  string `json:"chargeAmount"`
	ServiceNode   string `json:"serviceNode"`
	ServiceId     string `json:"serviceId"`
	Category      string `json:"category"`
	ValidityDays  string `json:"validityDays"`
	// HttpStatus string `json:"httpStatus"`
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
	// if subscriptionResponse.

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

func DataSyncNotification(datasync string, arg DataSyncRequestParams) (int32, error) {

	log.Println("Inside the send out function")
	log.Println(datasync)
	log.Println(arg)
	// return http.StatusOK, nil
	url := "https://sdp.broadbased.net/TestSP/public/index.php/api/messages/sms/subscription/callback"
	payload := map[string]interface{}{
		"serviceType":   arg.ServiceType,
		"chargingMode":  arg.ChargingMode,
		"appliedPlan":   arg.AppliedPlan,
		"contentId":     arg.ContentId,
		"resultCode":    arg.ResultCode,
		"renFlag":       arg.RenFlag,
		"result":        arg.Result,
		"validityType":  arg.ValidityType,
		"sequenceNo":    arg.SequenceNo,
		"callingParty":  arg.CallingParty,
		"bearerId":      arg.BearerId,
		"operationId":   arg.OperationId,
		"requestedPlan": arg.RequestedPlan,
		"chargeAmount":  arg.ChargeAmount,
		"serviceNode":   arg.ServiceNode,
		"serviceId":     arg.ServiceId,
		"category":      arg.Category,
		"validityDays":  arg.ValidityDays,
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
	// fmt.Println("response Body:", string(body))
	fmt.Println("API Response as String:\n" + string(body))

	// var subscriptionResponse SubscriptionResponse
	// json.Unmarshal(body, &subscriptionResponse)
	// response := authorization.AccessToken

	return http.StatusAccepted, nil

}
