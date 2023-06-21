package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type LogMessage struct {
	Message string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса на домашнюю страницу")
	renderTemplate(w, "templates/home.html", nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {

	var logMessage LogMessage
	err := json.NewDecoder(r.Body).Decode(&logMessage)
	if err != nil {
		log.Println("Ошибка при разборе лог-сообщения:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("Лог с фронтенда:", logMessage.Message)

	w.WriteHeader(http.StatusOK)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Println("Ошибка при загрузке шаблона:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println("Ошибка при рендеринге шаблона:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	log.Println("Запуск бэкенд-сервера")


	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/log", logHandler).Methods("POST")

	router.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	log.Println("Запуск сервера на порту 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}
}
