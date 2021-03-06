package httpGet

import (
	"chatroom/helper"
	"chatroom/utils/Constants"
	"chatroom/utils/JSON"
	"io/ioutil"
	"net/http"
	"strconv"
	"log"
)

type UserInfo struct {
	Id         int    `json:"id"` //进入直播的时候到wings获取用户信息
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	IsAuthor   bool   `json:"isAuthor"`
	Subscribed bool   `json:"subscribed"`
	FansNum		int		`json:"fansNum"`	//粉丝值
	FansTitle	string 	`json:"fansTitle"`
	FansValue	int		`json:"fansValue"`
	ShutUp		bool	`json:"shutUp"`
}

type Result struct {
	Code    int       `json:"code"`
	Type    string    `json:"type"`
	Message string    `json:"message"`
}

type UserResult struct {
	Result
	Data    *UserInfo `json:"data"`
}

const (
	SUCCESS = 1
	ERROR   = 0
)

//用于增删改直播，判断用户是否是作者
func CheckAuthorRight(cookies []*http.Cookie, bookId int) (*UserResult, helper.ErrorType) {
	info := &UserResult{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", Constants.HOST+"/system/room/author/check?bookId="+strconv.Itoa(bookId), nil)
	for _, v := range cookies {
		req.AddCookie(v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return info, helper.NetworkError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return info, helper.IOError
	}
	result := string(body)
	log.Println(result)
	if err = JSON.ParseToStruct(result, info); err != nil {
		log.Println(err)
		return info, helper.DataFormatError
	}
	if info.Code == ERROR {
		return info, helper.GetWingsErrorType(info.Type)
	}
	log.Println(info)
	return info, helper.NoError
}

//get logined user info
func GetLoginUserInfo(cookies []*http.Cookie, roomId int64) (*UserResult, helper.ErrorType) {
	info := &UserResult{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", Constants.HOST+"/system/room/login/info?roomId="+helper.Itoa64(roomId), nil)
	for _, v := range cookies {
		req.AddCookie(v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return info, helper.NetworkError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return info, helper.IOError
	}
	result := string(body)
	log.Println(result)
	if err = JSON.ParseToStruct(result, info); err != nil {
		log.Println(err)
		return info, helper.DataFormatError
	}
	if info.Code == ERROR {
		return info, helper.GetWingsErrorType(info.Type)
	}
	log.Println(info)
	return info, helper.NoError
}

func GetUserInfo(bookId int, userId int) (*UserResult, helper.ErrorType) {
	info := &UserResult{}
	resp, err := http.Get(Constants.HOST + "/system/room/user/info?userId=" + strconv.Itoa(userId) + "&bookId=" + strconv.Itoa(bookId))
	if err != nil {
		log.Println(err)
		return info, helper.NetworkError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return info, helper.IOError
	}
	result := string(body)
	log.Println(result)
	if err = JSON.ParseToStruct(result, info); err != nil {
		log.Println(err)
		return info, helper.DataFormatError
	}
	log.Println(info)
	if info.Code == ERROR {
		return info, helper.GetWingsErrorType(info.Type)
	}
	return info, helper.NoError
}

func BuyRoom(cookies []*http.Cookie, roomId int64, money int, bookId int) (*UserResult, helper.ErrorType) {
	info := &UserResult{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", Constants.HOST+"/system/room/buy?roomId="+helper.Itoa64(roomId)+"&money="+strconv.Itoa(money)+"&bookId=" + strconv.Itoa(bookId), nil)
	for _, v := range cookies {
		req.AddCookie(v)
	}
	resp, err := client.Do(req)
	if resp.Body == nil {
		return info, helper.NetworkError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return info, helper.IOError
	}
	result := string(body)
	log.Println(result)
	if err = JSON.ParseToStruct(result, info); err != nil {
		log.Println(err)
		return info, helper.DataFormatError
	}
	if info.Code == ERROR {
		return info, helper.GetWingsErrorType(info.Type)
	}
	log.Println(info)
	return info, helper.NoError
}

func GetUserIdFromCookie(cookies []*http.Cookie) (int, error) {
	for _, v := range cookies {
		if v.Name == "ud" {
			return strconv.Atoi(v.Value)
		}
	}
	return 0, nil
}
