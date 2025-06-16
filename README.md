# Uptime Monitor – Development Guide

This document outlines how to run the project, manage dependencies, and use the provided `Makefile`.

---

## ⚙️ Requirements

- [Go](https://golang.org/doc/install)
- [Node.js & npm](https://nodejs.org/)
- [Air](https://github.com/air-verse/air) for live reloading the Go backend
- [Make](https://www.gnu.org/software/make/)

---

## 🚀 Run Project

To start both frontend and backend with hot-reloading:

```sh
make run
```

This runs:

- `air` for live-reloading the Go backend
- `npm run dev` for the Vite React frontend

---

## 🔁 Dependency Management

### ✅ One-time install (auto-skipped if already done)

`make run` will check:

- If `frontend/node_modules` exists – if not, runs `npm install`
- If `backend/go.sum` exists – if not, runs `go mod tidy`

### 🔄 Force update dependencies

To refresh all dependencies manually:

```sh
make refresh-deps
```

This forcibly runs:

- `npm install` in `frontend`
- `go mod tidy` in `backend`

---

## 📁 File Structure (Partial)

```
uptime-monitor/
├── backend/
│   └── cmd/server/     # main.go entry point
├── frontend/           # React frontend
├── Makefile
└── README.md
```
---

## Updating Swagger Documentation (Backend)

The backend API uses Swagger (OpenAPI) documentation generated from Go code annotations.

### How to Update Swagger Docs

1. **Edit or Add Annotations:**
   - In your Go handler files (e.g., `backend/handler/`), update or add Swagger annotations above your handler functions.
   - Example:
     ```go
     // @Summary      Get all targets
     // @Description  Retrieve all uptime monitoring targets
     // @Tags         targets
     // @Accept       json
     // @Produce      json
     // @Success      200  {array}   model.Target
     // @Router       /targets [get]
     ```

2. **Generate the Swagger Spec:**
   - Make sure you have [swag](https://github.com/swaggo/swag) installed:
     ```sh
     go install github.com/swaggo/swag/cmd/swag@latest
     ```
   - From the `backend` directory, run:
     ```sh
     swag init -g cmd/server/main.go
     ```
   - This will generate or update the `docs/` folder with the OpenAPI spec.

3. **View the Updated Docs:**
   - Restart your backend server if it's running.
   - Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser to view the updated API documentation.

---

## 🧪 Helpful Commands

| Command              | Description                              |
|----------------------|------------------------------------------|
| `make run`           | Start both frontend & backend            |
| `make refresh-deps`  | Force re-install frontend/backend deps   |
| `make update-swagger`| Update swagger documentation (require swag installed)             |

---
