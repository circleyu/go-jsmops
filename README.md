# go-jsmops

一個簡單易用的 Go 語言庫，用於調用 Jira Service Management Operations REST API。

[中文](README.md) | [English](README_EN.md)

## 功能特性

- ✅ 完整的 API 覆蓋：支援所有 Jira Service Management Operations API 端點
- ✅ 類型安全：使用強類型的 Go 結構體定義請求和響應
- ✅ 易於使用：清晰的接口設計，符合 Go 語言慣例
- ✅ 日誌支援：可選的日誌記錄功能，方便調試
- ✅ 錯誤處理：統一的錯誤處理機制

## 支援的資源類別

- **Alerts** - 警報管理
- **Audit Logs** - 審計日誌
- **Contacts** - 聯絡人管理
- **Teams** - 團隊管理
- **Roles** - 自定義用戶角色
- **Escalations** - 升級規則
- **Forwarding Rules** - 轉發規則
- **Heartbeats** - 心跳監控
- **Integrations** - 整合管理
- **Integration Actions** - 整合操作
- **Integration Filters** - 整合過濾器
- **Maintenances** - 維護計劃（全局和團隊）
- **Notification Rules** - 通知規則
- **Notification Rule Steps** - 通知規則步驟
- **Policies** - 警報策略（全局和團隊）
- **Team Roles** - 團隊角色
- **Routing Rules** - 路由規則
- **Schedules** - 排程管理
- **Schedule On-calls** - 排程待命
- **Schedule Overrides** - 排程覆蓋
- **Schedule Rotations** - 排程輪換
- **Schedule Timelines** - 排程時間線
- **Syncs** - 同步管理
- **Sync Actions** - 同步操作
- **Sync Action Groups** - 同步操作組
- **JEC** - JEC 頻道管理

## 安裝

```bash
go get github.com/circleyu/go-jsmops
```

## 快速開始

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/circleyu/go-jsmops"
    "github.com/circleyu/go-jsmops/alert"
    "github.com/sirupsen/logrus"
)

func main() {
    // 初始化客戶端
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)
    
    options := &jsmops.ClientOptions{
        Level:  jsmops.LogDebug,
        Logger: logger,
    }
    
    client := jsmops.Init(
        "your-cloud-id",
        "your-api-token",
        "your-username",
        options,
    )
    
    // 使用 API
    // 範例：列出警報
    listReq := &alert.ListAlertsRequest{
        Limit: 10,
    }
    
    result, err := client.Alert.ListAlerts(listReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("找到 %d 個警報\n", len(result.Alerts))
}
```

### 不使用日誌

```go
client := jsmops.Init(
    "your-cloud-id",
    "your-api-token",
    "your-username",
    jsmops.EmptyOptions(),
)
```

## 使用範例

詳細的使用範例請參考 [examples](./examples) 目錄：

- [基本操作範例](./examples/basic/main.go) - 基本的 CRUD 操作
- [警報管理範例](./examples/alerts/main.go) - 警報的創建、查詢和更新
- [排程管理範例](./examples/schedules/main.go) - 排程和待命管理
- [整合管理範例](./examples/integrations/main.go) - 整合和操作管理

## API 文檔

### 認證

所有 API 請求使用基本認證（Basic Authentication）。您需要提供：
- `cloudID`: Jira Cloud ID
- `apiToken`: Jira API Token
- `userName`: Jira 用戶名或郵箱

### 錯誤處理

所有 API 方法都會返回錯誤。錯誤類型為 `APIError`，包含 HTTP 狀態碼和錯誤詳情。

```go
result, err := client.Alert.GetAlert(&alert.GetAlertRequest{ID: "alert-id"})
if err != nil {
    if apiErr, ok := err.(jsmops.APIError); ok {
        fmt.Printf("API 錯誤: %s (狀態碼: %d)\n", apiErr.Error(), apiErr.StatusCode)
    } else {
        fmt.Printf("其他錯誤: %v\n", err)
    }
}
```

## 開發

### 專案結構

```
go-jsmops/
├── main.go              # 主客戶端和初始化
├── http.go              # HTTP 請求處理
├── endpoints.go         # API 端點定義
├── alerts.go            # 警報管理
├── contacts.go          # 聯絡人管理
├── ...                  # 其他資源管理
├── alert/               # 警報相關結構
├── contacts/            # 聯絡人相關結構
└── examples/            # 使用範例
```

### 添加新功能

1. 在對應的資源目錄下創建 `request.go` 和 `result.go`
2. 創建或更新對應的 manager 文件（如 `alerts.go`）
3. 在 `endpoints.go` 中添加端點定義
4. 在 `main.go` 中註冊新的 manager

## 授權

本專案採用 MIT 授權。

## 貢獻

歡迎提交 Issue 和 Pull Request！

## 相關連結

- [Jira Service Management Operations API 文檔](https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-operations/)
