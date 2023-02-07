package test

import (
	models "models_zinc"
	"testing"
	"zinc_handler"
)

func TestSearchRecords(t *testing.T) {
	t.Setenv("ZINC_ADMIN_USER", "admin")
	t.Setenv("ZINC_ADMIN_PASSWORD", "Complexpass#123")

	recordQuery := models.RecordQueryRequest{
		SortFields: []string{"-Date"},
		IndexPage:  0,
	}
	body, err := zinc_handler.SearchRecords(&recordQuery)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(string(body))
}
