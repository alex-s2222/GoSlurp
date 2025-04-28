package app

import (
	"context"
	"httpServer/internals/cfg"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	config cfg.Cfg
	ctx context.Context
	srv *http.Server
	db *pgxpool.Pool
}

func NewServer(config cfg.Cfg, ctx context.Context) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	return server
}


func (server *Server) Serve() {
	log.Println("Starting server")
	var err error
	server.db, err = pgxpool.New(server.ctx, server.config.GetDBString())
	
	
	if err != nil {
		log.Fatal(err)
	}

	// методы с базой данных
	
	// Зависимости

	// DB

	// Prosess

	// handlers

	// routes

	server.srv = &http.Server{
		Addr: ":" + server.config.Port,
	}

	log.Println("Server started")
	
	err = server.srv.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}

	return 
}


func (server *Server) ShutDown() {
	log.Printf("server stoped")
	
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	server.db.Close()
	defer func(){
		cancel()
	}()
	
	var err error
	if err = server.srv.Shutdown(ctxShutDown); err != nil{
		log.Fatalln("server Shutdown Faild", err)
	}

	if err == http.ErrServerClosed{
		err = nil 
	}
}