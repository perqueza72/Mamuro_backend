package zinc_handler

import (
	. "constants_project"
	"io"
	"log"
	models "models_zinc"
	"net/http"
)

type InternalQuery struct {
	Term string `json:"term"`
}

type RecordQuery struct {
	SortFields []string      `json:"sort_fields"`
	From       int           `json:"from"`
	MaxResults int           `json:"max_results"`
	SearchType string        `json:"search_type"`
	Query      InternalQuery `json:"query"`
}

func searchRecords(query *models.IRequestData) ([]byte, error) {
	req, err := http.NewRequest("POST", ZINC_HOST+"/api/"+ZINC_EMAIL_INDEX+"/_search", *query)
	if err != nil {
		log.Default().Panic(err)
		return nil, err
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Default().Panic(err)
		return nil, err
	}

	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Default().Panic(err)
		return nil, err
	}

	return body, nil
}

// Return first records sorted by Sort Fields. Index starts to count in 0.
func SearchAllRecordsBy(recordQueryRequest *models.RecordQueryRequest) ([]byte, error) {

	recordQuery := SearchRecordsStandardStructure(recordQueryRequest)
	data := models.Model2IRequestData(recordQuery)

	return searchRecords(data)
}

// Return first records sorted by Sort fields with some petition text. Index start to count in 0.
func SearchLikeRecordsBy(recordQueryRequest *models.RecordQueryRequest) ([]byte, error) {

	recordQuery := SearchRecordsStandardStructure(recordQueryRequest)
	recordQuery.Query.Term = recordQueryRequest.SpecificText
	recordQuery.SearchType = "match"

	data := models.Model2IRequestData(recordQuery)
	return searchRecords(data)
}

// Get a RecordQueryRequest and return a standard RecordQuery object intance.
func SearchRecordsStandardStructure(recordQueryRequest *models.RecordQueryRequest) *RecordQuery {
	index_page := recordQueryRequest.IndexPage
	if index_page < 0 {
		index_page = 0
	}
	max_results := 10
	start_point := index_page * max_results

	return &RecordQuery{
		MaxResults: max_results,
		From:       start_point,
		SearchType: "alldocuments",
		SortFields: recordQueryRequest.SortFields,
	}
}
