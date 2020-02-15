package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World\n")
}

func todo(endpointName string) apiHandler {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, endpointName+" not implemented\n")
	}
}

func createConnection(ctx Context) {
	// result := db.CreateConnection()
	// log.Println(result)

	fmt.Fprint(ctx.Resp, "Ok")
}

func recipientApi(path string, router *httprouter.Router) {
	router.POST(path, handler("createConnection", createConnection))
	router.PUT(path+"/:connection", todo("secureConnection"))
	router.DELETE(path+"/:connection", todo("deleteConnection"))
	router.GET(path+"/:connection/messages", todo("getMessages"))
	router.GET(path+"/:connection/messages/:msgId", todo("getMessage"))
	router.DELETE(path+"/:connection/messages/:msgId", todo("deleteMessage"))
}

func senderApi(path string, router *httprouter.Router) {
	router.POST(path+"/:connection/messages", todo("sendMessage"))
}

// New returns instance of API router
func New(recipientPath string, senderPath string) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", hello)
	recipientApi(recipientPath, router)
	senderApi(senderPath, router)
	return router
}
