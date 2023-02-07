package router_functions

import (
	"encoding/json"
	models "models_zinc"
	"net/http"
	"zinc_handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func handlerFuncSearchRecords(w http.ResponseWriter, r *http.Request) {

	recordQuery := models.RecordQueryRequest{}
	if err := json.NewDecoder(r.Body).Decode(&recordQuery); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	records, err := zinc_handler.SearchRecords(&recordQuery)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	var result map[string]interface{}
	json.Unmarshal(records, &result)
	response_err := result["error"]
	if response_err != nil {
		json.NewEncoder(w).Encode(response_err)
		return
	}
	hitss := result["hits"].(map[string]interface{})
	hits := hitss["hits"].([]interface{})

	var response []interface{}
	for _, hit := range hits {
		response = append(response, hit.(map[string]interface{})["_source"])
	}

	json.NewEncoder(w).Encode(response)
}

func StartServer(port string) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Post("/search_records", handlerFuncSearchRecords)
	http.ListenAndServe(port, r)
}
