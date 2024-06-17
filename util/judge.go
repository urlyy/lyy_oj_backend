package util

import (
	pb "backend/proto/judge"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addressIndex int

func init() {
	addressIndex = 0
}

func Judge(submissionID string, addressList []string, code string, inputList []string, expectList []string, compiler string, timeLimit uint64, memoryLimit uint64, isSpecial bool, specialCode string) (*pb.JudgeReply, error) {
	address := addressList[addressIndex]
	// ctx1, cel := context.WithTimeout(context.Background(), time.Duration(timeLimit)*time.Millisecond*2)
	// defer cel()
	// conn, err := grpc.DialContext(ctx1, address, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewJudgeClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	req := pb.JudgeRequest{
		SubmissionID: submissionID,
		Special:      isSpecial,
		Code:         code,
		InputList:    inputList,
		ExpectList:   expectList,
		MemoryLimit:  memoryLimit,
		TimeLimit:    timeLimit,
		Compiler:     compiler,
		SpecialCode:  specialCode,
	}
	fmt.Println("in ", time.Now())
	r, err := c.SubmitJudge(ctx, &req)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
		return nil, err
	}
	return r, nil
}
