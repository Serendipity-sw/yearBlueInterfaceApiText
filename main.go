package main

import (
	"flag"
	"github.com/guotie/config"
	"strings"
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
)

var (
	configFn    = flag.String("config", "./config.json", "config file path") //配置文件地址
	interFaceApi string //创蓝接口地址
	account string //创蓝使用账号
	passWord string //创蓝使用密码
)

/**
构造函数
创建人:邵炜
创建时间:2016年8月22日15:12:00
 */
func main() {
	config.ReadCfg(*configFn)
	interFaceApi=strings.TrimSpace(config.GetString("interFaceApi"))
	account=strings.TrimSpace(config.GetString("account"))
	passWord=strings.TrimSpace(config.GetString("passWord"))
	sendMessage()
}

/**
短信下发接口
创建人:邵炜
创建时间:2016年8月22日15:17:42
 */
func sendMessage() {
	dataValue:=url.Values{}
	dataValue.Add("account",account)
	dataValue.Add("pswd",passWord)
	dataValue.Add("mobile","")
	dataValue.Add("msg","您好，您的验证码是123456")
	dataValue.Add("needstatus","false")
	resp,err:=http.PostForm(interFaceApi,dataValue)
	if err != nil {
		fmt.Printf("send message error! err: %s \n",err.Error())
		return
	}
	jsonStr,err:=ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("send message read request content error! err: %s \n",err.Error())
		return
	}
	fmt.Println(string(jsonStr))
}