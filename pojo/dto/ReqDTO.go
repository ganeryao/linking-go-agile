package dto

import (
	"linking/linking-go-agile/model"
	"linking/linking-go-agile/pojo"
)

type ReqDTO struct {
	Request *model.LRequest
	Api     *pojo.ApiBO
}
