package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Logging(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		date := time.Now()
		day := date.Day()
		month := date.Month()
		year := date.Year()

		hour, min, second := date.Clock()

		method := r.Method
		path := r.URL.Path

		// Date Format (DD-MM-YY)
		// Time Format (HH:MM:SS)
		fmt.Printf("[%v-%v-%v] - [%v:%v:%v] - HTTP request sent ot [%v] [%v]\n", day, month, year, hour, min, second, method, path)

		next(w, r, p)
	}
}
