package main

import (
	"fmt"
	"./nmapLib"
	"net/http"
	"encoding/json"
	"os/exec"
)



func main() {
	exec.Command("open", "http://127.0.0.1:3000")



	fmt.Println("go server running \n")
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:3000", nil)



}


func serveRest(w http.ResponseWriter, r *http.Request) {
  response, err := getJsonResponse()
  if err != nil {
    panic(err)
  }

  fmt.Fprintf(w, string(response))
  fmt.Printf("json served! \n")
}

func getJsonResponse() ([]byte, error) {
	var input string
	fmt.Println("Enter network subnet to scan: ")
	fmt.Scanln(&input)
	fmt.Println("scanning...")
	hosts, err := nmap.ScanOpenTcpPorts(input)
	fmt.Println("scanned!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hosts)

  return json.MarshalIndent(hosts, "", "  ")
}
