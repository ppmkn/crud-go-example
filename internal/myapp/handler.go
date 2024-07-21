package myapp
import (
	"fmt"
	"strconv"
	"net/http"
	"database/sql"
	"encoding/json"

	_"github.com/lib/pq"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbConnection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "JSON parsing error!", http.StatusBadRequest)
		return
	}

	err = CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "User creation failed!"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User successfully created!")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbConnection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter!", http.StatusBadRequest)
		return
	}

	user, err := GetUser(db, userID)
	if err != nil {
		http.Error(w, "User not found!"+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbConnection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter!", http.StatusBadRequest)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "JSON parsing error!", http.StatusBadRequest)
		return
	}

	err = UpdateUser(db, userID, user.Name, user.Email)
	if err != nil {
		http.Error(w, "User not found!", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User successfully updated!")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(DbDriver, DbConnection)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	vars := mux.Vars(r)
	idStr := vars["id"]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID parameter!", http.StatusBadRequest)
		return
	}
	
	user := DeleteUser(db, userID)
	if user != nil {
		http.Error(w, "User not found!", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User successfully deleted!")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}