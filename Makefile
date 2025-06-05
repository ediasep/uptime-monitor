.PHONY: run
run:
	npx concurrently "go run -C backend ./cmd/server" "npm --prefix frontend run dev"