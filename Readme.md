# Go Gin URL Shortener Backend

A simple and efficient URL shortener backend built with [Go](https://golang.org/) and [Gin](https://gin-gonic.com/). This service allows users to create short URLs that redirect to original long URLs, with support for basic analytics and RESTful API endpoints.

## Features
- Shorten long URLs to unique short codes
- Redirect short URLs to original destinations
- RESTful API endpoints for creating and retrieving URLs
- In-memory or persistent storage (configurable)
- Basic analytics (click counts, etc.)

## Getting Started

### Prerequisites
- Go 1.18 or higher

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-gin-url-shortener-backend.git
   cd go-gin-url-shortener-backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

## Usage

### API Endpoints

#### Create Short URL
- **POST** `/api/shorten`
- **Body:**
  ```json
  {
    "url": "https://example.com/long-url"
  }
  ```
- **Response:**
  ```json
  {
    "short_url": "http://localhost:8080/abc123"
  }
  ```

#### Redirect to Original URL
- **GET** `/{short_code}`
- Redirects to the original URL.

#### Get URL Analytics (if implemented)
- **GET** `/api/analytics/{short_code}`
- **Response:**
  ```json
  {
    "short_code": "abc123",
    "original_url": "https://example.com/long-url",
    "clicks": 42
  }
  ```

## Configuration
- Edit `config.yaml` or environment variables as needed for storage and server settings.

## Project Structure
- `main.go` - Entry point
- `handlers/` - HTTP handlers
- `models/` - Data models
- `storage/` - Storage logic (in-memory, database, etc.)
- `routes/` - API route definitions

## License
MIT

---

Feel free to contribute or open issues for improvements!