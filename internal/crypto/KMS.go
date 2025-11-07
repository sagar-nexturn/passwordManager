package crypto

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type KMSCrypto struct {
	Client  *kms.Client
	KeyID   string
	Context context.Context
}

func NewKMSCrypto() *KMSCrypto {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	keyID := os.Getenv("AWS_KMS_KEY_ID")

	if region == "" || accessKey == "" || secretKey == "" || keyID == "" {
		log.Fatal("Missing AWS credentials or KMS key info in environment variables")
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey, secretKey, "",
		)),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	client := kms.NewFromConfig(cfg)
	return &KMSCrypto{
		Client:  client,
		KeyID:   keyID,
		Context: context.TODO(),
	}
}

func (k *KMSCrypto) Encrypt(plaintext []byte) (ciphertext []byte, nonce []byte, err error) {
	resp, err := k.Client.Encrypt(k.Context, &kms.EncryptInput{
		KeyId:     aws.String(k.KeyID),
		Plaintext: plaintext,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("KMS encrypt failed: %v", err)
	}

	ciphertext = []byte(base64.StdEncoding.EncodeToString(resp.CiphertextBlob))
	return ciphertext, nil, nil
}

func (k *KMSCrypto) Decrypt(ciphertext []byte, nonce []byte) ([]byte, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, fmt.Errorf("invalid base64 ciphertext: %v", err)
	}

	resp, err := k.Client.Decrypt(k.Context, &kms.DecryptInput{
		CiphertextBlob: cipherBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("KMS decrypt failed: %v", err)
	}
	return resp.Plaintext, nil
}
