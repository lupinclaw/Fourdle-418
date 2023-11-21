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
    tile.style.border = "2px solid #565758"
    tile.innerHTML = input
}



function enterPress()
{
    var guessTable = document.getElementById("guess-table")

    if(letters < 3)
    {
        console.log("Guessed word is less than 4 letters.");
        return;
    }

    let wordGuessed = ""
    var tile;

    for(let i = 0; i < 4; i++) {
        tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(i)
        wordGuessed += tile.innerHTML;
    }

    for(let i = 0; i < 4; i++)
    {
        tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(i)
        if(wordOfTheDay[i] === wordGuessed[i])
        {
            tile.style.backgroundColor = "#538d4e";
            // style keyboard

        }
        else if(wordOfTheDay.includes(tile.innerHTML))
        {
            tile.style.backgroundColor = "#b59f3b";
            // style keyboard
        }
        else
        {
            tile.style.backgroundColor = "#3a3a3c"
            // style keyboard
        }
    }
    letters = -1;
    guesses++;
    if(wordOfTheDay === wordGuessed)
    {
        // Disable buttons (better way to do this?).
        const buttons = document.getElementsByClassName("key");
        for (let i = 0; i < buttons.length; i++) {
            buttons[i].disabled=true;
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
    tile.style.border = "2px solid #3a3a3c"
    tile.innerHTML = ""
    letters--;
}