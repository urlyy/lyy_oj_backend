package main

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tjfoc/gmsm/sm3"
	"gopkg.in/gomail.v2"
)

func Test(t *testing.T) {
	// claims, _ := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsImV4cCI6MTcxMDQ5NDQ1MH0.prNGSaJuMy_T1Ox04fiaZ3WXKDNIzhPg0pz1Mlqs_vE")
	// fmt.Println(claims.UserID)
	// t.Logf("发送成功执行正确...")
}

func Test2(t *testing.T) {
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", "2213732736@qq.com")
	//接收人
	m.SetHeader("To", "17873729520@163.com")
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "urlyy")
	//主题
	m.SetHeader("Subject", "subject")
	//内容
	m.SetBody("text/html", "content")
	//附件
	//m.Attach("./myPic.png")
	// 第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "2213732736@qq.com", "rwrblztzikywdjgi")
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}

func TestDB(t *testing.T) {
	params := "host=192.168.88.132 port=5432 user=postgres password=root dbname=lyy_oj sslmode=disable"
	db := sqlx.MustConnect("postgres", params)
	var ids = []int{1, 10}

	rows, err := db.Query(`SELECT * FROM "user" WHERE is_deleted=? AND  id IN (?);`, false, ids)
	// query, args, _ := sqlx.In(`SELECT * FROM "user" WHERE is_deleted=? AND  id IN (?);`, false, ids)
	// fmt.Println(query, args)
	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	// query = db.Rebind(query)
	// fmt.Println(query)
	// rows, err := db.Query(query, args...)
	fmt.Println(rows, err)
}

func TestSM3(t *testing.T) {
	src := []byte("sm3是我国国产的哈希算法")
	hash := sm3.New()
	hash.Write(src)
	hashed := hash.Sum(nil)
	hashString := hex.EncodeToString(hashed)

	// 打印哈希值的十六进制字符串
	fmt.Printf("哈希结果为：%s\n", hashString)
}
