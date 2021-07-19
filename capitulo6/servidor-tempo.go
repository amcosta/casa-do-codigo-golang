package capitulo6

import (
	"fmt"
	"net/http"
	"time"
)

func ServerInit() {
	http.HandleFunc("/tempo", tempo)
	http.ListenAndServe(":8054", nil)
}

func tempo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
}
