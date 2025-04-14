build:
	@go build -o bin/skyshift ./

run: build
	@./bin/skyshift

test:
	@go test ./ --cover

clean:
	@rm -rf bin/