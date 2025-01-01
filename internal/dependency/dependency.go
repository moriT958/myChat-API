package dependency

import (
	"database/sql"
	"log"
	"myChat-API2/internal/config"
	"myChat-API2/internal/repository"
	"myChat-API2/internal/server"
	"myChat-API2/internal/service"
	"os"
	"time"
)

func InitServer() *server.TodoServer {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Init Dependency
	userRepo := repository.NewUserRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	chatRepo := repository.NewChatRepository()

	authService := service.NewAuthService(userRepo)
	chatService := service.NewChatService(roomRepo, chatRepo, userRepo)

	srv := server.NewTodoServer(authService, chatService)

	srv.Addr = config.Address()
	srv.ReadTimeout = time.Duration(config.ReadTimeout() * int64(time.Second))
	srv.WriteTimeout = time.Duration(config.WriteTimeout() * int64(time.Second))

	return srv
}
