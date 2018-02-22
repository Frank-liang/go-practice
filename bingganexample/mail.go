package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"text/template"
)

var (
	tpl = template.Must(template.New("alarm").Parse(`From: <bingan@reboot.com>
To: <{{.To}}>
Subject: {{.Subject}}
Content-Type: text/html;charset=utf8

{{.Body}}
`))
)

type mailDesc struct {
	To      string
	Subject string
	Body    string
}

func sendmail(m *mailDesc) error {
	buf := new(bytes.Buffer)
	tpl.Execute(buf, m)

	cmd := exec.Command("sendmail", "-t")
	cmd.Stdin = buf
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s:%s", err, out)
	}

	return nil
}

func main() {
	m := &mailDesc{
		To:      "y1053419035@qq.com",
		Subject: "炸弹",
		Body:    "<h1>呵呵</h1>",
	}
	log.Print(sendmail(m))
}
