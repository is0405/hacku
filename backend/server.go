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

	"github.com/is0405/hacku/controller"
	"github.com/is0405/hacku/db"
	"github.com/is0405/hacku/middleware"
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

func (s *Server) Init(datasource string) error {
	// jwtSecretKeyFile, err := os.Open(jwtSecretKeyPath)
	jwtSecretKeyFile, err := os.Open(".hacku/jwt-secret.key")
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
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	
	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)
	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()

	 r.Methods(http.MethodOptions).PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	// r.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// })
	//# API　一覧
	//ログイン
	loginController := controller.NewLogin(s.db, s.jwtSecretKey)
	r.Methods(http.MethodPost).Path("/login").Handler(commonChain.Then(AppHandler{loginController.Login}))

	//アカウント情報
	UserControlloer := controller.NewUser(s.db)
	// r.Methods(http.MethodPost).Path("/users").Handler(commonChain.Then(AppHandler{UserControlloer.CreateSubUser}))
	r.Methods(http.MethodGet).Path("/users").Handler(authChain.Then(AppHandler{UserControlloer.GetUser}))
	r.Methods(http.MethodPatch).Path("/users").Handler(authChain.Then(AppHandler{UserControlloer.UpdateUser}))
	r.Methods(http.MethodDelete).Path("/users").Handler(authChain.Then(AppHandler{UserControlloer.DeleteUser}))
	r.Methods(http.MethodPost).Path("/users/create").Handler(commonChain.Then(AppHandler{UserControlloer.CreateUser}))
	// r.Methods(http.MethodGet).Path("/users-sub").Handler(commonChain.Then(AppHandler{UserControlloer.GetSubUser}))

	//雇用情報
	RecruitmentControlloer := controller.NewRecruitment(s.db)
	r.Methods(http.MethodPost).Path("/recruitment").Handler(authChain.Then(AppHandler{RecruitmentControlloer.CreateRecruitment}))
	r.Methods(http.MethodGet).Path("/recruitment/all/mine").Handler(authChain.Then(AppHandler{RecruitmentControlloer.GetMyAllRecruitment}))
	r.Methods(http.MethodGet).Path("/recruitment/all/other").Handler(authChain.Then(AppHandler{RecruitmentControlloer.GetOtherAllRecruitment}))
	r.Methods(http.MethodGet).Path("/recruitment/all/participation").Handler(authChain.Then(AppHandler{RecruitmentControlloer.GetMyParticipationAllRecruitment}))
	r.Methods(http.MethodGet).Path("/recruitment/{recruitment_id}").Handler(authChain.Then(AppHandler{RecruitmentControlloer.GetRecruitmentFromID}))
	r.Methods(http.MethodPatch).Path("/recruitment/{recruitment_id}").Handler(authChain.Then(AppHandler{RecruitmentControlloer.UpdateRecruitment}))
	r.Methods(http.MethodDelete).Path("/recruitment/{recruitment_id}").Handler(authChain.Then(AppHandler{RecruitmentControlloer.DeleteRecruitment}))
	r.Methods(http.MethodGet).Path("/recruitment/{recruitment_id}/participation").Handler(authChain.Then(AppHandler{RecruitmentControlloer.GetParticipation}))
	
	//申請情報
	HiredControlloer := controller.NewHired(s.db)
	r.Methods(http.MethodPost).Path("/hired/{recruitment_id}").Handler(authChain.Then(AppHandler{HiredControlloer.PostHired}))
	r.Methods(http.MethodDelete).Path("/hired/{recruitment_id}").Handler(authChain.Then(AppHandler{HiredControlloer.DeleteHired}))

	//日程調整
	CalenderControlloer := controller.NewCalender(s.db)
	// r.Methods(http.MethodPost).Path("/calender/{recruitment_id}").Handler(authChain.Then(AppHandler{CalenderControlloer.PostRecCalender}))
	r.Methods(http.MethodGet).Path("/calender/{recruitment_id}").Handler(authChain.Then(AppHandler{CalenderControlloer.GetCalender}))
	r.Methods(http.MethodPatch).Path("/calender/{recruitment_id}").Handler(authChain.Then(AppHandler{CalenderControlloer.PatchPartiCalender}))
	
	
	return r
}
