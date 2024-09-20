package user

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.HandleLogin).Method("POST")
	router.HandleFunc("/login",h.HandleLogin).Method("POST")

}

func (h *Handler) HandleLogin(w http.ResponseWriter,r *http.Request) {

}


func (h *Handler) HandlerRegister(w http.ResponseWriter,r *http.Request) {

}