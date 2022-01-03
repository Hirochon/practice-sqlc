FROM golang:1.17.5-buster

WORKDIR /go/src/work

COPY . /go/src/work/

RUN go mod tidy && \
    go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949 && \
    go install -v golang.org/x/tools/gopls@v0.7.3 && \
    go install -v github.com/kyleconroy/sqlc/cmd/sqlc@v1.11.0 && \
    go install -v github.com/rubenv/sql-migrate/...@v1.0.0 && \
    go get -v github.com/gorilla/mux && \
    go build -o ./ ./...

EXPOSE 8080
