package main

/*
Map 是一种无序的键值对的集合。
Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。
不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type res struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
}

func get()  {
	var clinet http.Client
	//resp, err := clinet.Get("http://localhost:11000/api?module=block&action=get_block_detail&apikey=a9564ebc3289b7a14551baf8ad5ec60a&block_num=18272475")
	resp, err := clinet.Get("http://localhost:11000/api?module=transaction&action=get_transaction_detail_info&apikey=a9564ebc3289b7a14551baf8ad5ec60a&trx_id=2d91c592d66e429134d6e4f70fbd4f0e8e6f27462b53a226aecec07337ed1c02")
	if err != nil {
		//json.Unmarshal([]byte(resp.Body),nil)


	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf(string(body))
	var resr res
	unmarshal := json.Unmarshal([]byte(string(body)), &resr)
	fmt.Println(unmarshal,resr)
}

func main() {
	//var countryCapitalMap map[string]string
	///* 创建集合 */
	//countryCapitalMap = make(map[string]string)
	//
	///* map 插入 key-value 对，各个国家对应的首都 */
	//countryCapitalMap["France"] = "Paris"
	//countryCapitalMap["Italy"] = "Rome"
	//countryCapitalMap["Japan"] = "Tokyo"
	//countryCapitalMap["India"] = "New Delhi"
	//
	///* 使用 key 输出 map 值 */
	//for country := range countryCapitalMap {
	//	fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	//}
	///* 查看元素在集合中是否存在 */
	//captial, ok := countryCapitalMap["United States"]
	///* 如果 ok 是 true, 则存在，否则不存在 */
	//
	//if ok {
	//	fmt.Println("Capital of United States is", captial)
	//} else {
	//	fmt.Println("Capital of United States is not present")
	//}
	///* delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key */
	///* 创建 map */
	//countryCapitalMap1 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	//
	//fmt.Println("原始 map")
	//
	///* 打印 map */
	//for country := range countryCapitalMap1 {
	//	fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	//}
	//
	///* 删除元素 */
	//delete(countryCapitalMap1, "France")
	//fmt.Println("Entry for France is deleted")
	//
	//fmt.Println("删除元素后 map")
	//
	///* 打印 map */
	//for country := range countryCapitalMap1 {
	//	fmt.Println("Capital of", country, "is", countryCapitalMap1[country])
	//}
	//
	////map 中的map
	//kk := make(map[string]map[string]string)
	//
	//kk["a"] = make(map[string]string)
	//kk["a"]["a"] = "a"
	//
	//kk["b"] = make(map[string]string)
	//kk["b"]["b"] = "b"
	//
	//kk1, ok := kk["c"]["c"]
	//
	//if !ok {
	//	kk["c"] = make(map[string]string)
	//}
	//kk["c"]["c"] = "c"
	//fmt.Println(kk1, kk) //map[a:map[a:a]]
	//
	//ll := make([]map[int]string, 3)
	//for i := range ll {
	//	ll[i] = make(map[int]string)
	//	ll[i][i] = "ok"
	//	fmt.Println(ll[i])
	//}
	//fmt.Println(ll)
	//
	//ii := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	//ii1 := make([]int, len(ii))
	//i := 0
	//for k, _ := range ii {
	//	ii1[i] = k
	//	i++
	//}
	//sort.Ints(ii1) //map 的 key排序
	//fmt.Println(ii1)
	//
	////map de key value转换
	//oldMap := map[int]string{1: "a", 2: "b", 3: "c"}
	//newMap := receiveMap(oldMap)
	//fmt.Println(newMap) //  map[a:1 b:2 c:3]
	//jsonStr := `[{"actor":"xiozoneoslo1","permission":"active"},{"actor":"xiozoneoslo11","permission":"active11"}]`

	//var a2 []map[string]string
	//json.Unmarshal([]byte(jsonStr), &a2)
	//fmt.Println(a2)
	//uu := "1,2,3,4,5,6,7"
	//zz := "1,2,3,4,5,6,7"
	//for i, v := range uu {
	//A:
	//	fmt.Println("1")
	//	for j := range zz {
	//		if j == 10 {
	//			goto A
	//		}
	//	}
	//
	//	fmt.Println(i, v)
	//}

	//get()
	var datetime = time.Now()
	datetime.Format("2006-01-02 15:04:05")
	fmt.Println(nowTime())
	parse1, i2 := time.Parse("2006-01-02", "2018-10-01")
	before := parse1.Before(time.Now().UTC())
	fmt.Println(parse1,i2,before)

	sub := time.Now().Sub(parse1)
	fmt.Println(7*24)
	fmt.Println(int(sub.Hours()/24))

	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())
}

func nowTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

type A1 struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}

func receiveMap(oldMap map[int]string) map[string]int {
	newMap := make(map[string]int)
	for k, v := range oldMap {
		newMap[v] = k
	}
	return newMap
}
