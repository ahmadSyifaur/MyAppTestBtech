 #MyAppTestBtech – Fullstack Authentication App

Aplikasi fullstack authentication sederhana menggunakan:

- Backend: Golang Gin v1.24.2
- Frontend: Vue 3 + Vite
- Database: MySQL
- Authentication: JWT
- Deployment: Docker & Docker Compose

Aplikasi menyediakan fitur:
- Register
- Login
- Protected Dashboard (JWT)
- Frontend API request via Axios
- Vue Router navigation
- Backend API v1

---

## 1. Features

### Backend (Golang Gin)
- REST API with Gin
- JWT Authentication middleware
- Password hashing
- API Endpoints:
  - POST /api/v1/register
  - POST /api/v1/login
  - GET /api/v1/protected (JWT required)

### Frontend (Vue 3 + Vite)
- Login & Register page
- Axios request client
- JWT stored in localStorage
- Dashboard (protected route)
- Logout

### Database
- MySQL
- Auto migrate table users (GORM)

### Docker
- Backend container
- Frontend container
- MySQL container

---

## 2. Project Structure
/backend
├── main.go
├── controllers/
├── configs/
├── helpers/
└── models/

/frontend
├── src/
├── vite.config.js
└── Dockerfile

docker-compose.yml
README.md


## 3. API Documentation

### Register
**POST /api/v1/register**

Request:
```json
{
  "email": "admin@admin.com",
  "password": "P@ssw0rd",
  "confirmPassword": "P@ssw0rd"
}

Response:
{
  "responseCode": "00",
  "responseMessage": "User registered successfully"
}

### Login
**POST /api/v1/login**

Request:
```json
{
  "email": "admin@admin.com",
  "password": "P@ssw0rd"
}

Response:
{
    "responseCode": "00",
    "responseMessage": "Successful login",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGFkbWluLmNvbSIsImV4cCI6MTc2NDE4NzQ5NiwidXNlcl9pZCI6Mn0.W-6Qe7RE6HjcNwZAb2Ba_2uyxYUW3-6Of1JAnoBG19A"
}

### DashBoard(Protected)
**GET /api/v1/dashboard**

Response:
{
  "message": "Hello admin@admin.com, welcome back"
}



docker compose up -d
