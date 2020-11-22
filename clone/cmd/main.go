package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	ScrapeHtml()
	laststr:= "获取网页后的字节流"
	result := ConvertToString(laststr, "gbk", "utf-8")

	fmt.Println(result)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//resp, err := http.Get("http://www.zhuoyachina.com" + r.RequestURI)
		//if err != nil {
		//	fmt.Println("http.Get err=",err)
		//	return
		//}
		//
		//bytes, err1 := ioutil.ReadAll(resp.Body)
		//
		//if err1 != nil {
		//	fmt.Println("ioutil.ReadAll err=",err)
		//	return
		//}
		//fmt.Println(string(bytes))

		res, _ := http.Get("https://www.vimetco.com.vn"+ r.RequestURI)
		fmt.Println(r.RequestURI)
		//doc, err := goquery.NewDocumentFromReader(res.Body)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//fmt.Println("xxxx:",doc.Find("title").Text())
		truebody, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		//ft, _ := os.Create("get1.html")
		//ft.Write(truebody)

		// 修改字符集
		//dec := mahonia.NewDecoder("utf-8")
		//ret := dec.ConvertString(string(truebody))
		//fmt.Println("GBK to UTF-8: ", ret)
		//fmt.Println(ConvertToString(string(truebody), "gbk", "utf-8"))
		//respNewStr := strings.Replace(ret, `<meta charset="gb2312">`, `<meta charset="gbk">`, -1)
		var respNewStr string
		if strings.Contains(string(truebody),`charset="gb2312"`) {
			respNewStr = ConvertToString(string(truebody), "gbk", "utf-8")
		}else{
			respNewStr = string(truebody)
		}
		//respNewTStr := strings.Replace(respNewStr, `https://www.vimetco.com.vn/`, `/`, -1)
		respNewTStr := strings.Replace(respNewStr, `https://www.vimetco.com.vn/`, `/`, -1)

		w.Write([]byte(respNewTStr))
	})

	http.ListenAndServe("127.0.0.1:7777", nil)
}


func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

func ScrapeHtml() {
	resp, err := http.Get("http://www.zhuoyachina.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status coder error: %d %s", resp.StatusCode, resp.Status)
	}
	//io.Copy(os.Stdout, resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Find("title").Text())
}