# Golang Community App API Documentation

## Overview
This repository provides an API backend built using the **Golang Echo framework**. It serves as the backend for a community-based application where users can create posts, leave comments, and like posts. The system also supports pagination, sorting, and user authentication via middleware.

### Features:
- **Post Management**: Create, retrieve, and delete posts.
- **Comments**: Add, view, and delete comments on posts.
- **Likes**: Add and remove likes on posts.
- **Pagination and Sorting**: Retrieve posts with customizable sorting and pagination.
- **Authentication**: Secure API endpoints using JWT middleware (Cognito).

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd golang-community-app
   ```
2. **Install dependencies**:
   ```bash
   go mod tidy
   ```
3. **Set up the database**:
    - Ensure PostgreSQL is running.
    - Configure the database connection in `config`.

4. **Run the server**:
   ```bash
   go run main.go
   ```

The server will start on **http://localhost:8080**.

---

## API Endpoints

### Post Management

#### 1. Get Paginated Posts
**Endpoint**: `GET /posts`

**Description**: Retrieve a list of posts with pagination and sorting options.

**Query Parameters**:
| Parameter   | Type   | Required | Default | Description                         |
|-------------|--------|----------|---------|-------------------------------------|
| `limit`     | int    | No       | 10      | Number of posts to retrieve.        |
| `offset`    | int    | No       | 0       | Starting point for pagination.      |
| `sortType`  | string | No       | trend   | Sort by `trend` or `latest`.        |
| `sort`      | string | No       | desc    | Sort order: `asc` or `desc`.        |

**Response**:
```json
[
    {
        "id": 1,
        "title": "First Post",
        "content": "This is the content of the post",
        "views": 5,
        "author": "username",
        "like_count": 10,
        "comment_count": 3,
        "created_at": "2 hours ago"
    }
]
```

#### 2. Get Post Details
**Endpoint**: `GET /posts/:id`

**Description**: Retrieve a single post's details along with its comments.

**Response**:
```json
{
    "post": {
        "id": 1,
        "title": "First Post",
        "content": "This is the content of the post",
        "views": 6,
        "author": "username",
        "like_count": 10,
        "comment_count": 3,
        "created_at": "2 hours ago"
    },
    "comments": [
        {
            "id": 1,
            "username": "commenter",
            "content": "Great post!",
            "created_at": "1 hour ago"
        }
    ]
}
```

#### 3. Create Post
**Endpoint**: `POST /posts`

**Description**: Create a new post (authenticated users only).

**Request Body**:
```json
{
    "title": "New Post Title",
    "content": "This is the post content."
}
```

**Headers**:
| Key         | Value       |
|-------------|-------------|
| `id_token`  | JWT Token   |

**Response**:
```json
{
    "id": 2,
    "title": "New Post Title",
    "content": "This is the post content.",
    "views": 0,
    "author": "username",
    "created_at": "just now"
}
```

#### 4. Delete Post
**Endpoint**: `DELETE /posts/:id`

**Description**: Delete a post by its ID (authenticated users only).

**Headers**:
| Key         | Value       |
|-------------|-------------|
| `id_token`  | JWT Token   |

**Response**:
```json
{
    "message": "Post deleted successfully"
}
```

---

### Comments

#### 5. Get Comments by Post ID
**Endpoint**: `GET /comments/post/:post_id`

**Description**: Retrieve all comments for a specific post.

**Response**:
```json
{
    "comment_count": 5
}
```

#### 6. Create Comment
**Endpoint**: `POST /comments/post/:post_id`

**Description**: Add a comment to a post (authenticated users only).

**Request Body**:
```json
{
    "content": "This is a comment."
}
```

**Headers**:
| Key         | Value       |
|-------------|-------------|
| `id_token`  | JWT Token   |

**Response**:
```json
{
    "id": 1,
    "username": "username",
    "content": "This is a comment.",
    "created_at": "just now"
}
```

#### 7. Delete Comment
**Endpoint**: `DELETE /comments/:id`

**Description**: Delete a comment by its ID (authenticated users only).

**Response**:
```json
{
    "message": "Comment deleted successfully"
}
```

---

### Likes

#### 8. Like a Post
**Endpoint**: `POST /likes/post/:post_id`

**Description**: Like a post (authenticated users only).

**Response**:
```json
{
    "message": "Post liked successfully"
}
```

#### 9. Unlike a Post
**Endpoint**: `DELETE /likes/post/:post_id`

**Description**: Unlike a post (authenticated users only).

**Response**:
```json
{
    "message": "Post unliked successfully"
}
```

#### 10. Get Post Likes Count
**Endpoint**: `GET /likes/post/:post_id`

**Description**: Retrieve the total likes count for a post.

**Response**:
```json
{
    "like_count": 15
}
```

---

## Middleware
- **CognitoJWTMiddleware**: Secures endpoints by verifying the `id_token` JWT header.

---

## Error Codes
| Status Code | Description                       |
|-------------|-----------------------------------|
| 400         | Bad Request - Invalid parameters  |
| 401         | Unauthorized - Missing/Invalid JWT|
| 404         | Not Found - Resource doesn't exist|
| 500         | Internal Server Error             |

---

## Running Tests
To run the tests:
```bash
go test ./...
```

---

## Contributing
1. Fork the repository.
2. Create a new branch.
3. Commit your changes.
4. Submit a pull request.

---

## License
This project is licensed under the MIT License.

---

## Contact
For any inquiries or issues, contact:
- **Email**: example@example.com
- **GitHub**: [YourGitHubProfile](https://github.com/YourGitHubProfile)
