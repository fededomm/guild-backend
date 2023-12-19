APP_NAME = guild-backend

.PHONY: swagger run build clean 

clean:
	rm -rf ./$(APP_NAME)

swagger:
	swag init

build:
	go build -o $(APP_NAME) ./main.go

run:
	./$(APP_NAME)

all: clean build swagger run 
