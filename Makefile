BINARY_NAME=myapp

build:
	@go mod vendor
	@echo "Building Brimestoneesan..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Brimestoneesan built!"

run: build
	@echo "Starting Brimestoneesan..."
	@./tmp/${BINARY_NAME} &
	@echo "Brimestoneesan started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Brimestoneesan..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Brimestoneesan!"

restart: stop start