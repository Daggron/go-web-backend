package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
	"html/template"
)

var templates *template.Template
var clients *redis.Client

func main(){
	templates = template.Must(template.ParseGlob("templates/*.html"));
	clients = redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
	})

	router := mux.NewRouter();
	router.HandleFunc("/",handleGetRequest).Methods("GET");
	router.HandleFunc("/",handlePostRequest).Methods("POST");

	http.Handle("/",router);
	http.ListenAndServe(":8000",nil);
}

func handleGetRequest(w http.ResponseWriter , r * http.Request){

	comments,err := clients.LRange("cooments",0,10).Result()

	if err!=nil{
		return ;
	}

	templates.ExecuteTemplate(w,"post.html",comments);

}


func handlePostRequest(w http.ResponseWriter, r * http.Request){

	r.ParseForm()
	comments := r.PostForm.Get("comment");
	clients.LPush("cooments",comments);
	http.Redirect(w,r,"/",302)

}