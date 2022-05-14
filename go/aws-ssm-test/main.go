package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("ap-northeast-1")},
		Profile: "default",
	})
	if err != nil {
		panic(err)
	}
	svc := ssm.New(sess)

	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: aws.String("/portfolio/dsn"),
		// `WithDecryption` を true にすれば、パラメータが暗号化している場合に復号して取得します
		// 平文のパラメータはそのまま取得するので、基本的には true でいいです
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", res)
	fmt.Printf("%#v\n", *res.Parameter.Value)
}
