package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  Crab_Positions_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  var Crab_Positions []int
  var Total_Position_Length int
  scanner := bufio.NewScanner(Crab_Positions_Input)
  for scanner.Scan() {
    Crab_Position_String_Row := scanner.Text()

    Crab_Position_Array := strings.Split(Crab_Position_String_Row, ",")

    for _, Crab_Position_String := range Crab_Position_Array {
      Crab_Position, _ := strconv.Atoi(Crab_Position_String)
      Crab_Positions = append(Crab_Positions, Crab_Position)
      Total_Position_Length += Crab_Position
    }
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Crab_Positions_Input.Close()

  log.Printf("Number of Crabs: %v - Total Crab Length: %v", len(Crab_Positions), Total_Position_Length)
  Avg_Position := Total_Position_Length / len(Crab_Positions)
  log.Printf("Avg Position: %v", Avg_Position)

  var Best_Fuel int
  var Total_Fuel int
  var Best_Position int
  for i := 0; i < 2048; i++ {
    Total_Fuel = 0
    for _, Position := range Crab_Positions {
      if Position > i {
        Total_Fuel += Position - i
      } else if Position < i {
        Total_Fuel += i - Position
      }
    }

    if Best_Fuel == 0 { Best_Fuel = Total_Fuel }

    if Total_Fuel < Best_Fuel {
      Best_Fuel = Total_Fuel
      Best_Position = i
    }
  }

  log.Printf("Fuel Used: %v", Best_Fuel)
  log.Printf("Best Position: %v", Best_Position)
}
