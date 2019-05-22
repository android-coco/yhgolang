## 闭包-defer-painc-recover()
##  闭包

```go
//返回值为函数  闭包
func closure(x int) func(int) int {
    fmt.Printf("%p\n", &x)
    return func(y int) int {
        fmt.Printf("%p\n", &x)
        return x + y
    }
}
```

## painc  recover()

```go
func main() {
    J()
    K()
    L()
}
func J() {
    fmt.Println("Func J")
}
func K() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recover in K")
        }
    }()
    panic("Panic in K")//相当于抛出异常
}
func L() {
    fmt.Println("Func C")
}
```

## defer 用defer定义的函数 后定义的先执行 和栈一样 先进后出


```go
//分析运行结果
func main() {
    var fs = [4]func(){}
    for i := 0; i < 4; i++ {
        defer fmt.Println("defer i = ", i)                           //3,2,1,0
        defer func() { fmt.Println("defer_closure1 i=", i) }()       //4,4,4,4
        defer func(i int) { fmt.Println("defer_closure2 i=", i) }(i) //3,2,1,0
        fs[i] = func() {
            fmt.Println("closure i = ", i) //4,4,4,4
        }
    }

    for _, f := range fs {
        f()
    }
}
```

## _import下划线的作用
在Golang里，import的作用是导入其他package，但是今天在看beego框架时看到了import 下划线，不知其意，故百度而解之。

　　import 下划线（如：import _ hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。
【实列】
 目录结构：
  GOPATH
      --bin
      --pkg
      --src
            --yhgolang
               --org.yh
                 --uitl
                     importDemo.go 
            --studyGo
              append.go
append.go
```go
package main

import (
    _ "yhgolang/org.yh/util"
    "fmt"
)

func main() {

    //util.Test()  编译报错，说：undefined: util in util.Test

    /**
            append主要用于给某个切片（slice）追加元素
            如果该切片存储空间（cap）足够，就直接追加，长度（len）变长；
            如果空间不足，就会重新开辟内存，并将之前的元素和新的元素一同拷贝进去
                第一个参数为切片，后面是该切片存储元素类型的可变参数
        基础用法：
    */
    slice := append([]int{1, 2, 3}, 4, 5, 6)
    fmt.Println(slice) //[1 2 3 4 5 6]
    //第二个参数也可以直接写另一个切片，将它里面所有元素拷贝追加到第一个切片后面。
    //要注意的是，这种用法函数的参数只能接收两个slice，并且末尾要加三个点
    slice1 := append([]int{1, 2, 3}, []int{4, 5, 6}...)
    fmt.Println(slice1) //[1 2 3 4 5 6]
    //还有种特殊用法，将字符串当作[]byte类型作为第二个参数传入
    //append函数返回值必须有变量接收，不然编译器会报错
    bytes := append([]byte("hello"), " world"...)
    fmt.Println(string(bytes))
}
输出结果：util-init() fadsfadsf.
```

```go
package util

import (
    "fmt"
)

func Test() int {
    return 1
}

func init() {
    fmt.Println("fadsfadsf")
}

```

```go
//language code
const (
	AF     = "af"     //Afrikaans	南非语
	AF_ZA  = "af-ZA"  //Afrikaans (South Africa)	南非语
	AR     = "ar"     //Arabic	阿拉伯语
	AR_AE  = "ar-AE"  //Arabic (U.A.E.)	阿拉伯语(阿联酋)
	AR_BH  = "ar-BH"  //Arabic (Bahrain)	阿拉伯语(巴林)
	AR_DG  = "ar-DZ"  //Arabic (Algeria)	阿拉伯语(阿尔及利亚)
	AR_EG  = "ar-EG"  //Arabic (Egypt)	阿拉伯语(埃及)
	AR_IQ  = "ar-IQ"  //Arabic (Iraq)	阿拉伯语(伊拉克)
	AR_JO  = "ar-JO"  //Arabic (Jordan)	阿拉伯语(约旦)
	AR_KW  = "ar-KW"  //Arabic (Kuwait)	阿拉伯语(科威特)
	AR_LB  = "ar-LB"  //Arabic (Lebanon)	阿拉伯语(黎巴嫩)
	AR_LY  = "ar-LY"  //Arabic (Libya)	阿拉伯语(利比亚)
	AR_MA  = "ar-MA"  //Arabic (Morocco)	阿拉伯语(摩洛哥)
	AR_OM  = "ar-OM"  //Arabic (Oman)	阿拉伯语(阿曼)
	AR_QA  = "ar-QA"  //Arabic (Qatar)	阿拉伯语(卡塔尔)
	AR_SA  = "ar-SA"  //Arabic (Saudi Arabia)	阿拉伯语(沙特阿拉伯)
	AR_SY  = "ar-SY"  //Arabic (Syria)	阿拉伯语(叙利亚)
	AR_TN  = "ar-TN"  //Arabic (Tunisia)	阿拉伯语(突尼斯)
	AR_YE  = "ar-YE"  //Arabic (Yemen)	阿拉伯语(也门)
	AZ     = "az"     //Azeri (Latin)	阿塞拜疆语
	AZ_AZ  = "az-AZ"  //Azeri (Cyrillic) (Azerbaijan)	阿塞拜疆语(西里尔文)
	BE     = "be"     //Belarusian	比利时语
	BE_BY  = "be-BY"  //Belarusian (Belarus)	比利时语
	BG     = "bg"     //Bulgarian	保加利亚语
	BG_BG  = "bg-BG"  //Bulgarian (Bulgaria)	保加利亚语
	BS_BA  = "bs-BA"  //Bosnian (Bosnia and Herzegovina)	波斯尼亚语(拉丁文，波斯尼亚和黑塞哥维那)
	CA     = "ca"     //Catalan	加泰隆语
	CA_ES  = "ca-ES"  //Catalan (Spain)	加泰隆语
	CS     = "cs"     //Czech	捷克语
	CS_CZ  = "cs-CZ"  //Czech (Czech Republic)	捷克语
	CY     = "cy"     //Welsh	威尔士语
	CY_GB  = "cy-GB"  //Welsh (United Kingdom)	威尔士语
	DA     = "da"     //Danish	丹麦语
	DA_DK  = "da-DK"  //Danish (Denmark)	丹麦语
	DE     = "de"     //German	德语
	DE_AT  = "de-AT"  //German (Austria)	德语(奥地利)
	DE_CH  = "de-CH"  //German (Switzerland)	德语(瑞士)
	DE_DE  = "de-DE"  //German (Germany)	德语(德国)
	DE_LI  = "de-LI"  //German (Liechtenstein)	德语(列支敦士登)
	DE_LU  = "de-LU"  //German (Luxembourg)	德语(卢森堡)
	DV     = "dv"     //Divehi	第维埃语
	DV_MV  = "dv-MV"  //Divehi (Maldives)	第维埃语
	EL     = "el"     //Greek	希腊语
	EL_GR  = "el-GR"  //Greek (Greece)	希腊语
	EN     = "en"     //English	英语
	EN_AU  = "en-AU"  //English (Australia)	英语(澳大利亚)
	EN_BZ  = "en-BZ"  //English (Belize)	英语(伯利兹)
	EN_CA  = "en-CA"  //English (Canada)	英语(加拿大)
	EN_CB  = "en-CB"  //English (Caribbean)	英语(加勒比海)
	EN_GB  = "en-GB"  //English (United Kingdom)	英语(英国)
	EN_IE  = "en-IE"  //English (Ireland)	英语(爱尔兰)
	EN_IN  = "en-IN"  //English (Indian)	英语(印度)
	EN_JM  = "en-JM"  //English (Jamaica)	英语(牙买加)
	EN_NZ  = "en-NZ"  //English (New Zealand)	英语(新西兰)
	EN_PH  = "en-PH"  //English (Republic of the Philippines)	英语(菲律宾)
	EN_TT  = "en-TT"  //English (Trinidad and Tobago)	英语(特立尼达)
	EN_US  = "en-US"  //English (United States)	英语(美国)
	EN_ZA  = "en-ZA"  //English (South Africa)	英语(南非)
	EN_ZW  = "en-ZW"  //English (Zimbabwe)	英语(津巴布韦)
	EO     = "eo"     //Esperanto	世界语
	ES     = "es"     //Spanish	西班牙语
	ES_AR  = "es-AR"  //Spanish (Argentina)	西班牙语(阿根廷)
	ES_BO  = "es-BO"  //Spanish (Bolivia)	西班牙语(玻利维亚)
	ES_CL  = "es-CL"  //Spanish (Chile)	西班牙语(智利)
	ES_CO  = "es-CO"  //Spanish (Colombia)	西班牙语(哥伦比亚)
	ES_CR  = "es-CR"  //Spanish (Costa Rica)	西班牙语(哥斯达黎加)
	ES_DO  = "es-DO"  //Spanish (Dominican Republic)	西班牙语(多米尼加共和国)
	ES_EC  = "es-EC"  //Spanish (Ecuador)	西班牙语(厄瓜多尔)
	ES_ES  = "es-ES"  //Spanish ( Spain )    西班牙语(国际)
	ES_GT  = "es-GT"  //Spanish (Guatemala)    西班牙语(危地马拉)
	ES_HN  = "es-HN"  //Spanish (Honduras)    西班牙语(洪都拉斯)
	ES_MX  = "es-MX"  //Spanish (Mexico)    西班牙语(墨西哥)
	ES_NI  = "es-NI"  //Spanish (Nicaragua)    西班牙语(尼加拉瓜)
	ES_PA  = "es-PA"  //Spanish (Panama)    西班牙语(巴拿马)
	ES_PE  = "es-PE"  //Spanish (Peru)    西班牙语(秘鲁)
	ES_PR  = "es-PR"  //Spanish (Puerto Rico)    西班牙语(波多黎各(美))
	ES_PY  = "es-PY"  //Spanish (Paraguay)    西班牙语(巴拉圭)
	ES_SV  = "es-SV"  //Spanish (El Salvador)    西班牙语(萨尔瓦多)
	ES_UY  = "es-UY"  //Spanish (Uruguay)    西班牙语(乌拉圭)
	ES_VE  = "es-VE"  //Spanish (Venezuela)    西班牙语(委内瑞拉)
	ET     = "et"     //Estonian    爱沙尼亚语
	ET_EE  = "et-EE"  // Estonian (Estonia)    爱沙尼亚语
	EU     = "eu"     //Basque    巴士克语
	EU_ES  = "eu-ES"  // Basque (Spain)    巴士克语
	FA     = "fa"     //Farsi    法斯语
	FA_IR  = "fa-IR"  // Farsi (Iran)    法斯语
	FI     = "fi"     //Finnish    芬兰语
	FI_FI  = "fi-FI"  // Finnish (Finland)    芬兰语
	FO     = "fo"     //Faroese    法罗语
	FO_FO  = "fo-FO"  // Faroese (Faroe Islands)    法罗语
	FR     = "fr"     //French    法语
	FR_BE  = "fr-BE"  // French (Belgium)    法语(比利时)
	FR_CA  = "fr-CA"  // French (Canada)    法语(加拿大)
	FR_CH  = "fr-CH"  // French (Switzerland)    法语(瑞士)
	FR_FR  = "fr-FR"  //French (France)    法语(法国)
	FR_LU  = "fr-LU"  //French (Luxembourg)    法语(卢森堡)
	FR_MC  = "fr-MC"  // French (Principality of Monaco)    法语(摩纳哥)
	GL     = "gl"     //Galician    加里西亚语
	GL_ES  = "gl-ES"  // Galician (Spain)    加里西亚语
	GU     = "gu"     // Gujarati    古吉拉特语
	GU_IN  = "gu-IN"  // Gujarati (India)    古吉拉特语
	HE     = "he"     //Hebrew    希伯来语
	HE_IL  = "he-IL"  // Hebrew (Israel)    希伯来语
	HI     = "hi"     //Hindi    印地语
	HI_IN  = "hi-IN"  //Hindi (India)    印地语
	HR     = "hr"     //Croatian    克罗地亚语
	HR_BA  = "hr-BA"  //Croatian (Bosnia and Herzegovina)    克罗地亚语(波斯尼亚和黑塞哥维那)
	HR_HR  = "hr-HR"  // Croatian (Croatia)    克罗地亚语
	HU     = "hu"     //Hungarian    匈牙利语
	HU_HU  = "hu-HU"  //Hungarian (Hungary)    匈牙利语
	HY     = "hy"     //Armenian    亚美尼亚语
	HY_AM  = "hy-AM"  //Armenian (Armenia)    亚美尼亚语
	ID     = "id"     //Indonesian    印度尼西亚语
	ID_ID  = "id-ID"  //Indonesian (Indonesia)    印度尼西亚语
	IS     = "is"     //Icelandic    冰岛语
	IS_IS  = "is-IS"  //Icelandic (Iceland)    冰岛语
	IT     = "it"     //Italian    意大利语
	IT_CH  = "it-CH"  //Italian (Switzerland)    意大利语(瑞士)
	IT_IT  = "it-IT"  //Italian (Italy)    意大利语(意大利)
	JA     = "ja"     // Japanese    日语
	JA_JP  = "ja-JP"  //Japanese (Japan)    日语
	KA     = "ka"     // Georgian    格鲁吉亚语
	KA_GE  = "ka-GE"  //Georgian (Georgia)    格鲁吉亚语
	KK     = "kk"     //Kazakh    哈萨克语
	KK_KZ  = "kk-KZ"  // Kazakh (Kazakhstan)    哈萨克语
	KN     = "kn"     //Kannada    卡纳拉语
	KN_IN  = "kn-IN"  //Kannada (India)    卡纳拉语
	KO     = "ko"     // Korean    朝鲜语
	KO_KR  = "ko-KR"  // Korean (Korea)    朝鲜语
	KOK    = "kok"    // Konkani    孔卡尼语
	KOK_IN = "kok-IN" // Konkani (India)    孔卡尼语
	KY     = "ky"     //Kyrgyz    吉尔吉斯语
	KY_KG  = "ky-KG"  //Kyrgyz (Kyrgyzstan)    吉尔吉斯语(西里尔文)
	LT     = "lt"     //Lithuanian    立陶宛语
	LT_LT  = "lt-LT"  //Lithuanian (Lithuania)    立陶宛语
	LV     = "lv"     //Latvian    拉脱维亚语
	LV_LV  = "lv-LV"  //Latvian (Latvia)    拉脱维亚语
	MI     = "mi"     //Maori    毛利语
	MI_NZ  = "mi-NZ"  // Maori (New Zealand)    毛利语
	MK     = "mk"     //FYRO Macedonian    马其顿语
	MK_MK  = "mk-MK"  //  FYRO Macedonian (Former Yugoslav Republic of Macedonia)    马其顿语(FYROM)
	MN     = "mn"     //Mongolian    蒙古语
	MN_MN  = "mn-MN"  // Mongolian (Mongolia)    蒙古语(西里尔文)
	MR     = "mr"     //Marathi    马拉地语
	MR_IN  = "mr-IN"  //Marathi (India)    马拉地语
	MS     = "ms"     //Malay    马来语
	MS_BN  = "ms-BN"  // Malay (Brunei Darussalam)    马来语(文莱达鲁萨兰)
	MS_MY  = "ms-MY"  // Malay (Malaysia)    马来语(马来西亚)
	MT     = "mt"     // Maltese    马耳他语
	MT_MT  = "mt-MT"  // Maltese (Malta)    马耳他语
	NB     = "nb"     //Norwegian (Bokm?l)    挪威语(伯克梅尔)
	NB_NO  = "nb-NO"  // Norwegian (Bokm?l) (Norway)    挪威语(伯克梅尔)(挪威)
	NL     = "nl"     //Dutch    荷兰语
	NL_BE  = "nl-BE"  // Dutch (Belgium)    荷兰语(比利时)
	NL_NL  = "nl-NL"  // Dutch (Netherlands)    荷兰语(荷兰)
	NN_NO  = "nn-NO"  // Norwegian (Nynorsk) (Norway)    挪威语(尼诺斯克)(挪威)
	NS     = "ns"     //Northern Sotho    北梭托语
	NS_ZA  = "ns-ZA"  // Northern Sotho (South Africa)    北梭托语
	PA     = "pa"     //Punjabi    旁遮普语
	PA_IN  = "pa-IN"  //Punjabi (India)    旁遮普语
	PL     = "pl"     // Polish    波兰语
	PL_PL  = "pl-PL"  // Polish (Poland)    波兰语
	PT     = "pt"     //Portuguese    葡萄牙语
	PT_BR  = "pt-BR"  // Portuguese (Brazil)    葡萄牙语(巴西)
	PT_PT  = "pt-PT"  // Portuguese (Portugal)    葡萄牙语(葡萄牙)
	QU     = "qu"     //Quechua    克丘亚语
	QU_BO  = "qu-BO"  // Quechua (Bolivia)    克丘亚语(玻利维亚)
	QU_EC  = "qu-EC"  // Quechua (Ecuador)    克丘亚语(厄瓜多尔)
	QU_PE  = "qu-PE"  // Quechua (Peru)    克丘亚语(秘鲁)
	RO     = "ro"     //Romanian    罗马尼亚语
	RO_RO  = "ro-RO"  //Romanian (Romania)    罗马尼亚语
	RU     = "ru"     //Russian    俄语
	RU_RU  = "ru-RU"  // Russian (Russia)    俄语
	SA     = "sa"     // Sanskrit    梵文
	SA_IN  = "sa-IN"  //Sanskrit (India)    梵文
	SE     = "se"     // Sami (Northern)    北萨摩斯语
	SE_FI  = "se-FI"  // Sami (Inari) (Finland)    伊那里萨摩斯语(芬兰)
	SE_NO  = "se-NO"  // Sami (Southern) (Norway)    南萨摩斯语(挪威)
	SE_SE  = "se-SE"  // Sami (Southern) (Sweden)    南萨摩斯语(瑞典)
	SK     = "sk"     //Slovak    斯洛伐克语
	SK_SK  = "sk-SK"  //Slovak (Slovakia)    斯洛伐克语
	SL     = "sl"     //Slovenian    斯洛文尼亚语
	SL_SI  = "sl-SI"  // Slovenian (Slovenia)    斯洛文尼亚语
	SQ     = "sq"     // Albanian    阿尔巴尼亚语
	SQ_AL  = "sq-AL"  //Albanian (Albania)    阿尔巴尼亚语
	SR_BA  = "sr-BA"  // Serbian (Cyrillic) (Bosnia and Herzegovina)    塞尔维亚语(西里尔文，波斯尼亚和黑塞哥维那)
	SR_SP  = "sr-SP"  // Serbian (Cyrillic) (Serbia and Montenegro)    塞尔维亚(西里尔文)
	SV     = "sv"     // Swedish    瑞典语
	SV_FI  = "sv-FI"  // Swedish (Finland)    瑞典语(芬兰)
	SV_SE  = "sv-SE"  // Swedish (Sweden)    瑞典语
	SW     = "sw"     //Swahili    斯瓦希里语
	SW_KE  = "sw-KE"  //Swahili (Kenya)    斯瓦希里语
	SYR    = "syr"    //Syriac    叙利亚语
	SYR_SY = "syr-SY" // Syriac (Syria)    叙利亚语
	TA     = "ta"     // Tamil    泰米尔语
	TA_IN  = "ta-IN"  //  Tamil (India)    泰米尔语(印度)
	TA_LK  = "ta-LK"  //  Tamil (Sri Lankan)    泰米尔语(斯里兰卡)
	TE     = "te"     //Telugu    泰卢固语
	TE_IN  = "te-IN"  // Telugu (India)    泰卢固语
	TH     = "th"     //Thai    泰语
	TH_TH  = "th-TH"  // Thai (Thailand)    泰语
	TL     = "tl"     //Tagalog    塔加路语
	TL_PH  = "tl-PH"  // Tagalog (Philippines)    塔加路语(菲律宾)
	TN     = "tn"     //Tswana    茨瓦纳语
	TN_ZA  = "tn-ZA"  //Tswana (South Africa)    茨瓦纳语
	TR     = "tr"     //Turkish    土耳其语
	TR_TR  = "tr-TR"  // Turkish (Turkey)    土耳其语
	TS     = "ts"     //Tsonga    宗加语
	TT     = "tt"     //Tatar    鞑靼语
	TT_RU  = "tt-RU"  //Tatar (Russia)    鞑靼语
	UK     = "uk"     //Ukrainian    乌克兰语
	UK_UA  = "uk-UA"  //Ukrainian (Ukraine)    乌克兰语
	UR     = "ur"     //Urdu    乌都语
	UR_PK  = "ur-PK"  //Urdu (Islamic Republic of Pakistan)    乌都语
	UZ     = "uz"     //Uzbek (Latin)    乌兹别克语
	UZ_UZ  = "uz-UZ"  // Uzbek (Cyrillic) (Uzbekistan)    乌兹别克语(西里尔文)
	VI     = "vi"     //Vietnamese    越南语
	VI_VN  = "vi-VN"  // Vietnamese (Viet Nam)    越南语
	XH     = "xh"     //Xhosa    班图语
	XH_ZA  = "xh-ZA"  // Xhosa (South Africa)    班图语
	ZH     = "zh"     //Chinese	中文
	ZH_CN  = "zh-CN"  //Chinese (Mainland)	中文(简体)
	ZH_HK  = "zh-HK"  //Chinese (Hong Kong)	中文(香港)
	ZH_MO  = "zh-MO"  //Chinese (Macau)	中文(澳门)
	ZH_SG  = "zh-SG"  //Chinese (Singapore)	中文(新加坡)
	ZH_TW  = "zh-TW"  //Chinese (Taiwan)	中文(繁体)
	ZU     = "zu"     //Zulu	祖鲁语
	ZU_ZA  = "zu-ZA"  //Zulu (South Africa)	祖鲁语
)


```

