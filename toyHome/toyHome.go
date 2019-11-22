package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"toyHome/api"
	"toyHome/entities"
)

const HeaderNameDiscovery = "DiscoverAppliancesRequest"
const HeaderNameRequest = "TurnOnRequest"

type server struct{}

/*api (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}*/

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var p entities.Request
	var res entities.Response
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))
	res.Head = p.Head
	switch p.Head.Name {
	case HeaderNameDiscovery:
		res.Head.Name = "DiscoverAppliancesResponse"
		res.Payload, err = api.ResponseDiscovery(p.Payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case HeaderNameRequest:
		res.Head.Name = "TurnOnConfirmation"
		res.Payload, err = api.ResponseRequest(p.Payload)
		//TODO: Not implement yet
	}

	out, _ := json.Marshal(res)
	// Do something with the Person struct...
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	//fmt.Fprintf(w, "Data: %+v", p)
}

//curl --header "Content-Type: application/json" --request POST --data "@discovery.json" http://localhost:8080/xqqqqqq
