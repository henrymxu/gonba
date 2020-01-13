package gonba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseAddress = "https://stats.nba.com/stats/"
const baseAddressV2 = "http://data.nba.com/data/5s/json/cms/noseason/"
const baseAddressV3 = "http://data.nba.net/prod/"

// A Client is required for api calls
type Client struct {
	baseURL string
	httpClient *http.Client
}

func NewClient() *Client {
	customHttp := &http.Client{
		Timeout:   time.Second * 10,
		Transport: buildCustomTransport(),
	}
	return &Client{
		baseURL:    baseAddress,
		httpClient: customHttp,
	}
}

func (c *Client) makeRequest(endpoint string, params map[string]string, schema interface{}) int {
	body, status := c.makeRequestWithoutJson(endpoint, params)
	json.Unmarshal(body, schema)
	return status
}

func (c *Client) makeRequestWithoutJson(endpoint string, params map[string]string) ([]byte, int) {
	requestUrl := c.baseURL+endpoint
	if val, ok := params["version"]; ok {
		if val == "2" {
			requestUrl = baseAddressV2 + endpoint
		} else if val == "3" {
			requestUrl = baseAddressV3 + endpoint
		}
		delete(params, "version")
	}
	request, _ := http.NewRequest("GET", requestUrl, nil)
	request.Header.Set("Content-Type", "application/json")
	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	//fmt.Println(request.URL)
	response, err := c.httpClient.Do(request)
	if err != nil {
		fmt.Printf("Error when making request %v\n", err.Error())
		return nil, 404
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, response.StatusCode
}

func buildCustomTransport() *http.Transport {
	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, success := defaultRoundTripper.(*http.Transport)
	if !success {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100
	return &defaultTransport
}

