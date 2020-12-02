const fs = require("fs");

function parseToPassword(inputLine) {
    const [rawLimit, rawLetter, password] = inputLine.split(" ");
    const [lowerLimit, higherLimit] = rawLimit.split("-");
    const letter = rawLetter[0];
    return {
        lowerLimit: parseInt(lowerLimit),
        higherLimit: parseInt(higherLimit),
        letter,
        password
    };
}

function isValidPassword(passwordInput) {
    const { lowerLimit, higherLimit, letter, password } = passwordInput;
    const occurrences = password.split(letter).length - 1;
    return occurrences >= lowerLimit && occurrences <= higherLimit;
}

(function main() {
    const output = fs.readFileSync("./input.txt")
        .toString()
        .split("\n")
        .map(parseToPassword)
        .filter(isValidPassword)
    console.log(output.length);
})();
