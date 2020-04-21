package cloudcalc

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func main() {

}
