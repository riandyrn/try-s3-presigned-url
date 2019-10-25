FROM golang:1.13.3-alpine3.10 as build
WORKDIR /go/src/github.com/riandyrn/try-s3-presigned-url

ENV GO111MODULE on
ADD . .
RUN go build -mod=vendor -o app

FROM alpine:3.10
RUN apk add ca-certificates &&\
    apk add -U tzdata
COPY --from=build /go/src/github.com/riandyrn/try-s3-presigned-url/app ./app
COPY --from=build /go/src/github.com/riandyrn/try-s3-presigned-url/web/ ./web/

ENTRYPOINT [ "./app" ]
EXPOSE 8765