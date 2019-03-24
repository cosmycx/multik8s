FROM  golang:latest as builder

WORKDIR /go/src/goapi

ENV GOPATH /go

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM scratch 

COPY --from=builder /go/src/goapi/goapi /goapi/

EXPOSE 4500
ENTRYPOINT [ "/goapi/goapi" ]