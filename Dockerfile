FROM golang:1.17 AS builder
WORKDIR /go/src/adversary16/armwatch
COPY . ./
RUN GOOS=linux CGO_ENABLED=0 go build .

FROM alpine:latest 
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/adversary16/armwatch/armwatch ./
COPY --from=builder /go/src/adversary16/armwatch/conf* ./
RUN ls
RUN chmod +x ./armwatch
RUN echo ./armwatch > entrypoint.sh
ENTRYPOINT ["sh entrypoint.sh"]

