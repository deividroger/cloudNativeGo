
FROM golang:1.9-alpine as builder


WORKDIR /go/src/github.com/deividroger/cloudNativeGo/

 COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /

COPY --from=builder /go/src/github.com/deividroger/cloudNativeGo/app /
 CMD ["/app"]