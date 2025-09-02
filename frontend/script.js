const board = document.getElementById("board")
console.log(board)

let selectedPiece = null;
let selectedSquare = null;

for (let row = 0; row < 8; row++) { 

    let colourP;

    if (row == 0 || row == 1){
        colourP = "black-piece"
    }else{
        colourP = "white-piece"
    }

    for (let col = 0; col < 8; col++) { 

        
        const square = document.createElement("div");
        square.classList.add("square");
        

        if ((row + col) % 2 === 0) { 
            square.classList.add("light");
         } else { 
            square.classList.add("dark");
        }

        const squareName = `${row}${col}`;

        square.setAttribute("data-square", squareName); 


        const piece = document.createElement("span");
        if (row == 1 || row == 6){

            piece.innerText = "♟";
            piece.classList.add(colourP); 
            square.appendChild(piece); 

        } else if (row == 0 ||  row == 7){
            
            if (col == 0 || col == 7){
                piece.innerText = "♜";
                piece.classList.add(colourP); 
                square.appendChild(piece);  
            } else if (col == 1 || col == 6){

                piece.innerText = "♞";
                piece.classList.add(colourP); 
                square.appendChild(piece);  
            } else if (col == 2 || col == 5){

                piece.innerText = "♝";
                piece.classList.add(colourP); 
                square.appendChild(piece);  
            } else if (col == 3){

                piece.innerText = "♛";
                piece.classList.add(colourP); 
                square.appendChild(piece);  
            } else{

                piece.innerText = "♚";
                piece.classList.add(colourP); 
                square.appendChild(piece);  
            }
        }

        square.addEventListener("click", () => {
            const pieceInSquare = square.querySelector("span");


            if (!selectedPiece && pieceInSquare) {
                selectedPiece = pieceInSquare;
                selectedSquare = square;
                square.classList.add("selected");


                const board = getBoard();
                const fromSquare = selectedSquare.getAttribute("data-square");

                const requestBody = {
                    from: fromSquare,
                    board: board
                };

                fetch("http://localhost:8080/vMoveCheck", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify(requestBody)
                })
                .then(res => res.json())
                .then(data => {
                    console.log(data.validSquares);
                    clearHighlights();

                    data.validSquares.forEach(coords => {
                        const square = board.querySelector(`[data-square='${coords}']`);
                        if (square) {
                            console.log("Found square:", coords, square);
                            square.classList.add("highlight");
                        } else {
                            console.warn("Square not found:", coords);
                        }

                    });
                })
                .catch(err => console.error(err));

                return; // done with this click
            }

            // --- CASE 2: clicked the same piece again to deselect ---
            if (square === selectedSquare) {
                clearHighlights();
                selectedSquare.classList.remove("selected");
                selectedPiece = null;
                selectedSquare = null;
                return; // done
            }

            // --- CASE 3: moving the selected piece to a new square ---
            if (selectedPiece) {
                clearHighlights();
                // Remove piece from old square
                selectedSquare.removeChild(selectedPiece);

                // Remove piece in destination if any (capturing)
                if (pieceInSquare) {
                    square.removeChild(pieceInSquare);
                }

                // Append the piece to new square
                square.appendChild(selectedPiece);

                // Clear selection
                selectedPiece = null;
                selectedSquare.classList.remove("selected");
                selectedSquare = null;
            }
        });

        


        board.appendChild(square);

     }
}


function clearHighlights() {
    document.querySelectorAll(".highlight").forEach(square => {
        square.classList.remove("highlight");
    });
}



function getBoard() {
    const state = [];

    const pieceMap = {
        "♟": "pawn",
        "♜": "rook",
        "♞": "knight",
        "♝": "bishop",
        "♛": "queen",
        "♚": "king"
    };

    for (let row = 0; row < 8; row++) {

        const newRow = [];
        for (let col = 0; col < 8; col++) {
            newRow.push(null);
        }
        state.push(newRow);
    }

    for (let row = 0; row < 8; row++) {
        for (let col = 0; col < 8; col++) {
            const square = board.querySelector(`[data-square='${row}${col}']`);
            const piece = square.querySelector("span");
            if (piece) {
                state[row][col] = {
                    piece: pieceMap[piece.innerText],      // or a type like 'bishop', 'pawn'
                    colour: piece.classList.contains("white-piece") ? "white" : "black"
                };
            }
        }
    }

    return state;
}



