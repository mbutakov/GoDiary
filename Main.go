package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	handleFunc()
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", addStudentPage)
	http.HandleFunc("/save", addStudent)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "index", nil)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	ags := r.FormValue("age")
	ag, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		panic(err)
	}
	fmt.Println(name + " " + ags)
	fmt.Println(ag)
	db, err := sql.Open("mysql", "matvejs0_test:ctyJRU6@tcp(matvejs0.beget.tech:3306)/matvejs0_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	insert, err := db.Query(fmt.Sprintf("insert into client (name, age) values ('%s', '%v')", name, ag))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

}

func addStudentPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/student.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "create", nil)
	// db, err := sql.Open("mysql", "matvejs0_test:ctyJRU6@tcp(matvejs0.beget.tech:3306)/matvejs0_test")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// insert, err := db.Query(fmt.Sprintf("insert into 'client' ('name', 'age') values ('%s', '%d')"), name, age)
	// defer insert.Close()
	// http.Redirect(w, r, "/", http.StatusSeeOther)
}

type User struct {
	Name string
	Age  uint16
}
