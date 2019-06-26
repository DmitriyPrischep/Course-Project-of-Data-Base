package handlers

import (
	"net/http"

	"github.com/DmitriyPrischep/Course-Project-of-Data-Base/database"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	result := database.GetStatusDB()
	resp, err := result.MarshalJSON()

	switch err {
	case nil:
		makeResponse(w, 200, resp)
	default:
		makeResponse(w, 500, []byte(err.Error()))
	}
}

func Clear(w http.ResponseWriter, r *http.Request) {
	database.ClearDB()
	makeResponse(w, 200, []byte("Очистка базы успешно завершена"))
}
