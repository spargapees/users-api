# Users API

This project provides a backend service to manage user creation, retrieval, and updates. The application uses Docker for containerization and includes migrations for database setup.

---

## **Instructions**

### **Build and Run the Application**

To build and run the application, use the following command:

```bash
docker-compose up --build
```

This command will:

1. Build the Docker images as defined in the `docker-compose.yml` file.
2. Start the application and its services (e.g., database).

---

### **Apply Migrations**

To apply database migrations, run:

```bash
make migrateup
```

This will execute the migration scripts to set up or update the database schema.

---

### **Endpoints**

#### **Create User**

To create a new user, send a POST request to:

```bash
POST localhost:8080/users
```

With the following JSON payload:

```json
{
    "name": "testName",
    "surname": "testSurname",
    "email": "test@gmail.com",
    "password": "testest"
}
```

#### **Get All Users**

To retrieve all users, send a GET request to:

```bash
GET localhost:8080/users
```

#### **Update User by ID**

To update a user by their ID, send a PUT request to:

```bash
PUT localhost:8080/users/{id}
```

Replace `{id}` with the user’s ID (e.g., `1`).

##### **Example Payloads:**

**Update multiple fields:**

```json
{
    "name": "newName",
    "surname": "newSurname"
}
```

**Update a single field:**

```json
{
    "name": "newName"
}
```

---

### **Notes**

- Ensure the application is running before making API requests.
- Use tools like `curl`, Postman, or any HTTP client to interact with the API.
- The application validates inputs and ensures data consistency (e.g., unique email addresses).

---

### **Common Commands**

- **Stop and Remove Containers:**
  ```bash
  docker-compose down
  ```
- **Rebuild Containers:**
  ```bash
  docker-compose up --build
  ```
- **Clean Up Resources:**
  ```bash
  docker system prune -a
  ```

