package util


import (
	"bytes"
	"html/template"
	// "fmt"
	"io"
	
	gomail "gopkg.in/gomail.v2"
)

const (
	from_mailaddress string = "hacku@na-cat.com"
)

func Send_mail( to_mailaddress string, user_name string , code string ) error {
	
	htmlTemplate, err := template.
		New( "mail.tmpl" ).
		ParseFiles( "/app/template/mail.tmpl" )

	if err != nil {
		return err
	}

	contentsValue := map[string]interface{}{
		"userName": user_name,
		"code": code,
	}

	htmlBytes, err := TemplateConvertToHtmlBytes( htmlTemplate, contentsValue )

	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader( "From", from_mailaddress )
	m.SetHeader( "To", to_mailaddress )
	m.SetHeader( "Subject", "会員登録手続きのお願い" )
	m.SetBody( "text/html", string( htmlBytes ) )

	d := gomail.Dialer{Host: "mailhog", Port: 1025}
    if err := d.DialAndSend(m); err != nil {
        return err
    }

	return err
}

func TemplateConvertToHtmlBytes(htmlTemplate *template.Template, contentsValues map[string]interface{}) ([]byte, error) {
	// 出力ファイルのバッファ
	buff := new( bytes.Buffer )
	fw := io.Writer( buff )

	// バッファに､HTMLファイルを書き込み
	err := htmlTemplate.ExecuteTemplate( fw, htmlTemplate.Name(), contentsValues )

	if err != nil {
		//log.Fatal( err )
		return nil, err
	}

	htmlBytes := buff.Bytes()

	// HTMLバイトを返す
	return htmlBytes, nil
}
