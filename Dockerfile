FROM golang:1.12-stretch AS builder
RUN mkdir -p /go/src/github.com/4geeks/pex-awardwallet
WORKDIR /go/src/github.com/4geeks/pex-awardwallet
COPY . .
RUN  go get -u github.com/golang/dep/cmd/dep && dep ensure -v
EXPOSE 8080

FROM alpine:3  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/4geeks/pex-awardwallet/.env .
COPY --from=builder /go/src/github.com/4geeks/pex-awardwallet/* .
CMD ["./pex-awardwallet"]
EXPOSE 8080
