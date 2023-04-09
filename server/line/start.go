package line

import (
	"github.com/gorilla/mux"
	"line-gpt/config"
	"line-gpt/global"
	"line-gpt/server/line/msgHandler"
	"log"
	"net/http"
	"strconv"
)

func Init() {
	global.LineInit()
	go msgHandler.TextChannelProcessor()
	go msgHandler.ImageChannelProcessor()
	createRouter()
}

func createRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/test", receiver).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(config.GetConfig().Server.Port, 10), router))
}
