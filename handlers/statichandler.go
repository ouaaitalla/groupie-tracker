package zone

import "net/http"

// HandleStatic serves CSS files and prevents direct access to /static/
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static" || r.URL.Path == "/static/" {
		HandleError(w, http.StatusForbidden, "Access Forbidden")
		return
	}

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	fs.ServeHTTP(w, r)
}
