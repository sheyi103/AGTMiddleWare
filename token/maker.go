package token

import "time"

//MAker is an interface for managing tokens
type Maker interface {

	//CreatesToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	//VerofyToken check if the token is valid or not
	verifyToken(token string) (*Payload, error)
}
