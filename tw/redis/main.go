package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/url"
	"time"
)

//r-wz9vn4znb60yxsqjtspd.redis.rds.aliyuncs.com
//r-wz9sf76l71hbu12bmwpd.redis.rds.aliyuncs.com
var ctx = context.Background()

func main() {

	params := url.Values{}
	params.Add("name", "@Rajeev")
	params.Add("phone", "+919999999999")
	fmt.Println(params.Encode())
	enEscapeUrl, _ := url.QueryUnescape(params.Encode())
	fmt.Println("解码:",enEscapeUrl)
	return

	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	q := u.Query()
	fmt.Println(q.Get("q"))
	u.Scheme = "https"
	u.Host = "google.com"
	q = u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
	return
	rdb := redis.NewClient(&redis.Options{
		Addr:         "r-wz9sf76l71hbu12bmwpd.redis.rds.aliyuncs.com:6379",
		Password:     "uJREJW9DNIk2H3I96ayz", // no password set
		DB:           1,                      // use default DB
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	})
	rdb1 := redis.NewClient(&redis.Options{
		Addr:         "r-wz9vn4znb60yxsqjtspd.redis.rds.aliyuncs.com:6379",
		Password:     "uJREJW9DNIk2H3I96ayz", // no password set
		DB:           1,                      // use default DB
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	})

	//mobile:labelCache:9*
	//2130525
	//0* 520153 ok
	//1*  1438281 ok
	// 2* 18489 ok
	// 3* 12293 ok
	// 4* 15002 ok
	// 5* 11736 ok
	// 6* 12497 ok
	//7* 6200 ok
	//8* 33313 ok
	// 9* 1286  ok
	//+86* 61272 ok
	// 86* 2 ok

	//val2, err := rdb.Get(ctx, "mobile:labelCache:+8601064571919").Result()
	//mobile:hCode: 123952
	//val2, err := rdb.Do(ctx, "keys", "mobile:labelCache:9*").Result()

	//mobile:hCode:

	//1* 388112

	//mobile:idCardLocation:110000
	val2, err := rdb.Do(ctx, "keys", "mobile:labelCache:0*").Result()
	//fmt.Println(401541 - (388112 + 607))
	fmt.Println(len(val2.([]interface{})))
	//return
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println("key2 does not exist", err)
	} else {
		for i, key := range val2.([]interface{}) {
			if key == nil || key.(string) == "" {
				continue
			}
			val := rdb.Get(ctx, key.(string)).Val()
			if val != "" {
				_, err = rdb1.Get(ctx, key.(string)).Result()
				if err == redis.Nil {
					r, err := rdb1.Do(ctx, "set", key.(string), val).Result()
					if err != nil {
						fmt.Println("写err：", err)
					} else {
						fmt.Println("写入key：", key.(string), r, i)
					}
				}

			}

		}
	}

}
