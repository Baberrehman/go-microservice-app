
Developer:  Baber Rehman
Email:      baberrehman26@gmail.com
Date:       23/08/2017
Modifier:

How to run:
run the server.go file.

Command:
go run server.go

Requirement:
Port: 8811
Database: MongoDB

About:
This is a test application. It exposes rest API's for CRUD operations on students
and teachers collection. This application follows a standard MVC model.

All the methods and structs related to teachers collection are in teacher package,
where as methods and structs related to students collection are in student package.

 If you want to modify the port, you can modify it in server.go file.

 If you wish to update the mongoDB, edit the file mongodb.go in mongo package.

 Endpoints for student:
 /student/add POST
 /student/delete/{id} DELETE
 /student/update/{id} PUT
 /student/get/{id} GET

 Sample to get a student:
 http://<localhost>:8811/student/get/1

 Endpoints for teacher:
 /teacher/add POST
 /teacher/delete/{id} DELETE
 /teacher/update/{id} PUT
 /teacher/get/{id} GET

 Sample to get a teacher:
  http://<localhost>:8811/teacher/get/1

*All endpoints are supported over HTTP requests.
*Please update mongodb.go in mongo package with the IP and credentials of your mongo.