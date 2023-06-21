package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kangness/shangdaren_server/config"
	"github.com/kangness/shangdaren_server/model"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// HandlerHttpRequest
func HandlerHttpRequest(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	headerStr, _ := json.Marshal(headers)
	fmt.Println("http request", string(headerStr))
	var response *model.SDRResponse
	response = &model.SDRResponse{}
	response.Code = 0
	response.Msg = "OK"
	response.Data = string(headerStr)
	/*
		requestUrl := r.URL.Path
		var body []byte
		var err error
		if r.Method == "POST" {
			body, err = ioutil.ReadAll(r.Body)
			if err != nil {
				log.Print(err)
				return
			}
		}
		switch requestUrl {
		case "/api/getCount":
			response, err = handlerGetCounter(nil, r, body)
		case "/api/setCount":
			response, err = handlerSetCounter(nil, r, body)
		default:
			err = fmt.Errorf("cmd not found")
		}
		if err != nil && response == nil {
			response = &model.SDRResponse{}
			response.Code = 500
			response.Msg = err.Error()
		}*/
	if response != nil {
		resp, _ := json.Marshal(response)
		log.Println("response ", string(resp))
		w.Write(resp)
		return
	}
}

func handlerGetCounter(ctx context.Context, r *http.Request, body []byte) (*model.SDRResponse, error) {
	request := &model.SDRGetCounterRequest{}
	if len(body) > 0 {
		if err := json.Unmarshal(body, request); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if request.Id <= 0 {
		request.Id = 1
	}
	tc := &model.TCounters{
		Id:        11,
		Count:     22,
		Createdat: time.Now().Format("2006-01-02 15:04:05"),
		Updatedat: time.Now().Format("2006-01-02 15:04:05"),
	}
	/*
		tc, err := model.FetchCounterInfoById(config.DB, request.Id)
		if err != nil {
			log.Error(err)
		}
	*/
	response := &model.SDRResponse{
		RequestId: "",
		Code:      0,
		Msg:       "OK",
		Data:      nil,
	}
	resp := &model.SDRGetCounterResponse{}
	if tc == nil || tc.Id <= 0 {
		log.Println("not found any data")
		return response, nil
	}
	resp.Id = tc.Id
	resp.Count = tc.Count
	resp.CreateTime = tc.Createdat
	resp.UpdateTime = tc.Updatedat
	headerBin, _ := json.Marshal(r.Header)
	resp.Ext = string(headerBin)
	response.Data = resp
	return response, nil
}

// handlerSetCounter ...
func handlerSetCounter(ctx context.Context, r *http.Request, body []byte) (*model.SDRResponse, error) {
	request := &model.SDRSetCounterRequest{}
	if len(body) > 0 {
		if err := json.Unmarshal(body, request); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	if request.Id <= 0 {
		request.Id = 1
	}
	if len(request.Action) <= 0 {
		request.Action = "inc"
	}
	if request.Value <= 0 {
		request.Value = int64(rand.Intn(100))
	}
	tc, _ := model.FetchCounterInfoById(config.DB, request.Id)
	if tc == nil {
		tc = &model.TCounters{}
	}
	tc.Id = request.Id
	if request.Action == "inc" {
		tc.Count = tc.Count + request.Value
	} else if request.Action == "clear" {
		tc.Count = 0
	} else if request.Action == "dec" {
		tc.Count = tc.Count - request.Value
	}
	response := &model.SDRResponse{}
	if err := model.UpdateOrderCreateCounter(config.DB, tc); err != nil {
		log.Println(err)
		response.Code = 500
		response.Msg = "数据库出错"
		return response, err
	}
	resp := &model.SDRSetCounterResponse{}
	if tc == nil || tc.Id <= 0 {
		log.Println("not found any data")
		return response, nil
	}
	resp.Id = tc.Id
	resp.Count = tc.Count
	resp.CreateTime = tc.Createdat
	resp.UpdateTime = tc.Updatedat
	headerBin, _ := json.Marshal(r.Header)
	resp.Ext = string(headerBin)
	response.Data = resp
	return response, nil
}
