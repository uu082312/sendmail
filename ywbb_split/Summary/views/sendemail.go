package views

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"mime"
	"regexp"
	"strings"
)

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func Send(msg, fromaddr, pwd, header string, toaddr, ccs, EnclosurePaths []string) error {
	var to_users []string

	for _, to := range toaddr {
		//fmt.Println(cc)
		if VerifyEmailFormat(to) {
			to_users = append(to_users, to)
		}

	}

	m := gomail.NewMessage()
	//m.SetHeader("From", m.FormatAddress("604134049@qq.com", "消息来自大帅比"))
	m.SetHeader("From", m.FormatAddress(fromaddr, "消息来自"+fromaddr))
	m.SetHeader("To", to_users...) //352269014@qq.com
	var ccs_users []string
	for _, cc := range ccs {
		//fmt.Println(cc)
		if VerifyEmailFormat(cc) {
			ccs_users = append(ccs_users, cc)
		}

	}
	m.SetHeader("Cc", ccs_users...)
	m.SetHeader("Subject", header)
	m.SetBody("text/html", msg) // 正文
	var name string
	for _, path := range EnclosurePaths {
		if len(strings.TrimSpace(path)) == 0 {
			continue
		}

		file_name := strings.Split(path, "/")
		file_name = strings.Split(file_name[len(file_name)-1], "\\")
		name = file_name[len(file_name)-1]

		m.Attach(path, gomail.SetHeader(map[string][]string{
			"Content-Disposition": []string{
				fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", name)),
			},
		}))

	}
	//m.Attach(filepath)
	//m.Attach(filepath)
	//d := gomail.NewDialer("smtp.qq.com", 25, "604134049@qq.com", "eqfalisxhxacbfie")

	d := gomail.NewDialer("smtphz.qiye.163.com", 465, fromaddr, pwd)
	err := d.DialAndSend(m)
	return err

	//d := gomail.NewDialer("smtphz.qiye.163.com", 465, fromaddr, "MUaZDrHgX5BqESjq")
	//d := gomail.NewDialer("smtp.163.com", 994, "scgopy@163.com", "LMHPRSOXTPLDXEEL")
	//err := d.DialAndSend(m)
	//return err
}
