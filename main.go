package main
import (
  "fmt"
  "net/http"
  "encoding/json"
  //"os"
  "os/user"
  //"github.com/vmware/govmomi"
)
 
const (
	//envURL      = "GOVMOMI_URL"
	envUserName = "GOVMOMI_USERNAME"
	envPassword = "GOVMOMI_PASSWORD"
	envInsecure = "GOVMOMI_INSECURE"
)



func baseDeploy(){

}

func vinfo(){
	var envURL string 
	fmt.Println("Enter your vSphere or ESXI instance for VM deployment")
	fmt.Scanln(&envURL)
	return
}


func Welcome(w http.ResponseWriter, req *http.Request) {
 
	fmt.Fprintf(w, "Welcome! Starting the Domain Deployer. \n")
   
  }
func main() {
	user, err := user.Current()
    if err != nil {
		panic(err)
	}
  fmt.Println("Hello " + user.Name)
  fmt.Println("Starting http server.")
  
  vinfo()

// Register handler function
 
http.HandleFunc("/welcome", Welcome)
 
  fmt.Println("Go to localhost:8080/welcome To terminate press CTRL+C")
 
// Start server
 
  http.ListenAndServe(":8080", nil)
}