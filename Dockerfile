FROM golang:1.23-alpine AS builder

WORKDIR /build

ADD . .


RUN mkdir -p out && go build -o out/GreatIterator

FROM scratch

COPY --from=builder /build/out/GreatIterator .

ENTRYPOINT ["/GreatIterator"]