package main

import(
	"OpenSoft-MT/database"
)
func main(){
	database.ConnectDb("root:<pass>@tcp(localhost:3306)/jwt_demo?parseTime=true") // later change this hardcoded string to non hardcoded
    database.Migrate()
}