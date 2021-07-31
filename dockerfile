FROM alpine:3.11

RUN mkdir /go && \
 cd /go && \
 wget https://golang.google.cn/dl/go1.16.6.linux-amd64.tar.gz && \
 tar -zxvf go1.16.6.linux-amd64.tar.gz && \
 rm -fr go1.16.6.linux-amd64.tar.gz

 RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld=linux-x86-65.so.2

 ENV GOPATH /go
 ENV PATH /usr/local/go/bin:$GOPATH/bin:$PATH
