package models

import (
	"bytes"
	"encoding/json"
)

type IRequestData interface {
	Read([]byte) (int, error)
}

type RecordQueryRequest struct {
	SortFields []string
	IndexPage  int
}

func Model2IRequestData(model interface{}) *IRequestData {
	var buf bytes.Buffer

	json.NewEncoder(&buf).Encode(model)
	request := IRequestData(&buf)
	return &request
}
