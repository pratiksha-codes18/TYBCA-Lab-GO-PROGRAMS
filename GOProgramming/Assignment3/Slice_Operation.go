package main
import "fmt"

func main() {
    s := []int{10, 20, 30}
    fmt.Println("Original slice:", s)

    // Append
    s = append(s, 40)
    fmt.Println("After append:", s)

    // Remove (2nd element)
    s = append(s[:1], s[2:]...)
    fmt.Println("After remove:", s)

    // Copy
    s2 := make([]int, len(s))
    copy(s2, s)
    fmt.Println("Copied slice:", s2)
}