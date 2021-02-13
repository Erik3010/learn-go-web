package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There is an Error", http.StatusInternalServerError)
		return
	}

	// passing data to the html views using map
	// data := map[string]interface{}{
	// 	"title":   "Home Page",
	// 	"content": "Hello this is the home page",
	// }

	data := entity.Product{ID: 1, Name: "Buku", Price: 10000, Stock: 10}

	err = tmpl.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World! My name is Erik"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprint(w, "Product id: %d", idNum)

	data := map[string]interface{}{
		"idNum": idNum,
	}

	tmplt, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There is an Error", http.StatusInternalServerError)
		return
	}

	tmplt.Execute(w, data)
}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	data := []entity.Product{
		{ID: 1, Name: "Buku", Price: 10000, Stock: 11},
		{ID: 2, Name: "Pensil", Price: 20000, Stock: 8},
		{ID: 3, Name: "Penghapus", Price: 30000, Stock: 1},
	}

	tmplt, err := template.ParseFiles(path.Join("views", "detail.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

	err = tmplt.Execute(w, data)

	if err != nil {
		log.Println(err)
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("You are using GET Method"))
	case "POST":
		w.Write([]byte("You are using POST method"))
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "GET" {

		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)

		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
}
