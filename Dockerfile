# --------------------
# Go build for local
# --------------------
FROM golang:1.21.0-alpine3.18 AS arm-build
WORKDIR /app
COPY /src/go.mod /src/go.sum ./
RUN go mod download
COPY /src/ ./
RUN go build -o main .

# --------------------
# Build the local container
# --------------------
FROM alpine:3.18 AS local
WORKDIR /root/
COPY --from=arm-build /app/main .
EXPOSE 8080
CMD ["./main"]