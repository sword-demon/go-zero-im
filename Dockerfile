# go环境测试

FROM alpine:3.18.4

RUN mkdir /go && cd /go \
    && wget --no-check-certificate https://golang.google.cn/dl/go1.24.0.linux-amd64.tar.gz \
    && tar -C /usr/local -zxf go1.24.0.linux-amd64.tar.gz \
    && rm -rf /go/go1.24.0.linux-amd64.tar.gz \
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH

CMD ["ping", "baidu.com"]
