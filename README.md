## Crisnet Task API
A simple task API to manage tasks, you can create, update and delete tasks and .

## Made with Golang

---

## Stacks used
* PostgreSQL
* JWT
* Bcrypt

## Build/Run

> Run
```bash
go run cmd/api/main.go

```

---

> Build
```bash
go build cmd/api/main.go -o /bin
```

> Output expected
* 2026/05/04 23:16:24 The server is running in: http://localhost:8080

## Endpoints
> user
* GET /user/ Get user profile
* DELETE /user/ delete user account

---

> Task
* GET /user/task/ Get All tasks
* DELETE /user/task/ Delete All tasks
* DELETE /user/task/{id} Get epecific task
* PUT /user/task/{id} Update especific task

--- 

> Auth
* POST /auth/login/ To make login
* POST /auth/signup/ To create new account
