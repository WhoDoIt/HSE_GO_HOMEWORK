task1: task1_build exec

task1_build:
	go build task1/cmd/task/main.go

exec:
	./main.exe