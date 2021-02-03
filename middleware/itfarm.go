package middleware

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/pretty66/gosdk"
//	"online_file/app/controller/base"
//	"online_file/utils"
//	"online_file/utils/constant"
//)
//
//func InitAppkeyChannel(c *gin.Context) {
//	server, err := gosdk.GetServerInstance(c.Request.Header)
//	if err != nil {
//		c.Abort()
//		panic(err)
//	}
//	base.Appkey = server.GetAppKey()
//	channelStr := server.GetChannel()
//	if base.Appkey == "" || channelStr == "" {
//		c.Abort()
//		panic(utiliy.NewCustomError(constant.InvalidAppkeyChannelMsg, constant.InvalidAppkeyChannelCode))
//	}
//	channel, err := utiliy.StringToInt(channelStr)
//	if err != nil {
//		c.Abort()
//		panic(err)
//	}
//	base.Channel = channel
//	base.AccountId = server.GetAccountId()
//	c.Next()
//}
