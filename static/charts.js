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
function drawWeekChart() {
    let date = new Date();
    let dates =  [];
    for (let i = 6; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Games Won');
    data.addColumn('number', 'Games Lost');
    data.addRows([
        [dates[0], 3, 2],
        [dates[1], 0, 3],
        [dates[2], 4, 1],
        [dates[3], 3, 8],
        [dates[4], 10, 3],
        [dates[5], 8, 2],
        [dates[6], 0, 0]
    ]);

    // Set chart options
    var options = {
        'title':'Games Played Over Time (Last 7 Days)',
        'titleTextStyle': {'color':'#dadada', 'fontSize':24},
        'width':700,
        'height':500,
        'legend': {'textStyle': {'color':'#dadada'} },
        'backgroundColor':'#101018',
        'hAxis': {'title':'Date', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color':'#dadada'} },
        'vAxis': {'title':'Number of Games', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color':'#dadada'}, 'gridlines': {'color':'#101018'} },
        'isStacked':'true',
        'series': {'0': {'color':'#458296'}, '1': {'color':'#6edbf9'}}
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.ColumnChart(document.getElementById('weekBarChart'));
    chart.draw(data, options);
}

function drawMonthChart() {
    let date = new Date();
    let dates =  [];
    for (let i = 30; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Games Won');
    data.addColumn('number', 'Games Lost');
    data.addRows([
        [dates[0], 1, 3],
        [dates[1], 1, 0],
        [dates[2], 1, 4],
        [dates[3], 1, 3],
        [dates[4], 1, 10],
        [dates[5], 1, 8],
        [dates[6], 1, 0],
        [dates[7], 1, 3],
        [dates[8], 1, 0],
        [dates[9], 1, 4],
        [dates[10], 1, 3],
        [dates[11], 1, 10],
        [dates[12], 1, 8],
        [dates[13], 1, 3],
        [dates[14], 1, 0],
        [dates[15], 1, 4],
        [dates[16], 1, 3],
        [dates[17], 1, 10],
        [dates[18], 1, 8],
        [dates[19], 1, 4],
        [dates[20], 1, 3],
        [dates[21], 1, 10],
        [dates[22], 1, 8],
        [dates[23], 1, 3],
        [dates[24], 1, 0],
        [dates[25], 1, 4],
        [dates[26], 1, 3],
        [dates[27], 1, 10],
        [dates[28], 1, 8],
        [dates[29], 1, 10]
    ]);

    var options = {
        'title':'Games Played Over Time (Last 30 Days)',
        'titleTextStyle': {'color':'#dadada', 'fontSize':24},
        'width':700,
        'height':500,
        'legend': {'textStyle': {'color':'#dadada'} },
        'backgroundColor':'#101018',
        'hAxis': {'title':'Date', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color':'#dadada'} },
        'vAxis': {'title':'Number of Games', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color':'#dadada'}, 'gridlines': {'color':'#101018'} },
        'isStacked':'true',
        'series': {'0': {'color':'#458296'}, '1': {'color':'#6edbf9'}}
    };

    var chart = new google.visualization.ColumnChart(document.getElementById('monthBarChart'));
    chart.draw(data, options);

    document.getElementById("monthBarChart").style.display = "none";
}

function drawWeekLineChart() {
    let date = new Date();
    let dates =  [];
    for (let i = 6; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Win Rate Percentage');
    data.addRows([
        [dates[0], 33.3],
        [dates[1], 25.0],
        [dates[2], 0.0],
        [dates[3], 100.0],
        [dates[4], 100.0],
        [dates[5], 45.0],
        [dates[6], 0.0]
    ]);

    // Set chart options
    var options = {
        'title':'Win Rate Over Time (Last 7 Days)',
        'titleTextStyle': {'color':'#dadada', 'fontSize':24},
        'width':700,
        'height':500,
        'legend': {'textStyle': {'color':'#dadada'} },
        'backgroundColor':'#101018',
        'hAxis': {'title':'Date', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color':'#dadada'} },
        'vAxis': {'title':'Percentage', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color':'#dadada'}, 'gridlines': {'color':'#101018'}, 'viewWindow': {'min':0, 'max':100 } },
        'isStacked':'true',
        'series': {'0': {'color':'#458296'}, '1': {'color':'#6edbf9'}}
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.LineChart(document.getElementById('weekLineChart'));
    chart.draw(data, options);
}

function drawMonthLineChart() {
    let date = new Date();
    let dates = [];
    for (let i = 30; i >= 0; i--) {
        dates[i] = (date.getMonth() + 1) + "/" + date.getDate();
        date.setDate(date.getDate() - 1);
    }

    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Win Rate Percentage');
    data.addRows([
        [dates[0], 3],
        [dates[1], 0],
        [dates[2], 4],
        [dates[3], 3],
        [dates[4], 10],
        [dates[5], 8],
        [dates[6], 0],
        [dates[7], 3],
        [dates[8], 0],
        [dates[9], 4],
        [dates[10], 3],
        [dates[11], 10],
        [dates[12], 8],
        [dates[13], 3],
        [dates[14], 0],
        [dates[15], 4],
        [dates[16], 3],
        [dates[17], 10],
        [dates[18], 8],
        [dates[19], 4],
        [dates[20], 3],
        [dates[21], 10],
        [dates[22], 8],
        [dates[23], 3],
        [dates[24], 0],
        [dates[25], 4],
        [dates[26], 3],
        [dates[27], 10],
        [dates[28], 8],
        [dates[29], 10]
    ]);

    var options = {
        'title': 'Win Rate Over Time (Last 30 Days)',
        'titleTextStyle': {'color': '#dadada', 'fontSize': 24},
        'width': 700,
        'height': 500,
        'legend': {'textStyle': {'color': '#dadada'}},
        'backgroundColor': '#101018',
        'hAxis': {'title':'Date', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color': '#dadada'}},
        'vAxis': {'title':'Percentage', 'titleTextStyle': {'color':'#dadada'}, 'textStyle': {'color': '#dadada'}, 'gridlines': {'color': '#101018'}, 'viewWindow': {'min':0, 'max':100 }},
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