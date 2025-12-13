# 使用範例

本目錄包含 go-jsmops 庫的使用範例。

## 環境設置

在運行範例之前，請設置以下環境變數：

```bash
export JIRA_CLOUD_ID="your-cloud-id"
export JIRA_API_TOKEN="your-api-token"
export JIRA_USERNAME="your-username@example.com"
```

### 獲取 Jira Cloud ID

1. 登入您的 Jira 實例
2. 訪問 `https://admin.atlassian.com/`
3. 選擇您的組織
4. 在設置中找到 Cloud ID

### 獲取 API Token

1. 訪問 `https://id.atlassian.com/manage-profile/security/api-tokens`
2. 點擊「Create API token」
3. 複製生成的 token

## 範例說明

### basic/main.go

基本操作範例，展示如何：
- 初始化客戶端
- 列出團隊
- 列出和獲取聯絡人

運行方式：
```bash
cd examples/basic
go run main.go
```

### alerts/main.go

警報管理範例，展示如何：
- 列出警報
- 創建新警報
- 獲取警報詳情
- 添加和列出警報備註

運行方式：
```bash
cd examples/alerts
go run main.go
```

**注意**：創建警報需要有效的團隊 ID，請在程式碼中替換 `your-team-id`。

### schedules/main.go

排程管理範例，展示如何：
- 列出排程
- 獲取排程詳情
- 列出待命響應者
- 列出下一個待命響應者

運行方式：
```bash
cd examples/schedules
go run main.go
```

### integrations/main.go

整合管理範例，展示如何：
- 列出整合
- 獲取整合詳情
- 列出整合操作
- 獲取整合警報過濾器

運行方式：
```bash
cd examples/integrations
go run main.go
```

## 自定義範例

您可以參考這些範例來創建自己的程式。主要步驟：

1. 導入必要的包
2. 設置環境變數或直接提供認證信息
3. 初始化客戶端
4. 使用對應的 manager 調用 API 方法

## 錯誤處理

所有 API 方法都會返回錯誤。建議始終檢查錯誤：

```go
result, err := client.Alert.ListAlerts(req)
if err != nil {
    log.Fatalf("操作失敗: %v", err)
}
```

## 更多資源

- [主 README](../README.md)
- [Jira Service Management Operations API 文檔](https://developer.atlassian.com/cloud/jira/service-desk/rest/api-group-operations/)

