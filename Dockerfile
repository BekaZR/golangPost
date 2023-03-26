FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go build -o main main.go

FROM golang:alpine
WORKDIR /app
ENV GIN_MODE=release
COPY --from=builder /app/main .
ENTRYPOINT [ "./main" ]
