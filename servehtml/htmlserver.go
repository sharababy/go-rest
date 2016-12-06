package main


import(
	"fmt"
	"net/http"
	"httprouter"
	"html/template"

)


func ServeHTML(w http.ResponseWriter, r *http.Request, _ httprouter.Params){


			file:= r.URL.Path

			filename:=file[1:]

			
			
			t,err:= template.ParseFiles(filename+".html")

			if(err!=nil){
				panic(err)
			}
			
			t.Execute(w,nil)

			fmt.Println(" requested ",filename," page...")	
	
		
	}



func main() {
		
		server := httprouter.New()

		server.GET("/addUser",ServeHTML)

		server.ServeFiles("/resources/*filepath", http.Dir("./resources"))


		fmt.Println("waiting at :4747");

		http.ListenAndServe(":4747",server)


}