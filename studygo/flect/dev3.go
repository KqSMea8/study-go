package main

import (
	"fmt"
	"reflect"
)


type Data struct {
	AppMsType    int16    //AppMsType 是端rd同学确认填写的值
	MessageQuote *Message //appmsContent 消息数据
}

type MsgData struct {
	Mid          int64         //消息id，要唯一
	Sound        string        //声音
	AppMsType    int16
	AppInContent string
	AppOutInfo   refineApnsMsg
}

//RefineAppMsg 重新定义端外数据 ［特殊］
//端外结构体
type refineApnsMsg struct {
	Title string
	Text string
	Content string
}

//Message 消息体格式，最终会以二进制往msggate推送
//此结构体的包含业务信息，这是一种约定
type Message struct {
	Title      string `json:"title,omitempty"`      //大标题
	SubTitle   string `json:"subTitle,omitempty"`   //小标题
}

func main(){
	u := Data{}
	uu := MsgData{}
	t := reflect.TypeOf(u)
	v := reflect.TypeOf(uu)
	if v.Name()=="MsgData" {
		fmt.Println("ok")
	}
	fmt.Println(t.Name(),v.Name())

}
