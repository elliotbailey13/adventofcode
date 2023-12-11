package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
  "strings"
)

var (
  Full_Time int
  Full_Distance int
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

  log.Printf("Finding best time. Time: %v | Distance: %v", Full_Time, Full_Distance)

  var Wins int
  Wins, err = Find_Best_Times(Full_Time, Full_Distance)
  if err != nil {
    log.Printf("Failed finding best times for %v | %v. (%v)", Full_Time, Full_Distance, err)
    os.Exit(1)
  }

  log.Printf("Time: %v - Wins: %v", Full_Time, Wins)
}

func Load_Data(Line string) error {
  var Mode int
  var Line_Values []string
  if strings.Contains(Line, "Time:") {
    Line = strings.Replace(Line, "Time:", "", -1)
    Mode = 1
  }

  if strings.Contains(Line, "Distance:") {
    Line = strings.Replace(Line, "Distance:", "", -1)
    Mode = 2
  }

  var err error
  var Full_Int_String string
  Line_Values = strings.Fields(Line)
  for _, Value := range Line_Values { Full_Int_String = Full_Int_String + Value }

  if Mode == 1 {
    Full_Time, err = strconv.Atoi(Full_Int_String)
  }

  if Mode == 2 {
    Full_Distance, err = strconv.Atoi(Full_Int_String)
  }

  if err != nil { return err }

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

    if Distance_Moved > Distance { Total_Wins++ }
  }

  return Total_Wins, nil
}
