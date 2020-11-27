package dto

import (
	"github.com/ganeryao/linking-go-agile/pojo"
	"github.com/ganeryao/linking-go-agile/protos"
)

type ReqDTO struct {
	Request *protos.LRequest
	Api     *pojo.ApiBO
}
