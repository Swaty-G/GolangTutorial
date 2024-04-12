This a code repo for learning Golang. It contains the basic concepts of Golang.

```markdown
# GolangTutorial

This is a tutorial project for learning Go programming language.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Usage

- Go version 1.22.1
- GoLand IDE 2023.3.5 or any other IDE with Go support

### Installing

Clone the repository:

git clone git@github.com:Swaty-G/GolangTutorial.git
```

Navigate to the project directory:

```bash
cd GolangTutorial
```

Run the main.go file:

```bash
go run main.go
```

## Important Links
- https://pkg.go.dev/  (Go packages)   
- https://golang.org/doc/  (Go documentation)
- https://go.dev/play/  (Go playground)
- https://go.dev/ref/mod (Go Modules reference documentation link to understand the Go modules in detail and how to use them in Go)
- https://pkg.go.dev/github.com/gorilla/mux (Gorilla Mux package documentation link to understand the package and its functions in detail)


## Project Structure

The project has several Go files each demonstrating different concepts in Go:

- `01hello/main.go`: A simple program that prints "Hello from swaty"
- `02variables/main.go`: A program demonstrating the use of variables in Go
- `03userinput/main.go`: A program demonstrating how to take user input in Go
- `04conversion/main.go`: A program demonstrating type conversion in Go
- `05mymaths/main.go`: A program demonstrating math operations in Go
- `06mytime/main.go`: A program demonstrating time operations in Go
- `07mypointers/main.go`: A program demonstrating the use of pointers in Go
- `08myarray/main.go`: A program demonstrating the use of arrays in Go
- `09myslices/main.go`: A program demonstrating the use of slices in Go
- `10mymaps/main.go`: A program demonstrating the use of maps in Go
- `11mystructs/main.go`: A program demonstrating the use of structs in Go
- `12myif/main.go`: A program demonstrating the use of if statements in Go
- `13myswitch/main.go`: A program demonstrating the use of switch statements in Go
- `14loops/main.go`: A program demonstrating the use of loops in Go
- `15functions/main.go`: A program demonstrating the use of functions in Go
- `16methods/main.go`: A program demonstrating the use of methods in Go
- `17defer/main.go`: A program demonstrating the use of defer in Go
- `18files/main.go`: A program demonstrating file operations in Go
- `19webrequests/main.go`: A program demonstrating handling web requests in Go
- `20urls/main.go`: A program demonstrating handling URLs in Go
- `lcowebserver`: A program creating a simple web server in Go
- `21webreqverbs/main.go`: A program demonstrating handling get, post, post form requests in Go. Start the lco web server before running this program. 
- `22bitmorejson/main.go`: A program demonstrating how to consume json data by encoding and decoding in Go
- `23mymodules/main.go`: A program demonstrating how to route requests in Go using gorilla/mux package 
- `24buildapp/main.go`: A program demonstrating how to build APIs(GET, POST, PUT, DELETE) in Go using gorilla/mux package. Used Postman to test the APIs.
- `25mongoapi/main.go`: A program demonstrating how to connect to MongoDB and perform CRUD operations in Go by creating Netflix and performs CRUD operations on it. Used Postman to test the APIs.
- `26goroutines/main.go`: A program demonstrating how to use goroutines in Go to run multiple tasks concurrently using the go keyword and WaitGroup. Also, mutex is used. 
- `27mutexAndAwaitGroups/main.go`: A program demonstrating how to use mutex and WaitGroup in Go to run multiple tasks concurrently using the go keyword and WaitGroup. First, the program runs without mutex and check race condition(go run --race .) and then with mutex to resolve the race condition.

## Authors

- **Swaty Gupta**

### Acknowledgments 
- https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa (Hitesh Choudhary Go lang tutorial)
- https://github.com/hiteshchoudhary/golang (github repo for the above tutorial)

### Prerequisites to run lco web server
To install node-js and npm and run the following command in the terminal:


1. Open your `.zshrc` file in a text editor.
   ```bash
   vim ~/.zshrc
    ```
2. Add the following lines to the end of the `.zshrc` file:

    ```bash
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" # This loads nvm
    ```

3. Save and exit the editor.


4. Load the new `.zshrc` configuration** by either closing and reopening your terminal or by sourcing the `.zshrc` file with the following command:
    ```bash
    source ~/.zshrc
    ```

5. Verify the installation by typing:

    ```bash
    command -v nvm
    ```
This should output `nvm` if the installation was successful.

6. Install Node.js using NVM by running:
    ```bash
       nvm install node # "node" is an alias for the latest version
    ```

7. Check the Node version to confirm it's installed:

    ```bash
       node -v
    ```
   
### Prerequisites to run MongoDB API 
Create a MongoDB account and follow the instructions to create a cluster and connect to it.
Login in to cloud.mongodb.com
- Create a cluster
- Connect to the cluster
   



