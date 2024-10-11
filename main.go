package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// ToggleRequest represents a request to toggle a cell
type ToggleRequest struct {
    Row    int `json:"row"`
    Column int `json:"column"`
    Color string `json:"color"`
}

// GameGrid represents the grid (NxN matrix)
type GameGrid struct {
    Cells [][]bool `json:"cells"` // true = black, false = white
	Colors [][]string `json:"colors"`
    Rows  int      `json:"rows"`
    Cols  int      `json:"cols"`
}

// NewGameGrid creates a new game grid with default white cells
func NewGameGrid(rows, cols int) *GameGrid {
    cells := make([][]bool, rows)
    colors := make([][]string, rows)
    for i := range cells {
        cells[i] = make([]bool, cols)
        colors[i] = make([]string, cols) // Initialize color grid
    }
    return &GameGrid{
        Cells:  cells,
        Colors: colors,
        Rows:   rows,
        Cols:   cols,
    }
}
// func NewGameGrid(rows, cols int) *GameGrid {
//     cells := make([][]bool, rows)
//     for i := range cells {
//         cells[i] = make([]bool, cols)
//     }
//     return &GameGrid{
//         Cells: cells,
//         Rows:  rows,
//         Cols:  cols,
//     }
// }

// ToggleCell toggles the state of a cell between white and black
func (g *GameGrid) ToggleCell(row, col int, color string) {
    if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
        g.Cells[row][col] = !g.Cells[row][col]
        if g.Cells[row][col] { // Only assign color when the cell is toggled
            g.Colors[row][col] = color
        } else {
            g.Colors[row][col] = "" // Clear color when untoggled
        }
    }
}
// func (g *GameGrid) ToggleCell(row, col int) {
//     if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
//         g.Cells[row][col] = !g.Cells[row][col]
//     }
// }

// WebSocket upgrader
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

var (
    clients  = make(map[*websocket.Conn]bool) // Connected clients
    gameGrid = NewGameGrid(20, 60)            // Create a 10x10 grid
    mu       sync.Mutex                      // Mutex to protect shared state
)

// Broadcast grid state to all connected clients
func broadcastGridState() {
    mu.Lock()
    defer mu.Unlock()

    gridData, err := json.Marshal(gameGrid)
    if err != nil {
        log.Println("Error marshalling grid:", err)
        return
    }

    for client := range clients {
        err := client.WriteMessage(websocket.TextMessage, gridData)
        if err != nil {
            log.Printf("Error sending message to client: %v", err)
            client.Close()
            delete(clients, client)
        }
    }
}
// func broadcastGridState() {
//     mu.Lock()
//     defer mu.Unlock()

//     gridData, err := json.Marshal(gameGrid)
//     if err != nil {
//         log.Println("Error marshalling grid:", err)
//         return
//     }

//     for client := range clients {
//         err := client.WriteMessage(websocket.TextMessage, gridData)
//         if err != nil {
//             log.Printf("Error sending message to client: %v", err)
//             client.Close()
//             delete(clients, client)
//         }
//     }
// }

// Send the current grid state to a newly connected client
func sendCurrentGridState(conn *websocket.Conn) {
    mu.Lock()
    defer mu.Unlock()

    gridData, err := json.Marshal(gameGrid)
    if err != nil {
        log.Println("Error marshalling grid:", err)
        return
    }

    err = conn.WriteMessage(websocket.TextMessage, gridData)
    if err != nil {
        log.Printf("Error sending grid data to new client: %v", err)
        conn.Close()
        delete(clients, conn)
    }
}

// Handle WebSocket connections
func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()

    clients[conn] = true

    // Send the current grid state to the new client
    sendCurrentGridState(conn)

    for {
		var req ToggleRequest
        fmt.Println("new req")
		err := conn.ReadJSON(&req)
		if err != nil {
			log.Printf("Error reading from WebSocket: %v", err)
			delete(clients, conn)
			break
		}
	
		// Update the grid by toggling the cell with the player's color
		mu.Lock()
        fmt.Println("toggeling cell")
		gameGrid.ToggleCell(req.Row, req.Column, req.Color)
		mu.Unlock()
	
		// Broadcast the updated grid state to all clients
		broadcastGridState()
	}
}

func main() {
    // Serve WebSocket connections
    http.HandleFunc("/ws", handleConnections)

    // Serve the frontend (if needed)
    fs := http.FileServer(http.Dir("./public"))
    http.Handle("/", fs)

    log.Println("Server starting on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}





// var addr = flag.String("addr", ":8080", "http service address")
// flag.Parse()
// fmt.Println(*addr)