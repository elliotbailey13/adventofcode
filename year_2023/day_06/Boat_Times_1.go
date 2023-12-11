package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "strings"
)

var (
  Times []int
  Distances []int
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

  for _, Full_Input_Line := range Full_Input {
    log.Printf(Full_Input_Line)
    err = Load_Data(Full_Input_Line)
    if err != nil {
      log.Printf("Loading data file failed. (%v)", err)
      os.Exit(1)
    }
  }

  log.Printf("Times: %+v", Times)
  log.Printf("Distances: %+v", Distances)

  var Wins []int
  var Current_Wins int
  for i := 0; i != len(Times); i++ {
    Current_Wins, err = Find_Best_Times(Times[i], Distances[i])
    if err != nil {
      log.Printf("Failed finding best times for %v | %v. (%v)", Times[i], Distances[i], err)
      os.Exit(1)
    }

    log.Printf("Time: %v - Wins: %v", Times[i], Current_Wins)

    Wins = append(Wins, Current_Wins)
  }

  Final_Calculation := Wins[0]
  for Index, Win := range Wins {
    if Index == 0 { continue }
    Final_Calculation = Final_Calculation * Win
  }

  log.Printf("Final Calucation: %v", Final_Calculation)
}

func Load_Data(Line string) error {
  var Array *[]int
  var Line_Values []string
  if strings.Contains(Line, "Time:") {
    Line = strings.Replace(Line, "Time:", "", -1)
    Array = &Times
  }

  if strings.Contains(Line, "Distance:") {
    Line = strings.Replace(Line, "Distance:", "", -1)
    Array = &Distances
  }

  var err error
  var New_Int int
  Line_Values = strings.Fields(Line)
  for _, Value := range Line_Values {
    New_Int, err = strconv.Atoi(Value)
    if err != nil { return err }

    *Array = append(*Array, New_Int)
  }

  return nil
}

func Find_Best_Times(Time int, Distance int) (int, error) {
  log.Printf("Time: %v - Distance: %v", Time, Distance)

  var Total_Wins int
  var Time_Moving int
  var Distance_Moved int
  for i := 1; i <= Time; i++ {
    Time_Moving = Time - i

    Distance_Moved = Time_Moving * i

    log.Printf("Speed: %v - Distance: %v", i, Distance_Moved)

    if Distance_Moved > Distance { Total_Wins++ }
  }

  return Total_Wins, nil
}
