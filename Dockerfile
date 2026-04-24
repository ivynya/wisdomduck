FROM golang:1.26-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /wisdomduck

FROM alpine:3.23 as production
WORKDIR /app
COPY --from=builder /wisdomduck .
COPY ./duck ./duck

EXPOSE 5500

CMD [ "./wisdomduck" ]
