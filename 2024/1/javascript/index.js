const fs = require('fs');

function main() {
    const input = fs.readFileSync("../input-1", 'utf-8');

    const [left, right] = parseLeftAndRight(input);

    // const left = [1, 2, 3, 4];
    // const right = [1, 1, 1, 1];

    left.sort();
    right.sort();

    const sum = sumDifferences(left, right);
    console.log("Sum of differences: ", sum)

    const similarityScore = calculateSimularityScore(left, right);
    console.log("Simularity score: ", similarityScore);
}

function parseLeftAndRight(s) {
    return s.split('\n')
        .map(tupleString => tupleString.split("   "))
        .map(tuple => [parseInt(tuple[0]), parseInt(tuple[1])])
        .reduce((acc, curr) => {
            acc[0].push(curr[0]);
            acc[1].push(curr[1]);
            return acc;
        }, [[], []]);
}

function sumDifferences(left, right) {
    let sum = 0;
    for (let i = 0; i < left.length; i++) {
        const distance = Math.abs(left[i] - right[i])
        sum += distance;
    }
    return sum;
}

function calculateSimularityScore(left, right) {
    // Find frequencies from right
    const rightFrequencies = right.reduce((acc, curr) => {
        const key = parseInt(curr);
        if(acc[key] === undefined) {
            acc[key] = 1;
        } else {
            acc[key] += 1
        }
        return acc;
    }, {});

    return left.reduce((sum, curr) => {
        const frequency = rightFrequencies[parseInt(curr)] || 0;
        return sum + (curr * frequency)
    }, 0);
}

main();