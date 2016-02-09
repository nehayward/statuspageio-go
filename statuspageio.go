package statuspageio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	BaseUrlV2 = "https://api.statuspage.io/v1/pages/"

	STATUS_OPERATIONAL                = 100
	STATUS_DEGRADED_PERFORMANCE       = 300
	STATUS_PARTIAL_SERVICE_DISRUPTION = 400
	STATUS_SERVICE_DISRUPTION         = 500
	STATUS_SECURITY_EVENT             = 600
)

type StatusPageIOApi struct {
	ApiKey     string
	PageID     string
	HttpClient *http.Client
	baseUrl    string
}

func (a StatusPageIOApi) apiRequest(method string, resource string, data interface{}, result interface{}) error {
	var (
		req *http.Request
		err error
	)

	if data != nil {
		if str, ok := data.(string); ok {
			/* act on str */
			status := strings.NewReader("component%5Bstatus%5D=" + str)
			req, err = http.NewRequest(method, fmt.Sprintf("%s%s", a.baseUrl, resource), status)
		} else {
			/* not string */
			b, err := json.Marshal(data)
			if err != nil {
				return err
			}
			req, err = http.NewRequest(method, fmt.Sprintf("%s%s", a.baseUrl, resource), bytes.NewReader(b))
		}
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", a.baseUrl, resource), nil)
	}
	// url := "https://api.statuspage.io/v1/pages/17jhzflc8qjp/components/237p399shf5q.json"
	//
	// payload := strings.NewReader("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"component[status]\"\r\n\r\noperational\r\n-----011000010111000001101001--")
	//
	// req, _ = http.NewRequest("PATCH", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("authorization", "OAuth 00b870eb-95b9-4e39-96db-5dc60bac1462")
	// req.Header.Add("cache-control", "no-cache")
	// req.Header.Add("postman-token", "4ef51566-2b8d-4a79-21f9-c37baa320343")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
	// return err
	//
	if err != nil {

		return err
	}
	//
	// resp, err := a.HttpClient.Do(req)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()
	//
	// if resp.StatusCode < 200 || resp.StatusCode >= 300 {
	// 	fmt.Println(resp.Status)
	// 	return errors.New(fmt.Sprintf("Response status code is %d", resp.StatusCode))
	// }
	//
	return json.NewDecoder(res.Body).Decode(result)
}

func NewStatusioPageApi(apiKey string, pageID string) *StatusPageIOApi {
	c := &StatusPageIOApi{
		ApiKey:     apiKey,
		PageID:     pageID,
		HttpClient: http.DefaultClient,
		baseUrl:    BaseUrlV2 + pageID + "/",
	}
	return c
}
