package nist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Get latest random value from beacon.nist.gov
func GetLatest() (string, error) {
	url := "https://beacon.nist.gov/beacon/2.0/chain/last/pulse/last"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return "", err
	}
	pulse := responseBody["pulse"].(map[string]interface{})
	// hexValue is a 512-bit hexadecimal string (128 hex characters)
	outputValue := pulse["outputValue"].(string)
	return outputValue, nil
}
