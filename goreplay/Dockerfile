FROM golang:1.19 as builder

RUN apt-get update && apt-get install ruby vim-common -y

RUN apt-get install flex bison -y
RUN wget http://www.tcpdump.org/release/libpcap-1.10.0.tar.gz && tar xzf libpcap-1.10.0.tar.gz && cd libpcap-1.10.0 && ./configure && make install

WORKDIR /go/src/github.com/buger/goreplay/
RUN wget https://github.com/buger/goreplay/archive/refs/tags/v2.0.0-rc2.tar.gz -O gor.tar.gz && tar xzf gor.tar.gz -C . --strip=1
RUN go mod vendor
RUN go build -mod=vendor -tags netgo -o gor -ldflags "-extldflags \"-static\""


FROM alpine:3.17

RUN apk add bash
RUN apk add --no-cache ca-certificates openssl

COPY --from=builder /go/src/github.com/buger/goreplay/gor /
ENTRYPOINT ["./gor"]
