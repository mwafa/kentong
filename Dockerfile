FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go build main.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/main main
EXPOSE 39000
CMD ["./main"]