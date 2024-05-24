run:
	go run main.go
unit:
	go test -v --tags=unit
integration:
	go test -v --tags=integration
testall:
	go test -v --tags=all -coverprofile=coverage.out
