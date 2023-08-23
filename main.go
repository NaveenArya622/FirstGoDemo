package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//if err := r.ParseForm(); err != nil {
		//	fmt.Fprintf(w, "ParseForm() err: %v", err)
		//	return
		//}
		var emp Employee
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err != nil {
			fmt.Printf("error occured while decoding")
		}
		EmpName := emp.Name
		EmpAge := emp.Age
		fmt.Printf("Name of Employee is : %s", EmpName)
		fmt.Printf("\nAge of Employee is : %d", EmpAge)
		fmt.Fprintf(w, "Hello Mr./Miss./Mrs. %s - %dY !", EmpName, EmpAge)
	})
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
