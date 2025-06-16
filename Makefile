.PHONY: run setup frontend backend refresh-deps update-swagger

run: setup
	npx concurrently "cd backend && air" "npm --prefix frontend run dev"

setup: frontend backend

frontend:
	@if [ ! -d frontend/node_modules ]; then \
		echo "Installing frontend dependencies..."; \
		cd frontend && npm install; \
	else \
		echo "Frontend dependencies already installed."; \
	fi

backend:
	@if [ ! -f backend/go.sum ]; then \
		echo "Running go mod tidy for backend..."; \
		cd backend && go mod tidy; \
	else \
		echo "Backend go.sum already exists, skipping go mod tidy."; \
	fi

refresh-deps:
	@echo "Refreshing frontend and backend dependencies..."
	cd frontend && npm install
	cd backend && go mod tidy

update-swagger:
	@echo "Generating Swagger docs..."
	cd backend && swag init -g ./cmd/server/main.go
