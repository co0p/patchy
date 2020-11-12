FROM golang:1.14-alpine AS build

RUN apk add --no-cache git gcc musl-dev

COPY . /go/src/github.com/co0p/patchy
WORKDIR /go/src/github.com/co0p/patchy

RUN go build -o /bin/patchy cmd/server/main.go

FROM alpine:3.7 AS final
RUN apk add --no-cache
COPY --from=build /bin/patchy /bin/patchy
CMD ["/bin/patchy"]
