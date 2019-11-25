package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/wubingwei/veigar/config"
	"net/http"
	"strings"
)

func init() {
	if config.DevMode {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	http.HandleFunc("/", HelloVeigar)
	log.Debug("veigar is dancing ================>")
	err := http.ListenAndServe(":1994", nil)
	if err != nil {
		log.Fatal("Listen and serve", err.Error())
	}
}

func HelloVeigar(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", strings.Join(v, " "))
	}
	fmt.Fprintf(w, "Hello verigar")
}
