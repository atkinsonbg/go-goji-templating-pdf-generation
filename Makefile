tests:
	go test -v ./... -coverpkg ./... -coverprofile cover.out

run:
	go run github.com/atkinsonbg/go-goji-templating-pdf-generation

docker:
	docker build -t github.com/atkinsonbg/go-goji-templating-pdf-generation/api:latest .

dockerrun:
	docker run -p 8000:8000 github.com/atkinsonbg/go-goji-templating-pdf-generation/api:latest