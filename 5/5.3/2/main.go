package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/smtp"
	"net/textproto"

	"golang.org/x/text/encoding/japanese"
)

func split(s string, n int) []string {
	ss := []string{}
	rs := []rune(s)
	l := len(rs)
	sub := ""
	for i, r := range rs {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			ss = append(ss, sub)
			sub = ""
		} else if i+1 == l {
			ss = append(ss, sub)
		}
	}
	return ss
}

func header(s string) string {
	var buf bytes.Buffer
	iso2022jp := japanese.ISO2022JP.NewEncoder()
	for i, line := range split(s, 13) {
		if i != 0 {
			buf.WriteString("\r\n")
		}
		buf.WriteString("=?iso-2022-jp?B?")
		input := []byte(line)
		b, err := iso2022jp.Bytes(input)
		if err != nil {
			b = input
		}
		buf.WriteString(base64.StdEncoding.EncodeToString(b))
		buf.WriteString("?=")

	}
	return buf.String()
}

func encodeBody(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func main() {
	smtpHost := "my-mail-server:25"
	smtpAuth := smtp.PlainAuth(
		"exampl.com",
		"example-user",
		"example-password",
		"auth.example.com",
	)
	var baf bytes.Buffer
	mw := multipart.NewWriter(&baf)

	mh := textproto.MIMEHeader{}
	mh.Set("Content-Type", "text/plain; charset=iso-2022-jp")
	mh.Set("Content-Transfer-Encoding", "base64")
	part, err := mw.CreatePart(mh)
	if err != nil {
		log.Fatal(err)
	}
	_, err = part.Write(encodeBody([]byte("メール本文")))
	if err != nil {
		log.Fatal(err)
	}
	mw.Close()

	boundary := mw.Boundary()
	var body bytes.Buffer
	body.WriteString(fmt.Sprintf("From: %s <%s>\n", header("送信太郎"), "taro@example.com"))
	body.WriteString(fmt.Sprintf("To: %s <%s>\n", header("宛先花子"), "hanako@example.com"))
	body.WriteString(fmt.Sprintf("Content-type: multipart/alternative; boundary=%s\n", boundary))
	body.WriteString(fmt.Sprintf("Subject: %s\n", header("メールのタイトル")))
	body.WriteString("\n")
	io.Copy(&body, &baf)

	err = smtp.SendMail(
		smtpHost,
		smtpAuth,
		"taro@example.com",
		[]string{"hanako@example.com"},
		body.Bytes())

	if err != nil {
		log.Fatal(err)
	}
}
