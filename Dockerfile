FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -a -o solaredge -buildvcs=false .


FROM alpine:latest

RUN apk update && \
    apk add --no-cache tzdata

ENV TZ Europe/Warsaw

COPY --from=builder /build/solaredge .

ENTRYPOINT [ "./solaredge" ]
CMD [ "--config", "solaredge.yaml" ]
