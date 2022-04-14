package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Stores all functions to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function withing a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// OnboardTxParams contains the input parameters of the Onboard transaction
type OnboardTxParams struct {
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	ShortcodeID   int32  `json:"shortcode_id"`
	UserID        int32  `json:"user_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	ContactPerson string `json:"contact_person"`
	ShortCode     int32  `json:"short_code"`
	// ServiceName             sql.NullString `json:"service_name"`
	// ServiceID               sql.NullString `json:"service_id"`
	// Service                 sql.NullString `json:"service"`
	// ServiceType             sql.NullString `json:"service_type"`
	// ProductID               sql.NullString `json:"product_id"`
	// NodeID                  sql.NullString `json:"node_id"`
	// SubscriptionID          sql.NullString `json:"subscription_id"`
	// SubscriptionDescription sql.NullString `json:"subscription_description"`
	// BaseUrl                 sql.NullString `json:"base_url"`
	// DatasynEndpoint         sql.NullString `json:"datasyn_endpoint"`
	// NotificationEndpoint    sql.NullString `json:"notification_endpoint"`
	// NetworkType             sql.NullString `json:"network_type"`
}

// type OnboardTxResult struct {
// 	UserCredential UserCredential `json:"user_credential"`
// 	ShortCode      ShortCode      `json:"short_code"`
// 	User           User           `json:"user"`
// }

// OnboardTx allows us to create a user,ShortCode and merge it into user Credential tables.
// Create user record, creates shortcode and merge into User credential Table

// func (store *Store) OnboardTx(ctx context.Context, arg OnboardTxParams) (OnboardTxResult, error) {
// 	var result OnboardTxResult

// 	userArg := CreateUserParams{

// 	}

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		result, err = q.CreateUser(ctx, arg)
// 	})

// }
