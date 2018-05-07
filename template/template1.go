package template

import (
	"log"
	"html/template"
	"net/http"
	"fmt"
)
type Package struct {
	Name string
	NumFuncs int
	NumVars int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("go-web").Parse(`
package name: {{.Name}}
Number of functions:{{.NumFuncs}}
Number of varaibles: {{.NumVars}}`)
		if err != nil {
			//打印到终端，研究一下区别
			fmt.Fprint(w, "Parse %v", err)
			return
		}
		err = tmpl.Execute(w, &Package{
			Name:"go-web",
			NumFuncs : 12,
			NumVars: 1200,
		})
		if err != nil {
			log.Fatalf("Execute: %v", err)
		}
	})

	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
