package jsmops

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/circleyu/go-jsmops/v2/jec"
)

type JECManager interface {
	ListJECChannels(*jec.ListJECChannelsRequest) (*jec.ListJECChannelsResult, error)
	CreateJECChannel(*jec.CreateJECChannelRequest) (*jec.JECChannel, error)
	GetJECChannel(*jec.GetJECChannelRequest) (*jec.JECChannel, error)
	DeleteJECChannel(*jec.DeleteJECChannelRequest) error
	SendJECAction(*jec.SendJECActionRequest) (*jec.SendJECActionResponse, error)
}

type jecManager struct {
	*APIClient
}

func newJECManager(client *APIClient) *jecManager {
	return &jecManager{
		client,
	}
}

func (manager *jecManager) ListJECChannels(data *jec.ListJECChannelsRequest) (*jec.ListJECChannelsResult, error) {
	output := &jec.ListJECChannelsResult{}
	_, err := manager.get(endpoints.jec.ListJECChannels, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *jecManager) CreateJECChannel(data *jec.CreateJECChannelRequest) (*jec.JECChannel, error) {
	output := &jec.JECChannel{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.jec.CreateJECChannel, jsonb, output, nil, http.StatusOK, http.StatusCreated)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *jecManager) GetJECChannel(data *jec.GetJECChannelRequest) (*jec.JECChannel, error) {
	output := &jec.JECChannel{}
	_, err := manager.get(endpoints.jec.GetJECChannel(data.ID), output, nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (manager *jecManager) DeleteJECChannel(data *jec.DeleteJECChannelRequest) error {
	return manager.delete(endpoints.jec.DeleteJECChannel(data.ID), nil, http.StatusNoContent, http.StatusOK)
}

func (manager *jecManager) SendJECAction(data *jec.SendJECActionRequest) (*jec.SendJECActionResponse, error) {
	output := &jec.SendJECActionResponse{}
	jsonb, err := sonic.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = manager.postJSON(endpoints.jec.SendJECAction, jsonb, output, nil, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return output, nil
}

