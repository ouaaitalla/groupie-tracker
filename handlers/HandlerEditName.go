package zone

import "net/http"

func HandlerEditName(w http.ResponseWriter, r *http.Request) {
	if len(AllArtists) == 0 {
		http.Error(w, "No artists available", http.StatusInternalServerError)
		return
	}

	for i := 0; i < 1000; i++ {
		go func() {
			AllArtists[0].Name = "New Name"
		}()
	}

	w.Write([]byte("Name changed!"))
}
