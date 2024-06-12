func fizzBuzz(n int) []string {
    
    var ans []string
    for i := 1; i <= n; i++ {
        if isDivisibleBy(i, 15) {
            ans = append(ans, "FizzBuzz")
        } else if isDivisibleBy(i, 3) {
            ans = append(ans, "Fizz")
        } else if isDivisibleBy(i, 5) {
            ans = append(ans, "Buzz")
        } else {
            ans = append(ans, strconv.Itoa(i))
        }
    }
    return ans

}


func isDivisibleBy(v, div int) bool {
    return v % div == 0 
}
