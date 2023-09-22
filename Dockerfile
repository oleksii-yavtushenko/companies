FROM golang:1.21.1-alpine as builder

WORKDIR /app

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

# COPY All things inside the project and build
COPY . .
RUN go build -o /app/myapp .

FROM alpine:latest
COPY --from=builder /app/myapp /app/myapp
COPY --from=builder /app/configs /app/configs


EXPOSE 8080
ENTRYPOINT [ "/app/myapp" ]