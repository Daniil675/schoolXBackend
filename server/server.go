package server

import (
	"log"
	"net/http"
)

func (s Server) Start(port string) {
	//router := httprouter.New()
	//http.HandleFunc("/", indexHandler)
	//http.HandleFunc("/city", cityHandler)
	//http.HandleFunc("/wh/ym", paymentYandexHandler)
	//router.POST("/event", Wrap(cors.Then(s.eventAddHandler)))
	//router.OPTIONS("/event", s.optionsHandler)

	//http.HandleFunc("/", indexHandler)

	//city
	http.HandleFunc("/city", cityHandler)
	http.HandleFunc("/cities", cityGetAllHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln("Error at ListenAndServe: ", err)
	} else {
		log.Printf("Server stared at localhost:%s\n", port)
	}
}
