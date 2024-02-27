FROM golang:1.22 AS builder

ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/milvus-migration

From alpine:3.17
WORKDIR /app
COPY --from=builder /app/milvus-migration .
COPY --from=builder /app/configs ./configs
EXPOSE 8080
ENTRYPOINT ["./milvus-migration", "server"]