package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "net/smtp"
)

func main() {
    host := "smtp.exmail.qq.com"
    port := 465
    email := "发送邮箱"
    password := "密码"
    toEmail := "接收邮箱"

    header := make(map[string]string)
    header["From"] = "Dockboard " + "<" + email + ">"
    header["To"] = toEmail
    header["Subject"] = "Dockboard 帐户激活邮件"
    header["Content-Type"] = "text/html; charset=UTF-8"

    body := `
    <table border="0" width="100%" cellpadding="0" cellspacing="0" bgcolor="#E6E5E7">
        <tr>
            <td height="50"></td>
        </tr><!-- ======= end top header ======= -->
        <!-- ======= header ======= -->

        <tr>
            <td align="center">
                <table border="0" align="center" width="590" cellpadding="0" cellspacing="0" bgcolor="#FFFFFF" class="container590">
                    <tr>
                        <td align="center">
                            <table border="0" align="center" width="590" cellpadding="0" cellspacing="0" class="container590">
                                <tr>
                                    <td>
                                        <table border="0" align="left" cellpadding="0" cellspacing="0" width="190" bgcolor="#2C3E50" class="logo">
                                            <tr>
                                                <td height="25" style="font-size: 25px; line-height: 25px;">&nbsp;</td>
                                            </tr>

                                            <tr>
                                                <td align="center">
                                                    <table border="0" cellpadding="0" cellspacing="0">
                                                        <tr>
                                                            <td align="center"><a href="" style="display: block; border-style: none !important; border: 0 !important;"><img width="194" height="42" border="0" style="display: block; width: 194px; height: 42px;" src="http://docker.u.qiniudn.com/dockboard_logo_194_42_email.png" alt="dockboard.org" /></a></td>
                                                        </tr>
                                                    </table>
                                                </td>
                                            </tr>

                                            <tr>
                                                <td height="25" style="font-size: 25px; line-height: 25px;">&nbsp;</td>
                                            </tr>
                                        </table>

                                        <table border="0" align="left" cellpadding="0" cellspacing="0" class="hideforiphone">
                                            <tr>
                                                <td height="20" width="20" style="font-size: 20px; line-height: 20px;">&nbsp;</td>
                                            </tr>
                                        </table>

                                        <table border="0" align="right" cellpadding="0" cellspacing="0" class="date">
                                            <tr>
                                                <td align="center" valign="middle">
                                                    <table width="300" border="0" cellpadding="0" cellspacing="0" align="center" class="date-inside">
                                                        <tr>
                                                            <td height="30" style="font-size: 30px; line-height: 40px;">&nbsp;</td>
                                                        </tr>
                                                        <tr>
                                                            <td align="right" style="color: #959da6; font-size: 16px; font-weight: normal; font-family:'Source Sans Pro' Arial, sans-serif;">Dockboard Hub / Build / CI / CD&nbsp;&nbsp;&nbsp;</td>
                                                        </tr>
                                                        <tr>
                                                            <td height="30" style="font-size: 30px; line-height: 30px;">&nbsp;</td>
                                                        </tr>                                                        

                                                    </table>
                                                </td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
            </td>
        </tr><!-- ======= end header ======= -->

        <tr>
            <td height="30" style="font-size: 30px; line-height: 30px;">&nbsp;</td>
        </tr>
        <!-- ======= Main section ======= -->
        <!-- ======= CTA ======= -->

        <tr>
            <td align="center">
                <table border="0" align="center" width="590" cellpadding="0" cellspacing="0" bgcolor="#FFFFFF" style="border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;" class="container590">
                    <tr>
                        <td height="45" style="font-size: 45px; line-height: 45px;">&nbsp;</td>
                    </tr>

                    <tr>
                        <td align="center" style="color: #222222; font-size: 24px; font-family: 'Ubuntu', Arial, sans-serif; mso-line-height-rule: exactly;" class="cta-header">
                            <div>
                                请激活您的 Dockboard 帐户
                            </div>
                        </td>
                    </tr>

                    <tr>
                        <td height="25" style="font-size: 25px; line-height: 25px;">&nbsp;</td>
                    </tr>

                    <tr>
                        <td>
                            <table border="0" align="center" width="490" cellpadding="0" cellspacing="0" bgcolor="#FFFFFF" class="container580">
                                <tr>
                                    <td align="center" style="color: #adb3ba; font-size: 16px; font-weight: 300; font-family: 'Source Sans Pro', Arial, sans-serif; mso-line-height-rule: exactly; line-height: 24px;">
                                        <div style="line-height: 24px;">
                                            感谢你注册 Dockboard 帐户！请点击下面的激活按钮激活帐号并登录到控制台设置个人信息。
                                        </div>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>

                    <tr>
                        <td height="40" style="font-size: 40px; line-height: 40px;">&nbsp;</td>
                    </tr>

                    <tr>
                        <td align="center"><a href="" style="display: block; width: 140px; height: 40px; border-style: none !important; border: 0 !important;"><img width="140" height="40" border="0" style="display: block; width: 140px; height: 40px;" src="http://docker.u.qiniudn.com/email_active-btn.png" alt="激活帐号" /></a></td>
                    </tr>

                    <tr>
                        <td height="45" style="font-size: 45px; line-height: 45px;">&nbsp;</td>
                    </tr>
                </table>
            </td>
        </tr><!-- ======= end CTA ======= -->

        <tr>
            <td height="10" style="font-size: 10px; line-height: 10px;">&nbsp;</td>
        </tr>

        <!-- ======= footer ======= -->

        <tr>
            <td align="center">
                <table border="0" align="center" width="590" cellpadding="0" cellspacing="0" bgcolor="#FFFFFF" class="container590">
                    <tr>
                        <td align="center">
                            <table border="0" align="center" width="560" cellpadding="0" cellspacing="0" class="container580">
                                <tr>
                                    <td height="20" style="font-size: 20px; line-height: 20px;">&nbsp;</td>
                                </tr>

                                <tr>
                                    <td></td>

                                    <td align="center">
                                        <table border="0" align="left" cellpadding="0" cellspacing="0" style="border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;" class="container580">
                                            <tr>
                                                <td align="center" style="color: #adb3ba; font-size: 13px; font-family: 'Source Sans pro', Arial, sans-serif; line-height: 30px;"><span style="color: #222222;">dockboard</span> © Copyright 2014 . All Rights Reserved</td>
                                            </tr>
                                        </table>

                                        <table border="0" width="10" align="left" cellpadding="0" cellspacing="0" style="border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;">
                                            <tr>
                                                <td height="20" width="10" style="font-size: 20px; line-height: 20px;">&nbsp;</td>
                                            </tr>
                                        </table>

                                        <table border="0" align="right" cellpadding="0" cellspacing="0" style="border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;" class="container580">
                                            <tr>
                                                <td align="center" style="color: #adb3ba; font-size: 13px; font-family: 'Source Sans pro', Arial, sans-serif; line-height: 30px;"><a href="" style="color: #adb3ba; text-decoration: none;">Privacy Policy</a>&nbsp;&nbsp;<span style="font-weight: 700; color: #2c3e50;">/</span>&nbsp;&nbsp;<a href="" style="color: #adb3ba; text-decoration: none;">Terms of Use</a>&nbsp;&nbsp;<span style="font-weight: 700; color: #2c3e50;">/</span>&nbsp;&nbsp;<a href="" style="color: #adb3ba; text-decoration: none;">Contact</a></td>
                                            </tr>
                                        </table>
                                    </td>
                                </tr>

                                <tr>
                                    <td height="20" style="font-size: 20px; line-height: 20px;">&nbsp;</td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>

        <tr>
            <td height="30" style="font-size: 30px; line-height: 30px;">&nbsp;</td>
        </tr>

        <!-- ======= end footer ======= -->
    </table>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     
  `

    message := ""
    for k, v := range header {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    auth := smtp.PlainAuth(
        "",
        email,
        password,
        host,
    )

    err := SendMailUsingTLS(
        fmt.Sprintf("%s:%d", host, port),
        auth,
        email,
        []string{toEmail},
        []byte(message),
    )

    if err != nil {
        panic(err)
    }
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
    conn, err := tls.Dial("tcp", addr, nil)
    if err != nil {
        log.Println("Dialing Error:", err)
        return nil, err
    }
    //分解主机端口字符串
    host, _, _ := net.SplitHostPort(addr)
    return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
    to []string, msg []byte) (err error) {

    //create smtp client
    c, err := Dial(addr)
    if err != nil {
        log.Println("Create smpt client error:", err)
        return err
    }
    defer c.Close()

    if auth != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(auth); err != nil {
                log.Println("Error during AUTH", err)
                return err
            }
        }
    }

    if err = c.Mail(from); err != nil {
        return err
    }

    for _, addr := range to {
        if err = c.Rcpt(addr); err != nil {
            return err
        }
    }

    w, err := c.Data()
    if err != nil {
        return err
    }

    _, err = w.Write(msg)
    if err != nil {
        return err
    }

    err = w.Close()
    if err != nil {
        return err
    }

    return c.Quit()
}
