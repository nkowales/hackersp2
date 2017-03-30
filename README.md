# hackersp2

A project using a Hipster programing language for Peter Bui's CSE 40842 class 'Hackers in the Bazaar'.

### Technologies
Bootstrap
Font-Awesome
Go
[Gorilla Websockets](https://github.com/gorilla/websocket)
jQuery

### Use
To use run the following command, this will start the server on localhost port 8080.
`go run server.go`

The program is simple. It takes a query from the user and returns one of many possible 
fortunes. Going to the webpage enables the user to connect to the server and constantly get fortunes.

On the serverside the use of goroutines allows the program to have multiple open websockets 
until they close allowing it to scale. It mainly has many threads concurrently running to 
keep giving all users their fortunes without suffering from performance issues. 

Use of go specific goroutines was part of this projects use of a feature specific to the go language. 
