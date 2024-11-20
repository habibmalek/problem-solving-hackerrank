package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
    // Find the longest common prefix between s and t
    commonLength := 0
    minLength := len(s)
    if len(t) < minLength {
        minLength = len(t)
    }

    for i := 0; i < minLength; i++ {
        if s[i] == t[i] {
            commonLength++
        } else {
            break
        }
    }

    // Operations needed to remove differing parts from s and add parts of t
    totalChanges := int32(len(s) - commonLength + len(t) - commonLength)

    // Check if we can match t with exactly k operations
    if totalChanges == k || (totalChanges < k && (k-totalChanges)%2 == 0) || (int32(len(s)+len(t)) <= k) {
        return "Yes"
    }

    return "No"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    t := readLine(reader)

    kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    k := int32(kTemp)

    result := appendAndDelete(s, t, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
