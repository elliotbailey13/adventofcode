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

  Full_String_Array := []string{}
  for _, Full_Input_Line := range Full_Input {
    Full_String_Array = strings.Split(Full_Input_Line, "")
  }

  var Rolling_Array []string
  for Index, Data := range Full_String_Array {
    Rolling_Array = append(Rolling_Array, Data)

    log.Printf("Index: %v - Data: %v - Rolling: %v", Index, Data, Rolling_Array)
    if len(Rolling_Array) < 4 { continue }

    Match := false
    for Index_Data, Rolling_Data := range Rolling_Array {
      if Match == true { continue }
      for Index_Data_Search, Rolling_Data_Search := range Rolling_Array {
        if Index_Data == Index_Data_Search { continue }
        if Rolling_Data == Rolling_Data_Search {
          Match = true
          Rolling_Array = Rolling_Array[1:]
          continue
        }
      }
    }

    if Match == false {
      log.Printf("Marker found at %v", Index + 1)
      log.Printf("Full Length: %v", len(Full_String_Array))
      break
    }
  }
}
