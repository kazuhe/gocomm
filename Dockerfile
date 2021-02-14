# Dockerに対するgolangイメージから起動させる指示
FROM golang

# ローカルのパッケージファイルをコンテナの
# ワークスペースにコピーして作業ディレクトリに設定
ADD . /go/src/github.com/kazuhe/gocomm
WORKDIR /go/src/github.com/kazuhe/gocomm

# PostgreSQLのドライバを取得
RUN go get github.com/lib/pq
# サービスをビルドし実行可能バイナリファイルを/go/binに配置
RUN go install github.com/kazuhe/gocomm

# コンテナ起動時に必ず/go/bin/gocommを実行するように指示
ENTRYPOINT /go/bin/gocomm

# 他のコンテナに対して8080番ポートを開放(一般に公開する訳ではない)
EXPOSE 8080
