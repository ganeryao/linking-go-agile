package service

import "github.com/ganeryao/linking-go-agile/protos"

type Handler interface {
	TestFunc(req *protos.LRequest) *protos.LResult
	DoFunc(req *protos.LRequest) *protos.LResult
}
