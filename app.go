package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/nfnt/resize"
	"github.com/tucnak/store"
)

// App struct
type App struct {
	ctx  context.Context
	conf map[string]interface{}
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	loadConfig(a)
	testBilibliApi()
}

func testBilibliApi() {
	danmuInfo := GetDanmuInfo(19864)
	fmt.Println(danmuInfo)
}

func (a *App) API_GetDanmuInfo(realRoomId int) ResultDanmuInfo {
	return GetDanmuInfo(realRoomId)
}
func (a *App) API_GetRoomInfo(realRoomId int) ResultRoomInfo {
	return GetRoomInfo(realRoomId)
}
func (a *App) API_GetRoomInit(realRoomId int) ResultRoomInit {
	return GetRoomInit(realRoomId)
}
func (a *App) API_GetLiveUserInfo(uid int) ResultLiveUserInfo {
	return GetLiveUserInfo(uid)
}
func (a *App) API_GetQRLoginStatus(oauthKey string, contentType string, body map[string]interface{}) ResultQRLoginStatus {
	return GetQRLoginStatus(oauthKey, contentType, body)
}
func (a *App) API_GetQRLoginInfo() ResultQRLoginInfo {
	return GetQRLoginInfo()
}

func (a *App) SavePic(faceImg string) string {
	base64String, err := saveFileFromRemote(faceImg)
	if err != nil {
		fmt.Println("Error:", err)
		return base64String
	}
	return base64String
}

func saveFileFromRemote(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return "", err
	}

	// Resize the image to desired dimensions
	resizedImg := resize.Resize(256, 0, img, resize.Lanczos2)

	// Convert the resized image to bytes
	var compressedImageBytes bytes.Buffer
	err = jpeg.Encode(&compressedImageBytes, resizedImg, nil)
	if err != nil {
		return "", err
	}

	// Convert to Base64
	base64String := base64.StdEncoding.EncodeToString(compressedImageBytes.Bytes())
	return base64String, nil
}

func loadConfig(a *App) {
	// load config
	store.Init("CatCat")
	var config = make(map[string]interface{})
	if err := store.Load("config.json", &config); err != nil {
		fmt.Println(err)
	}
	a.conf = config
	// print config
	fmt.Println("config:", config)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetConfig() map[string]interface{} {
	return a.conf
}

func (a *App) SetConfig(config map[string]interface{}) {
	a.conf = config
	store.Save("config.json", config)
}

func (a *App) OnLive() {
	a.conf["Started"] = true
	store.Save("config.json", a.conf)
}

func (a *App) Log(str string, level string) {
	if level == "info" {
		log.Println(str)
	} else if level == "error" {
		log.Fatalln(str)
	}
}
func (a *App) shutdown(ctx context.Context) {
	config := a.conf
	a.SetConfig(config)
}
