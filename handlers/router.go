package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/forum/create", CreateForum).Methods("POST")
	r.HandleFunc("/api/forum/{slug}/create", CreateForumThread).Methods("POST")
	r.HandleFunc("/api/forum/{slug}/details", GetForum).Methods("GET")
	r.HandleFunc("/api/forum/{slug}/threads", GetForumThreads).Methods("GET")
	r.HandleFunc("/api/forum/{slug}/users", GetForumUsers).Methods("GET")
	r.HandleFunc("/api/user/{nickname}/create", CreateUser).Methods("POST")
	r.HandleFunc("/api/user/{nickname}/profile", GetUser).Methods("GET")
	r.HandleFunc("/api/user/{nickname}/profile", UpdateUser).Methods("POST")
	r.HandleFunc("/api/post/{id}/details", GetPost).Methods("GET")
	r.HandleFunc("/api/post/{id}/details", UpdatePost).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/details", GetThread).Methods("GET")
	r.HandleFunc("/api/thread/{slug_or_id}/details", UpdateThread).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/create", CreatePost).Methods("POST")
	r.HandleFunc("/api/service/status", GetStatus).Methods("GET")
	r.HandleFunc("/api/service/clear", Clear).Methods("POST")
	r.HandleFunc("/api/thread/{slug_or_id}/posts", GetThreadPosts).Methods("GET")
	r.HandleFunc("/api/thread/{slug_or_id}/vote", MakeThreadVote).Methods("POST")
	r.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello USER"))
	})
	return r
}
