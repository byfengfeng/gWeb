package common

import (
	"errors"
	"github.com/Byfengfeng/gWeb/enum"
)

type WebBase struct {
	ReqMap map[string]map[enum.ReqType]func(pkt map[string][]string)interface{}
}

func NewWebBase() WebBase {
	return WebBase{ReqMap: make(map[string]map[enum.ReqType]func(pkt map[string][]string)interface{})}
}

func (w *WebBase) Bind(url string,reqFn map[enum.ReqType]func(pkt map[string][]string) interface{})  {
	_,ok := w.ReqMap[url]
	if ok {
		panic(errors.New("err Repeat binding request"))
	}
	reqMap := make(map[enum.ReqType]func(pkt map[string][]string)interface{})
	for reqType,fun := range reqFn {
		reqMap[reqType] = fun
	}
	if len(reqMap) > 0 {
		w.ReqMap[url] = reqMap
	}

}
