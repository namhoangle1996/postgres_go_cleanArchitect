Available Scripts
In the project directory, you can run:

**go run main.go**

Then, we have CRUD request using "curl" :

  _Fetch all_
  
 **curl localhost:9090/users**    
 
 _Add Users_ 
 
 **curl -d '{"email":"namnd@gmail.com","password":"123456"}' -H "Content-Type: application/json" -X POST  localhost:9090/users**

 _Get user by Id_
 
 **curl  POST localhost:9090/users/4**
 
 _DELETE By ID_
 
**curl -X DELETE  localhost:9090/users/18**

