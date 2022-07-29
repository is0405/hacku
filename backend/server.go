package server

import (
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"github.com/is0405/docker-env/controller"
	"github.com/is0405/docker-env/db"
	"github.com/is0405/docker-env/middleware"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

type Server struct {
	db           *sqlx.DB
	router       *mux.Router
	jwtSecretKey []byte
	authClient   *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string, jwtSecretKeyPath string) error {
	jwtSecretKeyFile, err := os.Open(jwtSecretKeyPath)
	if err != nil {
		return fmt.Errorf("failed read jwt secret key file. %s", err)
	}
	defer jwtSecretKeyFile.Close()

	jwtSecretKeyBuf, err := ioutil.ReadAll(jwtSecretKeyFile)
	if err != nil {
		return fmt.Errorf("failed init auth client. %s", err)
	}
	s.jwtSecretKey = jwtSecretKeyBuf

	cs := db.NewDB(datasource)
	dbcon, err := cs.Open()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	s.db = dbcon

	s.router = s.Route()
	return nil
}

func (s *Server) Run(port int) {
	log.Printf("Listening on port %d", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.jwtSecretKey, s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)
	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	//# モック API　一覧

	loginController := controller.NewLogin(s.db, s.jwtSecretKey)
	r.Methods(http.MethodPost).Path("/login").Handler(commonChain.Then(AppHandler{loginController.login}))
	return r
}
