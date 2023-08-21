package main

func main() {
	DB_connect()
	defer DB_disconnect()
	ServerStart()
}
