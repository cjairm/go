/**
 * Get each number from the full number
 * Pass them to a function
 * Get Product Result
 * Get Sum Result
 * Get Product and Sum Result and Sub Sum from Product
 */
package main

func main () {

}

func subtractProductAndSum(n int) int {
    var p, s int = 1, 0

    for (n > 0) {
        p *= n % 10
        s += n % 10
        n /= 10
    }

    return p - s
}
