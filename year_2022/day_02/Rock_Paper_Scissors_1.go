package main

import (
  "os"
  "log"
  "bufio"
  "strings"
)

func main() {
  Raw_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  var Full_Input []string
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Full_Input = append(Full_Input, scanner.Text())
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

/*
  A X -> Rock
  B Y -> Paper
  C Z -> Scissors
*/

  var Total_Score int
  for _, Full_Input_Line := range Full_Input {
    Line_Split := strings.Split(Full_Input_Line, " ")

    switch Line_Split[1] {
      case "X": // Rock
        Total_Score += 1
      case "Y": // Paper
        Total_Score += 2
      case "Z": // Scissors
        Total_Score += 3
      default:
        log.Printf("I shouldn't be here. Data is wrong. (%v)", Line_Split[1])
    }

    switch Line_Split[0] {
      case "A": // Rock
        if Line_Split[1] == "X" { Total_Score += 3 }
        if Line_Split[1] == "Y" { Total_Score += 6 }
//        if Line_Split[1] == "Z" { Total_Score += 0 }
      case "B": // Paper
//        if Line_Split[1] == "X" { Total_Score += 0 }
        if Line_Split[1] == "Y" { Total_Score += 3 }
        if Line_Split[1] == "Z" { Total_Score += 6 }
      case "C": // Scissors
        if Line_Split[1] == "X" { Total_Score += 6 }
//        if Line_Split[1] == "Y" { Total_Score += 0 }
        if Line_Split[1] == "Z" { Total_Score += 3 }
      default:
        log.Printf("I shouldn't be here. Data is wrong. (%v)", Line_Split[0])
    }
  }

  log.Printf("Total Score: %v", Total_Score)
}
