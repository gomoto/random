FROM golang:1.12-stretch as compiler
WORKDIR /go/src/github.com/gomoto/random
COPY ./nist ./nist
COPY ./numbers ./numbers
COPY ./main.go ./main.go
RUN go build -o /go/bin/random .
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -a \
    -ldflags '-w -extldflags "-static"' \
    -o /go/bin/random \
    .

FROM alpine:3.9
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
WORKDIR /app
COPY --from=compiler /go/bin/random ./random
ENTRYPOINT [ "./random" ]
