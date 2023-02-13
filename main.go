
package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: Iv", err)
		return
	}
	fmt.Fprintf(w, "POST request succesful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
        fmt.Fprintf(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request){
/*
	if r.URL.Path != "/hello"{
		http.Error(w, "404 File Not found", http.StatusNotFound)
		return
	}
*/
	if r.Method != "GET"{
		http.Error(w, "Method us not Supported", http.StatusNotFound)	
		return
	}
	fmt.Fprintf(w, "hello!")

}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)	
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting the Server at port 9090\n")
	if err := http.ListenAndServe(":9090", nil); 
	err != nil {
		log.Fatal(err)
	}
}


