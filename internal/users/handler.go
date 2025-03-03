package users

import (
	"encoding/json"
	"github.com/dolev7/spin-it/pkg/logger"
	"net/http"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message" example:"User created successfully"`
}

// SignupHandler handles user signup requests.
//
// @Summary Create a new user
// @Description Create a new user in the system
// @Tags user
// @Accept json
// @Produce json
// @Param request body SignupRequest true "User signup data"
// @Success 201 {object} SignupResponse
// @Failure 400 {string} - "Invalid request"
// @Router /users [post]
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Invalid request payload: ", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	logger.Log.Infof("Signup request received for email: %s", req.Email)
	err := CreateUser(req.Email, req.Password)
	if err != nil {
		logger.Log.Error("Signup failed: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Log.Infof("User %s created successfully", req.Email)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
	if err != nil {
		return
	}
}
