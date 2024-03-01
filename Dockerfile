# Build
FROM golang:latest as build

WORKDIR /app
ADD . /app

RUN cd /app && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /http-service .

CMD /http-service

# Test
FROM build as test

# Production build
FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

COPY --from=build /http-service ./

RUN chmod +x ./http-service

ENTRYPOINT [ "./http-service" ]

EXPOSE 8080