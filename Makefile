run:
	go run github.com/atkinsonbg/go-goji-templating-pdf-generation

docker:
	docker build -t github.com/atkinsonbg/go-goji-templating-pdf-generation/api:latest .

dockerrun:
	docker run -it github.com/atkinsonbg/go-goji-templating-pdf-generation/api:latest