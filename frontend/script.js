const board = document.getElementById("board")
console.log(board)

let selectedPiece = null;
let selectedSquare = null;

let currColour = "white-piece"
let enPassant;
let validMoves = [];

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


            if (!selectedPiece && pieceInSquare && pieceInSquare.classList.contains(currColour)) {
                selectedPiece = pieceInSquare;
                selectedSquare = square;


                square.classList.add("selected");


                const boardState = getBoard();
                const fromSquare = selectedSquare.getAttribute("data-square");
                const requestBody = {
                    from: fromSquare,
                    board: boardState,
                    enPassant: enPassant
                };

                fetch("http://localhost:8080/vMoveCheck", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify(requestBody)
                })
                .then(res => res.json())
                .then(data => {
                    clearHighlights();

                    validMoves = data.validSquares

                    data.validSquares.forEach(coords => {
                        const squareEl = board.querySelector(`[data-square='${coords}']`);

                        console.log("Found square:", coords, squareEl);
                        squareEl.classList.add("highlight");

                    });
                })
                .catch(err => console.error(err));

                return;
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

                const fromSquare = selectedSquare.getAttribute("data-square");
                const toSquare = square.getAttribute("data-square");           // destination
                const pieceSymbol = selectedPiece.innerText;

                if (validMoves.includes(toSquare)) {



                    clearHighlights();
                    // Remove piece from old square
                    selectedSquare.removeChild(selectedPiece);

                    // Remove piece in destination if any (capturing)
                    if (pieceInSquare) {
                        square.removeChild(pieceInSquare);
                    } else if (enPassant != "" && pieceSymbol == "♟" && toSquare[1] != fromSquare[1] && !pieceInSquare){


                        const enPassantSquare = board.querySelector(`[data-square='${enPassant}']`);
                        const capturedPawn = enPassantSquare.querySelector("span");
                        enPassantSquare.removeChild(capturedPawn);


                    }

                    if (pieceSymbol == "♟" && ((currColour == "white-piece" && toSquare[0] == "4" && fromSquare[0] == "6") || (currColour == "black-piece" && toSquare[0] == "3" && fromSquare[0] == "1") )){
                        enPassant = toSquare
                    } else {
                        enPassant = ""
                    }

                    // Append the piece to new square
                    square.appendChild(selectedPiece);

                    // Clear selection
                    selectedPiece = null;
                    selectedSquare.classList.remove("selected");
                    selectedSquare = null;
                    if (currColour == "white-piece"){
                        currColour = "black-piece"
                    } else{
                        currColour = "white-piece"
                    }


                }


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
                    piece: pieceMap[piece.innerText],
                    colour: piece.classList.contains("white-piece") ? "white" : "black"
                };
            }
        }
    }

    return state;
}



