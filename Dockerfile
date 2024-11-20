FROM golang:1.23.1-alpine3.20 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/src/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN chown 1001:1001 ./main && chmod 700 ./main

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /bin/main
# EXPOSE 8080
USER 1001
ENTRYPOINT ["/bin/main"]
