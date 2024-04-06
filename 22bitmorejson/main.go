package main

import (
	"encoding/json"
	"fmt"
)

type course struct { // defining a struct type course to store the details of a course; in lowercase, the struct is not exported and can be accessed only within the same package
	Name     string   `json:"coursename"` // json tag is used to specify the key name in the JSON for the field of the struct type otherwise the key name in the JSON will be the same as the field name of the struct type
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // - tag is used to omit the field from the JSON while encoding the struct to JSON format
	Tags     []string `json:"tags,omitempty"` // omitempty tag is used to omit the field from the JSON if the field is empty or nil
}

func main() {
	fmt.Println("More JSON in Go")
	EncodeJson() // calling the EncodeJson function to encode a struct to JSON
	DecodeJson() // calling the DecodeJson function to decode a JSON to a struct
}

// EncodeJson function encodes a struct to JSON format; encoding means converting from one format to another format  -- here, encoding a struct to JSON format
func EncodeJson() {
	// encoding a struct to JSON
	lcoCourses := []course{{"ReactJS", 299, "LearnCodeOnline", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "LearnCodeOnline", "bcd123", []string{"full-stack", "js"}},
		{"Angular", 199, "LearnCodeOnline", "gtc123", nil},
	} // creating a slice of course struct type to store the details of multiple courses with different fields and values

	// package this data into JSON format
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // MarshalIndent function from the encoding/json package encodes the data into JSON format with indentation and prefix for each line and returns the JSON as a slice of bytes and any error in encoding the data to JSON format
	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON: %s\n", string(finalJson)) // prints the JSON as a string;  %s is used to format the string and string(finalJson) converts the slice of bytes to a string
}

/* the output of the program will be:
More JSON in Go
JSON: [
        {
                "coursename": "ReactJS",
                "price": 299,
                "website": "LearnCodeOnline",
                "tags": [
                        "web-dev",
                        "js"
                ]
        },
        {
                "coursename": "MERN",
                "price": 199,
                "website": "LearnCodeOnline",
                "tags": [
                        "full-stack",
                        "js"
                ]
        },
        {
                "coursename": "Angular",
                "price": 199,
                "website": "LearnCodeOnline"
        }
]

the JSON output is formatted with indentation and for better readability and understanding of the JSON data structure
*/

func DecodeJson() { //consume JSON data from the web and decode it into a struct type
	// decoding a JSON to a struct
	jsonDataFromWeb := []byte(`
		{"coursename": "ReactJS", "price": 299, "website": "LearnCodeOnline", "tags": ["web-dev", "js"]}
	`) // declaring a variable jsonDataFromWeb of type slice of bytes to store the JSON data to be decoded into a struct type course -- the JSON data is stored as a string and converted to a slice of bytes using type conversion

	var decodedCourses course // declaring a variable decodedCourses of type course struct to store the decoded JSON data

	checkValid := json.Valid(jsonDataFromWeb) // Valid function from the encoding/json package checks if the JSON data is valid or not and returns a boolean value
	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &decodedCourses)       // Unmarshal function from the encoding/json package decodes the JSON data into the decodedCourses variable of type slice of course struct type and returns any error in decoding the JSON data; the second argument is a pointer to the variable where the decoded JSON data will be stored because I do not want to create a new variable or a copy to store the decoded JSON data and hence passing the address of the variable
		fmt.Printf("Decoded JSON data: %#v\n", decodedCourses) // prints the decoded JSON data as a slice of course struct type; %#v is used to format the struct type with field names and values for better understanding
	} else {
		fmt.Println("JSON data is invalid")
	}

	//some cases where you just want to add data to key value pairs in a map and not struct
	var myOnlineCourses map[string]interface{} // declaring a variable myOnlineCourses of type map with key as string and value as interface{} as I do not know the type of the value yet and will be decided at runtime based on the JSON data to be decoded into this map variable -- interface{} is an empty interface in Go that can hold values of any type and is used when the type of the value is not known at compile time and will be decided at runtime based on the data to be stored in the variable -- it is similar to the Object type in other languages like Java

	err := json.Unmarshal(jsonDataFromWeb, &myOnlineCourses) // Unmarshal function from the encoding/json package decodes the JSON data into the myOnlineCourses variable of type map with key as string and value as interface{} and returns any error in decoding the JSON data
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded JSON data into map: %#v\n", myOnlineCourses) // prints the decoded JSON data as a map with key as string and value as interface{}; %#v is used to format the map with key value pairs for better understanding

	for k, v := range myOnlineCourses {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}

/* decode Json() is used to decode the JSON data into a struct type course and a map with key as string and value as interface{}. The JSON data is stored as a slice of bytes and then decoded into the struct type course and the map myOnlineCourses. The JSON data is first checked for validity using the Valid function from the encoding/json package. If the JSON data is valid, it is decoded into the struct type course and the map myOnlineCourses using the Unmarshal function from the encoding/json package. The decoded JSON data is then printed as a slice of course struct type and a map with key as string and value as interface{} using the Printf function with the %#v format specifier for better understanding of the data structure. The key value pairs in the map myOnlineCourses are then printed using a for loop to iterate over the map and print the key, value, and type of the value for each key value pair.
 */
