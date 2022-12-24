
run: build
	clear && ./sandbox

build: *.go
	go build .

clean: sandbox
	rm -f sandbox
