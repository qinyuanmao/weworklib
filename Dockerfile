FROM golang:1.15

RUN mv lib/libWeWorkFinanceSdk_C.so /usr/local/lib
RUN echo "/usr/local/lib" >> /etc/ld.so.conf
RUN ldconfig

#TODO 在企业微信后台添加 IP 地址

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-linkmode external -extldflags -static" -o /test_tool ./cmd/