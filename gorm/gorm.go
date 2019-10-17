package main

import (
	"encoding/json"
	"fmt"
	"github.com/BlockABC/eospark_backend/model"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"

	"yhgolang/gorm/db"
)

type Loss struct {
	StoreID  int64
	Dateline string
	Type     int64
}

func (*Loss) TableName() string {
	return "loss"
}

func main() {
	//bpJson, _ := getBpCountry("https://bp.fish")
	//fmt.Println(bpJson.Org.Location.Country, bpJson.Org.Branding.LogoSvg)

	//bytes, err := util.YHGet(bpJson.Org.Branding.LogoSvg)
	//fmt.Println(bytes, err)
	GetRankAndCountry()
}

//获取排名 更新历史排名和国家
func GetRankAndCountry() {
	db := db.GetDb()
	var bpInfo []model.BpInfo
	//查询当前排名
	err := db.Order("score" + " desc").Find(&bpInfo).Error

	if err != nil {
		fmt.Errorf("getRankAndCountry  get rank %v", err)
		return
	}
	//更新国家和排名
	for i, bp := range bpInfo {
		//更新排名
		e := db.Exec("UPDATE t_bp_info SET his_rank = ? WHERE account = ? ", int64(i+1), bp.Account).Error
		if e != nil {
			fmt.Errorf("getRankAndCountry  update his_rank %v", e)
			return
		}
	}
	UpdateCountry(db, bpInfo)
}

func UpdateCountry(db *gorm.DB, bpInfo []model.BpInfo) {
	for _, bp := range bpInfo {
		//db.LogMode(true)
		if bp.Website == "" {
			continue
		}
		//var bpTemp model.BpInfo
		if bp.GuoJia != "" {
			continue
		}
		//获取国家
		if country, e := getBpCountry(bp.Website); e.ErrCode == 0 {
			var logoUrl string
			var countryStr = country.Org.Location.Country
			if country.Org.Branding.LogoSvg != "" {
				logoUrl = country.Org.Branding.LogoSvg
			}
			if logoUrl == "" {
				logoUrl = country.Org.Branding.Logo256
			}
			if logoUrl == "" {
				logoUrl = country.Org.Branding.Logo1024
			}

			if bp.GuoJia == countryStr {
				continue
			}
			var e error
			if countryStr != "" && logoUrl != "" {
				e = db.Exec("UPDATE t_bp_info SET guojia = ?,logo_url = ? WHERE account = ?", countryStr, logoUrl, bp.Account).Error
			} else {
				if countryStr != "" {
					e = db.Exec("UPDATE t_bp_info SET guojia = ?  WHERE account = ?", countryStr, bp.Account).Error
				}
				if logoUrl != "" {
					e = db.Exec("UPDATE t_bp_info SET logo_url = ? WHERE account = ?", logoUrl, bp.Account).Error
				}
			}
			//db.LogMode(true)
			if e != nil {
				fmt.Errorf("getRankAndCountry  update his_rank %v", e)
				return
			}
		}
	}
}

type BpJson struct {
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

type Branding struct {
	Logo256  string `json:"logo_256"`
	Logo1024 string `json:"logo_1024"`
	LogoSvg  string `json:"logo_svg"`
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

type Error struct {
	ErrCode int64
	ErrMsg  error
}

//获取国家log
func getBpCountry(website string) (*BpJson, Error) {
	bpJson := &BpJson{}
	if website == "" {
		return nil, Error{103200, nil}
	}
	bpUrl := website + "/bp.json"
	resp, err := http.Get(bpUrl)
	if err != nil {
		fmt.Errorf("get_bp_json_str no response. addr:%s err: %s", bpUrl, err)
		return nil, Error{1, err}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Errorf("get_bp_json_str no response. addr:%s err: %s", bpUrl, err)
			return nil, Error{1, err}
		}
		if unmarshal := json.Unmarshal(all, bpJson); unmarshal != nil {
			return nil, Error{1, nil}
		}
	}
	return bpJson, Error{0, nil}
}
