FROM golang:alpine as builder
RUN apk update && apk add git && apk add binutils 
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/github.com/kostiamol/go-rest-api-template
WORKDIR $GOPATH/src/github.com/kostiamol/go-rest-api-template
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o $GOPATH/bin/go-rest-api-template ./cmd/go-rest-api-template
RUN cd $GOPATH/bin \
    strip --strip-unneeded go-rest-api-template

FROM scratch
ARG port=127.0.0.1
EXPOSE $port
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/centerms /go/bin/centerms
USER appuser
ENTRYPOINT ["/go/bin/go-rest-api-template"]

