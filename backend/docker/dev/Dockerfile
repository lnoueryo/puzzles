FROM golang:1.18-alpine
RUN apk add build-base
WORKDIR /app
# COPY ./backend/go.mod /app/go.mod
# COPY ./backend/go.sum /app/go.sum
COPY . /app/.
RUN go mod download
EXPOSE 8080
ENV HOST 0.0.0.0