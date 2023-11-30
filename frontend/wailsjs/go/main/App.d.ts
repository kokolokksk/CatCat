// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function API_GetDanmuInfo(arg1:number):Promise<main.ResultDanmuInfo>;

export function API_GetLiveUserInfo(arg1:number):Promise<main.ResultLiveUserInfo>;

export function API_GetQRLoginInfo():Promise<main.ResultQRLoginInfo>;

export function API_GetQRLoginStatus(arg1:string):Promise<main.ResultQRLoginStatus>;

export function API_GetRoomInfo(arg1:number):Promise<main.ResultRoomInfo>;

export function API_GetRoomInit(arg1:number):Promise<main.ResultRoomInit>;

export function GetConfig():Promise<{[key: string]: any}>;

export function Greet(arg1:string):Promise<string>;

export function LoadDanmakuEvents(arg1:{[key: string]: any},arg2:{[key: string]: any}):Promise<string>;

export function Log(arg1:string,arg2:string):Promise<void>;

export function SavePic(arg1:string):Promise<string>;

export function SetConfig(arg1:{[key: string]: any}):Promise<void>;
