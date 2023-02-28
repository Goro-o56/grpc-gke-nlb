package secretmanager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"log"
	"os"
)

func AccessVersion() {
	projectID := "237602680776"

	// クライアント生成
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	// シークレット、test-appの最新バージョンにアクセス
	resourceName := "projects/" + projectID + "/secrets/" + "API_KEY" + "/versions/1"
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: resourceName,
	}

	// シークレット上にアクセスする
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err) //何かあった時はここで終わってほしい
	}
	log.Printf("%v", result.Payload)
	prefix := []byte("API_KEY=")

	bytes := append(prefix, result.Payload.Data...)

	//result.Payloadを書き込んでみる
	err = os.WriteFile(".env", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File written successfully.")
}
