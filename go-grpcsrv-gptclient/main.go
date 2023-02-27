package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"go-grpcsrv-gptclient/health"
	gptpb "go-grpcsrv-gptclient/pkg/grpc"
	"go-grpcsrv-gptclient/secretmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	//1. 8080 portのlistener
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	//2. gRPCサーバーを作成
	s := grpc.NewServer()

	//環境変数のロード
	secretmanager.AccessVersion()
	//3.登録
	gptpb.RegisterGeneratePersonaCompletionServiceServer(s, NewGptServer())
	grpc_health_v1.RegisterHealthServer(s, &health.Server{})
	reflection.Register(s)
	//4. 作成したgRPCサーバーを8080番ポートで動かす
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := s.Serve(listener)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}

	}()

	//5.Ctrl + C が入力されたらGraceful Shutdownされる
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

}

// **ServiceServerを実装する構造体の定義
func NewGptServer() *gptCompletionServer {
	return &gptCompletionServer{}
}

type gptCompletionServer struct {
	gptpb.UnimplementedGeneratePersonaCompletionServiceServer
}

func (s *gptCompletionServer) GeneratePersonaCompletion(ctx context.Context, prompt *gptpb.Prompt) (*gptpb.PersonaResponse, error) {
	value := prompt.GetValue()
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	client := gpt3.NewClient(apiKey)
	resp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:           value,
		MaxTokens:        gpt3.IntPtr(206),
		Temperature:      gpt3.Float32Ptr(1),
		TopP:             gpt3.Float32Ptr(0.5),
		PresencePenalty:  0.6,
		FrequencyPenalty: 0.25,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &gptpb.PersonaResponse{Message: resp.Choices[0].Text}, nil
}
