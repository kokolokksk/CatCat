package main

// roomid: number;
//   clientId?: string;
//   ttsDanmu?: boolean;
//   ttsGift?: boolean;
//   ttsKey?: string;
//   ttsServerUrl?: string;
//   ttsServerToken?: string;
//   alwaysOnTop?: boolean;
//   catdb?: boolean;
//   dmTs?: string;
//   SESSDATA?: string;
//   csrf?: string;
//   v1?: string;
//   v2?: string;
//   fansDisplay?: string;
//   darkMode?: boolean;
//   proxyApi?: boolean;
//   sessionId?: string;
//   started?: boolean;
//   count: number;
//   wave?: boolean;
//   theme: string;
//   danmuDir?: string;
//   [propName: string]: number | object | 0 | boolean | string | undefined;
type Config struct {
	Roomid         int
	ClientId       string
	TtsDanmu       bool
	TtsGift        bool
	TtsKey         string
	TtsServerUrl   string
	TtsServerToken string
	AlwaysOnTop    bool
	Catdb          bool
	DmTs           string
	SESSDATA       string
	Csrf           string
	V1             string
	V2             string
	FansDisplay    string
	DarkMode       bool
	ProxyApi       bool
	SessionId      string
	Started        bool
	Count          int
	Wave           bool
	Theme          string
	DanmuDir       string
}
