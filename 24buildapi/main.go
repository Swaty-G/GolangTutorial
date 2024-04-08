package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Model for course - file
type Course struct { // Course struct to represent a course with its details like course ID, name, price, and author
	CourseId    string  `json:"courseid"`   // CourseId field of type string to store the unique identifier of the course; courseid is the key field for the course in the database
	CourseName  string  `json:"coursename"` // CourseName field of type string to store the name of the course
	CoursePrice int     `json:"price"`      // CoursePrice field of type int to store the price of the course ; `json:"-"` is used to ignore the field in the JSON response; price will not be shown in response
	Author      *Author `json:"author"`     // Author field of type Author struct to store the details of the author of the course like fullname and website
}

type Author struct { // Author struct to represent the author of a course with their details like fullname and website
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course // courses slice of Course struct to store the list of courses in the fake database

// middleware, helper functions - file
func (c *Course) isEmpty() bool { // isEmpty method on Course struct to check if the course is empty or not based on the course ID and name fields; *
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //Id will be generated by the server not the client
}

func main() {
	fmt.Println("API - Build RESTful APIs in Go")
	r := mux.NewRouter() // creating a new router using the NewRouter function from the mux package to handle the incoming HTTP requests and route them to the appropriate handler functions based on the request URL and method

	// seed data
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "John Doe", Website: "https://johndoe.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Jane Doe", Website: "https://janedoe.com"}})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")                     // registering a handler function serveHome for the home page URL "/" with the GET method using the HandleFunc method of the router r from the mux package
	r.HandleFunc("/courses", getAllCourses).Methods("GET")          // registering a handler function getAllCourses for the "/courses" URL with the GET method using the HandleFunc method of the router r from the mux package
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")       // registering a handler function getOneCourse for the "/courses/{id}" URL with the GET method using the HandleFunc method of the router r from the mux package
	r.HandleFunc("/course", createOneCourse).Methods("POST")        // registering a handler function createOneCourse for the "/courses" URL with the POST method using the HandleFunc method of the router r from the mux package
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")    // registering a handler function updateOneCourse for the "/courses/{id}" URL with the PUT method using the HandleFunc method of the router r from the mux package
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE") // registering a handler function deleteOneCourse for the "/courses/{id}" URL with the DELETE method using the HandleFunc method of the router r from the mux package

	// listen to a port and serve
	log.Fatal(http.ListenAndServe(":8000", r)) // starting the HTTP server on port 8000 and passing the router r to the ListenAndServe function to handle the incoming HTTP requests and route them to the appropriate handler functions; log.Fatal is used to log any error returned by the ListenAndServe function and exit the program if there is an error; the server runs on the local machine or localhost and listens on port 8000; ListenAndServe function is from the net/http package in Go and it starts an HTTP server with the specified address and handler to handle the incoming HTTP requests and route them to the appropriate handler functions based on the request URL and method using the router r in this case from the mux package
}

// controllers - file

// serve home page route
func serveHome(w http.ResponseWriter, r *http.Request) { // serveHome function to serve the home page of the web application to the users; it takes the response writer and the request as input parameters and returns nothing as output
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>")) // writes the welcome message to the response writer to display on the web page when the user visits the home page of the web application using the web browser
}

// get all courses route - GET /courses - returns all courses in the database as JSON response
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // setting the content type of the response to application/json to indicate that the response body is in JSON format
	json.NewEncoder(w).Encode(courses)                 // encoding the courses slice of Course struct to JSON format and writing it to the response writer w to send the list of courses as a JSON response to the client who made the request
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json") // setting the content type of the response to application/json to indicate that the response body is in JSON format

	//grab the course id from the URL
	params := mux.Vars(r) // mux.Vars function from the mux package to get the URL parameters from the request r and store them in the params map

	// loop through the courses and find the course with the given id and return it
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) // encoding the course with the given ID to JSON format and writing it to the response writer w to send the course details as a JSON response to the client who made the request
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given ID") // if no course is found with the given ID, then return a message as a JSON response to the client who made the request
	return
}

// create one course route - POST /courses - creates a new course in the database and returns the details of the new course as JSON response
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json") // setting the content type of the response to application/json to indicate that the response body is in JSON format

	// what if: body is empty
	if r.ContentLength == 0 { // checking if the request body is nil or empty //if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {} - empty json
	var course Course                              // creating a new course of type Course struct to store the details of the new course
	err := json.NewDecoder(r.Body).Decode(&course) // decoding the request body to the course struct and storing any error in decoding the request body in the err variable of type error in Go
	if err != nil {
		json.NewEncoder(w).Encode("Error in decoding the request body")
		return
	}
	if course.isEmpty() { // checking if the course is empty or not using the isEmpty method on the course struct
		json.NewEncoder(w).Encode("Please send the course details")
		return
	}

	// what if: course name is duplicate
	// loop through the courses and find the course with the given name
	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	// generate a unique id, string
	rand.Seed(time.Now().UnixNano())                // seeding the random number generator with the current time in nanoseconds to generate a unique random number each time the program is run
	course.CourseId = strconv.Itoa(rand.Intn(1000)) // generating a random number as the course ID and converting it to a string using strconv.Itoa function from the strconv package; the course ID is a unique identifier for the course in the database; strconv.Itoa function converts the integer to a string

	// add the course to the courses slice
	courses = append(courses, course) // appending the new course to the courses slice of Course struct to add the new course to the fake database
	json.NewEncoder(w).Encode(course) // encoding the new course to JSON format and writing it to the response writer w to send the new course details as a JSON response to the client who made the request
	return
}

// update one course route - PUT /courses/{id} - updates the course with the given id in the database and returns the updated course details as JSON response
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json") // setting the content type of the response to application/json to indicate that the response body is in JSON format

	// first - grab the course id from the URL
	params := mux.Vars(r) // mux.Vars function from the mux package to get the URL parameters from the request r and store them in the params map

	// loop through the courses and find the course with the given id and remove it from the slice of courses and add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] { //`params` is a map of route variables extracted from the URL
			courses = append(courses[:index], courses[index+1:]...) // removing the course with the given ID from the courses slice of Course struct using the append function and slicing the courses slice to remove the course at the index `index` and then appending the remaining courses to the slice
			var course Course                                       // creating a new course of type Course struct to store the details of the updated course
			err := json.NewDecoder(r.Body).Decode(&course)          // decoding the request body to the course struct and storing any error in decoding the request body in the err variable of type error before updating the course details
			if err != nil {
				json.NewEncoder(w).Encode("Error in decoding the request body")
				return
			}
			if course.isEmpty() { // checking if the course is empty or not using the isEmpty method on the course struct
				json.NewEncoder(w).Encode("Please send the course details")
				return
			}
			course.CourseId = params["id"]    // setting the course ID to the ID extracted from the URL parameters
			courses = append(courses, course) // appending the updated course to the courses slice of Course struct to update the course details in the fake database
			json.NewEncoder(w).Encode(course) // encoding the updated course to JSON format and writing it to the response writer w to send the updated course details as a JSON response to the client who made the request
			return
		}
	}

	// if no course is found with the given ID, then return a message as a JSON response to the client who made the request
	json.NewEncoder(w).Encode("No course found with the given ID")
	return
}

// delete one course route - DELETE /courses/{id} - deletes the course with the given id from the database and returns a success message as JSON response
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json") // setting the content type of the response to application/json to indicate that the response body is in JSON format

	// first - grab the course id from the URL
	params := mux.Vars(r) // mux.Vars function from the mux package to get the URL parameters from the request r and store them in the params map

	// loop through the courses and find the course with the given id and remove it from the slice of courses
	for index, course := range courses {
		if course.CourseId == params["id"] { //`params` is a map of route variables extracted from the URL
			courses = append(courses[:index], courses[index+1:]...)  // removing the course with the given ID from the courses slice of Course struct using the append function and slicing the courses slice to remove the course at the index `index` and then appending the remaining courses to the slice
			json.NewEncoder(w).Encode("Course deleted successfully") // encoding a success message to JSON format and writing it to the response writer w to send the success message as a JSON response to the client who made the request
			return
		}
		if course.CourseId != params["id"] { // if no course is found with the given ID, then return a message as a JSON response to the client who made the request
			json.NewEncoder(w).Encode("No course found with the given ID")
			return
		}
	}
	if len(courses) == 0 { // if the courses slice is empty, then return a message as a JSON response to the client who made the request
		json.NewEncoder(w).Encode("No courses found")
	}

}
