// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"fmt"
)

type UserCredentialsNetworkType string

const (
	UserCredentialsNetworkTypeMTN     UserCredentialsNetworkType = "MTN"
	UserCredentialsNetworkTypeAIRTEL  UserCredentialsNetworkType = "AIRTEL"
	UserCredentialsNetworkTypeGLO     UserCredentialsNetworkType = "GLO"
	UserCredentialsNetworkType9MOBILE UserCredentialsNetworkType = "9MOBILE"
)

func (e *UserCredentialsNetworkType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserCredentialsNetworkType(s)
	case string:
		*e = UserCredentialsNetworkType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserCredentialsNetworkType: %T", src)
	}
	return nil
}

type UserCredentialsService string

const (
	UserCredentialsServiceSMS   UserCredentialsService = "SMS"
	UserCredentialsServiceUSSD  UserCredentialsService = "USSD"
	UserCredentialsServiceVOICE UserCredentialsService = "VOICE"
)

func (e *UserCredentialsService) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserCredentialsService(s)
	case string:
		*e = UserCredentialsService(s)
	default:
		return fmt.Errorf("unsupported scan type for UserCredentialsService: %T", src)
	}
	return nil
}

type UserCredentialsServiceType string

const (
	UserCredentialsServiceTypeDAILY    UserCredentialsServiceType = "DAILY"
	UserCredentialsServiceTypeWEEKLY   UserCredentialsServiceType = "WEEKLY"
	UserCredentialsServiceTypeMONTHLY  UserCredentialsServiceType = "MONTHLY"
	UserCredentialsServiceTypeONDEMAND UserCredentialsServiceType = "ON_DEMAND"
)

func (e *UserCredentialsServiceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserCredentialsServiceType(s)
	case string:
		*e = UserCredentialsServiceType(s)
	default:
		return fmt.Errorf("unsupported scan type for UserCredentialsServiceType: %T", src)
	}
	return nil
}

type Role struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type ShortCode struct {
	ID        int32        `json:"id"`
	ShortCode string       `json:"short_code"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type User struct {
	ID            int32        `json:"id"`
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	PhoneNumber   string       `json:"phone_number"`
	ContactPerson string       `json:"contact_person"`
	RoleID        int32        `json:"role_id"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
}

type UserCredential struct {
	ID                      int32                      `json:"id"`
	ClientID                string                     `json:"client_id"`
	ClientSecret            string                     `json:"client_secret"`
	ShortcodeID             int32                      `json:"shortcode_id"`
	UserID                  int32                      `json:"user_id"`
	ServiceName             sql.NullString             `json:"service_name"`
	ServiceID               sql.NullString             `json:"service_id"`
	Service                 UserCredentialsService     `json:"service"`
	ServiceType             UserCredentialsServiceType `json:"service_type"`
	ProductID               sql.NullString             `json:"product_id"`
	NodeID                  sql.NullString             `json:"node_id"`
	SubscriptionID          sql.NullString             `json:"subscription_id"`
	SubscriptionDescription sql.NullString             `json:"subscription_description"`
	BaseUrl                 sql.NullString             `json:"base_url"`
	DatasynEndpoint         sql.NullString             `json:"datasyn_endpoint"`
	NotificationEndpoint    sql.NullString             `json:"notification_endpoint"`
	NetworkType             UserCredentialsNetworkType `json:"network_type"`
	CreatedAt               sql.NullTime               `json:"created_at"`
	UpdatedAt               sql.NullTime               `json:"updated_at"`
}