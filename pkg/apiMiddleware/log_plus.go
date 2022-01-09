package apiMiddleware

import (
	"ControlCenter/pkg/utils"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type reqLog struct {
	Header http.Header `json:"header"`
	Body   string      `json:"body"`
	URL    string      `json:"URL"`
	Resp   string      `json:"resp"`
	ReqID  string      `json:"req-id"`
	In     time.Time   `json:"in"`
	Out    time.Time   `json:"out"`
	Method string      `json:"method"`
}

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LogPlusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		var r reqLog
		r.ReqID = utils.RandomString()
		r.Header = c.Request.Header
		rawReqData, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawReqData))

		r.Body = string(rawReqData)
		r.URL = c.Request.URL.RequestURI()
		customWriter := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = customWriter
		// 处理请求
		c.Next()

		r.Resp = customWriter.body.String()
		r.Out = time.Now()
		log(&r)
	}
}

func log(req *reqLog) {
	logByte, err := jsoniter.Marshal(req)
	if err != nil {
		return
	}
	fmt.Printf("[%s] %s %s\n%s", time.Now().Format("2006-01-02 03:04:05"), req.Method, req.URL, string(logByte))
}
