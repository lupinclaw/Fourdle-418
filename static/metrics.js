async function getMetrics() {
    let gamesPlayed = await getNumberOfGames();
    let gamesWon = await getNumberOfGamesWon();
    getWinRate(gamesPlayed, gamesWon);
    getRecentWords();
}

function getNumberOfGames() {
    console.log("[INFO] sending http request for getNumberOfGames");
    let gamesPlayed = document.getElementById("gamesPlayed");
    // Make a request to your Go server
    return fetch('http://localhost:8080/get-number-of-games')
        .then(response => response.text())
        .then(data => {
            // Update the result div with the response from the server
            gamesPlayed.textContent = data;
            return parseFloat(data);
        })
        .catch(error => console.error('Error:', error));
}

function getNumberOfGamesWon() {
    console.log("[INFO] sending http request for getNumberOfGamesWon");
    let gamesWon = document.getElementById("gamesWon");
    // Make a request to your Go server
    return fetch('http://localhost:8080/get-number-of-games-won')
        .then(response => response.text())
        .then(data => {
            // Update the result div with the response from the server
            gamesWon.textContent = data;
            return parseFloat(data);
        })
        .catch(error => console.error('Error:', error));
}

function getWinRate(gamesPlayed, gamesWon) {
    let winRate = document.getElementById("winRate");
    let rate = (gamesWon / gamesPlayed) * 100;
        winRate.textContent = rate.toFixed(1) + "%";
}

function getRecentWords() {
    let div = document.getElementById("seenWords");
    fetch('http://localhost:8080/get-recent-words')
        .then(response => response.text())
        .then(data => {
            // Update the result div with the response from the server
            div.textContent = data;
        })
        .catch(error => console.error('Error:', error));
}

function getWords() {
    let div = document.getElementById("seenWords");
    fetch('http://localhost:8080/get-words')
        .then(response => response.text())
        .then(data => {
            // Update the result div with the response from the server
            div.textContent = data;
        })
        .catch(error => console.error('Error:', error));
}