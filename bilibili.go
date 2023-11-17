package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BILIBILI_API_GET_DANMU_INFO = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%d"
const BILIBILI_API_GET_ROOM_INFO = "http://api.live.bilibili.com/room/v1/Room/get_info?room_id=%d"
const BILIBILI_API_GET_ROOM_INIT = "https://api.live.bilibili.com/room/v1/Room/room_init?id=%d"
const BILIBILI_API_LIVE_USER_INFO = "https://api.live.bilibili.com/live_user/v1/Master/info?uid=%d"
const BILIBILI_API_QR_LOGIN_INFO = "https://passport.bilibili.com/qrcode/getLoginInfo?oauthKey=%d"
const BILIBILI_API_QR_LOGIN_URL = "https://passport.bilibili.com/qrcode/getLoginUrl"
type ResultDanmuInfo struct {
	Code    int
	Message string
	Ttl     int
	Data    struct {
		Group              string
		Bussiness_id       int
		Refresh_row_factor float64
		Refresh_rate       int
		Max_delay          int
		Token              string
		Host_list          []Host
	}
}
type Host struct {
	Host     string
	Port     int
	Wss_port int
	Ws_port  int
}

func GetDanmuInfo(realRoomId int) ResultDanmuInfo {
	url := fmt.Sprintf(BILIBILI_API_GET_DANMU_INFO, realRoomId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var danmuInfo ResultDanmuInfo
	json.Unmarshal(body, &danmuInfo)
	return danmuInfo
}

type ResultRoomInfo struct {
	Code    int
	Message string
	Data    struct {
		Uid                     int
		Room_id                 int
		Short_id                int
		Attention               int
		Online                  int
		Is_portrait             bool
		Description             string
		Live_status             int
		Area_id                 int
		Parent_area_id          int
		Parent_area_name        string
		Old_area_id             int
		Background              string
		Title                   string
		User_cover              string
		Keyframe                string
		Is_strict_room          bool
		Live_time               string
		Tags                    string
		Is_anchor               int
		Room_silent_type        string
		Room_silent_level       int
		Room_silent_second      int
		Area_name               string
		Pendants                string
		Area_pendants           string
		Hot_words               []string
		Hot_words_status        int
		Verify                  string
		New_pendants            NewPendants
		Up_session              string
		Pk_status               int
		Pk_id                   int
		Battle_id               int
		Allow_change_area_time  int
		Allow_upload_cover_time int
		Studio_info             StudioInfo
	}
}
type NewPendants struct {
	Frame        Frame
	Badge        interface{}
	Mobile_frame MobileFrame
	Mobile_badge interface{}
}
type Frame struct {
	Name         string
	Value        string
	Position     int
	Desc         string
	Area         int
	Area_old     int
	Bg_color     string
	Bg_pic       string
	Use_old_area bool
}
type MobileFrame struct {
	Name         string
	Value        string
	Position     int
	Desc         string
	Area         int
	Area_old     int
	Bg_color     string
	Bg_pic       string
	Use_old_area bool
}
type StudioInfo struct {
	Status      int
	Master_list []interface{}
}

func GetRoomInfo(realRoomId int) ResultRoomInfo {
	url := fmt.Sprintf(BILIBILI_API_GET_ROOM_INFO, realRoomId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var roomInfo ResultRoomInfo
	json.Unmarshal(body, &roomInfo)
	return roomInfo
}

type ResultRoomInit struct {
	Code    int
	Message string
	Data    struct {
		Room_id      int
		Short_id     int
		Uid          int
		Need_p2p     int
		Is_hidden    bool
		Is_locked    bool
		Is_portrait  bool
		Live_status  int
		Hidden_till  int
		Lock_till    int
		Encrypted    bool
		Pwd_verified bool
		Live_time    int
		Room_shield  int
		Is_sp        int
		Special_type int
	}
}

func GetRoomInit(realRoomId int) ResultRoomInit {
	url := fmt.Sprintf(BILIBILI_API_GET_ROOM_INIT, realRoomId)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result ResultRoomInit
	json.Unmarshal(body, &result)
	return result
}

type ResultLiveUserInfo struct {
	Code    int
	Message string
	Data    struct {
		Info struct {
			Uid             int
			Uname           string
			Face            string
			Official_verify struct {
				Type int
				Desc string
			}
			Gender int
		}
		Exp struct {
			Master_level struct {
				Level   int
				Color   int
				Current []int
				Next    []int
			}
		}
		Follower_num   int
		Room_id        int
		Medal_name     string
		Glory_count    int
		Pendant        string
		Link_group_num int
		Room_news      struct {
			Content    string
			Ctime      string
			Ctime_text string
		}
	}
}

func GetLiveUserInfo(uid int) ResultLiveUserInfo {
	url := fmt.Sprintf(BILIBILI_API_LIVE_USER_INFO, uid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result ResultLiveUserInfo
	json.Unmarshal(body, &result)
	return result
}


type ResultQRLoginInfo struct {
	Code    int
	Status  bool
	Ts      int
	Data  struct {
		Url string
		OauthKey string
	}
}

func GetQRLoginInfo() ResultQRLoginInfo {
	resp, err := http.Get(BILIBILI_API_QR_LOGIN_URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result ResultQRLoginInfo
	json.Unmarshal(body, &result)
	return result
}

{
    "status": false,
    "data": -4,
    "message": "Can't scan~"
}
or
{
	"code": 0,
	"status": true,
	"ts": 1583315474,
	"data": {
		"url": "https://passport.biligame.com/crossDomain?DedeUserID=***&DedeUserID__ckMd5=***&Expires=***&SESSDATA=***&bili_jct=***&gourl=http%3A%2F%2Fwww.bilibili.com"
	}
}

type ResultQRLoginStatus struct {
	Status  bool   `json:"status"`
	Data    interface{} `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Ts      int64  `json:"ts"`
}

type Data struct {
	Url string `json:"url"`
}
func GetQRLoginStatus(oauthKey int) ResultQRLoginStatus {
	url := fmt.Sprintf(BILIBILI_API_QR_LOGIN_INFO, oauthKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result ResultQRLoginStatus
	json.Unmarshal(body, &result)
	return result
}