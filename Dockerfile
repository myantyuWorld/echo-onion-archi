FROM golang:1.22 AS develop

ENV TZ Asia/Tokyo
ENV GO111MODULE=on

WORKDIR /go/src/app

COPY go.sum go.mod ./

RUN go install github.com/air-verse/air@v1.52.3
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/golang/mock/mockgen@v1.6.0

RUN go mod download

CMD ["air", "-c", ".air.toml"]