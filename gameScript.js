guesses = 0;
letters = -1;
wordOfTheDay = "HOME";

document.addEventListener("keyup", (event) =>
{
    if(guesses >= 4)
    {
        console.log("No guesses left.")
        return
    }

    let keyPressValue = String(event.key)

    switch(keyPressValue)
    {
        case "Backspace":
            backspacePress()
            break
        case "Enter":
            enterPress()
            break
        case keyPressValue:
            let charValue = keyPressValue.charAt(0)
            if(charValue.toUpperCase() !== charValue)
            {
                keyPress(keyPressValue.toUpperCase())
            }
            break
        default:
            return
    }
})

function keyPress(input)
{
    if(letters >= 3)
    {
        // If we have reached the maximum amount of letters, return since we don't want the letter to change.
        return;
    }
    letters++;
    var guessTable = document.getElementById("guess-table")
    var tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(letters)
    tile.innerHTML = input;
}



function enterPress()
{ // seems to add last used letter sometimes?
    if(letters < 3)
    {
        console.log("Guessed word is less than 4 letters.");
        return;
    }

    let wordGuessed = ""

    for(let i = 0; i < 4; i++) {
        var guessTable = document.getElementById("guess-table")
        var tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(i)
        wordGuessed += tile.innerHTML;
    }
    console.log("hii this is the worddd " + wordGuessed)

    tile.style.backgroundColor = "green";

    for(let row of guessTable.rows) {
        for (let cell of row.cells) {
            //cell.getElementsByClassName("tile").item(0).innerHTML = input;

        }
    }
}

function backspacePress()
{
    if(letters === -1)
    {
        console.log("Returned since there were no letters.")
        return;
    }
    var guessTable = document.getElementById("guess-table")
    var tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(letters)
    tile.innerHTML = ""
    letters--;
}