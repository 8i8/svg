all:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/test

fire:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/fire.icon.120.svg

next:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/LIOT-LOUVEAU-Anais.rasi.svg

debug:
	clear
	-rm run
	go build -o run cmd/main.go
	dlv debug cmd/main.go -- test
