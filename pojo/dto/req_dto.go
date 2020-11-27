package dto

import (
	"linking/linking-go-agile/pojo"
	"linking/linking-go-agile/protos"
)

type ReqDTO struct {
	Request *protos.LRequest
	Api     *pojo.ApiBO
}
