package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type BiliBiliLive struct {
	Status  int                    `json:"status"`
	Config  map[string]interface{} `json:"config"`
	MuaConf map[string]interface{} `json:"muaconf"`
	Ws      struct {
		Url string `json:"url"`
	}
}

func OnBiliBiliLive(blive BiliBiliLive) chan string {

	OnBiliLiveChannel := make(chan string)

	go func() {
		if blive.Status == 0 {
			// live start
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)
			host := fmt.Sprintf("%v", blive.Config["host"])
			port := fmt.Sprintf("%v", blive.Config["port"])
			u := url.URL{Scheme: "ws", Host: host + ":" + port, Path: "/"}
			fmt.Println("connecting to ", u.String())
			c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
			if err != nil {
				log.Fatal("dial:", err)
			}
			defer c.Close()
			done := make(chan struct{})
			go func() {
				defer close(done)
				for {
					_, message, err := c.ReadMessage()
					if err != nil {
						log.Println("read:", err)
						return
					}
					fmt.Printf("Received: %s\n", message)
				}
			}()
			for {
				select {
				case <-interrupt:
					fmt.Println("Interrupt received, closing connection.")
					err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					if err != nil {
						log.Println("write close:", err)
						return
					}
					select {
					case <-done:
						OnBiliLiveChannel <- "done"
					}
					return
				}
			}
		} else {
			//get ws
		}
		for {
			OnBiliLiveChannel <- "Message every " + time.Second.String() + " seconds"
			time.Sleep(5 * time.Second)
		}
	}()
	return OnBiliLiveChannel
}
