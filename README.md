# interview-question-010

# Tech Stack

## Frontend

* Vue 3
* TypeScript
* Vite
* Axios

## Backend

* Golang
* Gin
* GORM
* SQLite

---

# Project Structure

```txt
online-exam-system/
├── backend/
├── frontend/
├── README.md
└── .gitignore
```

---

# Backend Structure

```txt
backend/
├── configs/
├── database/
├── handlers/
├── models/
├── routes/
├── seeders/
├── exam.db
├── go.mod
└── main.go
```

---

# Frontend Structure

```txt
frontend/
├── src/
├── public/
├── package.json
└── vite.config.ts
```

---

# Backend Setup

## 1. Go to backend folder

```bash
cd backend
```

## 2. Install dependencies

```bash
go mod tidy
```

## 3. Run backend server

```bash
go run .
```

Backend server will run on:

```txt
http://localhost:8080
```

---

# Frontend Setup

## 1. Go to frontend folder

```bash
cd frontend
```

## 2. Install dependencies

```bash
npm install
```

## 3. Run frontend server

```bash
npm run dev
```

Frontend server will run on:

```txt
http://localhost:5173
```

---

# API Endpoints

## Get Questions

```http
GET /api/questions
```

## Submit Exam

```http
POST /api/submit
```

Example request body:

```json
{
  "name": "John Doe",
  "answers": {
    "1": "B",
    "2": "C"
  }
}
```

---

# Database

This project uses SQLite database.

Database file:

```txt
backend/exam.db
```

Tables:

* questions
* choices
* exam_results

---

# Notes

* Mock questions are automatically seeded into database on first run.
* The project uses company-style naming convention with `example.com` module path.
* CORS is enabled for frontend-backend communication.

---

# Validates

* User can't submit test result if name input is empty.
* Name input will become disabled after user submited test result.
---

# Author

Panas Charoenruk
