# --------------------
# Go build for server
# --------------------
FROM --platform=linux/amd64 golang:1.21.0-alpine3.18 AS amd-build
WORKDIR /app
COPY /src/go.mod /src/go.sum ./
RUN go mod download
COPY /src/ ./
RUN go build -o main .

# --------------------
# Build the container for the server
# --------------------
FROM --platform=linux/amd64 alpine:3.18 AS server
WORKDIR /root/
COPY --from=amd-build /app/main .
EXPOSE 8080
CMD ["./main"]