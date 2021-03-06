package controllers

import (
	"github.com/astaxie/beego"
	"models"
	"fmt"
	"encoding/xml"
	//"time"
	"time"
)

type PubController struct  {
	beego.Controller
}


//TODO 这里负责回调模式的验证
func (c * PubController) Get() {
	var signature models.Signature
	if err := c.ParseForm(&signature) ; err != nil{
		Lg(err,beego.LevelNotice)
		c.Abort("400")
	}

	fmt.Println(signature.Echostr)
	c.Ctx.WriteString(signature.Echostr)

}




func (c * PubController) Post(){
	var msgIn models.PubTextMsg
	err := xml.Unmarshal(c.Ctx.Input.RequestBody,&msgIn)
	if err != nil {
		Lg(err)
		c.Abort("400")
		return
	}
	if(msgIn.MsgType == "event"){
		msgback := "感谢您的关注(O w O)～～这里是gochat框架机器人"
		_ = c.PubSendBack(msgback,msgIn)

	}

	msgback := "这里是自动回复（O w O）"


	_ = c.PubSendBack(msgback,msgIn)

}

func (this * PubController)PubSendBack(backMsg string,msgIn models.PubTextMsg)error{
	msgOut := models.PubTextOut{
		ToUserName:msgIn.FromUserName,
		FromUserName:msgIn.ToUserName,
		CreateTime:time.Now().Unix(),
		MsgType:"text",
		Content:fmt.Sprint(backMsg),
	}

	xmlData ,err := msgOut.ToXml()
	this.Ctx.WriteString(string(xmlData))
	return err
}



