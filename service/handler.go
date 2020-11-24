package service

import "linking/linking-go-agile/model"

type Handler interface {
	TestFunc(req *model.LRequest) *model.LResult
	DoFunc(req *model.LRequest) *model.LResult
}
