package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	// Authenticate with AWS
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		panic(err)
	}

	ks := kms.New(sess, aws.NewConfig().WithRegion("us-east-1"))

	listOutput, err := ks.ListKeys(&kms.ListKeysInput{})
	if err != nil {
		panic(err)
	}

	keyID := listOutput.Keys[len(listOutput.Keys)-1].KeyId

	// mnemonic := "hollow luxury rely minimum when shiver clarify galaxy prosper film float question"

	// encryptOutput, err := ks.Encrypt(&kms.EncryptInput{
	// 	KeyId:     keyID,
	// 	Plaintext: []byte(mnemonic),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// encrypted := encryptOutput.CiphertextBlob
	// fmt.Println(encrypted)

	// encodedData := base64.StdEncoding.EncodeToString(encrypted)
	// fmt.Println(encodedData)

	en := os.Getenv("ENCRYPTED_MNEMONIC")

	// fmt.Println(en)

	decodedData, err := base64.StdEncoding.DecodeString(en)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	decryptedOutput, err := ks.Decrypt(&kms.DecryptInput{
		KeyId:          keyID,
		CiphertextBlob: decodedData,
	})
	if err != nil {
		panic(err)
	}

	decr := decryptedOutput.Plaintext
	fmt.Println(string(decr))
}
