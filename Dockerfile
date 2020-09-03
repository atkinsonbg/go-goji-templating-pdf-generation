FROM golang:latest as BUILD
WORKDIR builddir
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-goji-templating-pdf-generation-api

FROM alpine:3.12.0
RUN wget http://dl-cdn.alpinelinux.org/alpine/edge/community/x86_64/wkhtmltopdf-0.12.6-r0.apk
RUN apk update && \
    apk add --allow-untrusted xvfb \
    ttf-freefont \
    fontconfig \
    wkhtmltopdf-0.12.6-r0.apk \
    ghostscript
COPY --from=BUILD ./go/builddir/go-goji-templating-pdf-generation-api .
COPY templates templates
ENTRYPOINT ["./go-goji-templating-pdf-generation-api"]