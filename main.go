package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func main() {
    // Обработка корневого маршрута "/" для отдачи index.html
    http.HandleFunc("/", indexHandler)

    // Обработка статических файлов (CSS, JS, изображения и т.д.)
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Запуск сервера на порту 8080
    fmt.Println("Запуск сервера на порту :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Ошибка при запуске сервера:", err)
    }
}

// indexHandler обрабатывает запросы к корневому URL и отдает HTML-шаблон index.html
func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("./templates/index.html")
    if err != nil {
        http.Error(w, "Не удалось загрузить HTML шаблон", http.StatusInternalServerError)
        return
    }

    // Рендер шаблона index.html
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Ошибка при рендеринге HTML", http.StatusInternalServerError)
    }
}
