FROM golang:alpine3.14 AS builder
RUN apk update && apk add git
COPY /app.go /app/
COPY go.mod go.sum /app/
WORKDIR /app
RUN go mod tidy
RUN go build .../..
CMD ["./k8s-demo"]

FROM alpine:3.14
COPY --from=builder /app/k8s-demo /app/
EXPOSE 8080
CMD ["/app/k8s-demo"]
