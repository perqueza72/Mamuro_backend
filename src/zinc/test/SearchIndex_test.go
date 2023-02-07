package test

import (
	mapping "mapping_zinc"
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
	body, err := zinc_handler.SearchAllRecordsBy(&recordQuery)

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	_, err = mapping.GetZincRecords(&body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Logf("Search all records works.")
}

func TestSearchRecordsBadAuth(t *testing.T) {
	t.Setenv("ZINC_ADMIN_USER", "unexisting_admin")
	t.Setenv("ZINC_ADMIN_PASSWORD", "unexisting_password")

	recordQuery := models.RecordQueryRequest{
		SortFields: []string{"-Date"},
		IndexPage:  0,
	}
	body, err := zinc_handler.SearchAllRecordsBy(&recordQuery)

	if err != nil {
		t.Log(err)
		t.Fail()
	}
	_, err = mapping.GetZincRecords(&body)
	if err != nil {
		t.Logf("Authentication process works!!\n%v", err)
		return
	}

	t.Fail()
}
