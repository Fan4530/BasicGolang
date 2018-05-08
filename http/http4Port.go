
package main
import (
	"net/http"
	"log"
	"os"
)


//【未解决】这个getEv到底干啥用的？？总是返回不了port number啊
// 【未记录】
func main() {
	//这俩的区别是
	http.HandleFunc("/test", handlerHello)
	log.Println("Starting the version 1..")


	log.Fatal(http.ListenAndServe(":8080", nil))
	//为什么这里可以传一个nil
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
func handlerHello(w http.ResponseWriter, r *http.Request) {
	//message := r.URL.Path;
	// For this case, we will always pipe "Hello World" into the response writer
	w.Write([]byte("Hello, this is the version 1"))
	//fmt.Fprintf(w, message)
	//fmt.Fprintf(w, "Hello dddWorld!")
}

