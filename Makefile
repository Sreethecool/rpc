 
all: build test
 
build:
 go build main.go
 
test:
 go test -v ./..
 
run:
 go build main.go
 ./main
 
clean:
 go clean
 rm main
