/**
 * @Title  控制器
 * @Description 控制接收HTTP请求后的处理
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:27
 */
package controller

import (
	"github.com/ganeryao/linking-go-agile/common"
	"github.com/ganeryao/linking-go-agile/pojo"
	"github.com/ganeryao/linking-go-agile/pojo/dto"
	"github.com/ganeryao/linking-go-agile/protos"
	"strings"
)

type BaseController struct {
}

func (c *BaseController) ConvertRequest(param string) *dto.ReqDTO {
	var lRequest = &protos.LRequest{}
	common.ConvertRequest(param, lRequest)
	arr := strings.Split(lRequest.Api, ".")
	api := &pojo.ApiBO{Module: arr[len(arr)-2], Name: arr[len(arr)-1]}
	return &dto.ReqDTO{Request: lRequest, Api: api}
}
