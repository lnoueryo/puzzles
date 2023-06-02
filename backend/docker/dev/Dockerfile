FROM golang:1.20-alpine
RUN apk add build-base
WORKDIR /app

# COPY wait-for-it.sh /usr/local/bin/wait-for-it.sh
# RUN chmod +x /usr/local/bin/wait-for-it.sh
# ENTRYPOINT ["wait-for-it.sh", "192.168.0.2:3306", "--"]

COPY . /app/.
RUN go mod download
EXPOSE 8200
ENV HOST 0.0.0.0