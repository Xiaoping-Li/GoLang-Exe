/*
Task: 
To complete this challenge, you must save a line of input from stdin to a variable, print Hello, World. on a single line, and 
finally print the value of your variable on a second line.

Input Format:
A single line of text denoting  (the variable whose contents must be printed).

Output Format:
Print Hello, World. on the first line, and the contents of  on the second line.

Sample Input:
Welcome to 30 Days of Code!
Sample Output

Hello, World. 
Welcome to 30 Days of Code!
*/

package main
import (
   "fmt" 
   "bufio"
   "os"
)


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')

    fmt.Printf("Hello, World. \n")
    fmt.Printf(text)
}
