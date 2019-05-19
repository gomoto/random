FROM golang:1.12-stretch as compiler
WORKDIR /go/src/github.com/gomoto/random
COPY ./nist ./nist
COPY ./numbers ./numbers
COPY ./main.go ./main.go
RUN go build -o /go/bin/random .

FROM buildpack-deps:stretch-scm as runner
WORKDIR /app
COPY --from=compiler /go/bin/random ./random
ENTRYPOINT [ "./random" ]
