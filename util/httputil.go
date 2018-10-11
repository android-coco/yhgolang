package util

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"errors"
	"strconv"
)

func YHGet(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("yh get no response. addr:%s err: %s", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("yh get no response. addr:%s err: %s", url, err)
			return nil, err
		}
		return all,nil
	}
	return nil,errors.New("stauscode is no ok, code is " + strconv.Itoa(resp.StatusCode))
}
