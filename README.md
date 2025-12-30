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
- **Integration Events** - Integration Events API（創建、確認、關閉警報、添加備註）
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
go get github.com/circleyu/go-jsmops/v2@v2.1.0
```

## 認證方式

本庫支援兩種認證方式：

1. **Basic Authentication**：用於普通 API（`/jsm/ops/api/{cloudId}/v1/...`）
   - 使用 `userName`（電子郵件）和 `apiToken`（API Token）
   - 適用於所有標準的 JSM Operations API

2. **API Integration (GenieKey)**：用於 Integration Events API（`/jsm/ops/integration/v2/...`）
   - 使用 `apiKey`（GenieKey）
   - 僅用於 Integration Events API（創建、確認、關閉警報、添加備註）
   - 可選：如果不需要使用 Integration Events API，可以傳入空字串

## 快速開始

### 基本使用（僅 Basic Auth）

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/circleyu/go-jsmops/v2"
    "github.com/circleyu/go-jsmops/v2/alert"
    "github.com/sirupsen/logrus"
)

func main() {
    // 初始化客戶端（僅使用 Basic Auth）
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)
    
    options := &jsmops.ClientOptions{
        Level:  jsmops.LogDebug,
        Logger: logger,
    }
    
    client := jsmops.Init(
        "your-cloud-id",
        "your-api-token",  // API Token for Basic Auth
        "your-email@example.com",  // Username/Email for Basic Auth
        "",  // apiKey (空字串表示不使用 Integration Events API)
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

### 使用 Integration Events API

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/circleyu/go-jsmops/v2"
    "github.com/circleyu/go-jsmops/v2/alert"
    "github.com/sirupsen/logrus"
)

func main() {
    // 初始化客戶端（同時支援 Basic Auth 和 Integration Events API）
    logger := logrus.New()
    logger.SetLevel(logrus.DebugLevel)
    
    options := &jsmops.ClientOptions{
        Level:  jsmops.LogDebug,
        Logger: logger,
    }
    
    client := jsmops.Init(
        "your-cloud-id",
        "your-api-token",  // API Token for Basic Auth
        "your-email@example.com",  // Username/Email for Basic Auth
        "your-genie-key",  // GenieKey for Integration Events API
        options,
    )
    
    // 使用普通 API（使用 Basic Auth）
    listReq := &alert.ListAlertsRequest{
        Limit: 10,
    }
    
    result, err := client.Alert.ListAlerts(listReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("找到 %d 個警報\n", len(result.Alerts))
    
    // 使用 Integration Events API（使用 GenieKey）
    createReq := &alert.IntegrationCreateAlertRequest{
        Message: "CPU usage exceeded 80%",
        Details: map[string]string{
            "server": "server-01",
            "cpu":    "85%",
        },
        Priority: alert.P1,
        Source:   "MonitoringTool",
    }
    
    createResult, err := client.IntegrationEvents.CreateAlert(createReq)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("警報創建請求已提交: %s\n", createResult.RequestID)
    
    // 範例：確認警報
    ackReq := &alert.IntegrationAcknowledgeAlertRequest{
        IdentifierType:  alert.ALERTID,
        IdentifierValue: "alert-id-here",
        User:            "John Smith",
        Source:          "MonitoringTool",
        Note:            "正在處理此警報",
    }
    
    ackResult, err := client.IntegrationEvents.AcknowledgeAlert(ackReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("警報確認請求已提交: %s\n", ackResult.RequestID)
    
    // 範例：關閉警報
    closeReq := &alert.IntegrationCloseAlertRequest{
        IdentifierType:  alert.ALERTID,
        IdentifierValue: "alert-id-here",
        User:            "John Smith",
        Source:          "MonitoringTool",
        Note:            "問題已解決",
    }
    
    closeResult, err := client.IntegrationEvents.CloseAlert(closeReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("警報關閉請求已提交: %s\n", closeResult.RequestID)
    
    // 範例：添加備註
    noteReq := &alert.IntegrationAddNoteRequest{
        IdentifierType:  alert.ALERTID,
        IdentifierValue: "alert-id-here",
        User:            "John Smith",
        Source:          "MonitoringTool",
        Note:            "這是通過 Integration Events API 添加的備註",
    }
    
    noteResult, err := client.IntegrationEvents.AddNote(noteReq)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("備註添加請求已提交: %s\n", noteResult.RequestID)
}
```

### 不使用日誌

```go
client := jsmops.Init(
    "your-cloud-id",
    "your-api-token",
    "your-email@example.com",
    "",  // apiKey (空字串表示不使用 Integration Events API)
    jsmops.EmptyOptions(),
)
```

## Integration Events API

Integration Events API 是一個特殊的 API 端點集合，用於從外部系統（如監控工具）創建和管理警報。這些 API 使用 GenieKey 認證，不需要 cloudID。

### 支援的操作

- **CreateAlert** - 創建新警報
- **AcknowledgeAlert** - 確認警報
- **CloseAlert** - 關閉警報
- **AddNote** - 添加備註到警報

### 重要差異

Integration Events API 與普通 Alerts API 的主要差異：

1. **認證方式**：使用 GenieKey 而不是 Basic Auth
2. **路徑**：`/jsm/ops/integration/v2/...` 而不是 `/jsm/ops/api/{cloudId}/v1/...`
3. **請求結構**：
   - `IntegrationCreateAlertRequest` 使用 `details`（`map[string]string`）而不是 `extraProperties`（`map[string]interface{}`）
   - 所有請求都包含 `user`, `source`, `note` 等欄位
4. **響應**：返回 202 Accepted，表示請求已提交處理（異步操作）

### 使用場景

Integration Events API 特別適合：
- 監控系統自動創建警報
- CI/CD 工具報告構建失敗
- 自動化腳本管理警報生命週期
- 第三方整合系統

## 使用範例

詳細的使用範例請參考 [examples](./examples) 目錄：

- [基本操作範例](./examples/basic/main.go) - 基本的 CRUD 操作
- [警報管理範例](./examples/alerts/main.go) - 警報的創建、查詢和更新
- [排程管理範例](./examples/schedules/main.go) - 排程和待命管理
- [整合管理範例](./examples/integrations/main.go) - 整合和操作管理

## API 文檔

### 認證

本庫支援兩種認證方式：

#### Basic Authentication（用於普通 API）

所有標準的 JSM Operations API 使用 Basic Authentication。您需要提供：
- `cloudID`: Jira Cloud ID
- `apiToken`: API Token（從 [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens) 獲取）
- `userName`: 您的電子郵件地址

要獲取 API Token：
1. 前往 [Atlassian Account Settings](https://id.atlassian.com/manage-profile/security/api-tokens)
2. 點擊 **Create API token**
3. 複製生成的 API Token

#### API Integration (GenieKey)（用於 Integration Events API）

Integration Events API 使用 GenieKey 認證。您需要提供：
- `apiKey`: API Integration 的 API Key（GenieKey）

要獲取 GenieKey，請在 Jira Service Management 中設置 API Integration：
1. 前往團隊的 Operations 頁面
2. 選擇 **Integrations** > **Add integration**
3. 搜索並選擇 "API"
4. 輸入整合名稱並完成設置
5. 複製生成的 API Key

API Key 格式為 UUID，例如：`g4ff854d-a14c-46a8-b8f0-0960774319dd`

**注意**：如果不需要使用 Integration Events API，可以將 `apiKey` 參數設為空字串。

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
├── main.go                        # 主客戶端和初始化
├── http.go                        # HTTP 請求處理
├── endpoints.go                   # API 端點定義
├── alerts.go                      # 警報管理
├── integration_events.go          # Integration Events API 管理
├── contacts.go                    # 聯絡人管理
├── ...                            # 其他資源管理
├── alert/                         # 警報相關結構
│   ├── integration_events_request.go  # Integration Events API 請求結構
│   └── ...                        # 其他警報相關結構
├── contacts/                      # 聯絡人相關結構
└── examples/                      # 使用範例
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

## 版本歷史

### v2.1.0
- ✅ 恢復 Basic Authentication 用於普通 API
- ✅ 添加 Integration Events API 支持（使用 GenieKey 認證）
- ✅ 支援 4 個 Integration Events API 端點：CreateAlert, AcknowledgeAlert, CloseAlert, AddNote

### v2.0.0
- ✅ 更新模組路徑為 `/v2` 以支持 Go 模組版本管理
- ✅ 改用 API Integration 認證方式（GenieKey）

## 相關連結

- [Jira Service Management Operations API 文檔](https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-operations/)
- [Integration Events API 文檔](https://developer.atlassian.com/cloud/jira/service-desk-ops/rest/v1/api-group-integration-events/)
