## healthcare-Access
Healthcare-Access is a simple Role-Based Access Control (RBAC) app for managing healthcare appointments. It showcases how different users(admins, staff, and patients) have access to specific actions based on their roles.

### Roles

- Admins can assign staff to appointments.

- User (Patient - the default role assigned during registration) can book appointments.

- Staff can view all assigned appointments.

### Technologies Used

- Go (Gin) 
- Postgres
- Bun ORM
- Docker and docker-compose

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
 docker-compose up --build
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
```sh
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

### Example Request

- assign a staff to an appointment
```sh
{
	"user_id" : 6
	"status_" : "assigned"
}	
```

### Example Response
```sh
{
 "appointment_id" : 2,
 "patient_id" : 1,
 "staff_id" : 3,
 "status_" : "assigned"
 "created_at" : 2025-03-05 14:30:15.123456789 +0000 UTC
 "assigned_by" 5
}
```