package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/tjfoc/gmsm/sm3"
	"gopkg.in/gomail.v2"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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

	_, err := db.Exec(`
		INSERT INTO 
		submission(problem_id,domain_id,from_type,user_id,submit_time,last_judge_time,lang,code,from_id,status,max_time,max_memory,pass_percent)
		VALUES($1,$2,$3,$4,$5,$5,$6,$7,$8,$9,$10,$10,$11)
		RETURNING id
		`, 1, 1, "qwer", 1, time.Now(), "java", "codecode", 1, 6, 0, 0.0,
	)
	if err != nil {
		fmt.Println(err)
	}
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

func TestDocker(t *testing.T) {
	cli, err := client.NewClientWithOpts(client.WithHost("tcp://192.168.88.132:2375"), client.WithVersion("1.44"))
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Printf("Container ID: %s\n", container.ID)
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Command: %s\n", container.Command)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Printf("Created: %d\n", container.Created)
		fmt.Printf("Ports: %+v\n", container.Ports)
		fmt.Printf("Labels: %+v\n", container.Labels)
		fmt.Printf("--------------------------------------------------\n")
	}
	// cbytes, _ := json.Marshal(containers)
	// fmt.Println(string(cbytes))
}

func TestRPC(t *testing.T) {
	// Set up a connection to the server.
	// conn, err := grpc.Dial("192.168.88.132:8800", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()
	// c := pb.NewGreeterClient(conn)
	// // Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.GetMessage())
}
