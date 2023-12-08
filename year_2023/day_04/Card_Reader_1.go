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

  var Card_Errors int
  var Card_Winning_Sum int
  var Total_Card_Winning_Sum int
  for Card_ID, Full_Input_Line := range Full_Input {
    log.Printf(Full_Input_Line)
    Card_Winning_Sum, err = Check_Card(Full_Input_Line)
    if err != nil {
      Card_Errors += 1
      log.Printf("Card %v - check error. (%v).", Card_ID, err)
    }

    Total_Card_Winning_Sum += Card_Winning_Sum
  }

  log.Printf("Total Errors: %v", Card_Errors)
  log.Printf("Sum of winning scores: %v", Total_Card_Winning_Sum)
}

func Check_Card(Line string) (int, error) {
  var Winners []string
  var Numbers []string
  Card_Split := strings.Split(Line, ":")
  Number_Split := strings.Split(Card_Split[1], "|")

  Winners = strings.Fields(Number_Split[0])
  Numbers = strings.Fields(Number_Split[1])

  var Winning_Sum int
  for _, Number := range Numbers {
    for _, Winner := range Winners {
      if Number != Winner { continue }

      if Winning_Sum == 0 { Winning_Sum += 1
      } else { Winning_Sum = Winning_Sum * 2 }
    }
  }

  return Winning_Sum, nil
}
