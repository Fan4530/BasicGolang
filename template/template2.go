package template

import (
	"log"
	"html/template"
	"net/http"
	"fmt"
	"strconv"
)
type Package struct {
	Name string
	NumFuncs int
	NumVars int
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("template3.html")
		if err != nil {
			//打印到终端，研究一下区别
			fmt.Fprint(w, "Parse %v", err)
			return
		}
		//从request的URL中获取score
		score := r.FormValue("score")
		num, _ := strconv.Atoi(score)
		//我怎么才能两个都打印在上面呢？？？score和package
		err = tmpl.Execute(w, num)
		// Execute两个参数，一个w一个r，r是 struct可以自己定义也可以用object
		//err = tmpl.Execute(w, r)
		//不是很懂r 里面有一个get是怎么搞的？？？

		//err = tmpl.Execute(w, &Package{
		//	Name: "go web",
		//	NumFuncs: 12,
		//	NumVars: 1200,
		//})
		if err != nil {
			log.Fatalf("Execute: %v", err)
			return
		}
	})

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
