build:
	go build -o ./bin/dynarr

run: build
	./bin/dynarr

clean:
	rm -r ./bin

