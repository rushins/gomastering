//Build a web server

package main
import (
   "fmt"
   "net/http"
)



func main() {
   http.ListenAndServe(":8080",nil)
   fmt.Println("Server is listening at port 8080")
   http.HandleFunc("/", handler)
}

func handler (write http.ResponseWriter, req *http.Request) {
   fmt.Fprint(write, "<h1>Login</h1><form action='/log-in/' method='POST'> Password:<br> <input type='password' name='pass'><br> <input type='submit' value='Go!'></form>")
}
