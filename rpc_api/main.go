package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Estructura del servicio RPC
type HelloService struct{}

// Método del servicio RPC
func (h *HelloService) SayHello(request string, reply *string) error {
	*reply = "¡Hola Mundo desde RPC!"
	return nil
}

func main() {
	// Registro del servicio RPC
	rpc.Register(new(HelloService))

	// Configuración del servidor
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor RPC ejecutándose en el puerto 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		go rpc.ServeConn(conn) // Maneja cada conexión en una goroutine
	}
}
