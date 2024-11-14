package main

import (
    "net/http"
    "log"
)

func main() {
    // Обработка статических файлов
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/public/", http.StripPrefix("/public/", fs))

    // Основные маршруты
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/success", successHandler)
    http.HandleFunc("/404", notFoundHandler)

    // Запуск сервера
    log.Println("Сервер запущен на порту 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./templates/index.html")
}

func successHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./templates/success.html")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./templates/404.html")
}