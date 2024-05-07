FROM golang:1.18 AS server
ENV GO111MODULE=on
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /root/itshujia
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o "server" .

FROM scratch
WORKDIR /root/itshujia/
COPY --from=server /root/itshujia/server ./
COPY config.yaml ./
ENTRYPOINT ["./server"]