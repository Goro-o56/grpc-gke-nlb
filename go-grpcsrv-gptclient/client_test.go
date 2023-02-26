package main

import (
	"fmt"
	"testing"
)

func BenchmarkGptClient(b *testing.B) {
	input := "以下は自己紹介の為のフォーマットです。かわいい日本人のラノベキャラっぽい名前を生成してください。\n【仕事】アイドル\n【趣味】イラスト\n【性別】女性\n【名前】\n【自己紹介文:4行】\n"
	instruction := "自己紹介のフォーマットの空いている欄を補完して"

	resp := GptEditClient(input, instruction)

	fmt.Println(resp)
	return
}
