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
  Symble_Arr []Symble_Data
  Number_Arr []Number_Data
)

func main() {
  Raw_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  var Full_Input []string
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Full_Input = append(Full_Input, scanner.Text())
  }

  if err = scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

  var Total_Errors int
  for Line_Number, Full_Input_Line := range Full_Input {
    log.Print(Full_Input_Line)
    err = Map_Locations(Line_Number, Full_Input_Line)
    if err != nil {
      Total_Errors += 1
      log.Printf("Error - %v", Full_Input_Line)
      log.Printf("Error - Line: %v | %v", Line_Number, err)
    }
  }

  var Part_Sum int
  Part_Sum, err = Match_Locations()
  if err != nil {
    Total_Errors += 1
    log.Printf("Error - %v", err)
  }

  log.Printf("Total Errors: %v", Total_Errors)
  log.Printf("Part Sum: %v", Part_Sum)
}

func Map_Locations(Line_Number int, Line string) error {
  var err error

  var Final_Number int
  var Current_Index int
  var Number_String string
  Line_Arr := strings.Split(Line, "")
  for Index, Charater := range Line_Arr {
    Current_Index = Index
    _, err = strconv.Atoi(Charater)
    if err != nil {
      if Charater != "." && Charater == "*" {
        Symble_Arr = append(Symble_Arr, Symble_Data{
          Line: Line_Number,
          Position: Index,
        })
      }

      if Number_String != "" {
        Final_Number, err = strconv.Atoi(Number_String)
        if err != nil {
          return errors.New("Error converting string to number. (" + Number_String + ")")
	}

        Number_Arr = append(Number_Arr, Number_Data{
          Value: Final_Number,
          Line: Line_Number,
          Start_Position: ( Index - 1 ) - len(Number_String) + 1,
          End_Position: ( Index - 1 ),
        })
      }

      Number_String = ""

      continue
    }

    Number_String = Number_String + Charater
  }

  if Number_String != "" {
    Final_Number, err = strconv.Atoi(Number_String)
    if err != nil {
      return errors.New("Error converting string to number. (" + Number_String + ")")
    }

    Number_Arr = append(Number_Arr, Number_Data{
      Value: Final_Number,
      Line: Line_Number,
      Start_Position: ( Current_Index - 1 ) - len(Number_String) + 1,
      End_Position: ( Current_Index - 1 ),
    })
  }

  return nil
}

func Match_Locations() (int, error) {
  var Match_Sum int
  var Matches []Number_Data
  for _, Symble_Info := range Symble_Arr {
    log.Printf("%v", Symble_Info)

    for _, Number_Info := range Number_Arr {
      if Number_Info.Line != Symble_Info.Line &&
          Number_Info.Line != ( Symble_Info.Line - 1 ) &&
          Number_Info.Line != ( Symble_Info.Line + 1 )  {
        continue
      }

      if Symble_Info.Position < Number_Info.Start_Position - 1 {
        continue
      }

      if Symble_Info.Position > Number_Info.End_Position + 1 {
        continue
      }

      Matches = append(Matches, Number_Info)
    }

    if len(Matches) != 2 {
      Matches = []Number_Data{}
      continue
    }

    for _, Match_Info := range Matches {
      log.Printf("Match. Number: %v | %v | %v | %v - Symble: %v | %v", Match_Info.Value, Match_Info.Line, Match_Info.Start_Position, Match_Info.End_Position, Symble_Info.Line, Symble_Info.Position)
    }

    Match_Sum += Matches[0].Value * Matches[1].Value
    Matches = []Number_Data{}
  }

  return Match_Sum, nil
}

type Number_Data struct {
  Value int
  Line int
  Start_Position int
  End_Position int
}

type Symble_Data struct {
  Line int
  Position int
}
