# Auth & Audit 🔐

A modern web application for user authentication and activity trail tracking, built with Go backend and vanilla JavaScript frontend. This project was developed as a hands-on learning experience for **CI/CD practices** and modern web development workflows.

## 🎯 Purpose

This application was built specifically to learn and practice:
- **Continuous Integration (CI)** workflows
- **Continuous Deployment (CD)** pipelines  
- Modern backend development with Go
- Database design and management
- RESTful API development
- Frontend-backend integration

## ✨ Features

### 🔐 Authentication System
- **User Registration** with secure password storage
- **User Login** with credential verification
- **API Key Generation** for authenticated requests
- **Stateless Authentication** using API keys

### 📊 Activity Tracking
- **Real-time Activity Logs** for all user actions
- **Trail Types**: Sign up, Login, Logout events
- **Timestamp Tracking** with precise date/time logging
- **User-specific Trails** - each user sees only their activity

### 🎨 Modern UI
- **Dark Theme** with glassmorphism effects
- **Tab-based Interface** for Sign Up and Login
- **Responsive Design** that works on all devices
- **Smooth Animations** and hover effects
- **Real-time Notifications** for user feedback

## 🛠️ Tech Stack

### Backend
- **Go** - Primary backend language
- **Chi Router** - HTTP router and middleware
- **PostgreSQL** - Database for user and trail data
- **SQLC** - Type-safe SQL code generation
- **Goose** - Database migration management
- **godotenv** - Environment variable management
- **lib/pq** - PostgreSQL driver

### Frontend
- **Vanilla JavaScript** - No frameworks, pure JS
- **HTML5 & CSS3** - Modern web standards
- **Fetch API** - For backend communication
- **Embedded Static Files** - Go embed for serving frontend

### Database
- **PostgreSQL** - Primary database
- **UUID** - For unique identifiers
- **Timestamps** - For precise activity tracking
- **Foreign Keys** - For data integrity

## 📁 Project Structure

```
authtrail/
├── main.go                 # Application entry point
├── handler_user.go         # User authentication handlers
├── handler_authtrail.go    # Activity trail handlers
├── middleware.go           # Authentication middleware
├── json.go                 # JSON response utilities
├── models.go              # Data model conversions
├── static/
│   └── index.html         # Frontend application
├── sql/
│   ├── schema/            # Database migrations
│   │   ├── 001_users.sql
│   │   └── 002_authtrails.sql
│   └── queries/           # SQL queries for SQLC
│       ├── users.sql
│       └── authtrails.sql
├── internal/
│   ├── database/          # Generated SQLC code
│   └── auth/              # Authentication utilities
├── sqlc.yaml              # SQLC configuration
├── go.mod                 # Go module dependencies
└── README.md              # This file
```

## 🚀 Getting Started

### Prerequisites
- Go 1.24+
- PostgreSQL 12+
- SQLC
- Goose (for migrations)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd authtrail
   ```

2. **Set up the database**
   ```bash
   createdb authtrail
   ```

3. **Run migrations**
   ```bash
   goose -dir sql/schema postgres "postgres://localhost:5432/authtrail?sslmode=disable" up
   ```

4. **Generate Go code from SQL**
   ```bash
   sqlc generate
   ```

5. **Install dependencies**
   ```bash
   go mod tidy
   ```

6. **Create environment file**
   ```bash
   echo "PORT=8080" > .env
   echo "DATABASE_URL=postgres://localhost:5432/authtrail?sslmode=disable" >> .env
   ```

7. **Run the application**
   ```bash
   go run .
   ```

8. **Access the application**
   Open http://localhost:8080 in your browser

## 📖 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /v1/user
Content-Type: application/json

{
  "name": "username",
  "password": "password"
}
```

#### Login User
```http
POST /v1/user/login
Content-Type: application/json

{
  "name": "username", 
  "password": "password"
}
```

#### Get User Info
```http
GET /v1/user
Authorization: Bearer <api_key>
```

### Activity Trail Endpoints

#### Get User Activity Trails
```http
GET /v1/{userID}/authtrails
Authorization: Bearer <api_key>
```

## 🎓 Learning Outcomes

Through building this project, I gained hands-on experience with:

### CI/CD Practices
- Setting up automated testing pipelines
- Implementing continuous integration workflows
- Managing deployment strategies
- Environment configuration management
- Database migration automation

### Backend Development
- RESTful API design and implementation
- Database schema design and optimization
- SQL query optimization with SQLC
- Authentication and authorization patterns
- Error handling and logging strategies

### Frontend Development
- Modern vanilla JavaScript techniques
- Responsive design principles
- API integration patterns
- User experience optimization
- State management without frameworks

### DevOps Skills
- Environment variable management
- Static file serving with Go embed
- RESTful API design patterns
- Database connection handling

## 🔧 Configuration

### Environment Variables
- `PORT` - Server port (default: 8080)
- `DATABASE_URL` - PostgreSQL connection string

### Database Schema
The application uses two main tables:
- `users` - User authentication data
- `authtrails` - Activity tracking records

## 🤝 Contributing

This is a learning project, but contributions are welcome! Please feel free to:
- Report bugs
- Suggest improvements
- Share CI/CD best practices
- Contribute to documentation

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- Built as part of my journey to learn modern CI/CD practices
- Inspired by the need to understand full-stack development workflows
- Thanks to the Go and PostgreSQL communities for excellent documentation

---

**Note**: This application was built primarily for educational purposes to understand CI/CD workflows and modern web development practices. It demonstrates real-world patterns and technologies commonly used in production applications.
