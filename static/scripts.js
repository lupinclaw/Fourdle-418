function onClickNav() {
    var menu = document.getElementById("menuLinks");
    if (menu.style.display === "block") {
        menu.style.display = "none";
    }
    else {
        menu.style.display = "block";
    }
}

function loadGeneralMetrics() {
    var httpRequest = new XMLHttpRequest();
    httpRequest.load= function() {

    };
    httpRequest.open("GET", "http://localhost:8080", true);
    httpRequest.send();
}

// Load the Visualization API and the corechart package.
google.charts.load('current', {'packages':['corechart']});

// Set a callback to run when the Google Visualization API is loaded.
google.charts.setOnLoadCallback(drawChart);

// Callback that creates and populates a data table,
// instantiates the pie chart, passes in the data and
// draws it.
function drawChart() {

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Date');
    data.addColumn('number', 'Games Won');
    data.addColumn('number', 'Games Lost');
    data.addRows([
        ['11/16', 3, 2],
        ['11/17', 0, 3],
        ['11/18', 4, 1],
        ['11/19', 3, 8],
        ['11/20', 10, 3],
        ['11/21', 8, 2],
        ['11/22', 0, 0]
    ]);

    // Set chart options
    var options = {
        'title':'Games Played Over Time',
        'titleTextStyle': {'color':'#dadada', 'fontSize':24},
        'width':600,
        'height':500,
        'legend': {'textStyle': {'color':'#dadada'} },
        'backgroundColor':'#101018',
        'hAxis': {'textStyle': {'color':'#dadada'} },
        'vAxis': {'textStyle': {'color':'#dadada'}, 'gridlines': {'color':'#101018'} },
        'isStacked':'true',
        'series': {'0': {'color':'#458296'}, '1': {'color':'#6edbf9'}}
    };

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.ColumnChart(document.getElementById('chart_div'));
    chart.draw(data, options);
}