package main

import (
	"fmt"
	"math"
	"time"
)

type NokiaPhone struct {
}

type Iphone struct {
}

type I_Phone interface {
	call()
	sms()
}

func (iphone Iphone) call() {
	fmt.Println("my is Iphone call")
}

func (iphone NokiaPhone) sms() {
	fmt.Println("my is NokiaPhone sms")
}

func (iphone NokiaPhone) call() {
	fmt.Println("my is NokiaPhone call")
}

func (iphone Iphone) sms() {
	fmt.Println("my is Iphone sms")
}

//排序
type Person struct {
	Name string
	Age  int64
}
type Persons []Person

// 获取此 slice 的长度
func (p Persons) Len() int { return len(p) }

// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (p Persons) Less(i, j int) bool {
	return p[i].Age > p[j].Age
}

// 交换数据
func (p Persons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

const emailReg = `^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$`

var jsonStr1 = `{"account":"betdicelucky","abi":"0e656f73696f3a3a6162692f312e30010474696d650675696e7433320807647261776c6f6700050b6163636f756e744e616d65046e616d65066e756d6265720675696e74333205636f756e740675696e743332067265776172640561737365740474696d650474696d650975736572706f696e7400030b6163636f756e744e616d65046e616d6505706f696e740675696e7433320c6c6173746672656574696d650474696d65047573657200030b6163636f756e744e616d65046e616d6504736565640b636865636b73756d3235360474696d650474696d650b64726177726563656970740006046e616d65046e616d6504736565640b636865636b73756d323536097369676e6174757265097369676e61747572650b6c75636b794e756d6265720675696e7433320877696e4173736574056173736574086472617754696d650474696d650a6472617772657665616c0002046e616d65046e616d65097369676e6174757265097369676e6174757265046472617700030466726f6d046e616d6501620675696e74363401630675696e7436340a63616e63656c6472617700020475736572046e616d6506726561736f6e06737472696e67056f6e62657400030475736572046e616d6508636f6e7472616374046e616d65056173736574056173736574050072750aa9cbcd4d0b6472617772656365697074000040346aabcbcd4d0a6472617772657665616c000000000000c0cd4d046472617700000037374585a6410a63616e63656c64726177000000000080accea4056f6e626574000300000080d1c8cd4d03693634010b6163636f756e744e616d6501046e616d6507647261776c6f670000c8d3d17a15d603693634010b6163636f756e744e616d6501046e616d650975736572706f696e7400000000007015d603693634010b6163636f756e744e616d6501046e616d650475736572000000"}`

type wait struct {
	WaitSec int `json:"wait_sec"`
	Weight  int `json:"weight"`
}

type key struct {
	Key    string `json:"key"`
	Weight int    `json:"weight"`
}
type permission struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
type account struct {
	Permission permission `json:"permission"`
	Weight     int        `json:"weight"`
}
type requiredAuth struct {
	Threshold int       `json:"threshold"`
	Keys      []key     `json:"keys"`
	Accounts  []account `json:"accounts"`
	Waits     []wait    `json:"waits"`
}

type permissionItem struct {
	PermName     string       `json:"perm_name"`
	Parent       string       `json:"parent"`
	RequiredAuth requiredAuth `json:"required_auth"`
}
type linkedPermissions struct {
	ActionKey      string `json:"action_key"`
	PermissionName string `json:"permission_name"`
}

type permissions struct {
	Permissions       []permissionItem    `json:"permissions"`
	LinkedPermissions []linkedPermissions `json:"linked_permissions"`
}

var json_permission = `[
        {
            "perm_name": "active",
            "parent": "owner",
            "required_auth": {
                "threshold": 4,
                "keys": [],
                "accounts": [
                    {
                        "permission": {
                            "actor": "eoscanadaaaa",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaab",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaac",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaad",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaae",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaaf",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaag",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaah",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaai",
                            "permission": "active"
                        },
                        "weight": 1
                    }
                ],
                "waits": [
                    {
                        "wait_sec": 10800,
                        "weight": 1
                    }
                ]
            }
        },
        {
            "perm_name": "blacklistops",
            "parent": "active",
            "required_auth": {
                "threshold": 1,
                "keys": [
                    {
                        "key": "EOS7idX86zQ6M3mrzkGQ9MGHf4btSECmcTj4i8Le59ga7CpSpZYy5",
                        "weight": 1
                    }
                ],
                "accounts": [],
                "waits": []
            }
        },
        {
            "perm_name": "claimer",
            "parent": "active",
            "required_auth": {
                "threshold": 1,
                "keys": [
                    {
                        "key": "EOS7NFuBesBKK5XHHLtzFxm7S57Eq11gUtndrsvq3Mt3XZNMTHfqc",
                        "weight": 1
                    }
                ],
                "accounts": [],
                "waits": []
            }
        },
        {
            "perm_name": "day2day",
            "parent": "active",
            "required_auth": {
                "threshold": 1,
                "keys": [],
                "accounts": [
                    {
                        "permission": {
                            "actor": "eoscanadaaaa",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaac",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaaf",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaag",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaah",
                            "permission": "active"
                        },
                        "weight": 1
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaai",
                            "permission": "active"
                        },
                        "weight": 1
                    }
                ],
                "waits": []
            }
        },
        {
            "perm_name": "eosforumdapp",
            "parent": "active",
            "required_auth": {
                "threshold": 1,
                "keys": [
                    {
                        "key": "EOS7YNS1swh6QWANkzGgFrjiX8E3u8WK5CK9GMAb6EzKVNZMYhCH3",
                        "weight": 1
                    }
                ],
                "accounts": [],
                "waits": []
            }
        },
        {
            "perm_name": "owner",
            "parent": "",
            "required_auth": {
                "threshold": 5,
                "keys": [],
                "accounts": [
                    {
                        "permission": {
                            "actor": "eoscanadaaaa",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaab",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaac",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaad",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaae",
                            "permission": "active"
                        },
                        "weight": 2
                    },
                    {
                        "permission": {
                            "actor": "eoscanadaaaf",
                            "permission": "active"
                        },
                        "weight": 1
                    }
                ],
                "waits": [
                    {
                        "wait_sec": 86400,
                        "weight": 1
                    },
                    {
                        "wait_sec": 604800,
                        "weight": 2
                    }
                ]
            }
        }
    ]`

type AutoGenerated []struct {
	PermName     string `json:"perm_name"`
	Parent       string `json:"parent"`
	RequiredAuth struct {
		Threshold int `json:"threshold"`
		Keys      []struct {
			Key    string `json:"key"`
			Weight int    `json:"weight"`
		} `json:"keys"`
		Accounts []struct {
			Permission struct {
				Actor      string `json:"actor"`
				Permission string `json:"permission"`
			} `json:"permission"`
			Weight int `json:"weight"`
		} `json:"accounts"`
		Waits []struct {
			WaitSec int `json:"wait_sec"`
			Weight  int `json:"weight"`
		} `json:"waits"`
	} `json:"required_auth"`
}

type AutoGenerated1 struct {
	PermName  string `json:"name"`
	Parent    string `json:"parent"`
	Threshold int    `json:"threshold"`
	Keys      []struct {
		Key    string `json:"key"`
		Weight int    `json:"weight"`
	} `json:"keys"`
	Accounts []struct {
		Permission struct {
			Actor      string `json:"account"`
			Permission string `json:"permission"`
		} `json:"permission"`
		Weight int `json:"weight"`
	} `json:"accounts"`
	Waits []struct {
		WaitSec int `json:"second"`
		Weight  int `json:"weight"`
	} `json:"waits"`
	LinkedAction []string `json:"linked_action"`
}

func CreateCmd(cmd byte, context []byte) []byte {
	msg := make([]byte, 0)
	// 包头
	msg = append(msg, 0x7f)
	msg = append(msg, cmd)

	msg = append(msg, context...)
	//校验和
	var check int
	for _, v := range msg {
		check += int(v)
	}
	check = check + 0x8f
	msg = append(msg, byte(check%(math.MaxUint8+1)))

	//包尾
	msg = append(msg, 0x8f)
	return msg
}

/**
 *
 * @Title: byteToInt
 * @Description: byte转int
 * @param @param higth
 * @param @param low
 * @param @return
 * @author hjl
 * @return int 返回类型
 * @throws
 */
func ByteToInt(high, low int) int {
	if low < 0 {
		low = low + 256
	}
	return high*256 + low
}

func main() {
	var xz = make(chan int)
	var out = make(chan int)
	go func() {
		for {
			select {
			case <-xz:
				fmt.Println("===算奖====")
			case <-out:
				fmt.Println("===退出====")
				goto E
			}
		}
		E:
		fmt.Println("===退出111====")
	}()
	time.Sleep(1 * time.Second)
	xz <- 0
	time.Sleep(1 * time.Second)
	xz <- 0
	time.Sleep(1 * time.Second)
	xz <- 0

	time.Sleep(2 * time.Second)
	out <- 1
	time.Sleep(5 * time.Second)
	xz <- 0
	//fmt.Print(ByteToInt(8, 8))
	return

	//心跳包
	//var xt = `{"errno":0,"errmsg":"","data":{"heart_beat":"2019-10-19 19:00:56"}}`
	//bytes, _ := rsa.Encrypt([]byte(xt))
	//
	//fmt.Println(base64.StdEncoding.EncodeToString(bytes))
	//cmds := CreateCmd(0x00, bytes)
	//fmt.Println(len(cmds))
	//fmt.Println(cmds)
	//v := cmds[1 : len(cmds)-1]
	//fmt.Println(v)
	//var check = 0
	//for _, x := range v {
	//	check += int(x)
	//}
	//check = check + 0x8f
	//fmt.Println(cmds[len(cmds)-1])
	//if cmds[len(cmds)-1] == byte(check%(math.MaxUint8+1)) {
	//	fmt.Println("ok")
	//}

	return
	//fmt.Printf("pid= %#v\n", os.Getpid())
	//var iphone I_Phone
	//iphone = new(NokiaPhone)
	//iphone.call()
	//iphone.sms()
	//
	//iphone = new(Iphone)
	//iphone.call()
	//iphone.sms()
	//
	//a := 100
	//b := 200
	//
	//fmt.Print("转换前", a, b)
	//fmt.Print("\n")
	//util.Swap(&a, &b)
	//fmt.Print("转换后", a, b)
	//fmt.Print("\n")
	//
	////定义一个切片
	////1,类型  2,长度  3,预计长度
	//c := make([]int, 3, 8) //数组不能用make
	//d := [3]int{}          //数组
	//e := []int{}           //没有限定长度的数组就是切片
	//util.PrintSlice(c)
	////util.PrintSlice(d);//报错  类型不一致
	//for i, v := range d {
	//	fmt.Printf("[%d->%d]\n", i, v)
	//}
	//util.PrintSlice(e)
	////map
	//f := map[string]string{"Youhao": "Lyl"} //一步到位初始化到赋值
	//
	//var f1 map[string]string     //定义
	//f1 = make(map[string]string) //初始化
	//f1["Youhao"] = "Lyl"         //赋值
	//
	////打印map
	//for v := range f {
	//	fmt.Printf("[%s => %s]\n", v, f[v])
	//}
	//for v := range f1 {
	//	fmt.Printf("[%s => %s]\n", v, f1[v])
	//}
	//
	//fmt.Printf("%#v\n", strings.HasPrefix(strings.ToUpper("eos:111就斤斤计较"), strings.ToUpper("EOS")))
	//
	//persons := Persons{
	//	{
	//		Name: "test1",
	//		Age:  20,
	//	},
	//	{
	//		Name: "test2",
	//		Age:  22,
	//	},
	//	{
	//		Name: "test3",
	//		Age:  21,
	//	},
	//}
	//for i := 0; i < 100000; i++ {
	//	persons = append(persons, Person{Name: "test" + strconv.Itoa(i), Age: int64(i)})
	//}
	//fmt.Println("排序前")
	//for _, person := range persons {
	//	fmt.Println(person.Name, ":", person.Age)
	//}
	//sort.Sort(persons)
	//fmt.Println("排序后")
	//for _, person := range persons {
	//	fmt.Println(person.Name, ":", person.Age)
	//}
	//fmt.Println(1547 % 120)
	//fmt.Println((1547 - 107) / 120)
	//
	//size := 1547
	//step := size / 10
	//tempSize := size - step*10
	//var temp []int
	//for i := tempSize; i < size; i += step {
	//	temp = append(temp, i)
	//}
	//if len(temp) > 0 {
	//	temp[len(temp)-1] = 888
	//}
	//fmt.Println(len(temp), temp)
	//for i := 0; i < 100000; i++ {
	//	fmt.Println(strings.Contains(strings.ToUpper("Eosx"), strings.ToUpper("eos")))
	//}
	//v := strings.Split("eosbetdice11:8871964,", ",")
	//fmt.Println(len(v[0]))
	//
	//WhiteList := "eosbetdice11:887196"
	//var isTotals = true
	//var total int64
	//if strings.Contains(WhiteList,":"){
	//	split := strings.Split(WhiteList, ":")
	//	if len(split) >= 2{
	//		if split[0] == "eosbetdice11" {
	//			x ,_:= strconv.Atoi(split[1])
	//			total = int64(x)
	//			isTotals = false
	//		}
	//	}
	//}
	//fmt.Println(total,isTotals,time.Now())
	//now := time.Now()
	//next := now.Add(time.Hour * 10)
	//fmt.Println("xxxxx:" ,next)
	//next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), 0, 0, 0, next.Location())
	//end  := time.Date(next.Year(), next.Month(), next.Day(), 23, 59, 59, 0, next.Location())
	//
	//fmt.Println(next.Format("2006-01-02 15:04:05"),end.Format("2006-01-02 15:04:05"))
	//jsonstr := `{"status":"executed","cpu_usage_us":2841,"net_usage_words":19,"trx":{"id":"685550f6d82f30265087df37eebb2ed92fcee3276b6fd6895f2296d5a6303858","signatures":["SIG_K1_K6yQ7sRh5hvDU9LjfKfPQv572Eb9tcWNmSvC7ezfosMf8uf7i7LhyaijLwdLWkHoKR6w9ZsFiuTEc2E3jiGwGMBn9XtDin"],"compression":"none","packed_context_free_data":"","context_free_data":[],"packed_trx":"2d3dc45bc1e603dab9a3000000000100a6823403ea3055000000572d3ccdcd01a09865f94e93876600000000a8ed32323aa09865f94e938766401dce8db90931556e0000000000000004454f5300000000196d61743a3333313939393a313a6a676e6a676e6a676e31323300","transaction":{"expiration":"2018-10-15T07:09:33","ref_block_num":59073,"ref_block_prefix":2746866179,"max_net_usage_words":0,"max_cpu_usage_ms":0,"delay_sec":0,"context_free_actions":[],"actions":[{"account":"eosio.token","name":"transfer","authorization":[{"actor":"gu3tanrtgqge","permission":"active"}],"data":{"from":"gu3tanrtgqge","to":"eosknightsio","quantity":"0.0110 EOS","memo":"mat:331999:1:jgnjgnjgn123"},"hex_data":"a09865f94e938766401dce8db90931556e0000000000000004454f5300000000196d61743a3333313939393a313a6a676e6a676e6a676e313233"}],"transaction_extensions":[]}}}`
	//fmt.Println(json.Unmarshal([]byte(jsonstr),&Person{}))
	//
	//matched, _ := regexp.MatchString(emailReg, "1111")
	//fmt.Println(matched)
	//
	//var xx map[string]interface{}
	//err := json.Unmarshal([]byte(jsonStr1),&xx)
	//fmt.Println(err)
	//fmt.Println(xx)
	//fmt.Println(json_permission)
	//var autoGenerated AutoGenerated
	//
	//json.Unmarshal([]byte(json_permission), &autoGenerated)
	//var autoGenerated1 []AutoGenerated1
	//for _, v := range autoGenerated {
	//	for _,vAccount := range v.RequiredAuth.Accounts {
	//		fmt.Println(vAccount)
	//	}
	//	autoGeneratedTemp := AutoGenerated1{
	//		PermName:     v.PermName,
	//		Parent:       v.Parent,
	//		Threshold:    v.RequiredAuth.Threshold,
	//		Keys:         v.RequiredAuth.Keys,
	//		Accounts:     v.RequiredAuth.Accounts,
	//		Waits:        v.RequiredAuth.Waits,
	//		LinkedAction: []string{"eosio@claimrewards", "eosio@claimrewards"},
	//	}
	//	autoGenerated1 = append(autoGenerated1, autoGeneratedTemp)
	//}
	//fmt.Println(json.Marshal(autoGenerated1))

	//

}
