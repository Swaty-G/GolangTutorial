package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Swaty-G/GolangTutorial/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://swaty:swatyrules@cluster0.h5dcfbv.mongodb.net/" //connectionString is a constant that holds the connection string to the MongoDB database
const dbName = "netflix"                                                                //dbName is a constant that holds the name of the database
const collectionName = "watchlist"                                                      //collectionName is a constant that holds the name of the collection

// MOST IMPORTANT
var collection *mongo.Collection //collection is a variable of type *mongo.Collection; it is a pointer to a mongo.Collection object from the mongo-driver package that represents a collection in a MongoDB database and is used to perform CRUD operations on the collection in the database (create, read, update, delete)

// connect with mongoDB
func init() { //init() is a function that is called automatically before the main function is executed when the program is run and is used to initialize the application by connecting to the MongoDB database
	// client options
	clientOption := options.Client().ApplyURI(connectionString) //clientOptions is a variable of type *options.ClientOptions; it is a pointer to a options.ClientOptions object from the mongo-driver package that holds the options for the MongoDB client; options.Client() returns a new options.ClientOptions object; ApplyURI(connectionString) sets the connection string for the MongoDB client // says this is the db I will be using

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption) //mongo.Connect() establishes a connection to the MongoDB database; context.TODO() returns a non-nil, empty Context; clientOptions is the options for the MongoDB client

	if err != nil {
		log.Fatal(err) //log.Fatal() logs the error message and exits the program
	}
	fmt.Println("MongoDB Connection Made") //prints a message to the console indicating that the connection to the MongoDB database was successful

	collection = client.Database(dbName).Collection(collectionName) //client.Database(dbName) returns a pointer to a mongo.Database object from the mongo-driver package that represents a database in a MongoDB database; Collection(collectionName) returns a pointer to a mongo.Collection object from the mongo-driver package that represents a collection in a MongoDB database; collection is assigned the value of the collection in the database // connect to the database and get the collection // says this is the collection I will be using

	//collection instance
	fmt.Println("Collection instance is ready") //prints a message to the console indicating that the collection instance was created
}

// MONGODB helpers - file

// insert 1 record
func insertOneMovie(movie model.Netflix) { //insertOneMovie() is a function that takes a Netflix object as a parameter and inserts the object into the MongoDB database; it is used to insert a single document into the collection in the database; movie is a parameter of type Netflix
	inserted, err := collection.InsertOne(context.Background(), movie) // collection.InsertOne() inserts a single document into the collection in the MongoDB database; context.Background() returns a non-nil, empty Context
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieID string) (*mongo.UpdateResult, error) {
	id, err := primitive.ObjectIDFromHex(movieID) //primitive.ObjectIDFromHex() converts a hexadecimal string to a primitive.ObjectID; movieID is the hexadecimal string that represents the unique identifier of the document in the MongoDB database
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}                                                     //bson.M is a type that represents a BSON document as a map; filter is a variable of type bson.M that holds the filter criteria for the update operation; {"_id": id} is the filter criteria to find the document with the specified _id field
	update := bson.M{"$set": bson.M{"watched": true}}                               //update is a variable of type bson.M that holds the update operation to be performed on the document; {"$set": bson.M{"watched": true}} is the update operation to set the watched field to true
	updateResult, err := collection.UpdateOne(context.Background(), filter, update) //collection.UpdateOne() updates a single document in the collection in the MongoDB database; context.Background() returns a non-nil, empty Context; filter is the filter criteria to find the document to update; update is the update operation to be performed on the document
	if err != nil {
		log.Fatal(err)
	}
	if updateResult.ModifiedCount == 0 { //if the movie is already watched or doesn't exist in the database; ModifiedCount is the number of documents modified by the update operation
		return updateResult, err
	}
	fmt.Println("modified count: ", updateResult.ModifiedCount)
	return updateResult, err
}

// delete 1 record
func deleteOneMovie(movieID string) (int64, error) {
	id, err := primitive.ObjectIDFromHex(movieID) //primitive.ObjectIDFromHex() converts a hexadecimal string to a primitive.ObjectID; movieID is the hexadecimal string that represents the unique identifier of the document in the MongoDB database
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}                                            //bson.M is a type that represents a BSON document as a map; filter is a variable of type bson.M that holds the filter criteria for the delete operation; {"_id": id} is the filter criteria to find the document with the specified _id field
	deleteCount, err := collection.DeleteOne(context.Background(), filter) //collection.DeleteOne() deletes a single document from the collection in the MongoDB database; context.Background() returns a non-nil, empty Context; filter is the filter criteria to find the document to delete
	if err != nil {
		log.Fatal(err)
	}
	if deleteCount == nil {
		fmt.Println("No movie found with the given ID")
	} else {
		fmt.Println("Movie got deleted with count: ", deleteCount)
	}
	return deleteCount.DeletedCount, err
}

// delete all records from mongoDB
func deleteAllMovies() (int64, error) {
	deleteAll, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil) //collection.DeleteMany() deletes multiple documents from the collection in the MongoDB database; context.Background() returns a non-nil, empty Context; bson.M{} is an empty BSON document; bson.M{} is the filter criteria to delete all documents in the collection
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", deleteAll.DeletedCount)
	return deleteAll.DeletedCount, err
}

// get all movies from mongoDB
func getAllMovies() []primitive.M { //getAllMovies() is a function that returns all the documents from the collection in the MongoDB database; it is used to retrieve all the documents from the collection
	cursor, err := collection.Find(context.Background(), bson.D{{}}) //collection.Find() returns a cursor to the documents that match the filter criteria in the collection in the MongoDB database; context.Background() returns a non-nil, empty Context; bson.M{} is an empty BSON document; bson.M{} is the filter criteria to find all documents in the collection
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M                //movies is a variable of type []primitive.M that holds the documents retrieved from the collection; []primitive.M is a slice of primitive.M objects that represent the documents in the collection in the MongoDB database // slice is a data structure that holds a sequence of elements of the same type and is used to store a collection of items of the same type in Go // primitive.M is a type that represents a BSON document as a map
	for cursor.Next(context.Background()) { //cursor.Next() moves the cursor to the next document in the result set
		var movie primitive.M       //movie is a variable of type primitive.M that holds the document retrieved from the collection;
		err = cursor.Decode(&movie) //cursor.Decode() decodes the current document in the cursor into the movie variable; &movie is a pointer to the movie variable that holds the document retrieved from the collection
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie) //append() adds the movie document to the movies slice
	}
	defer cursor.Close(context.Background()) //cursor.Close() closes the cursor
	return movies
}

//Actual controller - file

func GetMyAllMovies(writer http.ResponseWriter, request *http.Request) { //GetMyAllMovies() is a function that returns all the documents from the collection in the MongoDB database; it is a controller function that is called when a GET request is made to the /movies endpoint
	writer.Header().Set("Content-Type", "application/x-www-form-urlencoded") //or application/json
	allMovies := getAllMovies()
	json.NewEncoder(writer).Encode(allMovies)
}

func CreateMovie(writer http.ResponseWriter, request *http.Request) { //CreateMovie() is a function that inserts a document into the collection in the MongoDB database; it is a controller function that is called when a POST request is made to the /movie endpoint
	writer.Header().Set("Content-Type", "application/x-www-form-urlencoded") //or application/json
	writer.Header().Set("Access-Control-Allow-Methods", "POST")              //access control allow methods is a header that specifies the methods allowed when accessing the resource in response to a preflight request // POST is the method allowed when accessing the resource
	var movie model.Netflix                                                  //movie is a variable of type Netflix that holds the Netflix object
	err := json.NewDecoder(request.Body).Decode(&movie)                      //json.NewDecoder() creates a new decoder that reads from the request body; Decode() decodes the JSON data from the request body into the movie variable; &movie is a pointer to the movie variable
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("Error decoding the request body")
		return
	}

	// check if the movie already exists
	filter := bson.M{"movie": movie.Movie} //filter is a variable of type bson.M that holds the filter criteria to find the document with the specified movie field; {"movie": movie.Movie} is the filter criteria to find the document with the specified movie field
	var existingMovie model.Netflix
	err = collection.FindOne(context.Background(), filter).Decode(&existingMovie) //collection.FindOne() returns a single document that matches the filter criteria in the collection in the MongoDB database; context.Background() returns a non-nil, empty Context; filter is the filter criteria to find the document; Decode() decodes the document into the existingMovie variable
	if existingMovie.Movie != "" {                                                //if the movie already exists in the database
		writer.WriteHeader(http.StatusConflict) //writer.WriteHeader() writes the HTTP status code to the response; http.StatusConflict is the status code for a conflict
		json.NewEncoder(writer).Encode("Movie already exists")
		return
	}

	// insert the movie into the database if it does not exist already in the database
	insertOneMovie(movie)
	json.NewEncoder(writer).Encode(movie)
}

// update a movie to mark as watched
func MarkAsWatched(writer http.ResponseWriter, request *http.Request) { //MarkAsWatched() is a function that updates a document in the collection in the MongoDB database; it is a controller function that is called when a PUT request is made to the /movie/{id} endpoint
	writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(request) //mux.Vars() returns the route variables for the current request; params is a map that holds the route variables
	movieID := params["id"]     //movieID is a variable that holds the value of the id route variable from the request URL
	if movieID == "" {          //if the movieID is an empty string
		writer.WriteHeader(http.StatusBadRequest) //writer.WriteHeader() writes the HTTP status code to the response; http.StatusBadRequest is the status code for a bad request
		json.NewEncoder(writer).Encode("Invalid movie ID")
		return
	}
	modifiedResult, err := updateOneMovie(movieID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Error updating the movie")
		return
	}
	if modifiedResult.MatchedCount == 0 { //if the movie is not found; MatchedCount is the number of documents matched by the update operation
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("No movie found with the given ID")
		return
	}
	if modifiedResult.ModifiedCount == 0 { //if the movie is already watched; ModifiedCount is the number of documents modified by the update operation
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("Movie already watched")
	} else { //if the movie is not watched
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(movieID + " marked as watched")
	}
}

func DeleteAMovie(writer http.ResponseWriter, request *http.Request) { //DeleteAMovie() is a function that deletes a document from the collection in the MongoDB database; it is a controller function that is called when a DELETE request is made to the /movie/{id} endpoint
	writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(request)
	movieID := params["id"]

	//if id is not found, it will return a message
	if movieID == "" { //if the movieID is an empty string
		writer.WriteHeader(http.StatusBadRequest) //writer.WriteHeader() writes the HTTP status code to the response; http.StatusBadRequest is the status code for a bad request
		json.NewEncoder(writer).Encode("Invalid movie ID")
		return
	}
	deleteResult, err := deleteOneMovie(movieID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Error deleting the movie")
		return
	}
	//if movie id is not found, it will return a message
	if deleteResult == 0 {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("No movie found with the given ID")
		return
	} else { //if movie id is found, it will delete the movie
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(movieID + " deleted")
	}
}

func DeleteAllMovies(writer http.ResponseWriter, request *http.Request) { //DeleteAllMovies() is a function that deletes all documents from the collection in the MongoDB database; it is a controller function that is called when a DELETE request is made to the /movies endpoint
	writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	writer.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deleteCount, err := deleteAllMovies()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Error deleting all movies")
		return
	}
	if deleteCount == 0 { //if no movies are found to delete
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode("No movies found to delete")
		return
	} else { //if movies are found, it will delete all movies
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(deleteCount)
	}
}
