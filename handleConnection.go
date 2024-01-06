package handleConnection

import (
	"net"
)

func ConnectionWithServer(IP string, Port string) (conn net.Conn, err error) {
	ServerAddress := IP + ":" + Port

	connection, err := net.Dial("tcp", ServerAddress)
	if err != nil {
		return connection, err
	}
	return
}

func ConnectionWithWindows(IP string, Port string) (connection net.Conn,err error) {
	LocalAddr := IP + ":"+Port

	//建立监听
	Listener,err := net.Listen("tcp",LocalAddr)
	if err != nil {
		return connection,err
	}
	connection,err = Listener.Accept()
	if err != nil {
		return connection,err
		
		
	}else{
		return
	}
}
