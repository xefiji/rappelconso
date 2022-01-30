package datagouv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	dataset = "rappelconso0"
	url     = "https://data.economie.gouv.fr/api/v2/catalog/datasets/%s/records?timezone=UTC&limit=%d&offset=%d"
)

type errorResponse struct {
	Code      int
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func (e errorResponse) String() string {
	return fmt.Sprintf("%d: %s: %s", e.Code, e.ErrorCode, e.Message)
}

// GetRecords makes an http request to the api endpoint and returns records set.
func GetRecords(limit, offset int) (*Records, error) {
	resp, err := http.Get(fmt.Sprintf(url, dataset, limit, offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		errorResponse := &errorResponse{
			Code: resp.StatusCode,
		}
		err = json.Unmarshal(body, errorResponse)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.String())
	}

	records := &Records{}
	err = json.Unmarshal(body, records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// CountRecords makes a single hit on one record to get the total records count.
func CountRecords() (int, error) {
	records, err := GetRecords(1, 0)
	if err != nil {
		return 0, err
	}
	return int(records.Total), nil
}
