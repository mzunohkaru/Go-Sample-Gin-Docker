FROM golang:1.21.4

WORKDIR /app
# Dockerコンテナ内にディレクトリを作り、そこへ移動する

COPY src .
# ローカルのコードをコンテナ内のディレクトリ ( WORKDIRで指定した app ディレクトリ ) にコピー
COPY src/go.mod .
COPY src/go.sum .
# mod ファイルと sum ファイルをコンテナ内のディレクトリにコピー

RUN go mod download
# イメージ作成時に動かすコマンドで、mod ファイルにあるパッケージを全てダウンロードする

# CMD [ "go", "run", "main.go"]
# イメージ作成後、コンテナを動かした時に起動させる

