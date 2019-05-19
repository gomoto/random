package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/gomoto/random/numbers"
)

// Print random number between specified min and max
func main() {
	toRangeMin := flag.Int64("min", 0, "Minimum integer (default 0)")
	toRangeMax := flag.Int64("max", 0, "Maximum integer (default 0)")
	flag.Parse()

	// Make HTTP request
	url := "https://beacon.nist.gov/beacon/2.0/chain/last/pulse/last"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/json")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error 1")
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error 2")
		fmt.Println(err.Error())
		return
	}
	var responseBody map[string]interface{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error 3")
		fmt.Println(err.Error())
		return
	}
	pulse := responseBody["pulse"].(map[string]interface{})
	// hexValue is a 512-bit hexadecimal string (128 hex characters)
	hexValue := pulse["outputValue"].(string)
	hexMin := "0"
	hexMax := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	intValue := new(big.Int)
	intMin := new(big.Int)
	intMax := new(big.Int)
	intValue.SetString(hexValue, 16)
	intMin.SetString(hexMin, 16)
	intMax.SetString(hexMax, 16)
	fromRange := numbers.IntRange{Min: intMin, Max: intMax}
	toRange := numbers.IntRange{Min: big.NewInt(*toRangeMin), Max: big.NewInt(*toRangeMax)}
	outputValue, err := numbers.MapInt(*intValue, fromRange, toRange)
	if err != nil {
		fmt.Println("Error 4")
		fmt.Println(err.Error())
	}
	fmt.Println(outputValue)
}
