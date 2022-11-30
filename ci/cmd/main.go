package main

import (
	"context"
	"log"
	"os"

	"dagger.io/dagger"
)

var (
	repo   = "https://github.com/FlowingSPDG/Got5.git"
	branch = "unknown"
	commit = "unknown"
)

func main() {
	log.Println("repo:", repo)
	log.Println("branch:", branch)
	log.Println("commit:", commit)

	// Get context
	ctx := context.Background()

	// Dagger クライアントを初期化
	// dagger.WithLogOutput でログの出力先を設定することができる
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// リポジトリをクローン
	repo := client.Git(repo)
	src := repo.Branch(branch).Tree()

	// Docker イメージを読み込む
	golang := client.Container().From("golang:1.19")

	// クローンしたリポジトリのソースをマウントする
	// WithWorkdir で作業ディレクトリを設定できる
	golang = golang.WithMountedDirectory("/app", src).WithWorkdir("/app")

	// 実行するコマンドを追加
	golang = golang.
		Exec(dagger.ContainerExecOpts{
			Args: []string{"go", "test", "-v", "-count=1", "./..."},
		})

	// ExitCode でパイプラインが実行される
	if _, err := golang.ExitCode(ctx); err != nil {
		panic(err)
	}
}
