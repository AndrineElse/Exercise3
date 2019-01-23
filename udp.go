package main
import ("fmt"
  "net"
  "time"
  "strconv"
)
func reciever(port int){
  inBuf := make([]byte, 1024)
  serverAddr, err := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(port))
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }
  conn, err := net.ListenUDP("udp", serverAddr)
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }
  for {
    conn.Read(inBuf)
    fmt.Println(string(inBuf))
  }
}

func sender(){
  serverAddr, err := net.ResolveUDPAddr("udp", "10.100.23.242:20015")
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }

  conn, err := net.DialUDP("udp", nil, serverAddr)
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }
  defer conn.Close()

  string := "Hei Andrine"

  for {
    conn.Write([]byte(string))
    time.Sleep(100* time.Millisecond)
  }


}

func main() {

  go reciever(30000)
  go reciever(20015)
  go sender()

  select {}

}
