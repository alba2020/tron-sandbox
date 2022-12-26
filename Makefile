
run: build
	clear && ./sandbox

build: *.go
	go build .

mcprof:
	go test --bench MonteCarloAI -cpuprofile mc.prof

clean:
	rm -f sandbox mc.prof
