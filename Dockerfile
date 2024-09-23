FROM golang:1.22.1-alpine3.19 as builder
WORKDIR /app
COPY . /app
RUN go mod download && go build -o /main ./src/main/

FROM scratch
WORKDIR /app
COPY --from=builder main /bin/main
EXPOSE 8080
ENTRYPOINT ["/bin/main"]