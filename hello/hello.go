/* This is not supposed to be an example of a great Go application.
 * I just wanted to use a language with a compiler to show the
 * benefit of statically-linked binaries. */
package main
import (
	"flag"
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func main() {
	listenAddr := flag.String("listen-address", ":8080", "The address to bind the HTTP server to.")
	flag.Parse()
	http.HandleFunc("/hello", hello)
	fmt.Printf("Listening on %s\n", *listenAddr)
    http.ListenAndServe(*listenAddr, nil)
}
