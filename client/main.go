package main

// https://www.raylib.com/cheatsheet/cheatsheet.html
// https://github.com/gen2brain/raylib-go/blob/master/examples/games/snake/main.go
// https://medium.com/@viktordev/socket-programming-in-go-write-a-simple-tcp-client-server-c9609edf3671

import (
	"bufio"
	"net"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "Raylib with bufio and net")
	defer rl.CloseWindow()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		rl.TraceLog(rl.LogError, "Error connecting to server: %s", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeyD) {
			// Example of writing to the socket
			message := "Hello server!"
			_, err = writer.WriteString(message + "\n")
			if err != nil {
				rl.TraceLog(rl.LogError, "Error writing to server: %s", err)
				return
			}
			writer.Flush()

			// Example of reading from the socket
			response, err := reader.ReadString('\n')
			if err != nil {
				rl.TraceLog(rl.LogError, "Error reading from server: %s", err)
				return
			}
			rl.TraceLog(rl.LogInfo, "Server response: %s", response)
		}

		rl.EndDrawing()
	}
}
