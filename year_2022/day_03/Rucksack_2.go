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

  Group_Counter := 0
  Sack_Priority_Sum := 0
  var Sack_Arr_1, Sack_Arr_2, Sack_Arr_3 []string
  Priority_Index := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

  for _, Full_Input_Line := range Full_Input {
    Sack_Arr := strings.Split(Full_Input_Line, "")

    switch Group_Counter {
      case 0: Sack_Arr_1 = Sack_Arr
      case 1: Sack_Arr_2 = Sack_Arr
      case 2: Sack_Arr_3 = Sack_Arr
      default:
        log.Print("Shouldn't be here. Grouping switch")
        os.Exit(1)
    }

    if Group_Counter != 2 {
      Group_Counter += 1
      continue
    }

    Group_Counter = 0

    Match := ""
    Sack_Priority := 0
    for _, Sack_1_Type := range Sack_Arr_1 {
      for _, Sack_2_Type := range Sack_Arr_2 {
        for _, Sack_3_Type := range Sack_Arr_3 {
          if Sack_1_Type == Sack_2_Type && Sack_1_Type == Sack_3_Type {
            if Match != Sack_1_Type {
              Match = Sack_1_Type
              for Index, Value := range Priority_Index {
                if Sack_1_Type == Value { Sack_Priority = Index + 1 }
              }
            }
          }
        }
      }
    }

    if Sack_Priority == 0 {
      log.Printf("Sack priority not found\n%v\n%v\n%v", Sack_Arr_1, Sack_Arr_2, Sack_Arr_3)
      break
    }

    Sack_Priority_Sum += Sack_Priority
    log.Printf("Sack priority: %v / %v", Sack_Priority, Sack_Priority_Sum)
  }
}
