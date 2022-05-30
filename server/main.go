package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/*
rest api
get post put(полностью обновить) delete

GET /lists

{
   "lists":[
      {
         "id":1,
         "name":"список1"
      },
      {
         "id":2,
         "name":"список2"
      }
   ]
}


GET /lists/{id}
{
   "todos":[
      {
         "id":1,
         "todo":"умри",
         "done":false
      },
      {
         "id":2,
         "todo":"поешь",
         "done":false
      }
   ]
}

POST /lists - создание  списка
{
	"name":"spisok pisok"
}
{
	"id":3
}

POST /lists/{id}
{
	"todo":"покакай"
}
{
	"id":1
}

POST /lists/{list_id}/{id}/done


*/
func main() {
	fmt.Println("hello pidors")

	http.HandleFunc("/lists", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		rw.Write([]byte(`[
            {
               "id":1,
               "name":"список1"
            },
            {
               "id":2,
               "name":"список2"
            }
         ]`))
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Access-Control-Allow-Origin", "*")
		if strings.HasPrefix(r.RequestURI, "/lists") {
			var rparts []string = strings.Split(r.RequestURI, "/")
			i := rparts[len(rparts)-1]
			switch i {
			case "1":
				rw.Write([]byte(`[
            {
               "id":1,
               "todo":"умри",
               "done":false
            },
            {
               "id":2,
               "todo":"поешь",
               "done":false
            }
         ]`))
			case "2":
				rw.Write([]byte(`[
            {
               "id":1,
               "todo":"Отвезти кота в психлечебницу",
               "done":false
            },
            {
               "id":2,
               "todo":"Сходить на концерт Боба Марли",
               "done":false
            },
            {
               "id":3,
               "todo":"Оживить Моцарта",
               "done":false
            }
         ]`))
			default:
				rw.WriteHeader(http.StatusNotFound)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
