package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Akegarasu/blivedm-go/client"
	_ "github.com/Akegarasu/blivedm-go/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type BiliBiliLive struct {
	Status  int                    `json:"status"`
	client  *client.Client         `json:"client"`
	Config  map[string]interface{} `json:"config"`
	MuaConf map[string]interface{} `json:"muaconf"`
	Ws      struct {
		Url string `json:"url"`
	}
}

func OnBiliBiliLive(a *App, blive BiliBiliLive, even string) chan string {
	OnAllCmdChannel := make(chan string)
	OnDanmakuChannel := make(chan string)
	OnNoticeMsgChannel := make(chan string)

	go func() {
		if a.blive.Status == 0 {
			roomIDStr, ok := blive.Config["roomid"].(string)
			if !ok {
				fmt.Println("Error: roomid is not a string")
				return
			}

			// 将字符串转换为整数
			roomID, err := strconv.Atoi(roomIDStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			c := client.NewClient(roomID) // 房间号'
			config := blive.Config
			c.SetCookie(fmt.Sprintf("DedeUserID=%s; SESSDATA=%s; bili_jct=%s", config["DedeUserID"], config["SESSDATA"], config["bili_jct"]))
			c.RegisterCustomEventHandler("DANMU_MSG", func(s string) {
				fmt.Println("DANMU_MSG:", s)
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				//OnDanmakuChannel <- string(data)
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("SEND_GIFT", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				OnAllCmdChannel <- string(data)
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
			})
			c.RegisterCustomEventHandler("GUARD_BUY", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("USER_TOAST_MSG", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("NOTICE_MSG", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				//OnNoticeMsgChannel <- string(data)
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("SUPER_CHAT_MESSAGE", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("INTERACT_WORD", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			err = c.Start()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("started")
			a.blive.Status = 1
			a.blive.client = c
			select {}
		} else {
			c := a.blive.client
			c.Stop()
			roomIDStr, ok := blive.Config["roomid"].(string)
			if !ok {
				fmt.Println("Error: roomid is not a string")
				return
			}

			// 将字符串转换为整数
			roomID, err := strconv.Atoi(roomIDStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			c = client.NewClient(roomID) // 房间号'
			config := blive.Config
			c.SetCookie(fmt.Sprintf("DedeUserID=%s; SESSDATA=%s; bili_jct=%s", config["DedeUserID"], config["SESSDATA"], config["bili_jct"]))
			fmt.Println("BiliBiliLive is running")
			c.RegisterCustomEventHandler("DANMU_MSG", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				//OnDanmakuChannel <- string(data)
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("SEND_GIFT", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("GUARD_BUY", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("USER_TOAST_MSG", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("NOTICE_MSG", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				//OnNoticeMsgChannel <- string(data)
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("SUPER_CHAT_MESSAGE", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			c.RegisterCustomEventHandler("INTERACT_WORD", func(s string) {
				data, err := json.Marshal(s)
				if err != nil {
					fmt.Println("Error marshaling danmaku:", err)
				}
				runtime.EventsEmit(a.ctx, "OnMsg", string(data))
				OnAllCmdChannel <- string(data)
			})
			err = c.Start()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("started")
			a.blive.Status = 1
			a.blive.client = c
			select {}
		}

	}()
	switch even {
	case "OnDanmakuChannel":
		return OnDanmakuChannel
	case "NOTICE_MSG":
		return OnNoticeMsgChannel
	case "ALL":
		return OnAllCmdChannel
	default:
		return OnAllCmdChannel
	}
	//{"cmd":"DANMU_MSG","info":[[0,1,25,16777215,1701241864951,258547100,0,"87b6f788",0,0,0,"",0,{},{},{"mode":0,"show_player_type":0,"extra":{"send_from_me":false,"mode":0,"color":16777215,"dm_type":0,"font_size":25,"player_mode":1,"show_player_type":0,"content":"vvb","user_hash":"2276915080","emoticon_unique":"","bulge_display":0,"recommend_score":4,"main_state_dm_color":"","objective_state_dm_color":"","direction":0,"pk_direction":0,"quartet_direction":0,"anniversary_crowd":0,"yeah_space_type":"","yeah_space_url":"","jump_to_url":"","space_type":"","space_url":"","animation":{},"emots":null,"is_audited":false,"id_str":"64411cb3a589b4642c6d05846b6566e42","icon":null,"show_reply":true,"reply_mid":0,"reply_uname":"","reply_uname_color":"","hit_combo":0}},{"activity_identity":"","activity_source":0,"not_show":0},0],"vvb",[0,"菈***",0,0,0,10000,1,""],[23,"导盲猫","诺子喵呜",21654925,1725515,"",0,12632256,12632256,12632256,0,0,10276136],[20,0,6406234,"\\u003e50000",1],["",""],0,0,null,{"ts":1701241864,"ct":"E2BAE56A"},0,0,null,null,0,210,[22],null],"dm_v2":""}

}
