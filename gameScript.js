guesses = 0;
letters = 0;

function keyPress(input) {
    console.log(input)
    var guessTable = document.getElementById("guess-table")
    var tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(letters)
    console.log(tile);
    tile.innerHTML = input;
    letters++;
}

function enterPress() {
    for(let row of guessTable.rows) {
        for(let cell of row.cells) {
            //cell.getElementsByClassName("tile").item(0).innerHTML = input;

        }
    }
}

function backspacePress() {

}