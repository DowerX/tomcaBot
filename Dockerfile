FROM golang:alpine AS build

RUN apk --no-cache add git ca-certificates

RUN mkdir build
WORKDIR /build
COPY ./* /build/
RUN CGO_ENABLED=0 \
    go build \
    -installsuffix "static" \
    -o app

FROM scratch AS final
COPY --from=build /build/app /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app"]
