package zinc_handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	models "models_zinc"
	"net/http"
)

type RecordQuery struct {
	SortFields []string `json:"sort_fields"`
	From       int      `json:"from"`
	MaxResults int      `json:"max_results"`
	SearchType string   `json:"search_type"`
}

func searchRecords(query *models.IRequestData) ([]byte, error) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/email/_search", *query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return body, nil
}

// Index_Page starts in zero.
func SearchRecords(recordQueryRequest *models.RecordQueryRequest) ([]byte, error) {

	max_results := 10
	start_point := recordQueryRequest.IndexPage * max_results

	recordQuery := RecordQuery{
		SortFields: recordQueryRequest.SortFields,
		From:       start_point,
		MaxResults: max_results,
		SearchType: "alldocuments",
	}

	s, _ := json.Marshal(&recordQuery)
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(s)

	data := models.Model2IRequestData(&recordQuery)
	return searchRecords(data)

}
