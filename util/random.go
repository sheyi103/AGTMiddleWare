package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt generates a ramdon integer between min and maximum
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

//return base Url
func BaseUrl() string {
	baseUrl := "http://localhost:8080/"
	return baseUrl
}

//return Datasync Url
func DataSyncUrl() string {
	datasync := "http://localhost:8080/datasync"
	return datasync
}

//return notification Url
func NotificationUrl() string {
	notification := "http://localhost:8080/notification"
	return notification
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)

	}
	return sb.String()
}

//Generates random role from roles
func RandomRole() string {
	roles := []string{"ADMIN", "AGT", "SP"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

//Generates random protocol from interfaces
func RandomProtocol() string {
	protocol := []string{"HTTP", "SOAP", "SMPP"}
	n := len(protocol)
	return protocol[rand.Intn(n)]
}

//Generates random service from services
func RandomService() string {
	service := []string{"SMS", "USSD", "VOICE", "SUBCRIPTION"}
	n := len(service)
	return service[rand.Intn(n)]
}

//Generates random serviceType from serviceTypes
func RandomServiceType() string {
	serviceType := []string{"DAILY", "WEEKLY", "MONTHLY", "ON-DEMAND"}
	n := len(serviceType)
	return serviceType[rand.Intn(n)]
}

//Generates random network from networkTypes
func RandomNetworkType() string {
	serviceType := []string{"MTN", "AIRTEL", "GLO", "9MOBILE"}
	n := len(serviceType)
	return serviceType[rand.Intn(n)]
}

//Generates random phone number from numbers
func RandomPhoneNumber() string {
	PhoneNumber := []string{"2348026425250", "2348038024350", "23435156407"}
	n := len(PhoneNumber)
	return PhoneNumber[rand.Intn(n)]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// GetClientIP gets the correct IP for the end client instead of the proxy
func GetClientIP(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

// GetDurationInMillseconds takes a start time and returns a duration in milliseconds
func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
