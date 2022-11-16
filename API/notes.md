API notes
===

What is an API
---

Application Programming Interface (API)  
Java Script Object Notation (JSON)

MVC Framework
---

Model, View, Controller Architecture
![](/API\mvc_1321.webp)
View:  
- Where the clients can talk to the architecture
- Whatever the application displays comes from the view layer

Model:
- Where all the database updates happen
- Notifies Controller that an update has been received

Controller:
- Acts as intermediary between View and Model layer
- Updates the Model and View layers
- Lets the user know a change has occurs in the database and the state of the model has changed

HTTP Server
---

APIs are servers/packaged services that serve something to the client.  
Clients can request what they want to do to the server and if the server has a direct procedure defined, it will return a status code and a response.

HTTP Methods
---

- GET - Retrieval
- POST - Send Data
- PUT - Replace Data
- DELETE - Delete Data

[Postman](https://www.postman.com/) can be used to send HTTP requests