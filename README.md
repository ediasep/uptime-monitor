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

## 🧪 Helpful Commands

| Command             | Description                              |
|---------------------|------------------------------------------|
| `make run`          | Start both frontend & backend            |
| `make refresh-deps` | Force re-install frontend/backend deps   |