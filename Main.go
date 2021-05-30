package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB
var Glist []GroupS

func main() {
	handleFunc()
}

func handleFunc() {
	r := mux.NewRouter()
	setListGroup()
	r.HandleFunc("/", index)
	r.HandleFunc("/create", addStudentPage)
	r.HandleFunc("/save", addStudent)
	r.HandleFunc("/list", listClient)
	r.HandleFunc("/list/{id:[0-9]+}", listClientSelGroup)
	r.HandleFunc("/selectgroup", selectGroup)
	r.HandleFunc("/addGroup", addGroupPage)
	r.HandleFunc("/addGroupSite", addGroupSite)
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}

func refreshList() {
	Glist = nil
	setListGroup()
}

func addGroupPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/groupAdd.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "createGroup", nil)
}

func addGroupSite(w http.ResponseWriter, r *http.Request) {
	ags := r.FormValue("groupNumber")
	ag, err := strconv.Atoi(r.FormValue("groupNumber"))
	if err != nil {
		panic(err)
	}
	if ags == "" {
		fmt.Fprintf(w, "не все данные")
		return
	}
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", "matvejs0_test:ctyJRU6@tcp(matvejs0.beget.tech:3306)/matvejs0_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	insert, err := db.Query(fmt.Sprintf("insert into groups (idGroup) values ('%v')", ag))
	if err != nil {
		panic(err)
	}
	refreshList()
	defer insert.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func setListGroup() {
	db, err := sql.Open("mysql", "matvejs0_test:ctyJRU6@tcp(matvejs0.beget.tech:3306)/matvejs0_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from groups")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := GroupS{}
		err := rows.Scan(&p.GroupNumber)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Glist = append(Glist, p)
	}

}

func selectGroup(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/list/"+r.FormValue("group"), http.StatusSeeOther)
}
func listClientSelGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html", "templates/listgroup.html")
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("mysql", "matvejs0_test:ctyJRU6@tcp(matvejs0.beget.tech:3306)/matvejs0_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from client where groupId =" + id + "")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	user := []User{}
	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Group)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user = append(user, p)
	}
	t.ExecuteTemplate(w, "index", user)
	t.ExecuteTemplate(w, "listgroup", Glist)
}
func listClient(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html", "templates/listgroup.html")
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
		err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Group)
		if err != nil {
			fmt.Println(err)
			continue
		}
		user = append(user, p)
	}
	t.ExecuteTemplate(w, "index", user)
	t.ExecuteTemplate(w, "listgroup", Glist)
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
	gi, err := strconv.Atoi(r.FormValue("group"))
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
	insert, err := db.Query(fmt.Sprintf("insert into client (name, age, groupId) values ('%s', '%v','%v')", name, ag, gi))
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

type GroupS struct {
	GroupNumber int
}

type User struct {
	Id    uint16
	Name  string
	Age   int
	Group int
}
