package main

import (
  "os"
  "log"
  "sort"
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

  var Remaining_Chunks [][]string
  for Row_Id, Row := range Full_Input {
    log.Printf("Checking row %v - %v", Row_Id, Row)
    Remaining_Chunk, Corrupt := Line_Logic(Row)

    if Corrupt == false {
      Remaining_Chunks = append(Remaining_Chunks, Remaining_Chunk)
    }
  }

  var Total_Points []int
  for Chunk_Id, Chunk := range Remaining_Chunks {
    log.Printf("Calculating chunk %v points: %v", Chunk_Id, Chunk)
    Total_Points = append(Total_Points, Calculate_Points(Chunk))
  }

  sort.Ints(Total_Points)
  for Point_Id, Points := range Total_Points {
    log.Printf("Points for %v: %v", Point_Id, Points)
  }

  log.Printf("Middle Id: %v", len(Total_Points) / 2)
  log.Printf("Middle Score: %v", Total_Points[len(Total_Points) / 2])
}

func Line_Logic(Line string) ([]string, bool) {
  Split_Line := strings.Split(Line, "")

  var Chunk []string
  for _, Character := range Split_Line {
    if len(Chunk) == 0 {
//      log.Printf("Adding first character to chunck. %v", Character)
      Chunk = append(Chunk, Character)
      continue
    }

    switch Character {
      case "(", "[", "{", "<":
//        log.Printf("Adding next opening character to chunk. %v", Character)
        Chunk = append(Chunk, Character)

      case ")", "]", "}", ">":
        Top_Element := len(Chunk) - 1
        Last_Closing_Character := Chunk[Top_Element]

        Corrupted := true
        switch Character {
          case ")":
            if Last_Closing_Character == "(" { Corrupted = false }

          case "]":
            if Last_Closing_Character == "[" { Corrupted = false }

          case "}":
            if Last_Closing_Character == "{" { Corrupted = false }

          case ">":
            if Last_Closing_Character == "<" { Corrupted = false }
        }

        if Corrupted == false {
//          log.Printf("Removing chunk. %v", Character)
          Chunk = Chunk[:Top_Element]
        } else {
          log.Printf("Throwing out line")
          return []string{}, true
        }
    }
  }

  log.Printf("Returning remaining chunk: %v", Chunk)
  return Chunk, false
}

func Calculate_Points(Remaining_Chunk []string) int {
  var Total_Score int

//  var First_Closing_Character string
  for len(Remaining_Chunk) != 0 {
    Top_Element := len(Remaining_Chunk) - 1
    Last_Closing_Character := Remaining_Chunk[Top_Element]
//    First_Closing_Character, Remaining_Chunk = Remaining_Chunk[0], Remaining_Chunk[1:]

    switch Last_Closing_Character {
      case "(": Total_Score = (Total_Score * 5) + 1
      case "[": Total_Score = (Total_Score * 5) + 2
      case "{": Total_Score = (Total_Score * 5) + 3
      case "<": Total_Score = (Total_Score * 5) + 4
    }

    Remaining_Chunk = Remaining_Chunk[:Top_Element]
  }

  return Total_Score
}
