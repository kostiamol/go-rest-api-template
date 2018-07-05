FROM golang:alpine as builder
ENV http_proxy=http://135.245.192.7:8000 
RUN apk update && apk add git && apk add binutils 
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/github.com/kostiamol/go-rest-api-template
WORKDIR $GOPATH/src/github.com/kostiamol/go-rest-api-template
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o $GOPATH/bin/go-rest-api-template ./cmd/go-rest-api-template
COPY VERSION fixtures.json $GOPATH/bin/rsc/
RUN cd $GOPATH/bin \
    strip --strip-unneeded go-rest-api-template

FROM scratch
ENV ENV=PROD
EXPOSE 8080
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/go-rest-api-template /go/bin/go-rest-api-template
COPY --from=builder /go/bin/rsc/VERSION /go/bin/rsc/fixtures.json /go/bin/rsc/
WORKDIR /go/bin
ENTRYPOINT ["./go-rest-api-template"]
