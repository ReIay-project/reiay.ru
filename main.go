package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
	"log"
	"math/rand"
	"strings"
)

func main() {
	// Обработка статических файлов
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Обработка маршрутов
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/random", randomHandler)

	// Запуск сервера
	fmt.Println("Сервер запущен на порту 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/about.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон About", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/contact.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Contact", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка обработки формы", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	log.Printf("Получено сообщение от %s (%s): %s", name, email, message)

	tmpl, err := template.ParseFiles("./templates/thankyou.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Thank You", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, struct {
		Name string
	}{
		Name: name,
	})
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Fprintf(w, "Текущее время: %s", now.Format("02-01-2006 15:04:05"))
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	fmt.Fprintf(w, "Случайное число: %d", randomNumber)
}

// Дополнительные функции для расширения функционала

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/faq.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон FAQ", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/services.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Services", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/blog.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Blog", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func generateWelcomeMessage(name string) string {
	if strings.TrimSpace(name) == "" {
		return "Добро пожаловать, Гость!"
	}
	return fmt.Sprintf("Добро пожаловать, %s!", name)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	message := generateWelcomeMessage(name)
	fmt.Fprintf(w, "%s", message)
}

// Дополнительные обработчики маршрутов
func galleryHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/gallery.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Gallery", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func pricingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/pricing.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Pricing", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func testimonialsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/testimonials.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Testimonials", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Функция для обработки ошибок
func handleError(w http.ResponseWriter, err error, message string, code int) {
	log.Println("Ошибка:", err)
	http.Error(w, message, code)
}

// Логирование запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

// Маршрут с динамическим содержимым
func dynamicContentHandler(w http.ResponseWriter, r *http.Request) {
	content := "<h1>Добро пожаловать на динамическую страницу</h1><p>Это содержимое было сгенерировано программно.</p>"
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, content)
}

// Обработчик для загрузки файлов
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20) // Ограничение на размер файла до 10 MB
		file, _, err := r.FormFile("uploadfile")
		if err != nil {
			handleError(w, err, "Ошибка при загрузке файла", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "Файл успешно загружен!")
		return
	}

	tmpl, err := template.ParseFiles("./templates/upload.html")
	if err != nil {
		http.Error(w, "Не удалось загрузить шаблон Upload", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// Функция для отображения системного времени
func systemTimeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("02-01-2006 15:04:05 Monday")
	fmt.Fprintf(w, "Текущее системное время: %s", currentTime)
}

// Дополнительный обработчик для проверки статуса сервера
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Сервер работает нормально")
}

// Маршрут для вывода информации о сервере
func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	info := "<h2>Информация о сервере</h2><ul><li>Версия Go: 1.21.0</li><li>ОС: Ubuntu 20.04</li><li>Статус: работает</li></ul>"
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, info)
}

// Маршрут для получения IP-адреса клиента
func clientIPHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	fmt.Fprintf(w, "Ваш IP-адрес: %s", ip)
}
