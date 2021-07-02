package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	sjm := GetSjm(2120, 01, 03)
	start := time.Now().Nanosecond()
	i, i2, i3 := validationSjm(sjm)
	end := time.Now().Nanosecond()
	fmt.Println(end - start)
	fmt.Println("日期：", i, i2, i3)
}

func validationSjm(sjm int) (int, int, int) {
	var year int
	var moth int
	var day int
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
				//fmt.Println("=====:", sjZz, tempYear, j, i, sjm1, sjm)
				if sjm1 == sjm {
					year = tempYear
					moth = j
					day = i
					goto END
				}
			}
			currDay = 1
		}
	}
END:
	return year, moth, day
}

func GetSjm(year, moth, day int) int {
	//9999 随机的 内部定
	sj := fmt.Sprintf("%d%02d%02d%d", year, moth, day, 9999)
	sjZz, _ := strconv.Atoi(sj)
	r := rand.New(rand.NewSource(int64(sjZz)))
	//几位就几个9
	sjm := r.Intn(99999999)
	fmt.Println("随机码：", sjm)
	return sjm
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
