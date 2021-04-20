package v1

import (
	"github.com/gorilla/mux"
	"github.com/osetr/app/configs"
	"github.com/osetr/app/internal/dao"
	"github.com/osetr/app/internal/repository"
	"github.com/osetr/app/internal/service"
	"github.com/osetr/app/pkg/database"
)

func InitRoute() *mux.Router {

	var (
		config = configs.NewConfig("config")
		conn   = database.NewConnectionFactory(config.DBAddr, config.DBUser, config.DBPass, config.DBName)

		userDAO = dao.NewUserDAO(conn)
		postDAO = dao.NewPostDAO(conn)

		authRepo = repository.NewUserRepository(userDAO)
		postRepo = repository.NewPostRepository(postDAO)

		authService = service.NewAuthService(authRepo)
		postService = service.NewPostService(postRepo)

		authHandler = NewAuthHandler(authService)
		postHandler = NewPostHandler(postService)
	)

	router := mux.NewRouter()

	v1 := router.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/sign-in", authHandler.signIn).Methods("POST")
	v1.HandleFunc("/sign-up", authHandler.signUp).Methods("POST")

	posts := v1.PathPrefix("/posts").Subrouter()
	posts.Use(AuthMiddlware)

	posts.HandleFunc("", postHandler.createPost).Methods("POST")
	posts.HandleFunc("", postHandler.getAllPosts).Methods("GET")
	posts.HandleFunc("/{id}", postHandler.getPostById).Methods("GET")
	posts.HandleFunc("/{id}", postHandler.updatePost).Methods("PUT")
	posts.HandleFunc("/{id}", postHandler.deletePost).Methods("DELETE")

	return router
}
