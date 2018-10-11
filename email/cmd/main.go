package main

import (
	"fmt"
	"yhgolang/email"
)

func main() {
	// 邮箱账号
	user := "yhlovelyl@vip.qq.com"
	//注意，此处为授权码、不是密码
	password := "kwkwywnymrklbjfi"
	//smtp地址及端口
	host := "smtp.qq.com:587"
	//接收者，内容可重复，邮箱之间用；隔开
	to := "yhlovelyl@vip.qq.com"
	//邮件主题
	subject := "测试通过golang发送邮件"
	//邮件内容
	text := `<p>Good day!</p>

			<p>Thanks a lot for contacting with us!</p>

			<p>We are currently pre-filtering all advertisers so please tell us more about your product and advertising demands:</p>

			<p>1. Your official siteURL:</p>

			<p>2. Contact name:</p>

			<p>3. Desired location of advertisement (i.e HomePage, AddressPage, etc):</p> 

			<p>4. The advertisement duration (Start/End Date):</p> 

			<p>5. The advert text/banner that you want displayed (less than 22 words):</p> 

			<p>6. Desired budget:</p>

			<p>7. Desired mode of payment (EOS preferred)</p>

			<p>You can directly reply this email for further discussion.</p>

			<p>-----------------</p>

			<p>您好！</p>
			<p>非常感谢您与我们联系！预定特定广告版面，请提前告诉我们您的产品信息广告需及求：</p>
			<p>1.您的官方网站URL：</p>
			<p>2.联系人姓名：</p>
			<p>3.广告的期望版面（即主页，搜索结果页等）：</p>
			<p>4.广告持续时间（开始/结束日期）：</p>
			<p>5.您想要显示的广告文字/横幅（少于22个字）：</p>
			<p>6.您的广告预算：</p>
			<p>7.期望的付款方式（EOS最佳）</p>
			<p>请直接回复此邮件进行下一步合作探讨。谢谢。</p>



			<p>--</p>

			<p>EOSPark Team</p>

	`
	body := `
    <html>
    <body>
    ` + text + `
    </body>
    </html>
    `
	fmt.Println("send email")
	//执行逻辑函数
	err := email.SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("发送邮件失败!")
		fmt.Println(err)
	} else {
		fmt.Println("发送邮件成功!")
	}
}
