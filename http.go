package jsmops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/circleyu/go-jsmops/querybuilder"
)

const httpClientTimeout = time.Second * 10

func (c *APIClient) postFile(path string, fileName string, file io.Reader, out interface{}, expectedStatus ...int) error {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}
	// debug
	if c.logger != nil {
		c.logger.Debug(fileName)
	}

	// 实例化multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建multipart 文件字段
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return err
	}
	// 写入文件数据到multipart，和读取本地文件方法的唯一区别
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", c.cloudID, path), body)

	req.SetBasicAuth(c.userName, c.apiToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-ExperimentalApi", "opt-in")
	req.Header.Add("X-Atlassian-Token", "no-check")
	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logReq(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logRes(res)
	}

	isNotError := false

	for _, status := range expectedStatus {
		isNotError = res.StatusCode == status
		if isNotError {
			break
		}
	}

	if !isNotError {
		body, err := io.ReadAll(res.Body)
		apiError := ""
		if err == nil {
			var jsonBuffer bytes.Buffer
			err := json.Indent(&jsonBuffer, body, "", "\t")
			if err == nil {
				apiError = jsonBuffer.String()
			}
		}
		if res.StatusCode == http.StatusBadRequest && c.logger != nil {
			if c.logger != nil {
				c.logger.Error(apiError)
			}
		}
		return APIError{
			fmt.Errorf("received status code %d (%d expected)", res.StatusCode, expectedStatus),
			apiError,
		}
	}

	if out != nil {
		err = json.NewDecoder(res.Body).Decode(out)
		return err
	}

	return nil
}

func (c *APIClient) postJSON(path string, requestBody []byte, out interface{}, expectedStatus ...int) error {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logger.Debug(string(requestBody))
	}

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", c.cloudID, path), bytes.NewReader(requestBody))

	req.SetBasicAuth(c.userName, c.apiToken)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-ExperimentalApi", "opt-in")

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logReq(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logRes(res)
	}

	isNotError := false

	for _, status := range expectedStatus {
		isNotError = res.StatusCode == status
		if isNotError {
			break
		}
	}

	if !isNotError {
		body, err := io.ReadAll(res.Body)
		apiError := ""
		if err == nil {
			var jsonBuffer bytes.Buffer
			err := json.Indent(&jsonBuffer, body, "", "\t")
			if err == nil {
				apiError = jsonBuffer.String()
			}
		}
		if res.StatusCode == http.StatusBadRequest && c.logger != nil {
			if c.logger != nil {
				c.logger.Error(apiError)
			}
		}
		return APIError{
			fmt.Errorf("received status code %d (%d expected)", res.StatusCode, expectedStatus),
			apiError,
		}
	}

	if out != nil {
		err = json.NewDecoder(res.Body).Decode(out)
		return err
	}

	return nil
}

func (c *APIClient) put(path string, requestBody []byte, out interface{}, expectedStatus ...int) error {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logger.Debug(string(requestBody))
	}
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", c.cloudID, path), bytes.NewReader(requestBody))

	req.SetBasicAuth(c.userName, c.apiToken)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-ExperimentalApi", "opt-in")

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logReq(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logRes(res)
	}

	isNotError := false

	for _, status := range expectedStatus {
		isNotError = res.StatusCode == status
		if isNotError {
			break
		}
	}

	if !isNotError {
		body, err := io.ReadAll(res.Body)
		var apiError string
		if err == nil {
			var jsonBuffer bytes.Buffer
			err := json.Indent(&jsonBuffer, body, "", "\t")
			if err == nil {
				apiError = jsonBuffer.String()
			}
		}
		if res.StatusCode == http.StatusBadRequest && c.logger != nil {
			if c.logger != nil {
				c.logger.Error(apiError)
			}
		}
		return APIError{
			fmt.Errorf("received status code %d (%d expected)", res.StatusCode, expectedStatus),
			apiError,
		}
	}

	err = json.NewDecoder(res.Body).Decode(out)

	return err
}

func (c *APIClient) get(path string, out interface{}, query *querybuilder.Query) (http.Header, error) {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}
	var req *http.Request
	var err error

	if query != nil {
		req, err = http.NewRequest("GET", fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s?%s", c.cloudID, path, query.URLSafe()), nil)
	} else {
		req, err = http.NewRequest("GET", fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", c.cloudID, path), nil)
	}

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.userName, c.apiToken)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-ExperimentalApi", "opt-in")

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logReq(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logRes(res)
	}
	json.NewDecoder(res.Body).Decode(out)

	return res.Header, err
}

func (c *APIClient) getFile(path string) (*bytes.Reader, error) {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}
	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.userName, c.apiToken)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-ExperimentalApi", "opt-in")

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logReq(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logRes(res)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(data)

	return r, nil
}

func (c *APIClient) delete(path string, requestBody []byte, expectedStatus ...int) error {
	httpClient := &http.Client{
		Timeout: httpClientTimeout,
	}

	var body io.Reader
	if requestBody != nil {
		body = bytes.NewReader(requestBody)
	} else {
		body = nil
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("https://api.atlassian.com/jsm/ops/api/%s/%s", c.cloudID, path), body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.userName, c.apiToken)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("X-ExperimentalApi", "opt-in")

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logReq(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// debug
	if c.logger != nil && c.logLevel >= LogDebug {
		c.logRes(res)
	}

	isNotError := false

	for _, status := range expectedStatus {
		isNotError = res.StatusCode == status
		if isNotError {
			break
		}
	}

	if !isNotError {
		body, err := io.ReadAll(res.Body)
		var apiError string
		if err == nil {
			var jsonBuffer bytes.Buffer
			err := json.Indent(&jsonBuffer, body, "", "\t")
			if err == nil {
				apiError = jsonBuffer.String()
			}
		}
		if res.StatusCode == http.StatusBadRequest && c.logger != nil {
			if c.logger != nil {
				c.logger.Error(apiError)
			}
		}
		return APIError{
			fmt.Errorf("received status code %d (%d expected)", res.StatusCode, expectedStatus),
			apiError,
		}
	}

	return nil
}
