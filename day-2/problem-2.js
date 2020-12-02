const fs = require("fs");

function parseToPassword(inputLine) {
    const [rawLimit, rawLetter, password] = inputLine.split(" ");
    const [firstIndex, secondIndex] = rawLimit.split("-");
    const letter = rawLetter[0];
    return {
        firstIndex: parseInt(firstIndex),
        secondIndex: parseInt(secondIndex),
        letter,
        password
    };
}

function isValidPassword(passwordInput) {
    const { firstIndex, secondIndex, letter, password } = passwordInput;
    let counter = 0;
    if (password[firstIndex - 1] === letter) counter++;
    if (password[secondIndex - 1] === letter) counter++;
    return counter === 1;
}

(function main() {
    const output = fs.readFileSync("./input.txt")
        .toString()
        .split("\n")
        .map(parseToPassword)
        .filter(isValidPassword)
    console.log(output.length);
})();
