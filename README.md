Movie Management API

This is a simple RESTful API built with Go that manages a collection of movies. The API supports basic CRUD (Create, Read, Update, Delete) operations for movie records, allowing users to retrieve a list of movies, get details of a specific movie, add new movies, update existing ones, and delete movies from the collection.
Features

    Retrieve all movies: GET /movies - Returns a JSON array of all movies in the collection.
    Get a specific movie: GET /movies/{id} - Returns details of a movie with the specified ID.
    Create a new movie: POST /movies - Accepts a JSON object representing a movie and adds it to the collection. A random ID is generated for each new movie.
    Update an existing movie: PUT /movies/{id} - Updates the details of a movie with the specified ID.
    Delete a movie: DELETE /movies/{id} - Removes a movie with the specified ID from the collection.

Technologies Used

    Go: The primary programming language for the API.
    Gorilla Mux: A powerful URL router and dispatcher for building Go web applications.
    JSON: Data is exchanged in JSON format, with struct tags ensuring proper serialization and deserialization.

Getting Started

To run the API locally:

    Ensure you have Go installed on your machine.

    Clone the repository:

 git clone https://github.com/yourusername/movies-api.git

Install the Gorilla Mux package:

go get github.com/gorilla/mux

Run the application:

    go run main.go

    The server will start on localhost:8000, and you can access the API using tools like Postman or curl.

