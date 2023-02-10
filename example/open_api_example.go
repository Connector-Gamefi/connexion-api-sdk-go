package example

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Connector-Gamefi/connexion-api-sdk-go/connexionapisdk"
)

const OPEN_API_DOMAIN string = ""
const HEADER_API_KEY string = "X-CONNEXION-APIKEY"
const HEADER_SIGNATURE string = "X-CONNEXION-SIGNATURE"

func RequestGetOpenApi() ([]byte, error) {
	client := &http.Client{}
	//api url
	url := OPEN_API_DOMAIN + "/global/topList/roleLevel"
	//request parameters
	params := map[string]interface{}{
		"timestamp": "1675998834",
		"topNum":    "5",
	}
	//sercet
	sercet := "a05315753c2842598ee5daca4f7ef399"
	signature := connexionapisdk.Sign(params, sercet)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, fmt.Sprintf("%v", v))
	}

	//set header api key & signature
	req.Header.Set(HEADER_API_KEY, "fa61655a1aca4804b5e2c3c7a10c6257")
	req.Header.Set(HEADER_SIGNATURE, signature)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	readBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return readBytes, nil
}