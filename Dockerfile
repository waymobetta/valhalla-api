FROM golang:1.12rc1-alpine3.9 AS build

RUN apk --no-cache add ca-certificates
COPY . /go/src/github.com/waymobetta/valhalla-api
WORKDIR /go/src/github.com/waymobetta/valhalla-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o valhalla-api cmd/valhalla/main.go

FROM scratch

WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/waymobetta/valhalla-api/cmd/valhalla/db.json .
COPY --from=build /go/src/github.com/waymobetta/valhalla-api/valhalla-api .

# expose default port
EXPOSE 5000

# start app
CMD ["./valhalla-api"]
