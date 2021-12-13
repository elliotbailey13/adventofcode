package main

import (
  "os"
  "log"
  "bufio"
//  "strconv"
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

  var Illegal_Characters []string
  for Row_Id, Row := range Full_Input {
    log.Printf("Checking row %v - %v", Row_Id, Row)
    Illegal_Character, Incomplete := Line_Logic(Row)

    if Incomplete == false {
      Illegal_Characters = append(Illegal_Characters, Illegal_Character)
    }
  }

//  log.Printf("Illegal_Characters: %v", Illegal_Characters)
  var Total_Score int
  for _, Illegal_Character := range Illegal_Characters {
    switch Illegal_Character {
      case ")": Total_Score += 3
      case "]": Total_Score += 57
      case "}": Total_Score += 1197
      case ">": Total_Score += 25137
    }
  }

  log.Printf("Total Score: %v", Total_Score)
}

func Line_Logic(Line string) (string, bool) {
  Split_Line := strings.Split(Line, "")

  var Chunk []string
  for Location, Character := range Split_Line {
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
          log.Printf("Returning character from location: %v", Location)
          return Character, false
        }
    }
  }

  log.Printf("Throwing out line")
  return "", true
}
