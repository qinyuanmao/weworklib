FROM golang:1.15

COPY . /home/weworklib

COPY ./lib/libWeWorkFinanceSdk_C.so /usr/local/lib
RUN echo "/usr/local/lib" >> /etc/ld.so.conf
RUN ldconfig

#TODO 在企业微信后台添加 IP 地址
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /main ./cmd/

CMD ["/main"]