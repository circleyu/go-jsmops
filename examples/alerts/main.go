package main

import (
	"fmt"
	"log"
	"os"

	"github.com/circleyu/go-jsmops/v2"
	"github.com/circleyu/go-jsmops/v2/alert"
	"github.com/sirupsen/logrus"
)

func main() {
	// 從環境變數獲取認證信息
	cloudID := os.Getenv("JIRA_CLOUD_ID")
	apiToken := os.Getenv("JIRA_API_TOKEN")
	userName := os.Getenv("JIRA_USERNAME")
	apiKey := os.Getenv("JIRA_API_KEY") // Optional, for Integration Events API

	if cloudID == "" || apiToken == "" || userName == "" {
		log.Fatal("請設置環境變數: JIRA_CLOUD_ID, JIRA_API_TOKEN, JIRA_USERNAME")
	}

	// 初始化客戶端
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	options := &jsmops.ClientOptions{
		Level:  jsmops.LogInfo,
		Logger: logger,
	}

	client := jsmops.Init(cloudID, apiToken, userName, apiKey, options)

	// 範例 1: 列出警報
	fmt.Println("=== 列出警報 ===")
	listReq := &alert.ListAlertsRequest{
		Limit: 10,
	}

	alertsResult, err := client.Alert.ListAlerts(listReq)
	if err != nil {
		log.Fatalf("獲取警報列表失敗: %v", err)
	}

	fmt.Printf("找到 %d 個警報\n", len(alertsResult.Alerts))
	for i, alertItem := range alertsResult.Alerts {
		if i >= 5 { // 只顯示前 5 個
			break
		}
		fmt.Printf("  %d. %s (ID: %s, 優先級: %s)\n",
			i+1, alertItem.Message, alertItem.Id, alertItem.Priority)
	}

	// 範例 2: 創建新警報
	fmt.Println("\n=== 創建新警報 ===")
	createReq := &alert.CreateAlertRequest{
		Message:     "這是一個測試警報",
		Alias:       "test-alert-" + fmt.Sprintf("%d", os.Getpid()),
		Description: "由 go-jsmops 範例程式創建",
		Priority:    "P3",
		Responders: []alert.Responder{
			{
				Type: "team",
				Id:   "your-team-id", // 請替換為實際的團隊 ID
			},
		},
		Tags: []string{"test", "example"},
		Source: "go-jsmops-example",
	}

	createResult, err := client.Alert.CreateAlert(createReq)
	if err != nil {
		log.Printf("創建警報失敗: %v", err)
		log.Println("提示: 請確保提供了有效的團隊 ID 和其他必要參數")
	} else {
		fmt.Printf("警報創建成功！請求 ID: %s\n", createResult.RequestID)
		
		// 注意：創建警報後需要通過 RequestID 查詢狀態來獲取 AlertID
		// 這裡我們假設有一個警報 ID 用於演示
		// 實際使用中，您需要先查詢請求狀態
		fmt.Println("\n提示: 創建警報後，請使用 RequestID 查詢狀態以獲取 AlertID")
		
		// 如果我們有警報 ID，可以繼續其他操作
		// 這裡我們使用一個示例 ID（實際使用中應該從請求狀態中獲取）
		exampleAlertID := "example-alert-id"
		
		// 範例 3: 獲取警報詳情（如果有警報 ID）
		fmt.Println("\n=== 獲取警報詳情 ===")
		getReq := &alert.GetAlertRequest{
			ID: exampleAlertID,
		}

		alertDetail, err := client.Alert.GetAlert(getReq)
		if err != nil {
			log.Printf("獲取警報詳情失敗: %v (這是正常的，因為我們使用的是示例 ID)", err)
		} else {
			fmt.Printf("警報詳情:\n")
			fmt.Printf("  ID: %s\n", alertDetail.Id)
			fmt.Printf("  訊息: %s\n", alertDetail.Message)
			fmt.Printf("  優先級: %s\n", alertDetail.Priority)
			fmt.Printf("  狀態: %s\n", alertDetail.Status)
		}

		// 範例 4: 添加警報備註（如果有警報 ID）
		fmt.Println("\n=== 添加警報備註 ===")
		noteReq := &alert.AddNoteRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: exampleAlertID,
			Note:            "這是一個測試備註",
		}

		noteResult, err := client.Alert.AddAlertNote(noteReq)
		if err != nil {
			log.Printf("添加備註失敗: %v (這是正常的，因為我們使用的是示例 ID)", err)
		} else {
			fmt.Printf("備註添加成功！備註 ID: %s\n", noteResult.ID)
		}

		// 範例 5: 列出警報備註（如果有警報 ID）
		fmt.Println("\n=== 列出警報備註 ===")
		listNotesReq := &alert.ListAlertNotesRequest{
			ID:   exampleAlertID,
			Size: 10,
		}

		notesResult, err := client.Alert.ListAlertNotes(listNotesReq)
		if err != nil {
			log.Printf("列出備註失敗: %v (這是正常的，因為我們使用的是示例 ID)", err)
		} else {
			fmt.Printf("找到 %d 個備註:\n", len(notesResult.Notes))
			for i, note := range notesResult.Notes {
				fmt.Printf("  %d. %s (由 %s 於 %s 創建)\n",
					i+1, note.Note, note.Owner, note.CreatedAt.Format("2006-01-02 15:04:05"))
			}
		}
	}

	fmt.Println("\n完成！")
}

