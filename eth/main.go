package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/crypto/sha3"
)

func TextHash(data []byte) []byte {
	hash, _ := TextAndHash(data)
	return hash
}
func TextAndHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(msg))
	return hasher.Sum(nil), msg
}
//https://test-resource.ycandyz.com/webapi/resource/taskinfo?task_id=48
//127.0.0.1:10201
type msg struct {
	Code int64 `json:"code"`
}
func main() {
	//fmt.Println(strings.Compare("Jason", "As"))
	//go func() {
	//	for {
	//		req, _ := http.NewRequest("GET", "http://127.0.0.1:20201/webapi/resource/tasklist?page_size=1&search_key=&sort_type=&sort_name=&sort_rule=&page=1", nil)
	//		// 比如说设置个token
	//		req.Header.Set("PD-Token", "1eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjE2NzIyODg1MzE3NTk4MDAzMiIsInVzZXJuYW1lIjoiIiwiZXhwIjoxNjIxNTEwNTQxLCJpc3MiOiJhcGktY2VudGVyIn0.0ehd3iWR8LWtU2BNR4zfYiJxXPrCFix-tE4q1eHUncE")
	//		// 再设置个json
	//		req.Header.Set("Content-Type","application/json")
	//
	//
	//		resp, err := (&http.Client{}).Do(req)
	//		//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	//		if err != nil {
	//			fmt.Println("query topic failed", err.Error())
	//		}
	//		defer resp.Body.Close()
	//
	//		respByte, _ := ioutil.ReadAll(resp.Body)
	//		//fmt.Println("tasklist:",string(respByte))
	//
	//		var msgData msg
	//		json.Unmarshal(respByte,&msgData)
	//		if msgData.Code != 200 {
	//			fmt.Println(time.Now(),"tasklist:",string(respByte))
	//		}
	//		time.Sleep(10 * time.Millisecond)
	//
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		req, _ := http.NewRequest("GET", "http://127.0.0.1:20201/webapi/resource/tasklist?page_size=1&search_key=&sort_type=&sort_name=&sort_rule=&page=1", nil)
	//		// 比如说设置个token
	//		req.Header.Set("PD-Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjE2NzIyODg1MzE3NTk4MDAzMiIsInVzZXJuYW1lIjoiIiwiZXhwIjoxNjIxNTEwNTQxLCJpc3MiOiJhcGktY2VudGVyIn0.0ehd3iWR8LWtU2BNR4zfYiJxXPrCFix-tE4q1eHUncE")
	//		// 再设置个json
	//		req.Header.Set("Content-Type","application/json")
	//
	//
	//		resp, err := (&http.Client{}).Do(req)
	//		//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	//		if err != nil {
	//			fmt.Println("query topic failed", err.Error())
	//		}
	//		defer resp.Body.Close()
	//
	//		respByte, _ := ioutil.ReadAll(resp.Body)
	//		//fmt.Println("tasklist1:",string(respByte))
	//
	//		var msgData msg
	//		json.Unmarshal(respByte,&msgData)
	//		if msgData.Code != 200 {
	//			fmt.Println(time.Now(),"tasklist:",string(respByte))
	//		}
	//		time.Sleep(10 * time.Millisecond)
	//
	//	}
	//}()
	for {
		go func() {
			for i := 0; i < 100 ; i++ {
				start := time.Now()
				taskrefresh()
				fmt.Println("taskrefresh:",time.Now().Sub(start))

				//start = time.Now()
				//taskinfo()
				//fmt.Println("taskinfo:",time.Now().Sub(start))
			}
		}()

		//taskinfo()
		time.Sleep(1 * time.Second)
	}
}

func taskinfo () {

	//https://tuiguang.kehulian.cn
	//http://127.0.0.1:20201
	//https://uat-demande.ycandyz.com
	req, _ := http.NewRequest("GET", "https://uat-demande.ycandyz.com/webapi/resource/taskinfo/list?task_id=50&page=1&page_size=10", nil)
	// 比如说设置个token
	req.Header.Set("PD-Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjE3NTI2NDU1Mjk1MzMxOTQyNCIsInVzZXJuYW1lIjoiIiwiZXhwIjoxNjIyNTE1ODAyLCJpc3MiOiJhcGktY2VudGVyIn0.SH8TBgFzTDpF6oG4LyFOgGpmQ5fuspFZ0fXrQkDPlXM")
	// 再设置个j
	req.Header.Set("Content-Type","application/json")


	resp, err := (&http.Client{}).Do(req)
	//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	if err != nil {
		fmt.Println("query topic failed", err.Error())
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("taskinfo:",string(respByte))

	var msgData msg
	json.Unmarshal(respByte,&msgData)
	if msgData.Code != 200 {
		fmt.Println(time.Now(),"taskinfo:",string(respByte))
	}
}


func tasklist () {
	//https://tuiguang.kehulian.cn
	//http://127.0.0.1:20201
	//https://uat-demande.ycandyz.com
	req, _ := http.NewRequest("GET", "https://uat-demande.ycandyz.com/webapi/demander/task/list?search_key=&start_time=&end_time=&page_size=10&state=99&begin_time=", nil)
	// 比如说设置个token
	req.Header.Set("PD-Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjE3NTI2NDU1Mjk1MzMxOTQyNCIsInVzZXJuYW1lIjoiIiwiZXhwIjoxNjIyMTg5MzYwLCJpc3MiOiJhcGktY2VudGVyIn0.42uVf4XybWhUFmgmdRLZ9-pu9gP7L_bW4ofS7exRsEM")
	// 再设置个j
	req.Header.Set("Content-Type","application/json")


	resp, err := (&http.Client{}).Do(req)
	//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	if err != nil {
		fmt.Println("query topic failed", err.Error())
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("tasklist:",string(respByte))

	var msgData msg
	json.Unmarshal(respByte,&msgData)
	if msgData.Code != 200 {
		fmt.Println(time.Now(),"tasklist:",string(respByte))
	}
}



func taskrefresh () {
	//https://tuiguang.kehulian.cn
	//http://127.0.0.1:20201
	//https://uat-demande.ycandyz.com
	req, _ := http.NewRequest("GET", "http://127.0.0.1:20201/webapi/resource/taskrefresh?task_ids=68,67,66,63,62,61,60,59", nil)
	// 比如说设置个token
	req.Header.Set("PD-Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjE2NzIyODg1MzE3NTk4MDAzMiIsInVzZXJuYW1lIjoiIiwiZXhwIjoxNjIzMjU2NjIyLCJpc3MiOiJhcGktY2VudGVyIn0.Qu5mC3EX3D6o7rYYr4BwUSQxupAyzG63gqxHvdEkfoo")
	// 再设置个j
	req.Header.Set("Content-Type","application/json")


	resp, err := (&http.Client{}).Do(req)
	//resp, err := http.Get(serviceUrl + "/topic/query/false/lsj")
	if err != nil {
		fmt.Println("query topic failed", err.Error())
	}
	defer resp.Body.Close()

	respByte, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("tasklist:",string(respByte))

	var msgData msg
	json.Unmarshal(respByte,&msgData)
	if msgData.Code != 200 {
		fmt.Println(time.Now(),"tasklist:",string(respByte))
		panic(200)
	}
	//fmt.Println(time.Now(),"tasklist:",string(respByte))
}