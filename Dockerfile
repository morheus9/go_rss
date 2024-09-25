FROM golang:1.23.1-alpine3.20 as builder
WORKDIR /app
COPY . /app
RUN go mod download && go build -o /main ./src/cmd/

FROM scratch
WORKDIR /app
COPY --from=builder main /bin/main
EXPOSE 8080
ENTRYPOINT ["/bin/main"]