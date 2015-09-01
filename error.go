package main

<<<<<<< HEAD
// Error represents information regarding errors and status
=======
>>>>>>> a1e0e468b830fd9f6f6a7d69a13e9b162a07ce21
type Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

var codes = map[int]string{
	400: "Bad Request",
	404: "Not Found",
	405: "Invalid Input",
}
