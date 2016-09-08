package main

import (
	"encoding/json"
	"fmt"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func authenticate() (*gophercloud.ProviderClient, error) {
	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return nil, err
	}

	return openstack.AuthenticatedClient(authOpts)
}

func times(n int, f func(*sync.WaitGroup)) {
	var wg sync.WaitGroup

	wg.Add(n)
	for i := 0; i < n; i++ {
		go f(&wg)
	}
	wg.Wait()
}

func scenario1(wg *sync.WaitGroup) {
	defer log.Print("Scenario1 stopped.")
	log.Print("Scenario1 launched.")
	wg.Done()
}

var scenarioMap = map[string]func(*sync.WaitGroup){
	"scenario1": scenario1,
}

func yieldError(w http.ResponseWriter, errCode int, prefix string, err error) {
	errMsg := fmt.Sprintf("%s%s", prefix, err)
	log.Print(errMsg)
	w.WriteHeader(500)
	w.Write([]byte(errMsg))
}

type Command struct {
	Command string
	LoadSize *int
}

func goHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		log.Print("GET /go : Health-check passed.")
		w.Write([]byte("GET /go : Health-check passed."))
	} else if req.Method == "POST" {
		var cmd *Command

		log.Print("POST /go : received.")

		body, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			yieldError(w, 500, "ERROR: ", err)
			return
		}

		err = json.Unmarshal(body, &cmd)
		if err != nil {
			yieldError(w, 500, "ERROR: ", err)
			return
		}

		proc, ok := scenarioMap[cmd.Command]
		if !ok {
			yieldError(w, 500, "ERROR: ", fmt.Errorf("Command not supported: %s", cmd.Command))
			return
		}

		n := 1000
		if cmd.LoadSize != nil {
			n = *cmd.LoadSize
		} else {
			log.Print("POST /go : command received without LoadSize; assuming 1000.")
			w.Write([]byte("POST /go : command received without LoadSize; assuming 1000."))
		}
		log.Printf("POST /go : Kicking off %d workers", n);
		times(n, proc)
		log.Print("POST /go : completed.")
	} else {
		log.Print(fmt.Sprintf("%s /go : unknown method", req.Method))
		w.WriteHeader(500)
		io.WriteString(w, fmt.Sprintf("%s /go : unknown method", req.Method))
	}
}

func main() {
	log.Print("This is the example load generator for Exercissus\n")
	cl, err := authenticate()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Gophercloud's token: %s\n", cl.TokenID)
	http.HandleFunc("/go", goHandler)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
