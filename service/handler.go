package service

import "linking/linking-go-agile/protos"

type Handler interface {
	TestFunc(req *protos.LRequest) *protos.LResult
	DoFunc(req *protos.LRequest) *protos.LResult
}
