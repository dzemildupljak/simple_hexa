// **domain:** Defines business domain models and repository interfaces.
package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
