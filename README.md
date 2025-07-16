
## 🚀 Simple User Creation API in Go (Echo Framework)

This is a beginner-friendly REST API built in Go using the [Echo web framework](https://echo.labstack.com/). It allows you to:

* Create a user via POST request
* Validate user input (name, email, password)
* Hash the password securely using bcrypt
* Save the user details to a local JSON file
* Retrieve user details from a file via GET request

---

### 📦 Features

* ✅ POST `/creat` – Create a new user
* ✅ GET `/:filename` – Get user info from a file
* 🔐 Passwords are hashed with `bcrypt`
* 📁 File-based user storage (`os.WriteFile`)
* 📩 Email format validation (`net/mail`)

---

### 📂 Project Structure

```
main.go          # Entry point
user-json        # Created JSON file with user data
```

---

### 🧪 API Usage

#### ✅ Create a User

```
POST http://localhost:8080/creat
Content-Type: application/json

Request Body:
{
  "user_name": "bhushan",
  "user_email": "bhushan@example.com",
  "user_password": "secret123"
}
```

**Success Response:**

```json
{
  "message": "user created successfully"
}
```

---

#### 📄 Read User from File

```
GET http://localhost:8080/user-json
```

**Success Response:**

```json
{
  "data": {
    "user_name": "John Doe",
    "user_email": "john@example.com",
    "user_password": "$2a$10$...."
  }
}
```

---

### 🔧 Run Locally

1. **Clone the repository**

```bash
git clone https://github.com/bhushan-aruto/user-api-go.git
cd user-api-go
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Run the server**

```bash
go run main.go
```

4. **Test using Postman or curl**

---

### 📚 Dependencies

* [Echo v4](https://github.com/labstack/echo)
* [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
* [Go standard libraries: `os`, `json`, `mail`](https://pkg.go.dev/std)

---

### ✅ Future Improvements

* Store users in a DB (e.g., SQLite or PostgreSQL)
* Add login & JWT authentication
* Add validation using middleware
* Encrypt files instead of plain JSON

---