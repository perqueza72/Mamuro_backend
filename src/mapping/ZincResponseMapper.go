package mapping

import (
	"encoding/json"
	"fmt"
)

func GetZincRecords(zinc_response *[]byte) ([]interface{}, error) {
	var result map[string]interface{}

	if string(*zinc_response) == "" {
		return nil, fmt.Errorf("error trying to get data from zinc")
	}

	json.Unmarshal(*zinc_response, &result)
	err := result["auth"]
	if err != nil {
		return nil, fmt.Errorf("error in authentication process, %v", err)
	}

	response_err := result["error"]
	if response_err != nil {
		return nil, fmt.Errorf("error trying to get data from zinc. %v", response_err.(string))
	}

	if result["hits"] == nil {
		return nil, nil
	}

	hitss := result["hits"].(map[string]interface{})
	if result["hits"] == nil {
		return nil, nil
	}

	hits := hitss["hits"].([]interface{})
	var response []interface{}
	for _, hit := range hits {
		response = append(response, hit.(map[string]interface{})["_source"])
	}

	return response, nil
}
