package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/simplex-chat/simplex-server/db"
)

func todo(endpointName string) apiHandler {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, endpointName+" not implemented\n")
	}
}

func getRandomBase64(sizeBytes int8) string {
	b := make([]byte, sizeBytes)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

func createConnection(cx ApiContext) {
	recipientKeyStr, _ := cx.Body["recipient"].(string)
	recipientKey, err := base64.StdEncoding.DecodeString(recipientKeyStr)
	if err != nil {
		log.Println("Error:", err)
		cx.Resp.WriteHeader(http.StatusBadRequest)
		io.WriteString(cx.Resp, "Bad Request")
		return
	}
	simplex := db.NewSimplex{
		Recipient_id:  getRandomBase64(16),
		Sender_id:     getRandomBase64(16),
		Recipient_key: recipientKey,
	}
	result := db.CreateConnection(cx.Req.Context(), simplex)
	log.Println(result)

	fmt.Fprint(cx.Resp, "Ok")
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
	recipientApi(recipientPath, router)
	senderApi(senderPath, router)
	return router
}
