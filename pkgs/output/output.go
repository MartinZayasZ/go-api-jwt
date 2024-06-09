package output

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func JsonResponse(w http.ResponseWriter, data any, httpStatusCode int, cacheTime uint) {

	if httpStatusCode == 0 {
		httpStatusCode = 200
	}

	//Obtenemos el slide de bytes
	responseSB, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Ha ocurrido un error", err)
		return
	}

	w.Header().Set("httpStatusCode", strconv.Itoa(httpStatusCode))
	w.Header().Set("Content-Type", "application/json")

	if cacheTime > 0 {
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age: %v, must-revalidate", cacheTime))
	}

	_, err = w.Write(responseSB)

	if err != nil {
		fmt.Println("Ha ocurrido un error", err)
	}
}
