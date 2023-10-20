package helpers

import (
	"fmt"
	"github.com/AlejandroDelg/webgo/internal/config"
	"math/rand"
	"net/http"
	"runtime/debug"
)

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("client error with status of: ", status)
	http.Error(w, http.StatusText(status), status)

}

func ServerError(w http.ResponseWriter, e error) {

	trace := fmt.Sprintf("%s\n%s", e.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
