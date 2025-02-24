# 📚 Library Management System  

A **Library Management System** built with **Go**, **GORM**, **PostgreSQL**, and **Docker**. The system allows users to manage books and borrowing records efficiently, featuring authentication, book inventory management, and role-based access control.

---

## 🚀 Features  

✅ **User Management** (Admin & Regular Users)  
✅ **Book Management** (Add, List, Update, and Remove Books)  
✅ **Borrow & Return Books** (Track borrowed books)  
✅ **JWT-Based Authentication** (Secure login & access tokens)  
✅ **Role-Based Access Control** (Middleware for authorization)  
✅ **Automatic Database Seeding** (Initial users and books upon startup)  
✅ **Docker Support** (Run everything with `docker-compose`)  
✅ **GORM Integration** (ORM for PostgreSQL)  
✅ **Scalable & Modular Architecture**  

---

## 🛠️ Tech Stack  

- **Go** (Backend)  
- **GORM** (ORM for PostgreSQL)  
- **PostgreSQL** (Database)  
- **JWT** (Authentication)  
- **Docker & Docker Compose** (Containerization)  
- **Alpine Linux** (Minimal Base Image)  

---

## 📂 Project Structure  

```
/library-management
│── cmd/
│   └── main.go              # Main application entry point
│── internal/
│   ├── bootstrap/           # Application initialization (e.g., database, server setup)
│   ├── constants/           # Global constants used across the application
│   ├── dto/                 # Data Transfer Objects (Request/Response validation)
│   ├── handlers/            # API request handlers (Controllers)
│   ├── middleware/          # Authentication & role-based access control
│   ├── mocks/               # Mock implementations for testing
│   ├── models/              # Database models
│   ├── repository/          # Data access layer (Interacts with the database)
│   ├── routes/              # Route definitions for the application
│   ├── services/            # Business logic services
│   ├── utils/               # Helper functions & utilities
│── scripts/
│   └── seed.go              # Seeder script for database initialization
│── .env                     # Environment variables
│── go.mod                   # Go module dependencies
│── go.sum                   # Dependency checksums
│── Dockerfile               # Docker setup for application
│── docker-compose.yml       # Docker Compose setup
│── README.md                # Project documentation
```

---

## ⚙️ Setup Instructions  

### 1️⃣ Clone the Repository  
```sh
git clone https://github.com/EsraaFarhat/Library-Management-GO.git
cd Library-Management-GO
```

### 2️⃣ Configure Environment Variables  

Create a **`.env`** file in the root directory and define the required variables:  

```ini
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library
JWT_SECRET=your-secret-key
```

---

## 🚀 Running the Project  

You can run the project in **two ways**:

### ✅ Option 1: Using Docker  
Ensure **Docker** and **Docker Compose** are installed, then run:

```sh
docker compose up --build
```

This will:  
- Start the **PostgreSQL database**  
- Build and run the **Go application**  
- Automatically seed the database with initial users and books  

> **To stop the containers, press `CTRL+C` or run:**  
> ```sh
> docker compose down
> ```

---

### ✅ Option 2: Running Locally (Without Docker)  

If you prefer to run the application directly on your machine:

1. **Start PostgreSQL** manually  
2. **Load environment variables** from `.env`
3. **Run the application:**
   ```sh
   go run cmd/main.go
   ```

---

## 🔌 API Endpoints  

### 🔑 Authentication  
| Method | Endpoint       | Description                 |
|--------|---------------|-----------------------------|
| `POST` | `/auth/register` | Register a new user         |
| `POST` | `/auth/login`    | Authenticate & get JWT      |

### 👥 Users  
| Method | Endpoint       | Description                 | Access  |
|--------|---------------|-----------------------------|---------|
| `POST` | `/users/`     | Create a new user           | Admin   |
| `GET`  | `/users/`     | Get all users               | Admin   |
| `GET`  | `/users/:id`  | Get a specific user         | Admin   |
| `PUT`  | `/users/:id`  | Update a user               | Admin   |
| `DELETE` | `/users/:id` | Delete a user              | Admin   |

### 📚 Books  
| Method | Endpoint       | Description                 | Access |
|--------|---------------|-----------------------------|--------|
| `POST` | `/books/`     | Add a new book              | Admin  |
| `GET`  | `/books/`     | List all books              | Public |
| `GET`  | `/books/:id`  | Get details of a book       | Public |
| `PUT`  | `/books/:id`  | Update book details         | Admin  |
| `DELETE` | `/books/:id` | Remove a book              | Admin  |

### 📖 Borrowing  
| Method | Endpoint                 | Description                  | Access |
|--------|---------------------------|------------------------------|--------|
| `POST` | `/borrows/`               | Borrow a book                | Public   |
| `GET`  | `/borrows/`               | Get all borrow records       | Admin  |
| `POST` | `/borrows/return`         | Return a borrowed book       | Public   |
| `GET`  | `/borrows/my-borrows`     | Get borrow records for logged in user | Public   |
| `GET`  | `/borrows/user/:user_id`  | Get borrow records for a user | Admin   |

---

### 🔍 Assumptions and Decisions  

- The **pagination limit** is set to **10 items per page** by default and returns the **first page**.  
- Pagination can be modified using the **query parameters**:  
  - `page` → specifies the page number.  
  - `limit` → specifies the number of items per page.  
- All list endpoints have **default sorting by `created_at` in descending order**.  
---

### 🧪 Testing  

Run unit tests using:  

```sh
go test ./...
```  
- This will execute all unit tests across the project.

- **Unit testing** is applied **only to certain parts of the project**.