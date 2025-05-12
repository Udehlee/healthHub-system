## healthcare-Access
Healthcare-Access is a simple role-based access control (RBAC) app that shows how users can only perform specific tasks based on their assigned roles. During registration, a default role(patient) is assigned if not specified .

### Roles

### Admin

- assigns a staff to a specific appointment
- view all users and thier roles

### Staff

- can view all assigned appointments

### User (Patient)

- can book appointment

### Technologies Used

- Go (Gin) 
- Postgres
- Bun ORM

### Setup and installation

- Clone the repository:

```sh 
git clone https://github.com/Udehlee/healthcare-Access.git 
```
```sh
cd healthcare-Access
 ```
- Install dependencies 
```sh
go mod tidy
```

- Create .env file and fill it with your credentials as shown in the .env.example

- Start the application with
 ```sh
go run main.go
```
The sever is listening on http://localhost:8000

 starting the application will apply  the following migration files in the db/migrations folder and create:

- a users table to store all users(admin,patient,staff)
- an appointments table store all appointments
- an added role column to users table

### Api Endpoints

- admin
```sh
PATCH admin/appointments/:id
GET admin/users
```

- staff
```sh
GET staff/appointments/assigned

```

- user(patient)
```sh
POST /appointments/book-appointment
```

