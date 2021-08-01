all:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/chart.svg

fire:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/fire.icon.120.svg

next:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/test

debug:
	clear
	-rm run
	go build -o run cmd/main.go
	dlv debug cmd/main.go -- test
