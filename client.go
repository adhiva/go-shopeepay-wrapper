package shopeepay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	APIEnvType    Env
	APIEnvRegion  EnvRegion
	ClientID      string
	ClientSecret  string
	MerchantExtID string
	StoreExtID    string
	LogLevel      int
	Logger        *log.Logger
}

func NewClient() Client {
	return Client{
		APIEnvType:   Sandbox,
		APIEnvRegion: ID,
		LogLevel:     2,
		Logger:       log.New(os.Stderr, "", log.LstdFlags),
	}
}

var defHTTPTimeout = 80 * time.Second
var httpClient = &http.Client{Timeout: defHTTPTimeout}

func (c *Client) DoRequest(method, path string, body interface{}, v interface{}) error {
	byteJson, _ := json.Marshal(body)
	logLevel := c.LogLevel
	logger := c.Logger

	req, err := http.NewRequest(method, path, bytes.NewBuffer(byteJson))
	if err != nil {
		if logLevel > 0 {
			logger.Println("Request creation failed: ", err)
		}
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Airpay-ClientId", c.ClientID)
	req.Header.Add("X-Airpay-Req-H", CreateSignatureHeaders(byteJson, c.ClientSecret))
	start := time.Now()

	if logLevel > 1 {
		logger.Println("Request ", req.Method, ": ", req.URL.Host, req.URL.Path)
	}

	// Set Log Request
	RequestLog := RequestLog{
		Method:      req.Method,
		Host:        req.URL.Host,
		Path:        req.URL.Path,
		RequestBody: Stringify(body),
	}
	logger.Println(Stringify(RequestLog))
	res, err := httpClient.Do(req)
	if err != nil {
		logger.Println("Cannot send request : ", err)
		return err
	}

	// Println stdout for checking request body and response
	logger.Println("Completed in ", time.Since(start))
	defer res.Body.Close()

	if res.StatusCode != 200 {
		errMessage := fmt.Sprintf("Cannot send request and get status code %d from shopeepay", res.StatusCode)
		return errors.New(errMessage)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Println("Cannot read response body: ", err)
		return err
	}

	if v != nil {
		response := json.Unmarshal(resBody, v)
		RequestLog.ResponseBody = Stringify(v)
		RequestLog.TimeLength = int(time.Since(start))
		logger.Println(Stringify(RequestLog))

		return response
	}

	return nil
}
