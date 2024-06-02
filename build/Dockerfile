# ベースイメージを指定
FROM golang:1.22

# 作業ディレクトリを作成
WORKDIR /go/src/app

# ローカルのファイルをコンテナにコピー
COPY . .

# 8080ポートを公開
EXPOSE 8080

# Goのアプリケーションをビルドして実行
CMD ["go", "run", "./cmd/my-go-app/main.go"]
