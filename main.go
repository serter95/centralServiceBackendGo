package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	crcindservicego "github.com/serter95/crcindServiceGo"
	itunesservicego "github.com/serter95/itunesServiceGo"
	tvmazeservicego "github.com/serter95/tvmazeServiceGo"
)

// go get github.com/githubnemo/CompileDaemon
// export PATH=$PATH:$(go env GOPATH)/bin
// go get -u github.com/gorilla/mux
// CompileDaemon main.go
// CompileDaemon -command="./centralServiceBackendGo"

// StandardResponse struct ...
type StandardResponse struct {
	Category   string `json:"category"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	PreviewURL string `json:"previewUrl"`
	Origin     string `json:"origin"`
}

// searchData godoc
// @Summary search criteria in all services
// @Description search criteria in all services
// @Tags search
// @Accept  json
// @Produce  json
// @Success 200 {array} StandardResponse
// @Router /search/{criteria} [get]
// @Param criteria path string true "criteria that you want to find"
func searchData(w http.ResponseWriter, r *http.Request) {
	// seteo los datos de la respuesta por defecto
	w.Header().Set("Content-Type", "application/json")
	statusReturn := http.StatusOK
	var standardResponse []StandardResponse

	// obtengo las variables del path
	params := mux.Vars(r)
	// trato de convertir el id a int
	criteria := params["criteria"]
	// valido si hay errores
	if reflect.TypeOf(criteria).Kind() != reflect.String {
		fmt.Println("FALLO, el criterio debe ser string")
	}

	// channel := make(chan interface{})
	itunesArray := processItunes(criteria)
	// go processItunes(criteria, channel)
	tvmazeArray := processTvmaze(criteria)
	// go processTvmaze(criteria, channel)
	crcindArray := processCrcind(criteria)
	// go processCrcind(criteria, channel)

	for _, iterator := range itunesArray {
		standardResponse = append(standardResponse, StandardResponse{
			Category:   iterator.Category,
			Name:       iterator.Name,
			Author:     iterator.Author,
			PreviewURL: iterator.Author,
			Origin:     iterator.Origin,
		})
	}

	for _, iterator := range tvmazeArray {
		standardResponse = append(standardResponse, StandardResponse{
			Category:   iterator.Category,
			Name:       iterator.Name,
			Author:     iterator.Author,
			PreviewURL: iterator.Author,
			Origin:     iterator.Origin,
		})
	}

	for _, iterator := range crcindArray {
		standardResponse = append(standardResponse, StandardResponse{
			Category:   iterator.Category,
			Name:       iterator.Name,
			Author:     iterator.Author,
			PreviewURL: iterator.Author,
			Origin:     iterator.Origin,
		})
	}

	// // defer func() {
	// var errorCount int
	// // for index := 0; index < len(channel); index++ {
	// response := <-channel
	// switch response := response.(type) {
	// case error:
	// 	fmt.Println("caso 1 entro en case error")
	// 	errorCount = errorCount + 1
	// case []StandardResponse:
	// 	fmt.Println("caso 2 entro for", response)
	// 	for _, iterator := range response {
	// 		standardResponse = append(standardResponse, iterator)
	// 	}
	// }
	// // }

	// if errorCount == 3 {
	// 	fmt.Println("LOS ERRORES LLEGARON A 3")
	// 	statusReturn = http.StatusExpectationFailed
	// }
	fmt.Println("TERMINA TODO")
	// envio la respuesta al front en JSON
	w.WriteHeader(statusReturn)
	json.NewEncoder(w).Encode(standardResponse)
	// }()
}

func processItunes(criteria string) []StandardResponse {
	// channel chan interface{}
	var standardResponse []StandardResponse
	itunesSlice, err := itunesservicego.FindResults(criteria)
	fmt.Println("processItunes")

	if err != nil {
		fmt.Println("ITUNES FALLO", err)
		// channel <- err
	}

	if len(itunesSlice) > 0 {
		// recorro los resultados
		for _, iterator := range itunesSlice {
			standardResponse = append(standardResponse, StandardResponse{
				Category:   iterator.Category,
				Name:       iterator.Name,
				Author:     iterator.Author,
				PreviewURL: iterator.Author,
				Origin:     iterator.Origin,
			})
		}
	}

	// channel <- standardResponse
	return standardResponse
}

func processTvmaze(criteria string) []StandardResponse {
	// channel chan interface{}
	var standardResponse []StandardResponse
	tvmazeSlice, err := tvmazeservicego.FindResults(criteria)
	fmt.Println("processTvmaze")

	if err != nil {
		fmt.Println("TVMAZE FALLO", err)
		// channel <- err
	}

	if len(tvmazeSlice) > 0 {
		// recorro los resultados
		for _, iterator := range tvmazeSlice {
			standardResponse = append(standardResponse, StandardResponse{
				Category:   iterator.Category,
				Name:       iterator.Name,
				Author:     iterator.Author,
				PreviewURL: iterator.Author,
				Origin:     iterator.Origin,
			})
		}
	}

	// channel <- standardResponse
	return standardResponse
}

func processCrcind(criteria string) []StandardResponse {
	// channel chan interface{}
	var standardResponse []StandardResponse
	crcindSlice, err := crcindservicego.FindResults(criteria)
	fmt.Println("processCrcind")

	if err != nil {
		fmt.Println("CRCIND FALLO", err)
		// channel <- err
	}

	if len(crcindSlice) > 0 {
		// recorro los resultados
		for _, iterator := range crcindSlice {
			standardResponse = append(standardResponse, StandardResponse{
				Category:   iterator.Category,
				Name:       iterator.Name,
				Author:     iterator.Author,
				PreviewURL: iterator.Author,
				Origin:     iterator.Origin,
			})
		}
	}

	// channel <- standardResponse
	return standardResponse
}

// @title Central Service API
// @version 1.0
// @description Central Service API that consume 3 direfent services
// @termsOfService http://swagger.io/terms/
// @contact.name Sergei Teran
// @contact.email steran@tribalworldwide.gt
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	router := mux.NewRouter().StrictSlash(true)

	routes(router)

	port := ":3000"
	fmt.Println("\n servidor corriendo en puerto " + port)
	log.Fatal(http.ListenAndServe(port, router))
}
