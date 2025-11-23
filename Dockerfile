FROM golang:1.24.10 AS builder
WORKDIR /src
COPY src/go.mod ./
RUN go mod download
COPY src/* ./
RUN GOOS=linux GOARCH=amd64 go build -o demo

FROM scratch
COPY --from=builder /src/demo /demo
EXPOSE 8080
CMD ["/demo"]
