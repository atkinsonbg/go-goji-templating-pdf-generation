FROM golang:latest as BUILD
WORKDIR builddir
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-goji-templating-pdf-generation-api
RUN upx --best --ultra-brute go-goji-templating-pdf-generation-api


FROM scratch
COPY --from=BUILD ./go/builddir/go-goji-templating-pdf-generation-api .
ENTRYPOINT ["./go-goji-templating-pdf-generation-api"]