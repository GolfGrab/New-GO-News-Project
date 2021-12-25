package app

import "net/http"

// Mount mounts the application to the given http.Handler.
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", indexPage)
	mux.HandleFunc("/news/", newsView)

	adminMux := http.NewServeMux()

	adminMux.HandleFunc("/login", adminLogin)
	adminMux.HandleFunc("/list", adminList)
	adminMux.HandleFunc("/create", adminCreate)
	adminMux.HandleFunc("/delete", adminDelete)
	adminMux.HandleFunc("/edit", adminEdit)

	mux.Handle("/admin/", onlyAdmin(adminMux))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
