FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go mod tidy
RUN go build -ldflags="-s -w" -o /app/pg-backend ./pg-backend.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/pg-backend /app/pg-backend
COPY --from=builder /app/etc /app/etc

EXPOSE 80
EXPOSE 9091

CMD ["./pg-backend", "-f", "etc/pg-backend.yaml"]
