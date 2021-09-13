package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type FilterRequest struct {
	Text string `json:"text"`
}

type ComResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type Result struct {
	Suggestion      string   `json:"suggestion"`
	SensitiveWords  []string `json:"sensitiveWords,omitempty"`
	Desensitization string   `json:"desensitization,omitempty"`
}

func FilterHandler(c *gin.Context) {
	//httpserver底层将panic捕获处理掉了，这里不需要再处理了;
	body, _ := c.GetRawData()
	var request FilterRequest

	if err := json.Unmarshal(body, &request); err != nil {
		log.Warningf("httpServerHandler parse json fail : %s\n", err.Error())

		resp := &ComResp{
			Code: 300,
			Msg:  "Parse parameter error!",
		}
		c.JSON(200, resp)
		log.Warnf("body : %s, resp : %v", string(body), resp)
		return
	}

	if len(request.Text) == 0 {
		resp := &ComResp{
			Code: 300,
			Msg:  "request.Text is empty!",
		}
		c.JSON(200, resp)
		log.Warnf("body : %s, resp : %v", string(body), resp)
		return
	}

	var result Result
	result.Suggestion = "pass"
	result.SensitiveWords, result.Desensitization = GMatchService.Match(request.Text)
	if len(result.SensitiveWords) > 0 {
		result.Suggestion = "block"
	}

	resp := &ComResp{
		Code: 200,
		Msg:  "success",
		Data: &result,
	}

	log.Infof("request: %v, result: %v", request, resp.Data)

	c.JSON(200, resp)
}

func httpServerRun(port int) {
	router := gin.Default()
	router.POST("/words/filter", FilterHandler)

	address := ":" + strconv.Itoa(port)
	log.Info("Start http server at ", address)
	if err := router.Run(address); err != nil {
		fmt.Printf("Start http server %s fail : %s", "8088", err.Error())
		log.Fatal(err)
	}
}
