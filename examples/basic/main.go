package main

import (
	"fmt"
	"log"
	"os"

	"github.com/circleyu/go-jsmops"
	"github.com/circleyu/go-jsmops/contacts"
	"github.com/circleyu/go-jsmops/teams"
	"github.com/sirupsen/logrus"
)

func main() {
	// 從環境變數獲取認證信息
	cloudID := os.Getenv("JIRA_CLOUD_ID")
	apiToken := os.Getenv("JIRA_API_TOKEN")
	userName := os.Getenv("JIRA_USERNAME")

	if cloudID == "" || apiToken == "" || userName == "" {
		log.Fatal("請設置環境變數: JIRA_CLOUD_ID, JIRA_API_TOKEN, JIRA_USERNAME")
	}

	// 初始化客戶端（帶日誌）
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	options := &jsmops.ClientOptions{
		Level:  jsmops.LogInfo,
		Logger: logger,
	}

	client := jsmops.Init(cloudID, apiToken, userName, options)

	// 範例 1: 列出所有團隊
	fmt.Println("=== 列出所有團隊 ===")
	teamsResult, err := client.Teams.ListTeams(&teams.ListTeamsRequest{})
	if err != nil {
		log.Fatalf("獲取團隊列表失敗: %v", err)
	}

	fmt.Printf("找到 %d 個團隊:\n", len(teamsResult.Teams))
	for _, team := range teamsResult.Teams {
		fmt.Printf("  - %s (ID: %s)\n", team.Name, team.ID)
	}

	// 範例 2: 列出聯絡人
	fmt.Println("\n=== 列出聯絡人 ===")
	contactsReq := &contacts.ListContactsRequest{
		Offset: 0,
		Size:   10,
	}

	contactsResult, err := client.Contacts.ListContacts(contactsReq)
	if err != nil {
		log.Fatalf("獲取聯絡人列表失敗: %v", err)
	}

	fmt.Printf("找到 %d 個聯絡人:\n", len(contactsResult.Contacts))
	for _, contact := range contactsResult.Contacts {
		fmt.Printf("  - %s (ID: %s, 方法: %s, 啟用: %v)\n",
			contact.Value, contact.ID, contact.Method, contact.Enabled)
	}

	// 範例 3: 獲取特定聯絡人（如果有的話）
	if len(contactsResult.Contacts) > 0 {
		fmt.Println("\n=== 獲取特定聯絡人 ===")
		contactID := contactsResult.Contacts[0].ID
		contact, err := client.Contacts.GetContact(&contacts.GetContactRequest{ID: contactID})
		if err != nil {
			log.Printf("獲取聯絡人失敗: %v", err)
		} else {
			fmt.Printf("聯絡人詳情: %+v\n", contact)
		}
	}

	fmt.Println("\n完成！")
}

