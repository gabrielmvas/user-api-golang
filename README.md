Basic CRUD operations with MongoDB for study.

## Run the project
Start a MongoDB instance using docker:
```sh
docker run --name mongodb -d -p 27017:27017 mongo
```
Clone the repository:
```sh
git clone git@github.com:gabrielmvas/user-api-golang.git
```
Change the current directory to the repository:
```sh
cd user-api-golang
```
Install the dependencies:
```sh
go get ./...
```
Finally, run the app on port `9080`:
```sh
go run .
```

## Endpoints:
```sh
GET    /users
GET    /users/:email
POST   /users
PUT    /users/:email
DELETE /users/:email
```

### Get User
This endpoint retrieves a user given the email.  
Send a `GET` request to `/users/:email`:
```sh
curl -X GET 'http://127.0.0.1:9080/users/test@test.com'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "first_name": "Test",
    "last_name": "User",
    "email": "test@test.com",
    "password": "testpassword"
  }
}
```
### Create User
This endpoint inserts a document in the `users` collection of the `users` database.  
Send a `POST` request to `/users`:
```sh
curl -X POST 'http://127.0.0.1:9080/users' -H "Content-Type: application/json" -d '{"first_name": "Test", "last_name": "User" "email": "test@test.com", "password": "testpassword"}'
```
Response:  
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "TestUser",
    "email": "test@test.com",
    "password": "testpassword"
  }
}
```
### Update User
This endpoint updates the provided fields within the specified document filtered by email.  
Send a `PUT` request to `/users/:email`:
```sh
curl -X PUT 'http://127.0.0.1:9080/users/test@test.com' -H "Content-Type: application/json" -d '{"password": "testpassword"}'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "first_name": "Test",
    "last_name": "User",
    "email": "test@test.com",
    "password": "testpassword"
  }
}
```

### Delete User
This endpoint deletes the user from database given the email.  
Send a `DELETE` request to `/users/:email`:
```sh
curl -X DELETE 'http://127.0.0.1:9080/users/test@test.com'
```
Response:
```sh
{}
```

### Errors
All of the endpoints return an error in json format with a proper http status code, if something goes wrong:
```sh
{
  "error": "User not found"
}
```
