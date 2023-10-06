package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// Создаем новый маршрутизатор Chi
	r := chi.NewRouter()

	// Добавляем промежуточное ПО для логирования и обработки ошибок
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Определяем маршруты и функции-обработчики для них
	r.Get("/", HomeHandler)
	r.Get("/about", AboutHandler)

	// Запускаем веб-сервер на порту 8080
	port := ":8080"
	fmt.Printf("Сервер запущен на порту %s\n", port)
	http.Handle("/", r)
	http.ListenAndServe(port, nil)
}

// Обработчик для маршрута "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Добро пожаловать на домашнюю страницу!")
}

// Обработчик для маршрута "/about"
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Это страница о нас.")
}
