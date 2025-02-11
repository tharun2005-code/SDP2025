function appendValue(value) {
    document.getElementById('box').value += value;
}

function clearScreen() {
    document.getElementById('box').value = '';
}

function calculateResult() {
    try {
        document.getElementById('box').value = eval(document.getElementById('box').value);
    } catch {
        document.getElementById('box').value = 'Error';
    }
}
