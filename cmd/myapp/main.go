package main
import (
	"os"
	"fmt"
	"log"
	"net/http"

	"crud-go-example/internal/myapp"
	_"github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Присваиваем переменным значения из .env
func init() {
    if err := godotenv.Load("../../config/config.env"); err != nil {
        log.Fatal("Error loading .env file!")
    }
    myapp.Port = os.Getenv("PORT")
	myapp.DbDriver = os.Getenv("DB_DRIVER")
	myapp.DbUser = os.Getenv("DB_USER")
	myapp.DbPass = os.Getenv("DB_PASSWORD")
	myapp.DbHost = os.Getenv("DB_HOST")
	myapp.DbName = os.Getenv("DB_NAME")
	myapp.DbMode = os.Getenv("DB_SSLMODE")
	myapp.DbConnection = fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", myapp.DbUser, myapp.DbPass, myapp.DbHost, myapp.DbName, myapp.DbMode)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user", myapp.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", myapp.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", myapp.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", myapp.DeleteUserHandler).Methods("DELETE")

	log.Println("The server is running! Port", myapp.Port)
	log.Fatal(http.ListenAndServe(myapp.Port, r))
}