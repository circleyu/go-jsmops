package jsmops

import (
	"github.com/circleyu/go-jsmops/auditlogs"
)

type AuditLogsManager interface {
	GetAuditLogs(*auditlogs.GetAuditLogsRequest) (*auditlogs.ListAuditLogsResult, error)
}

type auditLogsManager struct {
	*APIClient
}

func newAuditLogsManager(client *APIClient) *auditLogsManager {
	return &auditLogsManager{
		client,
	}
}

func (manager *auditLogsManager) GetAuditLogs(data *auditlogs.GetAuditLogsRequest) (*auditlogs.ListAuditLogsResult, error) {
	output := &auditlogs.ListAuditLogsResult{}
	_, err := manager.get(endpoints.auditLogs.GetAuditLogs, output, data.RequestParams())
	if err != nil {
		return nil, err
	}
	return output, nil
}
