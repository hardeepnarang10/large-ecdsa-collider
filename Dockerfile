# Stage 1
FROM golang:1.17-alpine3.14 as builder

# Add git
RUN apk update && \
    apk add gcc && \
    apk add libc-dev

RUN mkdir /large-ecdsa-collider

ADD . /large-ecdsa-collider

WORKDIR /large-ecdsa-collider

RUN go mod tidy

RUN go build .

# Stage 2
FROM alpine:3.14

RUN apk update && \
    touch daemon_crash.log

COPY --from=builder /large-ecdsa-collider/large-ecdsa-collider /

CMD ["./large-ecdsa-collider -node=<Insert Node Address> -timeout=15 -wallets=5 -workers=20 -debug=true", "2>>", "daemon_crash.log"]
