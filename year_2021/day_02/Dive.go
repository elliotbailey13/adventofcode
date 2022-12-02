package main

import (
  "os"
  "log"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  Path_Readings, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }
  defer Path_Readings.Close()

  Depth := 0
  Horizontal := 0
  scanner := bufio.NewScanner(Path_Readings)

  for scanner.Scan() {
    Path_Reading_Array := strings.Split(scanner.Text(), " ")
    Path_Amount, _ := strconv.Atoi(Path_Reading_Array[1])
    Path_Type := Path_Reading_Array[0]

    switch Path_Type {
    case "forward":
      Horizontal += Path_Amount

    case "up":
      Depth -= Path_Amount

    case "down":
      Depth += Path_Amount
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  Position := Horizontal * Depth

  log.Printf("Submarine position: %v", Position)
}
