FROM golang:1.18.3-alpine3.16 AS development
ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOARCH=amd64 \
    GOOS=linux

WORKDIR /app
COPY . .
RUN apk add build-base
RUN go mod download && go mod tidy -go=1.18
EXPOSE 8001 9001
HEALTHCHECK --interval=5m --timeout=3s CMD curl --fail http://localhost:8001/ || exit 1
CMD ["go", "run", "."]

FROM golang:1.18.3-alpine3.16 AS build
ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOARCH=amd64 \
    GOOS=linux

COPY --from=development /app/ /app/
WORKDIR  /app
RUN apk add build-base
RUN go build -o app 

FROM alpine:3.16 AS production

COPY --from=build /app /usr/local/app
EXPOSE 8001 9001
USER nobody:nobody

HEALTHCHECK --interval=5m --timeout=3s CMD curl --fail http://localhost:8001/ || exit 1
ENTRYPOINT ["/usr/local/app"]