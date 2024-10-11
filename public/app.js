var user = {
  name: "",
  color: "",
};
function isLogged() {
  const strdata = localStorage.getItem("user");
  const data = JSON.parse(strdata);
  console.log(data);

  if (!data) {
    window.location.href = "/login.html";
  }
  user.name = data.name;
  user.color = data.color;
}
isLogged();
console.log("hellowo");

function getUser() {
  return user;
}
function getName() {
  return user.name;
}
function getColor() {
  return user.color;
}

function showPlayerName() {
  const spn = document.getElementById("playername");

  spn.innerHTML = user.name;
  spn.style.color = user.color;
}
showPlayerName();
// Establish WebSocket connection
const ws = new WebSocket("ws://localhost:8080/ws");

const gridContainer = document.getElementById("grid");
// const gridSize = 20; // Same size as the backend grid
const gridHeight = 20;
const gridWidth = 60;

// Initialize the grid in the frontend
function initializeGrid(cells) {
  gridContainer.innerHTML = ""; // Clear existing grid
  cells.forEach((row, rowIndex) => {
    row.forEach((cell, colIndex) => {
      const cellDiv = document.createElement("div");
      cellDiv.classList.add("grid-cell");
      if (cell) {
        cellDiv.classList.add("black");
      }

      // Add click event to toggle the cell
      cellDiv.addEventListener("click", () => {
        toggleCell(rowIndex, colIndex);
      });

      gridContainer.appendChild(cellDiv);
    });
  });
  // Update grid size in CSS based on the new dimensions
  gridContainer.style.gridTemplateRows = `repeat(${gridHeight}, 1fr)`;
  gridContainer.style.gridTemplateColumns = `repeat(${gridWidth}, 1fr)`;
}

// Toggle cell state (send to server)
function toggleCell(row, col) {
  const request = {
    row: row,
    column: col,
    color: getColor(), // Send player's color
  };
  console.log(request);

  ws.send(JSON.stringify(request)); // Send toggle request to WebSocket server
}
// function toggleCell(row, col) {
//   const request = {
//     row: row,
//     column: col,
//   };
//   console.log(request);

//   ws.send(JSON.stringify(request)); // Send toggle request to WebSocket server
// }

function updateGrid(cells, colors) {
  const cellDivs = document.querySelectorAll(".grid-cell");
  cells.forEach((row, rowIndex) => {
    row.forEach((cell, colIndex) => {
      const cellDiv = cellDivs[rowIndex * gridWidth + colIndex]; // Use gridWidth instead of gridSize
      if (cell) {
        cellDiv.classList.add("black");
        cellDiv.style.backgroundColor = colors[rowIndex][colIndex]; // Apply the correct color
      } else {
        cellDiv.classList.remove("black");
        cellDiv.style.backgroundColor = ""; // Clear the color
      }
    });
  });
}

ws.onmessage = (event) => {
  const gridData = JSON.parse(event.data);
  if (gridContainer.children.length === 0) {
    initializeGrid(gridData.cells);
  } else {
    updateGrid(gridData.cells, gridData.colors); // Pass colors to the updateGrid function
  }
};
// function updateGrid(cells) {
//   const cellDivs = document.querySelectorAll(".grid-cell");
//   cells.forEach((row, rowIndex) => {
//     row.forEach((cell, colIndex) => {
//       const cellDiv = cellDivs[rowIndex * gridSize + colIndex];
//       if (cell) {
//         cellDiv.classList.add("black");
//       } else {
//         cellDiv.classList.remove("black");
//       }
//     });
//   });
// }

// WebSocket connection handling
ws.onopen = () => {
  console.log("Connected to WebSocket");
};

// ws.onmessage = (event) => {
//   const gridData = JSON.parse(event.data);
//   if (gridContainer.children.length === 0) {
//     // Initialize the grid for the first time
//     initializeGrid(gridData.cells);
//   } else {
//     // Update grid after toggling a cell
//     updateGrid(gridData.cells);
//   }
// };

ws.onclose = () => {
  console.log("Disconnected from WebSocket");
};
