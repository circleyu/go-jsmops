package main

import (
	"fmt"
	"log"
	"os"

	"github.com/circleyu/go-jsmops"
	"github.com/circleyu/go-jsmops/schedules"
	"github.com/circleyu/go-jsmops/schedules/oncalls"
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

	// 範例 1: 列出所有排程
	fmt.Println("=== 列出所有排程 ===")
	listReq := &schedules.ListSchedulesRequest{
		Size:   10,
		Offset: 0,
	}

	schedulesResult, err := client.Schedules.ListSchedules(listReq)
	if err != nil {
		log.Fatalf("獲取排程列表失敗: %v", err)
	}

	fmt.Printf("找到 %d 個排程:\n", len(schedulesResult.Schedules))
	for i, schedule := range schedulesResult.Schedules {
		if i >= 5 { // 只顯示前 5 個
			break
		}
		fmt.Printf("  %d. %s (ID: %s, 時區: %s)\n",
			i+1, schedule.Name, schedule.ID, schedule.Timezone)
	}

	// 如果有排程，繼續其他操作
	if len(schedulesResult.Schedules) > 0 {
		scheduleID := schedulesResult.Schedules[0].ID

		// 範例 2: 獲取排程詳情
		fmt.Println("\n=== 獲取排程詳情 ===")
		getReq := &schedules.GetScheduleRequest{
			ID: scheduleID,
		}

		scheduleDetail, err := client.Schedules.GetSchedule(getReq)
		if err != nil {
			log.Printf("獲取排程詳情失敗: %v", err)
		} else {
			fmt.Printf("排程詳情:\n")
			fmt.Printf("  名稱: %s\n", scheduleDetail.Name)
			fmt.Printf("  描述: %s\n", scheduleDetail.Description)
			fmt.Printf("  時區: %s\n", scheduleDetail.Timezone)
			fmt.Printf("  啟用: %v\n", scheduleDetail.Enabled)
		}

		// 範例 3: 列出待命響應者
		fmt.Println("\n=== 列出待命響應者 ===")
		onCallReq := &oncalls.ListOnCallRespondersRequest{
			ScheduleID: scheduleID,
		}

		onCallResult, err := client.SchedulesOnCalls.ListOnCallResponders(onCallReq)
		if err != nil {
			log.Printf("獲取待命響應者失敗: %v", err)
		} else {
			fmt.Printf("找到 %d 個待命響應者:\n", len(onCallResult.Responders))
			for i, responder := range onCallResult.Responders {
				if i >= 5 { // 只顯示前 5 個
					break
				}
				fmt.Printf("  %d. %s (開始: %s, 結束: %s)\n",
					i+1, responder.Username, responder.Start, responder.End)
			}
		}

		// 範例 4: 列出下一個待命響應者
		fmt.Println("\n=== 列出下一個待命響應者 ===")
		nextOnCallReq := &oncalls.ListNextOnCallRespondersRequest{
			ScheduleID: scheduleID,
		}

		nextOnCallResult, err := client.SchedulesOnCalls.ListNextOnCallResponders(nextOnCallReq)
		if err != nil {
			log.Printf("獲取下一個待命響應者失敗: %v", err)
		} else {
			fmt.Printf("找到 %d 個下一個待命響應者:\n", len(nextOnCallResult.Responders))
			for i, responder := range nextOnCallResult.Responders {
				if i >= 5 { // 只顯示前 5 個
					break
				}
				fmt.Printf("  %d. %s (開始: %s, 結束: %s)\n",
					i+1, responder.Username, responder.Start, responder.End)
			}
		}
	} else {
		fmt.Println("\n沒有找到排程，跳過其他操作")
	}

	fmt.Println("\n完成！")
}

