package main

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

func BenchmarkGptClient(b *testing.B) {
	input := "以下は自己紹介の為のフォーマットです。かわいい日本人のラノベキャラっぽい名前を生成してください。\n【仕事】アイドル\n【趣味】イラスト\n【性別】女性\n【名前】\n【自己紹介文:4行】\n"
	instruction := "自己紹介のフォーマットの空いている欄を補完して"

	resp := GptEditClient(input, instruction)

	fmt.Println(resp)
	return
}

func TestAccess(t *testing.T) {
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
		t.Fatalf("failed to access secret version: %v", err)
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
