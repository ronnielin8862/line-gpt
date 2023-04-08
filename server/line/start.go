package line

import (
	"github.com/gorilla/mux"
	"line-gpt/config"
	"line-gpt/global"
	"log"
	"net/http"
	"strconv"
)

func Init() {
	global.LineInit()
	createRouter()
}

func createRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/test", testReceive).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(config.GetConfig().Server.Port, 10), router))
}
