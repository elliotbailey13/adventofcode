package main

import (
  "os"
  "log"
  "bufio"
  "errors"
  "strconv"
  "strings"
)

var (
  Card_Map map[int]int
)

func main() {
  Card_Map = make(map[int]int)

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
  var Card_Sum int
  var Total_Card_Sum int
  for Card_ID, Full_Input_Line := range Full_Input {
    Card_Sum, err = Check_Card(Full_Input_Line)
    if err != nil {
      Card_Errors += 1
      log.Printf("Card %v - check error. (%v).", Card_ID, err)
    }

    Total_Card_Sum += Card_Sum
  }

  log.Printf("Total Errors: %v", Card_Errors)
  log.Printf("Sum of cards: %v", Total_Card_Sum)
}

func Check_Card(Line string) (int, error) {
  var err error
  var Card_ID int
  var Winners []string
  var Numbers []string
  Card_Split := strings.Split(Line, ":")
  Number_Split := strings.Split(Card_Split[1], "|")

  Card_ID_Split := strings.Fields(Card_Split[0])
  Card_ID, err = strconv.Atoi(Card_ID_Split[1])
  if err != nil {
    return 0, errors.New("Card " + Card_ID_Split[1] + " couldn't be convert to a int")
  }

  Card_Map[Card_ID] = Card_Map[Card_ID] + 1

  Winners = strings.Fields(Number_Split[0])
  Numbers = strings.Fields(Number_Split[1])

  var Winning_Sum int
  for _, Number := range Numbers {
    for _, Winner := range Winners {
      if Number != Winner { continue }

      Winning_Sum += 1
    }
  }

  if Winning_Sum == 0 {
    log.Printf("No winners for card %v.", Card_ID)
  } else {
    log.Printf("%v winners for card %v.", Winning_Sum, Card_ID)
    for Loop_Count := 1; Winning_Sum >= Loop_Count; Loop_Count++ {
      Card_Map[Card_ID + Loop_Count] = Card_Map[Card_ID + Loop_Count] + Card_Map[Card_ID]
    }
  }

  return Card_Map[Card_ID], nil
}
