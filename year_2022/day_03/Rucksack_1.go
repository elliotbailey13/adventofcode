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

  Sack_Priority_Sum := 0
  Priority_Index := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

  for _, Full_Input_Line := range Full_Input {
    Sack_Arr := strings.Split(Full_Input_Line, "")

    Sack_Arr_Left	:= Sack_Arr[0 : (len(Sack_Arr) / 2 )]
    Sack_Arr_Right	:= Sack_Arr[len(Sack_Arr) / 2 : len(Sack_Arr)]

    Match := ""
    Sack_Priority := 0
    for _, Left_Type := range Sack_Arr_Left {
      for _, Right_Type := range Sack_Arr_Right {
        if Left_Type == Right_Type {
          if Match != Left_Type {
            Match = Left_Type
	    for Index, Value := range Priority_Index {
              if Left_Type == Value { Sack_Priority = Index + 1 }
            }
	  }
        }
      }
    }

    if Sack_Priority == 0 {
      log.Printf("Sack priority not found\n%v", Sack_Arr)
      break
    }

    Sack_Priority_Sum += Sack_Priority
    log.Printf("Sack priority: %v / %v", Sack_Priority, Sack_Priority_Sum)
  }
}
