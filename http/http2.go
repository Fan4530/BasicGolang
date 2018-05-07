// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package http
import (
	"net/http"
	"log"
)

type myHandler struct {

}
func (*myHandler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Hello, v2" + r.URL.String()))
}
func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})//[问题]为什么用{}
	mux.HandleFunc("/bye", sayBye)
	log.Println("Starting the version 2..")
	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,


	log.Fatal(http.ListenAndServe(":8080", mux))
	//为什么这里加个mux
	//因为在version 1里面最简单的http.HandlerFunc创建的时候 里面先默认会创造一个defaultmux
	//所以我猜如果你传的mux是nil，会将创建一个server给你
	//	server := &Server{Addr: addr, Handler: handler}
	//这句话的意思是什么，是如果是nil的话你传进去，否则传出来么
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
// 【接口】w是接口，后面是个指针地址【问题】什么是接口
func sayBye(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	w.Write([]byte("Hello, this is the version 1"))
	//fmt.Fprintf(w, "Hello dddWorld!")
}

