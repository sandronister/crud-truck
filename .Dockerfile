FROM golang:1.22 as builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main .


FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/
COPY --from=builder /app/.env /app/

CMD ["./main"]
