package main
import ("fmt"
  "net"
  "time"
  "strconv"
)
func checkError(err error){
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }

}


func reciever(port int){
  serverAddr, err := net.ResolveTCPAddr("tcp4", ":" + strconv.Itoa(port))
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }
  listener, err := net.ListenTCP("tcp", serverAddr)
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }

  for {
    time.Sleep(1*time.Millisecond)
    connRcv, err := listener.AcceptTCP();
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    } else {
      fmt.Println("Accepted connection")
    }
    go reader(connRcv)
  }

  /*
  err = listener.Close()
  if err!= nil {
    fmt.Printf("Some error %v", err)
  }
  err = connRcv.Close()
  if err!= nil {
    fmt.Printf("Some error %v", err)
  }
  */
}

func reader(conn *net.TCPConn){
  inBuf := make([]byte, 1024)
  for {
    time.Sleep(1*time.Millisecond)
    n, _ := conn.Read(inBuf)
    fmt.Println("Received:\n\t", string(inBuf[:n]), "\n")
  }
}

func sender(){
  serverAddr, err := net.ResolveTCPAddr("tcp4", "10.100.23.242:34933")
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }

  conn, err := net.DialTCP("tcp", nil, serverAddr)
  if err != nil {
      fmt.Printf("Some error %v", err)
      return
  }
  go reader(conn)
  /*n, err := conn.Write([]byte("asdf\x00"))
  fmt.Println(n, err)

  n, err = conn.Write([]byte("asdf mk2\x00"))
  fmt.Println(n, err)
  //conn.Write([]byte("sdfg\x00"))
*/
  //defer conn.Close()
  //conn.Write("Connect to: "10.100.23.156":"33546"\0")
  n, err := conn.Write([]byte("Connect to: 10.100.23.156:23453\x00"))
  fmt.Println(n, err)

  /*
  string := "Hei Andrine \x00"

  for {
    conn.Write([]byte(string))
    time.Sleep(500* time.Millisecond)
  }
  */


}

func main() {
  go sender()
  go reciever(23453)
  //go reciever(30000)


  select {}

}
