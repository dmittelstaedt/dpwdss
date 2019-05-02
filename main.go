package main

func main() {
	server := NewServer("root", "david", "pwdss")
	server.SetRoutes()
	server.Run()
}
