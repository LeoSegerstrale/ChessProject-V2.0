const board = document.getElementById("board")

let selectedPiece = null;
let selectedSquare = null;


let enPassant;
let validMoves = [];
let rookLocs = [];
let WCastleAv = [true,true,true]
let BCastleAv = [true,true,true]
let WKingLoc;
let BKingLoc;
let oppColour;

let currColour = prompt("What do u wanna promote to, (sleepy leo cant be asked to write a ui for this srry) :")

if (currColour === "white"){
    currcolour = "white-piece"
    oppColour = "black-piece"
} else{
    oppColour = "white-piece"
    currColour = "black-piece"
}

const pieceMap = {
    "♟": "pawn",
    "♜": "rook",
    "♞": "knight",
    "♝": "bishop",
    "♛": "queen",
    "♚": "king"
};

for (let row = 0; row < 8; row++) {

    let colourP;

    if (row === 0 || row === 1){
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
        if (row === 1 || row === 6){

            piece.innerText = "♟";
            piece.classList.add(colourP);
            square.appendChild(piece);

        } else if (row === 0 ||  row === 7){

            if (col === 0 || col === 7){
                piece.innerText = "♜";
                piece.classList.add(colourP);
                square.appendChild(piece);
                rookLocs.push(squareName);
            } else if (col === 1 || col === 6){

                piece.innerText = "♞";
                piece.classList.add(colourP);
                square.appendChild(piece);
            } else if (col === 2 || col === 5){

                piece.innerText = "♝";
                piece.classList.add(colourP);
                square.appendChild(piece);
            } else if (col === 3){

                piece.innerText = "♛";
                piece.classList.add(colourP);
                square.appendChild(piece);
            } else{

                piece.innerText = "♚";
                piece.classList.add(colourP);
                square.appendChild(piece);
                if (row===7){
                    WKingLoc = `${row}${col}`;
                } else{
                    BKingLoc = `${row}${col}`;
                }
            }
        }

        square.addEventListener("click", () => {
            const pieceInSquare = square.querySelector("span");


            if (!selectedPiece && pieceInSquare && pieceInSquare.classList.contains(currColour)) {
                selectedPiece = pieceInSquare;
                selectedSquare = square;

                let currCastleAv;
                let currKingLoc;

                if (currColour === "white-piece"){
                    currCastleAv = WCastleAv
                    currKingLoc = WKingLoc

                } else {
                    currCastleAv = BCastleAv
                    currKingLoc = BKingLoc
                }
                square.classList.add("selected");


                const boardState = getBoard(pieceMap);
                const fromSquare = selectedSquare.getAttribute("data-square");

                const requestBody = {
                    from: fromSquare,
                    board: boardState,
                    enPassant: enPassant,
                    CastleStatus: currCastleAv,
                    rookLocs: rookLocs,
                    kingLoc: currKingLoc
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


                            squareEl.classList.add("highlight");

                        });
                    })
                    .catch(err => console.error(err));

                return;
            }

            // dos
            if (square === selectedSquare) {
                clearHighlights();
                selectedSquare.classList.remove("selected");
                selectedPiece = null;
                selectedSquare = null;
                return; // done
            }

            // tres is the best
            if (selectedPiece) {

                const fromSquare = selectedSquare.getAttribute("data-square");
                const toSquare = square.getAttribute("data-square");           // destination
                const pieceSymbol = selectedPiece.innerText;

                if (validMoves.includes(toSquare)) {

                    const fromRLCol = parseInt(rookLocs[0][1])
                    const fromRRCol = parseInt(rookLocs[1][1])


                    clearHighlights();

                    selectedSquare.removeChild(selectedPiece);


                    if (pieceInSquare) {
                        square.removeChild(pieceInSquare);
                    } else if (enPassant !== "" && pieceSymbol === "♟" && toSquare[1] !== fromSquare[1] && !pieceInSquare){


                        const enPassantSquare = board.querySelector(`[data-square='${enPassant}']`);
                        const capturedPawn = enPassantSquare.querySelector("span");
                        enPassantSquare.removeChild(capturedPawn);


                    } else if (pieceSymbol === "♚") {

                        if (currColour === "white-piece"){
                            WKingLoc = toSquare
                        }else{
                            BKingLoc = toSquare
                        }

                        const fromCol = parseInt(fromSquare[1]);
                        const toCol = parseInt(toSquare[1]);
                        const row = fromSquare[0];

                        if (toCol - fromCol === 2){
                            const rookFrom = board.querySelector(`[data-square='${row}${fromRRCol}']`);
                            const rookTo = board.querySelector(`[data-square='${row}5']`);
                            const rook = rookFrom.querySelector("span");
                            rookFrom.removeChild(rook);
                            rookTo.appendChild(rook);
                        } else if (fromCol - toCol === 2){
                            const rookFrom = board.querySelector(`[data-square='${row}${fromRLCol}']`);
                            const rookTo = board.querySelector(`[data-square='${row}3']`);
                            const rook = rookFrom.querySelector("span");
                            rookFrom.removeChild(rook);
                            rookTo.appendChild(rook);
                        }

                    }

                    if (pieceSymbol === "♟" && ((currColour === "white-piece" && toSquare[0] === "4" && fromSquare[0] === "6") || (currColour === "black-piece" && toSquare[0] === "3" && fromSquare[0] === "1") )){
                        enPassant = toSquare
                    } else {
                        enPassant = ""
                    }

                    if (pieceSymbol === "♚"){
                        if (currColour === "white-piece"){
                            WCastleAv[1] = false
                        } else {
                            BCastleAv[1] = false
                        }
                    } else if (pieceSymbol === "♜"){
                        if (fromSquare === rookLocs[0] || fromSquare === rookLocs[2]){

                            if (currColour === "white-piece"){
                                WCastleAv[0] = false

                            } else {
                                BCastleAv[0] = false
                            }

                        } else if (fromSquare === rookLocs[1] || fromSquare === rookLocs[3]){

                            if (currColour === "white-piece"){
                                WCastleAv[2] = false
                            } else {
                                BCastleAv[2] = false
                            }

                        }
                    }

                    if (pieceSymbol === "♟"){
                        let promotionRow;

                        if (currColour === "white-piece"){
                            promotionRow = "0"
                        } else{
                            promotionRow = "7"
                        }
                        if (toSquare[0] === promotionRow){
                            const choice = prompt("What do u wanna promote to, (sleepy leo cant be asked to write a ui for this srry) :")
                            selectedPiece.innerText = pieceMap[choice.toLowerCase()] || "♛";
                        }


                    }

                    square.appendChild(selectedPiece);


                    selectedPiece = null;
                    selectedSquare.classList.remove("selected");
                    selectedSquare = null;

                    const boardState = getBoard(pieceMap);

                    if (currColour === "white-piece"){
                        currCastleAv = WCastleAv
                        currKingLoc = WKingLoc

                    } else {
                        currCastleAv = BCastleAv
                        currKingLoc = BKingLoc
                    }

                    const requestBody = {
                        board: boardState,
                        colour: oppColour,
                        enPassant: enPassant,
                        CastleStatus: currCastleAv,
                        rookLocs: rookLocs,
                        kingLoc: currKingLoc
                    };


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



function getBoard(pieceMap) {
    const state = [];



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



