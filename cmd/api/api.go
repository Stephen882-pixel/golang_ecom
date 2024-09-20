package api 

import(
	"database/sql",
	"net/http"
	"github.com/gorrila/mux"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string,db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: DB,
	}
}

func (s *APIServer) Run() error{
	router :=mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)


	log.Println("Listening on",s.addr)
	return http.ListenAndServe(s.addr,router)
}