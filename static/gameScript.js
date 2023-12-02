guesses = 0;
letters = -1;
const wordOfTheDay = "HOME";

document.addEventListener("keyup", (event) => {
    if(guesses >= 4) {
        console.log("No guesses left.")
        return
    }

    let keyPressValue = String(event.key)

    switch(keyPressValue) {
        case "Backspace":
            backspacePress()
            break
        case "Enter":
            enterPress()
            break
        case keyPressValue:
            let charValue = keyPressValue.charAt(0)
            if(charValue.toUpperCase() !== charValue) {
                keyPress(keyPressValue.toUpperCase())
            }
            break
        default:
            return
    }
})

function keyPress(input) {
    if(letters >= 3 || guesses >= 4) {
        // If we have reached the maximum amount of letters, return since we don't want the letter to change.
        return
    }
    letters++
    var guessTable = document.getElementById("guess-table")
    var tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(letters)
    tile.style.border = "2px solid #565758"
    tile.innerHTML = input
}

function colorKey(letter, colorID) {
    for(const key of document.getElementsByClassName("key")) {
        if(key.innerHTML === letter) {
            switch(colorID) {
                case 0:
                    key.style.backgroundColor = "rgb(83, 141, 78)"
                    return
                case 1:
                    if(key.style.backgroundColor === "rgb(83, 141, 78)") {
                        return
                    }
                    key.style.backgroundColor = "rgb(181, 159, 59)"
                    return
                case 2:
                    if(key.style.backgroundColor === "rgb(83, 141, 78)" || key.style.backgroundColor === "rgb(181, 159, 59)") {
                        return
                    }
                    key.style.backgroundColor = "rgb(58, 58, 60)"
                    return
            }
        }
    }
}

function enterPress() {
    var guessTable = document.getElementById("guess-table")

    if(letters < 3 || guesses >= 4) {
        console.log("Guessed word is less than 4 letters.");
        return
    }

    let wordGuessed = ""
    var tile
    for(let i = 0; i < 4; i++) {
        tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(i)
        wordGuessed += tile.innerHTML
    }

    // if wordGuessed is not in dictionary, shake row, send notification word isn't valid, then return.

    let letterGuess;
    let tempWordOfTheDay = wordOfTheDay
    let lettersNotInWord = new Array(4).fill(true);
    for(let i = 0; i < 4; i++) {
        letterGuess = wordGuessed[i]
        if(letterGuess === wordOfTheDay[i]) {
            guessTable.rows.item(guesses).getElementsByClassName("tile")
                .item(i).style.backgroundColor = "#538d4e" // Letter found in correct place.
            colorKey(letterGuess, 0)
            lettersNotInWord[i] = false

            tempWordOfTheDay = tempWordOfTheDay.split("")
            tempWordOfTheDay[i] = "!"
            tempWordOfTheDay = tempWordOfTheDay.join("")
        }
    }

    for(let guessIndex = 0; guessIndex < 4; guessIndex++) {
        tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(guessIndex)
        letterGuess = wordGuessed[guessIndex]
        if(tempWordOfTheDay.includes(letterGuess) && tempWordOfTheDay[guessIndex] !== "!") {
            tile.style.backgroundColor = "#b59f3b" // Letter belongs in word, but not in position.
            colorKey(letterGuess, 1)
            lettersNotInWord[guessIndex] = false

            tempWordOfTheDay = tempWordOfTheDay.split("")
            tempWordOfTheDay[tempWordOfTheDay.indexOf(letterGuess)] = "?"
            tempWordOfTheDay = tempWordOfTheDay.join("")
        }
    }

    for(let i = 0; i < 4; i++) {
        if(lettersNotInWord[i] === true) {
            tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(i)
            tile.style.backgroundColor = "#3a3a3c"
            colorKey(tile.innerHTML, 2)
        }
    }
    letters = -1
    guesses++
    if(wordOfTheDay === wordGuessed) {
        guesses += 5
    }
}

function backspacePress() {
    if(letters === -1 || guesses >= 4) {
        console.log("Returned since there were no letters.")
        return
    }
    var guessTable = document.getElementById("guess-table")
    var tile = guessTable.rows.item(guesses).getElementsByClassName("tile").item(letters)
    tile.style.border = "2px solid #3a3a3c"
    tile.innerHTML = ""
    letters--
}

window.onload = function() {
    fetch('/check-cookie')
        .then(response => {
            if (response.status === 200) {
                // The session is valid, change some elements on the document
                document.getElementById('dp1').style.display = 'none';
                document.getElementById('dp2').style.display = '';
                document.getElementById('dp3').style.display = '';
            } else {
                // The session is not valid, handle the error
                console.error('Invalid session');
                document.getElementById('dp1').style.display = '';
                document.getElementById('dp2').style.display = 'none';
                document.getElementById('dp3').style.display = 'none';
            }
        })
        .catch(error => {
            // There was an error with the fetch request, handle the error
            console.error('Fetch error: ', error);
        });
};