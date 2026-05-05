# 🚀 Go Task API

A modern, scalable RESTful Task Management API built with **Go (Golang)** following clean architecture principles.
Designed for performance, simplicity, and real-world backend development.

---

## ✨ Features

* ✅ Create, Read, Update, Delete (CRUD) Tasks
* 👤 User-based task management
* ⚡ Fast and lightweight (built with Go `net/http`)
* 🧱 Clean architecture (handlers, services, repository)
* 🔐 Middleware support (CORS, logging, etc.)
* 🗄️ Database integration (PostgreSQL)
* 🧪 Structured and maintainable codebase

---

## 📁 Project Structure

```
go-task-api/
│
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   ├── services/
│   ├── models/
│   ├── repository/
│   └── middleware/
│
├── pkg/
│   └── database/
│
├── .env
├── go.mod
└── README.md
```

---

## ⚙️ Tech Stack

* Language: Go (Golang)
* HTTP: net/http
* Database: PostgreSQL
* Architecture: Clean Architecture

---

## 🚀 Getting Started

### 1️⃣ Clone the repository

```
git clone https://github.com/ajilaries/GO-TASK-API.git
cd go-task-api
```

---

### 2️⃣ Install dependencies

```
go mod tidy
```

---

### 3️⃣ Setup environment variables

Create a `.env` file:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db
PORT=8080
```

---

### 4️⃣ Run the application

```
go run ./cmd/app
```

Server will start at:

```
http://localhost:8080
```

---

## 📌 API Endpoints

### 📝 Tasks

| Method | Endpoint    | Description     |
| ------ | ----------- | --------------- |
| GET    | /tasks      | Get all tasks   |
| POST   | /tasks      | Create new task |
| PUT    | /tasks/{id} | Update a task   |
| DELETE | /tasks/{id} | Delete a task   |

---

## 🧪 Example Request

### Create Task

```
POST /tasks
Content-Type: application/json

{
  "title": "Learn Go",
  "completed": false,
  "user_id": 1
}
```

---

## 📦 Build Executable (.exe)

To build a Windows executable:

```
go build -o go-task-api.exe ./cmd/app
```

For versioned build:

```
go build -o go-task-api-v1.1.exe ./cmd/app
```

---

## 🆕 Latest Update (v1.1.0)

* ✨ Added task update functionality
* 🐛 Fixed task fetching issues for users
* ⚡ Improved API performance
* 🧱 Code structure improvements

---

## 🔮 Future Improvements

* 🔐 Authentication & Authorization (JWT)
* 📊 Pagination & filtering
* 🌐 Deployment (Docker / Cloud)
* 📱 Frontend integration

---

## 🤝 Contributing

Contributions are welcome!
Feel free to fork this repo and submit a pull request.

---

## 📄 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

Built with ❤️ using Go.
