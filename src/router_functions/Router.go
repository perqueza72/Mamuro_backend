package router_functions

import (
	constants "constants_project"
	"encoding/json"
	. "mapping_zinc"
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	records, err := zinc_handler.SearchAllRecordsBy(&recordQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	response, err := GetZincRecords(&records)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handlerFuncSearchLikeRecords(w http.ResponseWriter, r *http.Request) {

	recordQuery := models.RecordQueryRequest{}
	if err := json.NewDecoder(r.Body).Decode(&recordQuery); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	records, err := zinc_handler.SearchLikeRecordsBy(&recordQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	response, err := GetZincRecords(&records)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func StartServer(port string) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{constants.VUE_URL},
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
	r.Post("/find_records", handlerFuncSearchLikeRecords)
	http.ListenAndServe(port, r)
}
