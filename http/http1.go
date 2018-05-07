
package http
import (
	"net/http"
	"log"
)



func main() {
	//这俩的区别是
	http.HandleFunc("/", handlerHello)
	log.Println("Starting the version 1..")

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye, this is the version 1"))

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	//为什么这里可以传一个nil
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
// 【接口】w是接口，后面是个指针地址【问题】什么是接口
func handlerHello(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	w.Write([]byte("Hello, this is the version 1"))
	//fmt.Fprintf(w, "Hello dddWorld!")
}

