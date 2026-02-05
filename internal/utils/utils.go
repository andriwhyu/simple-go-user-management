package utils

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetStringEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

// IsValidEmail performs basic email validation
func IsValidEmail(email string) bool {
	// Basic email validation
	if !strings.Contains(email, "@") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	if len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	if !strings.Contains(parts[1], ".") {
		return false
	}
	return true
}

// GetParamID performs extracted parameter from url and return it as int
func GetParamID(r *http.Request, param string) (int, error) {
	idStr := chi.URLParam(r, param)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
