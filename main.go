package jsmops

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// APIClient ...
type APIClient struct {
	cloudID  string
	userName string
	apiToken string
	logLevel LogLevel
	logger   *logrus.Logger
	Alert    AlertsManager
}

type LogLevel int

const (
	LogError LogLevel = iota
	LogInfo
	LogDebug
)

// ClientOptions ...
type ClientOptions struct {
	Level  LogLevel
	Logger *logrus.Logger
}

// EmptyOptions ...
func EmptyOptions() *ClientOptions {
	return nil
}

func NewOptions() *ClientOptions {
	return &ClientOptions{
		Level: LogError,
	}
}

// Init initializes the package
func Init(cloudID, apiToken, userName string, options *ClientOptions) *APIClient {
	client := &APIClient{
		cloudID:  cloudID,
		userName: userName,
		apiToken: apiToken,
	}
	if options != nil {
		client.logger = options.Logger
		client.logLevel = options.Level
	}
	if client.logger != nil {
		client.logger.Infof("JSM ops Client initializing..., authorization = %s", userName)
	}

	client.Alert = newAlertsManager(client)

	return client
}

// func (client *APIClient) logErr(err error) {
// 	if err != nil && client.logger != nil {
// 		client.logger.Error(err.Error())
// 	}
// }

func (client *APIClient) logReq(req *http.Request) {
	if client.logger != nil {
		client.logger.Debug("Headers")
		for k, v := range req.Header {
			client.logger.Debugf("%s: %s\n", k, v)
		}
		client.logger.Debugf("URL: %s", req.URL)
		if req.Body != nil {
			body, _ := io.ReadAll(req.Body)
			var jsonBuffer bytes.Buffer
			json.Indent(&jsonBuffer, body, "", "\t")
			client.logger.Debug(jsonBuffer.String())
		}
	}
}

func (client *APIClient) logRes(res *http.Response) {
	if client.logger != nil {
		client.logger.Debugf("Status: %d", res.StatusCode)
		if res.StatusCode != 200 {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				client.logger.Debug(err.Error())
			}
			var jsonBuffer bytes.Buffer
			json.Indent(&jsonBuffer, body, "", "\t")
			client.logger.Debug(jsonBuffer.String())
		}
	}
}

// BackupJSON ...
func (client *APIClient) BackupJSON(fileName string, data interface{}) error {
	backupJSON, _ := json.Marshal(data)
	return os.WriteFile(fileName, backupJSON, 0644)
}
