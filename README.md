# Real-Time Grid Game with WebSockets

This project is a **real-time grid-based game** where users can interact with a grid in their browser by toggling grid cells. The grid state is synchronized across multiple clients using WebSockets, allowing all connected users to see updates in real time.

## Features

- **Grid Interaction**: Users can toggle cells on the grid by clicking, and the state (color) of the cells is shared across all users.
- **Real-Time Updates**: Changes are instantly broadcasted to all connected clients using WebSockets.
- **User Personalization**: Each user has a name and color, which they can use to modify the grid.
- **Multi-User Support**: Multiple users can connect, interact with the grid, and see the updates from all participants.

## Tech Stack

- **Frontend**: HTML, CSS, JavaScript (Vanilla JS)
  - The grid is dynamically created in the frontend, allowing users to toggle cells via a simple click interface.
- **Backend**: Go (Golang)
  - A WebSocket server built with Go that handles real-time communication and broadcasts grid updates to all connected clients.
- **WebSocket Communication**: The client and server communicate over WebSockets to provide real-time updates on grid state changes.

## How It Works

1. **Client-side**: Users log in, interact with the grid by clicking on cells, and send requests to the WebSocket server to toggle cells.
2. **Server-side**: The WebSocket server listens for requests, updates the grid state, and broadcasts the updated state to all connected clients.
3. **Synchronization**: Every client receives the current state of the grid in real time, ensuring that all users see the same grid.

## How to Run the Project

### Backend

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/your-repository-name.git
   cd your-repository-name
   ```
2. Install Go dependencies .
3. Run the WebSocket server:

   ```bash
   go run main.go

   ```

### Frontend

1. Open **loocalhost:8080/login.html** in your browser.
2. Log in with your username and color.
3. Start interacting with the grid!

## Project Structure

```bash
/root-directory
│
├── /public
│   ├── index.html        # App HTML page
│   ├── login.html        # Login HTML page
│   ├── style.css         # Styling for the grid
│   └── app.js            # JavaScript logic for WebSocket connection and grid interaction
│   └── login.js
│
├── main.go          # WebSocket server in Go
│
└── README.md             # Project documentation

```

## Contributions

Feel free to open issues or submit pull requests if you find bugs or want to contribute.
