package coinspaid

import "fmt"

type APIError struct {
	StatusCode int
	// Errors contains api errors with descriptions.
	Errors map[string]string `json:"errors"`
}

// Error implements error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("api error [%d]: %v", e.StatusCode, e.Errors)
}

type AuthorizationError struct {
	Err  string `json:"error"`
	Code string `json:"code"`
}

// Error implements error interface.
func (e *AuthorizationError) Error() string {
	return fmt.Sprintf("authorization error: %s [%s]", e.Err, e.Code)
}
