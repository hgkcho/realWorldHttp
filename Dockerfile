FROM golang:1.15-buster
ENV CGO_ENABLED 0
ENV GO111MODULE on

RUN set -ex && \
    apt-get update && \
    apt-get install -y git curl less mariadb-client && \
    go get -v \
      github.com/go-delve/delve/cmd/dlv@latest \
      github.com/rogpeppe/godef@latest \
      golang.org/x/tools/cmd/goimports@latest \
      github.com/ramya-rao-a/go-outline@latest \
      github.com/cweill/gotests/... \
      github.com/mdempsky/gocode@latest \
      github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
      github.com/stamblerre/gocode@latest \
      github.com/sqs/goreturns@latest \
      golang.org/x/lint/golint@latest


