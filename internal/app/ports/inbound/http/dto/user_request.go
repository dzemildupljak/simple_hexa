package httpdto

// CreateUserRequest is used for user registration.
// @Description DTO for user creation containing username and email.
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
