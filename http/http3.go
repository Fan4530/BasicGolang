// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package http
import (
	"net/http"
	"log"
	"time"
	"os"
	"os/signal"
)

type myHandler struct {

}
func (*myHandler) ServeHTTP(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Hello, v3" + r.URL.String()))
}
func main() {
	server := &http.Server {
		Addr:":8080",
		WriteTimeout: 3 * time.Second,
	}

	quit := make(chan os.Signal)
	//研究一下signal和channel的关系
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})//[问题]为什么用{}
	mux.HandleFunc("/bye", sayBye)
	log.Println("Starting the version 3..")


	server.Handler = mux

	go func() {
		<-quit
		if err := server.Close(); err !=nil {
			log.Fatal("Close server", err)
		}
	}()
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}

	log.Print("server exit")
	// 【问题】log 的几个级别再仔细研究一下



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
	time.Sleep(2 * time.Second)
	w.Write([]byte("Hello, this is the version 3"))
	//fmt.Fprintf(w, "Hello dddWorld!")
}

