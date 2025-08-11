# SimpleForum API

A simple REST API for a forum platform built with Go. This project is a personal learning journey into backend development with the Go programming language.

## About The Project

SimpleForum API provides the backend services for a forum where users can create posts and comments. It uses JWT for authentication to secure its endpoints and connects to a MySQL database for data persistence. This is a pure API project without any frontend interface.

### Features

-   **User Authentication:** Secure sign-up and login using JSON Web Tokens (JWT).
-   **Posts:**
    -   Create a new post.
    -   View all posts.
-   **Comments:**
    -   Add a comment to a post.
    -   View all comments for a specific post.
-   **(Upcoming) Likes:**
    -   Like/Unlike a post.

### Built With

-   [Go (Golang)](https://golang.org/)
-   [Gin Web Framework](https://github.com/gin-gonic/gin)
-   [GORM (Go-ORM)](https://gorm.io/)
-   [MySQL](https://www.mysql.com/)
-   [JWT for Go](https://github.com/golang-jwt/jwt)

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

Make sure you have Go and MySQL installed on your machine.

-   [Go](https://golang.org/doc/install)
-   [MySQL](https://dev.mysql.com/downloads/installer/)

### Installation

1.  **Clone the repo**
    ```sh
    git clone [https://github.com/mhmmdFabiann/SimpleForum.git](https://github.com/mhmmdFabiann/SimpleForum.git)
    ```
2.  **Navigate to the project directory**
    ```sh
    cd SimpleForum
    ```
3.  **Install Go modules**
    ```sh
    go mod tidy
    ```
4.  **Database Setup**
    -   Create a new database in MySQL.
    -   Update the database connection string in the `database/database.go` file with your database credentials.
      ```go
      dsn := "username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
      ```

5.  **Run the application**
    ```sh
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

## API Endpoints

Here are the available API endpoints:

#### Authentication
-   `POST /register` - Register a new user.
-   `POST /login` - Login a user and get a JWT token.

#### Posts
-   `GET /posts` - Get all posts.
-   `POST /posts` - Create a new post (Authentication required).

#### Comments
-   `GET /posts/:postId/comments` - Get all comments for a specific post.
-   `POST /posts/:postId/comments` - Add a new comment to a post (Authentication required).


## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Muhammad Fabian - [@mhmmdfabian_](https://www.instagram.com/mhmmdfabian_/)

Project Link: [https://github.com/mhmmdFabiann/SimpleForum](https://github.com/mhmmdFabiann/SimpleForum)
