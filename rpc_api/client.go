package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// Conectar con el servidor RPC
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error al conectar con el servidor RPC:", err)
		return
	}
	defer client.Close()

	// Llamada al método SayHello en el servidor
	var request string
	var reply string
	err = client.Call("HelloService.SayHello", request, &reply)
	if err != nil {
		fmt.Println("Error al llamar al servicio RPC:", err)
		return
	}

	// Mostrar la respuesta del servidor
	fmt.Println("Respuesta del servidor:", reply)
}
