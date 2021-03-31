package main

const User string = "tarik"
const Password string = "Test123!"
const DatabaseName string = "tarik"

const port string = ":8080"

func main() {
	a := App{}
	a.Initialize(User, Password, DatabaseName)
	a.Run(port)
}
