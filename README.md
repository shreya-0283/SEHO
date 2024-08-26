# SEHO

SEHO (Self-Hosted) is a self-hosted music server/platform that allows you to build and manage your own music library. It reads music files from a specified directory, scrapes their metadata, and stores the information in a Redis database, making it easy to access and manage your music collection.

## Features

- Scans a directory for music files (`.mp3`, `.flac`, `.m4a`).
- Extracts metadata such as title, artist, album, and file path.
- Stores music metadata in a Redis database.
- Easily expandable with additional features such as a web frontend and APIs.

## Project Structure

```
SEHO/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── redis/
│   │   └── client.go
│   └── music/
│       ├── metadata.go
│       └── scanner.go
│
├── go.mod
└── go.sum
```

## Installation and Setup

### Prerequisites

1. **Go**: [Install Go](https://golang.org/dl/).
2. **Redis**: [Install Redis](https://redis.io/download).

### Steps

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/cursoroid/SEHO.git
   cd SEHO
   ```

2. **Initialize Go Modules**:
   ```bash
   go mod tidy
   ```

3. **Set Environment Variables**:
   ```bash
   export REDIS_ADDR="localhost:6379"
   export MUSIC_DIR="/path/to/music/directory"
   ```

4. **Run the Server**:
   ```bash
   go run cmd/server/main.go
   ```

5. **View Data in Redis**:
   Use the Redis CLI to inspect the stored music metadata.
   ```bash
   redis-cli
   keys music:*
   ```

## To-Do List

- **Frontend Development**:
  - Create a web-based frontend to browse and play music from the library.
  - Implement user authentication and music streaming functionality.

- **API Development**:
  - Expose RESTful or GraphQL APIs to interact with the music library.
  - Implement endpoints for querying music metadata and streaming music files.

- **Dockerization**:
  - Dockerize the application to make deployment easier.
  - Create a `Dockerfile` and `docker-compose.yml` to manage dependencies and deployment.

- **Additional Features**:
  - Add support for more music file formats.
  - Implement search functionality within the music library.
  - Create a mobile-friendly version of the frontend.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an Issue for any bugs or feature requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries, please contact [prathameshmudgale@gmail.com](mailto:prathameshmudgale@gmail.com).
```

### Explanation

- **Features**: Describes the current capabilities of SEHO.
- **Project Structure**: Outlines the organization of the project files and directories.
- **Installation and Setup**: Provides detailed instructions on how to set up and run the project.
- **To-Do List**: Lists future enhancements, including frontend development, API creation, Dockerization, and additional features.
- **Contributing**: Encourages others to contribute to the project.
- **License**: Notes that the project is under the MIT License.
- **Contact**: Provides a contact email for further inquiries.

Feel free to customize the contact information and repository URL as needed.
