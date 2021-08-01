all:
	clear
	-rm run
	go build -o run cmd/main.go
	./run assets/test.svg

ff:
	clear
	-rm run
	-rm temp.svg
	go build -o run cmd/main.go
	./run assets/test.svg > temp.svg
	firefox temp.svg

debug:
	clear
	-rm run
	go build -o run cmd/main.go
	dlv debug cmd/main.go -- assets/test.svg
