FROM golang:1.23.1-alpine3.20 as builder
WORKDIR /app
COPY . /app
RUN go mod download && go build -o main .
RUN chown 1001:1001 ./main && chmod 700 ./main

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /bin/main
# EXPOSE 8080
USER 1001
ENTRYPOINT ["/bin/main"]
