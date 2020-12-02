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
    return (password[firstIndex - 1] === letter) ^ (password[secondIndex - 1] === letter)
}

(function main() {
    const output = fs.readFileSync("./input.txt")
        .toString()
        .split("\n")
        .map(parseToPassword)
        .filter(isValidPassword)
    console.log(output.length);
})();
