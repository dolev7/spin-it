package users

import (
	"encoding/json"
	"github.com/dolev7/spin-it/pkg/auth"
	"github.com/dolev7/spin-it/pkg/logger"
	"net/http"
)

// UserAuthRequest represents the request body for both signup and login
type UserAuthRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword"`
}

// UserAuthResponse represents the response returned after authentication
type UserAuthResponse struct {
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	Message string `json:"message" example:"Authentication successful"`
}

// UserProfileResponse represents the user profile response
type UserProfileResponse struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
}

// SignupHandler handles user signup requests.
//
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags user
// @Accept json
// @Produce json
// @Param request body UserAuthRequest true "User signup data"
// @Success 201 {object} UserAuthResponse
// @Failure 400 {string} string ErrorResponse
// @Router /users [post]
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req UserAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Invalid request payload: ", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	logger.Log.Infof("Signup request received for email: %s", req.Email)

	// Create user in the database
	err := CreateUser(req.Email, req.Password)
	if err != nil {
		logger.Log.Error("Signup failed: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(req.Email)
	if err != nil {
		logger.Log.Error("Failed to generate JWT token: ", err)
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Return JWT token in response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserAuthResponse{
		Token:   token,
		Message: "User created successfully",
	})
}

// LoginHandler handles user authentication
//
// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param request body UserAuthRequest true "User credentials"
// @Success 200 {object} UserAuthResponse
// @Failure 401 {string} string ErrorResponse
// @Router /users/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req UserAuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate user in database
	user, err := GetUserByEmail(req.Email)
	if err != nil || !CheckPasswordHash(req.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(req.Email)
	if err != nil {
		logger.Log.Error("Failed to generate JWT token: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Could not generate token", http.StatusBadRequest)
		return
	}

	// Return JWT
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserAuthResponse{
		Token:   token,
		Message: "Login successful",
	})

}

// GetProfileHandler handles fetching the logged-in user's profile.
//
// @Summary Get user profile
// @Description Returns the authenticated user's profile information
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} UserProfileResponse
// @Failure 401 {string} string ErrorResponse
// @Router /api/users/me [get]
// @Security BearerAuth
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract email from JWT token stored in the request context
	email := r.Header.Get("X-User-Email")

	if email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthorized", http.StatusBadRequest)
		return
	}

	// Fetch user from database
	user, err := GetUserByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	// Return user profile
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserProfileResponse{
		Email: user.Email,
		ID:    user.ID,
	})
}
