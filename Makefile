all:
	clear
	go run cmd/main.go -- assets/fire.icon.120.svg

short:
	clear
	go run cmd/main.go -- assets/short.svg > temp.svg
car:
	clear
	go run cmd/main.go -- assets/car.svg > temp.svg

car2:
	clear
	go run cmd/main.go -- assets/tommek_Car.svg > temp.svg

anim:
	clear
	go run cmd/main.go -- assets/scimitar-anim.svg > temp.svg

ff:
	clear
	go run cmd/main.go -- assets/fire.icon.120.svg > temp.svg
	firefox temp.svg

test:
	clear
	go run cmd/main.go -- assets/test.svg > temp.svg
	firefox temp.svg

debug:
	clear
	-rm run
	go build -o run cmd/main.go
	dlv debug cmd/main.go -- assets/short.svg
