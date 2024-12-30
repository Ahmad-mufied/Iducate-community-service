# iducate-community-service

A robust REST API backend built with Go (Golang) and Echo framework for managing community interactions including posts, comments, and likes. Features secure authentication via AWS Cognito JWT.

## 🌟 Features

- **Post Management**
  - Create and retrieve posts
  - View tracking
  - Pagination and sorting
  - Delete functionality
  
- **Interactive Features**
  - Comment system
  - Like/Unlike functionality
  - Real-time counters
  
- **Security & Performance**
  - AWS Cognito JWT Authentication
  - Input validation
  - PostgreSQL database
  - Docker containerization
  - Graceful shutdown

## 🚀 Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL
- Docker & Docker Compose (optional)
- AWS Cognito setup

### Installation

1. Clone the repository
```bash
git clone https://github.com/Ahmad-mufied/iducate-community-service
cd iducate-community-service
```

2. Install dependencies
```bash
go mod tidy
```

3. Configure environment
Create `.env` file:
```env
APP_ENV=development
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_dbname
```

4. Run the application

Local development:
```bash
go run cmd/main.go
```

Using Docker:
```bash
docker-compose up --build
```

## 📚 API Documentation

### Authentication
All protected endpoints require a valid JWT token:
```
Authorization: Bearer <your_jwt_token>
```

### Posts Endpoints

#### Get All Posts
```http
GET /posts

Query Parameters:
- limit (int): Number of posts per page (default: 10)
- offset (int): Number of posts to skip (default: 0)
- sortType (string): Sort by ["trend", "latest"] (default: "latest")
- sort (string): Sort direction ["asc", "desc"] (default: "desc")

Response: 200 OK
{
    "posts": [
        {
            "id": 1,
            "title": "Post Title",
            "content": "Post content",
            "views": 10,
            "author": "John Doe",
            "like_count": 5,
            "comment_count": 3,
            "created_at": "17 hours ago"
        }
    ]
}
```

#### Get Post Detail
```http
GET /posts/:id

Response: 200 OK
{
    "id": 1,
    "title": "Post Title",
    "content": "Post content",
    "views": 10,
    "author": "John Doe",
    "like_count": 5,
    "comment_count": 3,
    "created_at": "17 hours ago",
    "comments": [
        {
            "id": 1,
            "username": "Jane Doe",
            "content": "Great post!",
            "created_at": "17 hours ago"
        }
    ]
}
```

#### Create Post
```http
POST /posts
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
id_token: <your_id_token>

Request Body:
{
    "title": "Post Title",
    "content": "Post content"
}

Response: 201 Created
{
    "id": 1,
    "title": "Post Title",
    "content": "Post content",
    "author": "John Doe",
    "created_at": "in about a second"
}
```

#### Delete Post
```http
DELETE /posts/:id
Authorization: Bearer <your_jwt_token>
id_token: <your_id_token>

Response: 200 Ok
{
    "message": "Post deleted successfully"
}
```

### Comments Endpoints

#### Get Post Comments
```http
GET /comments/post/:post_id

Response: 200 OK
{
    "comments": [
        {
            "id": 1,
            "username": "Jane Doe",
            "content": "Comment content",
            "created_at": "2024-03-20T16:00:00Z"
        }
    ]
}
```

#### Create Comment
```http
POST /comments/post/:post_id
Authorization: Bearer <your_jwt_token>
Content-Type: application/json
id_token: <your_id_token>

Request Body:
{
    "content": "Comment content"
}

Response: 201 Created
{
    "id": 1,
    "username": "Jane Doe",
    "content": "Comment content",
    "created_at": "in about a second"
}
```

#### Delete Comment
```http
DELETE /comments/:id
Authorization: Bearer <your_jwt_token>
id_token: <your_id_token>

Response: 200 Ok
{
    "message": "Comment deleted successfully"
}
```

### Likes Endpoints

#### Get Post Likes Count
```http
GET /likes/post/:post_id

Response: 200 OK
{
    "like_count": 5
}
```

#### Like Post
```http
POST /likes/post/:post_id
Authorization: Bearer <your_jwt_token>
id_token: <your_id_token>

Response: 200 OK
{
    "message": "Post liked successfully",
    "like_count": 6
}
```

#### Unlike Post
```http
DELETE /likes/post/:post_id
Authorization: Bearer <your_jwt_token>
id_token: <your_id_token>

Response: 200 OK
{
    "message": "Post unliked successfully",
    "like_count": 5
}
```

## 🔧 Development

### Database Migrations

The database schema is managed through SQL files in the `sql/` directory:
- `DDL.sql`: Contains table definitions
- `Seed.sql`: Contains sample data for development

### Hot Reloading

For development, use Air for hot reloading:
```bash
air
```

## 🐳 Docker Deployment

### Development
```bash
docker-compose up --build
```

### Production
```bash
docker-compose -f docker-compose-prod.yml up -d
```

## 🧪 Testing

Run tests:
```bash
go test ./...
```

## 📁 Project Structure
```
.
├── cmd/                    # Application entrypoint
├── config/                 # Configuration
├── constants/             # Global constants
├── data/                  # Data models and DB operations
├── server/                # HTTP server setup
│   ├── handler/           # Request handlers
│   └── middlewares/       # Custom middlewares
├── utils/                 # Utility functions
├── sql/                   # SQL migrations
├── Dockerfile             # Docker configuration
├── docker-compose.yml     # Docker Compose dev config
└── docker-compose-prod.yml # Docker Compose prod config
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

For any inquiries or issues, contact:
- **GitHub**: [Ahmad-mufied](https://github.com/Ahmad-mufied)

