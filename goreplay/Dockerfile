FROM alpine:3.17

RUN apk add bash
RUN apk add curl

ARG RELEASE_VERSION=1.3.3

RUN apk add --no-cache ca-certificates openssl
RUN wget https://github.com/buger/goreplay/releases/download/${RELEASE_VERSION}/gor_${RELEASE_VERSION}_x64.tar.gz -O gor.tar.gz
RUN tar xzf gor.tar.gz

ENTRYPOINT ["./gor"]
