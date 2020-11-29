package models

import (
	"encoding/json"
	"fmt"
	"github.com/OlegVanyaGreatBand/architecture-lab-3/server/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TelemetryHttpHandler http.HandlerFunc

func HttpHandler(store *Store) TelemetryHttpHandler {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handlerGetTelemetry(r, rw, store)
		} else if r.Method == "POST" {
			handleAddTelemetry(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handlerGetTelemetry(r *http.Request, rw http.ResponseWriter, store *Store) {
	path := strings.Split(r.URL.Path, "/")
	tabletId, err := strconv.Atoi(path[len(path) - 1])
	if err != nil {
		log.Printf("Error parsing tablet id: %s", path)
		utils.WriteJsonBadRequest(rw, "Invalid tablet id")
		return
	}

	telemetry, err := store.GetTelemetry(tabletId)
	if err != nil {
		message := fmt.Sprintf("Error making query to the db: %s", err)
		log.Printf(message)
		utils.WriteJsonInternalError(rw, message)
	}

	utils.WriteJsonResult(rw, telemetry)
}

func handleAddTelemetry(r *http.Request, rw http.ResponseWriter, store *Store) {
	var telemetry TelemetryData
	if err := json.NewDecoder(r.Body).Decode(telemetry); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		utils.WriteJsonBadRequest(rw, "Invalid telemetry data")
		return
	}

	previous, err := store.GetTelemetry(telemetry.TabletId)
	if err != nil {
		log.Printf("Error retrieving previous telemetry: %s", err)
		utils.WriteJsonBadRequest(rw, "Invalid tablet id")
		return
	}

	lastTime := time.Unix(0, 0)
	if len(previous.Telemetry) > 0 {
		if parsed, err := time.Parse(time.RFC3339, previous.Telemetry[0].Timestamp); err != nil {
			log.Printf("Error parsing time: %s", err)
			utils.WriteJsonInternalError(rw, "Error parsing time")
			return
		} else {
			lastTime = parsed
		}
	}

	currentTime := time.Now()
	diff := currentTime.Unix() - lastTime.Unix()
	// ignore if 10 seconds haven't passed
	if diff < 10 {
		log.Printf("Ignoring request: %d/10 seconds passed", diff)
		// still return 200 ok - user don't know that we've ignored him
		utils.WriteJsonResult(rw, struct {
			Message string `json:"message"`
		}{
			Message: "Йой, най буде",
		})
		return
	}

	if err := store.AddTelemetry(&telemetry); err != nil {
		log.Printf("Inserting error: %s", err)
		utils.WriteJsonInternalError(rw, "Так не буде")
		return
	}

	log.Printf("Inserted %d telemetry records", len(telemetry.Telemetry))
	utils.WriteJsonResult(rw, struct {
		Message string `json:"message"`
	}{
		Message: "Йой, най буде",
	})
}
