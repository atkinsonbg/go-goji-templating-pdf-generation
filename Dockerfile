FROM golang:latest as BUILD
WORKDIR builddir
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-goji-templating-pdf-generation-api

FROM alpine:3.12.0
RUN apk update && \
    apk add xvfb ttf-freefont fontconfig wkhtmltopdf ghostscript
COPY --from=BUILD ./go/builddir/go-goji-templating-pdf-generation-api .
COPY test test
ENTRYPOINT ["./go-goji-templating-pdf-generation-api"]