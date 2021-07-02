package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

var json1 = `
{
    "producer_account_name": "zbeosbp11111",
    "producer_public_key": "EOS7rhgVPWWyfMqjSbNdndtCK8Gkza3xnDbUupsPLMZ6gjfQ4nX81",
    "org": {
        "candidate_name": "zb eos",
        "website": "https://www.zbeos.com",
        "code_of_conduct": "https://steemit.com/zbeos/@zbcom/zbeos-or-code-of-conduct-statement",
        "ownership_disclosure": "https://steemit.com/zbeos/@zbcom/zbeos-or-ownership-disclosure",
        "email": "zbeos@zb.com",
        "branding": {
            "logo_256": "https://www.zbeos.com/img/zbeos-256.png",
            "logo_1024": "https://www.zbeos.com/img/zbeos-1024.png",
            "logo_svg": "https://www.zbeos.com/img/zbeoslogo.svg"
        },
        "location": {
            "name": "HK, China",
            "country": "CN",
            "latitude": 22.55,
            "longitude": 114.26
        },
        "social": {
            "steemit": "@zbcom",
            "twitter": "ZB__EOS",
            "youtube": "ZB EOS",
            "facebook": "ZBEOS",
            "github": "ZBEOS",
            "reddit": "ZBEOS",
            "keybase": "ZBEOS",
            "telegram": "ZBEOS",
            "wechat": "ZB_EOS"
        }
    },
    "nodes": [
        {
            "location": {
                "name": "ZBnode1",
                "country": "Singapore",
                "latitude": 1.53,
                "longitude": 103.74
            },
            "node_type": "full",
            "p2p_endpoint": "node1.zbeos.com:9876",
            "api_endpoint": "https://node1.zbeos.com",
            "ssl_endpoint": "https://node1.zbeos.com"
        },
        {
            "location": {
                "name": "ZBnode2",
                "country": "Singapore",
                "latitude": 1.53,
                "longitude": 103.74
            },
            "node_type": "full",
            "p2p_endpoint": "node2.zbeos.com:9876",
            "ssl_endpoint": "https://node2.zbeos.com"
        }
    ]
}
`

type A struct {
	ProducerAccountName string `json:"producer_account_name"`
	ProducerPublicKey   string `json:"producer_public_key"`
	Org                 Org    `json:"org"`
	Nodes               []Node `json:"nodes"`
}
type Org struct {
	CandidateName       string   `json:"candidate_name"`
	WebSite             string   `json:"website"`
	CodeOfConduct       string   `json:"code_of_conduct"`
	OwnershipDisclosure string   `json:"ownership_disclosure"`
	Email               string   `json:"email"`
	Branding            Branding `json:"branding"`
	Location            Location `json:"location"`
	//Social string `json:"social"`
}

type Node struct {
	NodeType    string   `json:"node_type"`
	P2pEndpoint string   `json:"p2p_endpoint"`
	ApiEndpoint string   `json:"api_endpoint"`
	SslEndpoint string   `json:"ssl_endpoint"`
	Location    Location `json:"location"`
}

type Location struct {
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
}

type Branding struct {
	Logo256  string `json:"logo_256"`
	Logo1024 string `json:"logo_1024"`
	LogoSvg  string `json:"logo_svg"`
}

func main() {

	//2021-12-01
	//随机后缀 80
	// 202201219999 69546313
	//1 202212019999 54353336
	//sj := fmt.Sprintf("%d%02d%02d%d", 2022, 01, 21, 9999)
	sj := fmt.Sprintf("%d%02d%02d%d", 2022, 12, 1, 9999)
	sjZz, _ := strconv.Atoi(sj)
	r := rand.New(rand.NewSource(int64(sjZz)))
	//几位就几个9
	sjm := r.Intn(99999999)
	fmt.Println(sj, sjm)
	//
	//去当前时间和日期
	//1年365天    10年3600天  100年36000
	//当前年月日
	currYear := time.Now().Year()
	currMoth := int(time.Now().Month())
	currDay := time.Now().Day()
	//var sjNew int64
	//30 要按照当前的月是多少天来 月和年自己实现
	var flag bool
	var tempYear, tmpMoth int
	for { //年
		if !flag {
			tempYear = currYear
			tmpMoth = currMoth
			flag = true
		} else {
			tmpMoth = 1
			tempYear = tempYear + 1
		}
		for j := tmpMoth; j <= 12; j++ {
			yueDays := count(tempYear, j)
			for i := currDay; i <= yueDays; i++ {
				sj := fmt.Sprintf("%d%02d%02d%d", tempYear, j, i, 9999)
				sjZz, _ := strconv.Atoi(sj)
				r := rand.New(rand.NewSource(int64(sjZz)))
				//几位就几个9
				sjm1 := r.Intn(99999999)
				//2022 12 1 99999
				//2022 12 1 99999
				//2022 1 21 99999
				fmt.Println("=====:", sjZz, tempYear, j, i, sjm1, sjm)
				if sjm1 == sjm {
					fmt.Println(tempYear, j, i)
					goto END
				}
			}
			currDay = 1
		}
	}
END:
	//for i := currDay; i < 30; i++ {
	//	sjNew = int64(currYear + currMoth + 1 + 80)
	//	r := rand.New(rand.NewSource(sjNew))
	//	if r.Intn(99999999) == sjm {
	//		break
	//	}
	//}
	return
	//f := Get("http://www.baidu.com")
	//fmt.Println(f)
	//fmt.Println(strings.NewReader("123").Len())
	//fmt.Println(strings.NewReader("123").Size())
	a := &Branding{}
	unmarshal := json.Unmarshal([]byte(json1), a)
	fmt.Println(unmarshal)
	//for _, node := range a.Nodes {
	//	if node.NodeType == "full" {
	//		fmt.Println(node.Location.Country)
	//		break
	//	}
	//}
	//fmt.Println(a.Org.Branding)
	//layout := "2006-01-02 15:04:05"

	// just one second
	//t1, _ := time.Parse(layout, "21:59:59")
	//t2, _ := time.Parse(layout, "00:00:00")
	//fmt.Println(timeSub(t1, t2))
	//fmt.Println(time.Now().Hour())
	Post("")
}

func count(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
			//fmt.Fprintln(os.Stdout, "The month has 31 days")
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return
}

type eosParkReq struct {
	Id   string `json:"id"`
	Json bool   `json:"json"`
}

func Post(url string) {
	eosReq := eosParkReq{
		Id:   "cce517a177944c32721a6effcbbd262a5aa9607a89e2d7e268cfc3f3165b90be",
		Json: true,
	}
	body, err := SendRequest("Post", eosReq, "https://mainnet-history.meet.one/v1/history/get_transaction")

	fmt.Println(string(body), err)
}
func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
func timeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func timeSubDays(t1, t2 time.Time) int {

	if t1.Location().String() != t2.Location().String() {
		return -1
	}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {

			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24)
		} else { // more than 24 hours
			return int(hours/24) + 1
		}
	}
}

var EOSClient = &http.Client{
}

func SendRequest(method string, t_req interface{}, req_url string) (body []byte, err error) {

	//通常是采用strings.NewReader函数，将一个string类型转化为io.Reader类型，或者bytes.NewBuffer函数，将[]byte类型转化为io.Reader类型。

	req_buf, _ := json.Marshal(t_req)
	req_byte := bytes.NewBuffer(req_buf)

	req, err := http.NewRequest("GET", "http://baidu.com", nil)
	if t_req == nil {
		req, err = http.NewRequest(method, req_url, nil)
	} else {
		req, err = http.NewRequest(method, req_url, req_byte)
	}
	if err != nil {
		//Logger.Error("SendRequest  failed.  req_buf:", string(req_buf), "req_url:", req_url, " err:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := EOSClient.Do(req)
	if err != nil && resp == nil {
		//Logger.Error("SendRequest  failed.  req_buf:", string(req_buf), " req_url:", req_url, " err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		//Logger.Error("SendRequest  failed.  req_buf:", string(req_buf), "req_url:", req_url, " err:", err)
		return nil, err
	}

	//Logger.Info(" SendRequest response:", string(body))

	return body, nil
}
