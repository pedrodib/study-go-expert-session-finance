package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody Structure
type HealthBody struct {
	Status string `json:"status"`
}

// Health is used to check if the application is alive
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	status := HealthBody{Status: "alive"}

	_ = json.NewEncoder(w).Encode(status)
}
