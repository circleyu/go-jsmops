package main

import (
	"fmt"
	"log"
	"os"

	"github.com/circleyu/go-jsmops/v2"
	"github.com/circleyu/go-jsmops/v2/integrations"
	"github.com/circleyu/go-jsmops/v2/integrations/actions"
	"github.com/circleyu/go-jsmops/v2/integrations/filters"
	"github.com/sirupsen/logrus"
)

func main() {
	// 從環境變數獲取認證信息
	cloudID := os.Getenv("JIRA_CLOUD_ID")
	apiKey := os.Getenv("JIRA_API_KEY")

	if cloudID == "" || apiKey == "" {
		log.Fatal("請設置環境變數: JIRA_CLOUD_ID, JIRA_API_KEY")
	}

	// 初始化客戶端
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	options := &jsmops.ClientOptions{
		Level:  jsmops.LogInfo,
		Logger: logger,
	}

	client := jsmops.Init(cloudID, apiKey, options)

	// 範例 1: 列出所有整合
	fmt.Println("=== 列出所有整合 ===")
	listReq := &integrations.ListIntegrationsRequest{
		Size:   10,
		Offset: 0,
	}

	integrationsResult, err := client.Integrations.ListIntegrations(listReq)
	if err != nil {
		log.Fatalf("獲取整合列表失敗: %v", err)
	}

	fmt.Printf("找到 %d 個整合:\n", len(integrationsResult.Integrations))
	for i, integration := range integrationsResult.Integrations {
		if i >= 5 { // 只顯示前 5 個
			break
		}
		fmt.Printf("  %d. %s (ID: %s, 類型: %s)\n",
			i+1, integration.Name, integration.ID, integration.Type)
	}

	// 如果有整合，繼續其他操作
	if len(integrationsResult.Integrations) > 0 {
		integrationID := integrationsResult.Integrations[0].ID

		// 範例 2: 獲取整合詳情
		fmt.Println("\n=== 獲取整合詳情 ===")
		getReq := &integrations.GetIntegrationRequest{
			ID: integrationID,
		}

		integrationDetail, err := client.Integrations.GetIntegration(getReq)
		if err != nil {
			log.Printf("獲取整合詳情失敗: %v", err)
		} else {
			fmt.Printf("整合詳情:\n")
			fmt.Printf("  名稱: %s\n", integrationDetail.Name)
			fmt.Printf("  類型: %s\n", integrationDetail.Type)
			fmt.Printf("  團隊 ID: %s\n", integrationDetail.TeamID)
		}

		// 範例 3: 列出整合操作
		fmt.Println("\n=== 列出整合操作 ===")
		listActionsReq := &actions.ListIntegrationActionsRequest{
			IntegrationID: integrationID,
			Size:          10,
			Offset:        0,
		}

		actionsResult, err := client.IntegrationActions.ListIntegrationActions(listActionsReq)
		if err != nil {
			log.Printf("獲取整合操作列表失敗: %v", err)
		} else {
			fmt.Printf("找到 %d 個整合操作:\n", len(actionsResult.Actions))
			for i, action := range actionsResult.Actions {
				if i >= 5 { // 只顯示前 5 個
					break
				}
				fmt.Printf("  %d. %s (ID: %s, 類型: %s, 順序: %d)\n",
					i+1, action.Name, action.ID, action.Type, action.Order)
			}
		}

		// 範例 4: 獲取整合警報過濾器
		fmt.Println("\n=== 獲取整合警報過濾器 ===")
		getFilterReq := &filters.GetIntegrationAlertFilterRequest{
			IntegrationID: integrationID,
		}

		filter, err := client.IntegrationFilters.GetIntegrationAlertFilter(getFilterReq)
		if err != nil {
			log.Printf("獲取整合警報過濾器失敗: %v", err)
		} else {
			fmt.Printf("過濾器詳情:\n")
			if filter.Filter != nil {
				fmt.Printf("  過濾器配置: %+v\n", filter.Filter)
			} else {
				fmt.Printf("  無過濾器配置\n")
			}
		}
	} else {
		fmt.Println("\n沒有找到整合，跳過其他操作")
	}

	fmt.Println("\n完成！")
}

