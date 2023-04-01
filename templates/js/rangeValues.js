document.getElementById('creationDateFrom').addEventListener('input', function functionName(e) {
    let end = e.target.value;
    console.log(end);
    document.getElementById('creationDateFromNr').innerHTML = end;
});
document.getElementById('creationDateTo').addEventListener('input', function functionName(e) {
    let start = e.target.value;
    console.log(start);
    document.getElementById('creationDateToNr').innerHTML = start;
});