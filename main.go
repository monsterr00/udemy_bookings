package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрация обработчик для пути "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello world!")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Bytes writen: %d", n))
	})

	// Запуск web-сервера
	_ = http.ListenAndServe(":8080", nil)
}
