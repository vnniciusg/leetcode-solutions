/**
 * @param {number[]} heights
 * @return {number}
 */
var heightChecker = function(heights) {
    
    let n = heights.length;

    let expected = heights.slice();
    expected.sort(function(a, b) {return a - b});

    let count = 0;
    for (i = 0; i < n; i++) {
        if (heights[i] != expected[i]) {
            count++;
        }
    }

    return count;
};
