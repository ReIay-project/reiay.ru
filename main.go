package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
)

type ParticleConfig struct {
	Particles struct {
		Number struct {
			Value int `json:"value"`
			Density struct {
				Enable     bool    `json:"enable"`
				ValueArea  float64 `json:"value_area"`
			} `json:"density"`
		} `json:"number"`
		Color struct {
			Value string `json:"value"`
		} `json:"color"`
		Shape struct {
			Type    string   `json:"type"`
			Stroke  struct {
				Width int    `json:"width"`
				Color string `json:"color"`
			} `json:"stroke"`
			Polygon struct {
				Sides int `json:"nb_sides"`
			} `json:"polygon"`
		} `json:"shape"`
	} `json:"particles"`
}

var particleConfig = `{
	"particles": {
		"number": {
			"value": 80,
			"density": {
				"enable": true,
				"value_area": 800
			}
		},
		"color": {
			"value": "#ffffff"
		},
		"shape": {
			"type": "circle",
			"stroke": {
				"width": 0,
				"color": "#000000"
			},
			"polygon": {
				"nb_sides": 5
			}
		}
	}
}`

func handler(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON конфигурации
	var config ParticleConfig
	err := json.Unmarshal([]byte(particleConfig), &config)
	if err != nil {
		http.Error(w, "Не удалось распарсить конфигурацию", http.StatusInternalServerError)
		return
	}

	// Рендеринг HTML страницы с конфигурацией для particles.js
	tmpl, err := template.New("index").Parse(`<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>ParticleJS in Go</title>
	<link rel="stylesheet" href="/static/css/main.css">
</head>
<body>
	<div id="particles-js"></div>
	<script src="/static/js/particles.min.js"></script>
	<script>
		document.addEventListener("DOMContentLoaded", function() {
			particlesJS('particles-js', {{.}});
		});
	</script>
</body>
</html>`)
	if err != nil {
		http.Error(w, "Не удалось загрузить HTML шаблон", http.StatusInternalServerError)
		return
	}

	configJSON, _ := json.Marshal(config)
	tmpl.Execute(w, template.JS(configJSON))
}

func main() {
	// Обработка маршрутов
	http.HandleFunc("/", handler)

	// Обработка статических файлов (JS, CSS, изображения)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Запуск сервера
	fmt.Println("Запуск сервера на порту :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
