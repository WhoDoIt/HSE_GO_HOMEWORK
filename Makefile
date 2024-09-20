all: build exec

build:
	go build cmd/task/main.go

exec:
	./main.exe