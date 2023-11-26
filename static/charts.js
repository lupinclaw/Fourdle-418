let gamesWon = {}
let gamesPlayed = {}

// Load the Visualization API and the corechart package.
google.charts.load('current', {'packages':['corechart']});

// Set a callback to run when the Google Visualization API is loaded.
google.charts.setOnLoadCallback(drawWeekChart);
google.charts.setOnLoadCallback(drawMonthChart);
google.charts.setOnLoadCallback(drawWeekLineChart);
google.charts.setOnLoadCallback(drawMonthLineChart);

// Callback that creates and populates a data table,
// instantiates the pie chart, passes in the data and
// draws it.
async function drawWeekChart() {
    let date = new Date();
    let dates = [];
    for (let i = 6; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Games Lost');
    data.addColumn('number', 'Games Won');

    await populateDictionaries()
    data.addRows([
        [dates[0], gamesPlayed[dates[0]], gamesWon[dates[0]]],
        [dates[1], gamesPlayed[dates[1]], gamesWon[dates[1]]],
        [dates[2], gamesPlayed[dates[2]], gamesWon[dates[2]]],
        [dates[3], gamesPlayed[dates[3]], gamesWon[dates[3]]],
        [dates[4], gamesPlayed[dates[4]], gamesWon[dates[4]]],
        [dates[5], gamesPlayed[dates[5]], gamesWon[dates[5]]],
        [dates[6], gamesPlayed[dates[6]], gamesWon[dates[6]]]
    ]);

    // Set chart options
    var options = {
        'title': 'Games Played Over Time (Last 7 Days)',
        'titleTextStyle': {'color': '#dadada', 'fontSize': 24},
        'width': 700,
        'height': 500,
        'legend': {'textStyle': {'color': '#dadada'}},
        'backgroundColor': '#101018',
        'hAxis': {'title': 'Date', 'titleTextStyle': {'color': '#dadada'}, 'textStyle': {'color': '#dadada'}},
        'vAxis': {
            'title': 'Number of Games',
            'titleTextStyle': {'color': '#dadada'},
            'textStyle': {'color': '#dadada'},
            'gridlines': {'color': '#101018'}
        },
        'isStacked': 'true',
        'series': {'0': {'color': '#458296'}, '1': {'color': '#6edbf9'}}
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.ColumnChart(document.getElementById('weekBarChart'));
    chart.draw(data, options);
}

async function drawMonthChart() {
    let date = new Date();
    let dates = [];
    for (let i = 29; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    var data = new google.visualization.DataTable();
    await populateDictionaries();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Games Lost');
    data.addColumn('number', 'Games Won');
    data.addRows([
        [dates[0], gamesPlayed[dates[0]], gamesWon[dates[0]]],
        [dates[1], gamesPlayed[dates[1]], gamesWon[dates[1]]],
        [dates[2], gamesPlayed[dates[2]], gamesWon[dates[2]]],
        [dates[3], gamesPlayed[dates[3]], gamesWon[dates[3]]],
        [dates[4], gamesPlayed[dates[4]], gamesWon[dates[4]]],
        [dates[5], gamesPlayed[dates[5]], gamesWon[dates[5]]],
        [dates[6], gamesPlayed[dates[6]], gamesWon[dates[6]]],
        [dates[7], gamesPlayed[dates[7]], gamesWon[dates[7]]],
        [dates[8], gamesPlayed[dates[8]], gamesWon[dates[8]]],
        [dates[9], gamesPlayed[dates[9]], gamesWon[dates[9]]],
        [dates[10], gamesPlayed[dates[10]], gamesWon[dates[10]]],
        [dates[11], gamesPlayed[dates[11]], gamesWon[dates[11]]],
        [dates[12], gamesPlayed[dates[12]], gamesWon[dates[12]]],
        [dates[13], gamesPlayed[dates[13]], gamesWon[dates[13]]],
        [dates[14], gamesPlayed[dates[14]], gamesWon[dates[14]]],
        [dates[15], gamesPlayed[dates[15]], gamesWon[dates[15]]],
        [dates[16], gamesPlayed[dates[16]], gamesWon[dates[16]]],
        [dates[17], gamesPlayed[dates[17]], gamesWon[dates[17]]],
        [dates[18], gamesPlayed[dates[18]], gamesWon[dates[18]]],
        [dates[19], gamesPlayed[dates[19]], gamesWon[dates[19]]],
        [dates[20], gamesPlayed[dates[20]], gamesWon[dates[20]]],
        [dates[21], gamesPlayed[dates[21]], gamesWon[dates[21]]],
        [dates[22], gamesPlayed[dates[22]], gamesWon[dates[22]]],
        [dates[23], gamesPlayed[dates[23]], gamesWon[dates[23]]],
        [dates[24], gamesPlayed[dates[24]], gamesWon[dates[24]]],
        [dates[25], gamesPlayed[dates[25]], gamesWon[dates[25]]],
        [dates[26], gamesPlayed[dates[26]], gamesWon[dates[26]]],
        [dates[27], gamesPlayed[dates[27]], gamesWon[dates[27]]],
        [dates[28], gamesPlayed[dates[28]], gamesWon[dates[28]]],
        [dates[29], gamesPlayed[dates[29]], gamesWon[dates[29]]]
    ]);

    var options = {
        'title': 'Games Played Over Time (Last 30 Days)',
        'titleTextStyle': {'color': '#dadada', 'fontSize': 24},
        'width': 700,
        'height': 500,
        'legend': {'textStyle': {'color': '#dadada'}},
        'backgroundColor': '#101018',
        'hAxis': {'title': 'Date', 'titleTextStyle': {'color': '#dadada'}, 'textStyle': {'color': '#dadada'}},
        'vAxis': {
            'title': 'Number of Games',
            'titleTextStyle': {'color': '#dadada'},
            'textStyle': {'color': '#dadada'},
            'gridlines': {'color': '#101018'}
        },
        'isStacked': 'true',
        'series': {'0': {'color': '#458296'}, '1': {'color': '#6edbf9'}}
    };

    var chart = new google.visualization.ColumnChart(document.getElementById('monthBarChart'));
    chart.draw(data, options);

    document.getElementById("monthBarChart").style.display = "none";
}

async function drawWeekLineChart() {
    let date = new Date();
    let dates = [];
    for (let i = 6; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Win Rate Percentage');
    await populateDictionaries();
    data.addRows([
        [dates[0], getWinRateNumber(dates[0])],
        [dates[1], getWinRateNumber(dates[1])],
        [dates[2], getWinRateNumber(dates[2])],
        [dates[3], getWinRateNumber(dates[3])],
        [dates[4], getWinRateNumber(dates[4])],
        [dates[5], getWinRateNumber(dates[5])],
        [dates[6], getWinRateNumber(dates[6])],
    ]);

    // Set chart options
    var options = {
        'title': 'Win Rate Over Time (Last 7 Days)',
        'titleTextStyle': {'color': '#dadada', 'fontSize': 24},
        'width': 700,
        'height': 500,
        'legend': {'textStyle': {'color': '#dadada'}},
        'backgroundColor': '#101018',
        'hAxis': {'title': 'Date', 'titleTextStyle': {'color': '#dadada'}, 'textStyle': {'color': '#dadada'}},
        'vAxis': {
            'title': 'Percentage',
            'titleTextStyle': {'color': '#dadada'},
            'textStyle': {'color': '#dadada'},
            'gridlines': {'color': '#101018'},
            'viewWindow': {'min': 0, 'max': 100}
        },
        'isStacked': 'true',
        'series': {'0': {'color': '#458296'}, '1': {'color': '#6edbf9'}}
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.LineChart(document.getElementById('weekLineChart'));
    chart.draw(data, options);
}

async function drawMonthLineChart() {
    let date = new Date();
    let dates = [];
    for (let i = 29; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Win Rate Percentage');
    await populateDictionaries();
    data.addRows([
        [dates[0], getWinRateNumber(dates[0])],
        [dates[1], getWinRateNumber(dates[1])],
        [dates[2], getWinRateNumber(dates[2])],
        [dates[3], getWinRateNumber(dates[3])],
        [dates[4], getWinRateNumber(dates[4])],
        [dates[5], getWinRateNumber(dates[5])],
        [dates[6], getWinRateNumber(dates[6])],
        [dates[7], getWinRateNumber(dates[7])],
        [dates[8], getWinRateNumber(dates[8])],
        [dates[9], getWinRateNumber(dates[9])],
        [dates[10], getWinRateNumber(dates[10])],
        [dates[11], getWinRateNumber(dates[11])],
        [dates[12], getWinRateNumber(dates[12])],
        [dates[13], getWinRateNumber(dates[13])],
        [dates[14], getWinRateNumber(dates[14])],
        [dates[15], getWinRateNumber(dates[15])],
        [dates[16], getWinRateNumber(dates[16])],
        [dates[17], getWinRateNumber(dates[17])],
        [dates[18], getWinRateNumber(dates[18])],
        [dates[19], getWinRateNumber(dates[19])],
        [dates[20], getWinRateNumber(dates[20])],
        [dates[21], getWinRateNumber(dates[21])],
        [dates[22], getWinRateNumber(dates[22])],
        [dates[23], getWinRateNumber(dates[23])],
        [dates[24], getWinRateNumber(dates[24])],
        [dates[25], getWinRateNumber(dates[25])],
        [dates[26], getWinRateNumber(dates[26])],
        [dates[27], getWinRateNumber(dates[27])],
        [dates[28], getWinRateNumber(dates[28])],
        [dates[29], getWinRateNumber(dates[29])]
    ]);

    var options = {
        'title': 'Win Rate Over Time (Last 30 Days)',
        'titleTextStyle': {'color': '#dadada', 'fontSize': 24},
        'width': 700,
        'height': 500,
        'legend': {'textStyle': {'color': '#dadada'}},
        'backgroundColor': '#101018',
        'hAxis': {'title': 'Date', 'titleTextStyle': {'color': '#dadada'}, 'textStyle': {'color': '#dadada'}},
        'vAxis': {
            'title': 'Percentage',
            'titleTextStyle': {'color': '#dadada'},
            'textStyle': {'color': '#dadada'},
            'gridlines': {'color': '#101018'},
            'viewWindow': {'min': 0, 'max': 100}
        },
        'isStacked': 'true',
        'series': {'0': {'color': '#458296'}, '1': {'color': '#6edbf9'}}
    };

    var chart = new google.visualization.LineChart(document.getElementById('monthLineChart'));
    chart.draw(data, options);

    document.getElementById("monthLineChart").style.display = "none";
}

function toggleBarCharts() {
    var weekChart = document.getElementById("weekBarChart");
    var monthChart = document.getElementById("monthBarChart");
    var button = document.getElementById("chartButton");

    if (weekChart.style.display !== "none") {
        weekChart.style.display = "none";
        monthChart.style.display = "block";
        button.textContent = "last 7 days";
    }
    else {
        weekChart.style.display = "block";
        monthChart.style.display = "none";
        button.textContent = "last 30 days";
    }
}

function toggleLineCharts() {
    var weekLineChart = document.getElementById("weekLineChart");
    var monthLineChart = document.getElementById("monthLineChart");
    var button = document.getElementById("chartButton");

    if ( weekLineChart.style.display !== "none") {
        weekLineChart.style.display = "none";
        monthLineChart.style.display = "block";
        button.textContent = "last 7 days";
    }
    else {
        weekLineChart.style.display = "block";
        monthLineChart.style.display = "none";
        button.textContent = "last 30 days";
    }
}

function getGames() {
    console.log("[INFO] sending http request for getGames");
    // Make a request to your Go server
    return fetch('http://localhost:8080/get-games')
        .then(response => response.text())
        .then(data => {
            // Update the result div with the response from the server
            return data;
        })
        .catch(error => console.error('Error:', error));
}

function getGamesWon() {
    console.log("[INFO] sending http request for getGames");
    // Make a request to your Go server
    return fetch('http://localhost:8080/get-games-won')
        .then(response => response.text())
        .then(data => {
            // Update the result div with the response from the server
            return data;
        })
        .catch(error => console.error('Error:', error));
}

async function populateDictionaries() {
    let gameDatesAndTotals = await getGames();
    let arrayPlayed = gameDatesAndTotals.split("\n");
    let gamesDatesAndWins = await getGamesWon();
    let arrayWon = gamesDatesAndWins.split("\n");
    let i = 0
    while (arrayWon[i] != null) {
        gamesWon[arrayWon[i].substring(5, 7) + "/" + arrayWon[i].substring(8, 10)] = parseInt(arrayWon[i + 1]);
        i += 2;
    }
    i = 0;
    while (arrayPlayed[i] != null) {
        gamesPlayed[arrayPlayed[i].substring(5, 7) + "/" + arrayPlayed[i].substring(8, 10)] = parseInt(arrayPlayed[i + 1]);
        i += 2;
    }
}

function getWinRateNumber(index) {
    var returnVal = (parseFloat(gamesWon[index]) / (parseFloat(gamesWon[index]) + parseFloat(gamesPlayed[index]))) * 100;
    if (!isNaN(returnVal)) {
        return returnVal;
    }
    return 0.00;
}