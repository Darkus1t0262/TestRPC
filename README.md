# TestRPC
This is a simple example demonstrating how to create a basic RPC (Remote Procedure Call) server and client using the Go programming language.

## Overview

In this project, we have two main components:

1. **RPC Server** (`main.go`): A simple server that listens for TCP connections on port `1234` and exposes a method `SayHello` that returns a greeting message.
2. **RPC Client** (`client.go`): A client that connects to the RPC server, calls the `SayHello` method, and displays the server's response.

The server and client communicate using the `net/rpc` package in Go.

## Prerequisites

- Go 1.18 or higher

## Getting Started

### 1. Clone the repository


1. Clone this repository to your local machine using:

    ```bash
    git clone https://github.com/Darkus1t0262/TestRPC
    cd rpc_api
2. Run the RPC Server
    ```bash
    go run main.go
3. Run the RPC Client
    ```bash
    go run client.go

## Screenshot



