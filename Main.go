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
	http.HandleFunc("/list", listClient)
	http.ListenAndServe(":8000", nil)

}

func listClient(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", "matvejs0_test:ctyJRU6@tcp(matvejs0.beget.tech:3306)/matvejs0_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from client")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	user := []User{}
	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user = append(user, p)
	}
	t.ExecuteTemplate(w, "index", user)
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
	if name == "" || ags == "" {
		fmt.Fprintf(w, "не все данные")
		return
	}

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
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func addStudentPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/student.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "create", nil)
}

type User struct {
	Id   uint16
	Name string
	Age  int
}
