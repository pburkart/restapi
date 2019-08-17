# Golang RESTful API

To use simply:

<ol>
	<li>Clone the Repository `git clone https://github.com/pburkart/restapi.git`</li>
	<li>Build the app `go build`</li>
	<li>Run the app `./restapi`</li>
	<li>Use postman to make requests to `localhost:8000/api/items`</li>
<ol>

Example Requests:

<ul>
	<li>Get all items: `GET http://localhost:8000/api/items`</li>
	<li>Get a single item: `GET http://localhost:8000/api/items/{id}</li>
	<li>Create an item: `POST http://localhost:8000/api/items` header `Content-Type: application/json`
	```
	{
		"item":"Item name",
		"price":"200.00",
		"age":"3"
	}
	```
	<li>Update an item: `PUT http://localhost:8000/api/items/{id}` header `Content-Type: application/json`
	```
	{
		"item": "New Item",
		"price": "250.00",
		"age": "4"
	}
	```
	<li>Delete an item: `DELETE http://localhost:8000/api/items/{id}`
</ul>
