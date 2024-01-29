package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"pchat/utils"
	"strings"
	"time"
)

type AccessLog struct {
	Method        string    `json:"method"`
	Url           string    `json:"url"`
	RemoteAddress string    `json:"remoteAddr"`
	RemotePort    string    `json:"remotePort"`
	StatusCode    int       `json:"responseStatus"`
	ResponseTime  int64     `json:"responseTime"`
	Referer       string    `json:"referer"`
	UserAgent     string    `json:"userAgent"`
	Body          string    `json:"body"`
	ContentLength int       `json:"responseBodySize"`
	RequestId     string    `json:"reqId"`
	Host          string    `json:"host,omitempty"`
	UserId        string    `json:"userId"`
	StartTime     time.Time `json:"-"`
	ResponseBody  string    `json:"responseBody"`
}

func initAccessLog(ctx *gin.Context) AccessLog {
	body, _ := io.ReadAll(ctx.Request.Body)
	io.NopCloser(ctx.Request.Body)
	return AccessLog{
		Body: string(body),
	}
}

func (accessLog AccessLog) Record(ctx *gin.Context) {
	// TODO: read from nginx header
	remoteAddr, remotePort, _ := net.SplitHostPort(strings.TrimSpace(ctx.Request.RemoteAddr))
	accessLog.Method = ctx.Request.Method
	accessLog.Url = ctx.Request.URL.RequestURI()
	accessLog.RemoteAddress = remoteAddr
	accessLog.RemotePort = remotePort
	accessLog.StatusCode = ctx.Writer.Status()
	accessLog.Referer = ctx.Request.Referer()
	accessLog.UserAgent = ctx.Request.UserAgent()
	accessLog.ContentLength = ctx.Writer.Size()
	accessLog.Host = ctx.Request.Host
	accessLog.ResponseTime = time.Now().UnixMilli() - accessLog.StartTime.UnixMilli()
	accessLog.RequestId = utils.GetRequestId(ctx)
	accessLog.UserId = utils.GetUserId(ctx)
	if accessLog.StatusCode >= 400 {
		accessLog.Body = utils.GetResponseBody(ctx)
	}
	accessLog.print()
}

func (a AccessLog) print() {
	logrus.Println(utils.MarshalInterfaceToString(a))
}

func accessLog(ctx *gin.Context) {
	accessLog := initAccessLog(ctx)
	ctx.Next()
	accessLog.Record(ctx)
}
