# Project variables
COMPOSE_FILE=compose.yml
APP_CONTAINER=instashop
DB_CONTAINER=instashop-db

# Default port for the application
PORT=15001

# Build the Docker images
.PHONY: build
build:
	@echo "Building Docker images..."
	docker-compose -f $(COMPOSE_FILE) build

# Start the containers in detached mode
.PHONY: start
start:
	@echo "Starting containers..."
	docker-compose -f $(COMPOSE_FILE) up -d

# Stop the containers
.PHONY: stop
stop:
	@echo "Stopping containers..."
	docker-compose -f $(COMPOSE_FILE) down

# Deploy: build and start the containers
.PHONY: deploy
deploy: build start
	@echo "Containers deployed successfully!"

# Clean up Docker images and containers
.PHONY: clean
clean:
	@echo "Stopping and removing containers..."
	docker-compose -f $(COMPOSE_FILE) down --rmi all --volumes --remove-orphans

# Tail logs for the app container
.PHONY: logs
logs:
	@echo "Tailing logs for app container..."
	docker-compose -f $(COMPOSE_FILE) logs -f app

# Check the status of containers
.PHONY: status
status:
	@echo "Checking status of containers..."
	docker-compose -f $(COMPOSE_FILE) ps

# Access the app container shell
.PHONY: shell
shell:
	@echo "Opening shell for app container..."
	docker exec -it $(APP_CONTAINER) sh
