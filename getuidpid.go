package main
 
import (
  "fmt"
  "os"
)
 
func main() {

  uid := os.Getuid()
  fmt.Println("UID: ",uid)

  pid := os.Getpid()
  fmt.Println("PID: ",pid)

}
