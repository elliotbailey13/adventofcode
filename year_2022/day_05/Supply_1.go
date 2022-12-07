package main

import (
  "os"
  "log"
  "bufio"
  "strconv"
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

/*
  Supply_Stacks := [][]string{
    []string{ "G", " ", " ", " ", " ", "D", "R", " ", " " },
    []string{ "W", " ", " ", "V", " ", "C", "T", "M", " " },
    []string{ "L", " ", " ", "P", "Z", "Q", "F", "V", " " },
    []string{ "J", " ", " ", "S", "D", "J", "M", "T", "V" },
    []string{ "B", " ", "M", "H", "L", "Z", "J", "B", "S" },
    []string{ "R", "C", "T", "C", "T", "R", "D", "R", "D" },
    []string{ "T", "W", "Z", "T", "P", "B", "B", "H", "P" },
    []string{ "D", "S", "R", "D", "G", "F", "S", "L", "Q" },
  }
*/
  Supply_Stacks := [][]string{
    []string{ "D", "T", "R", "B", "J", "L", "W", "G" },
    []string{ "S", "W", "C" },
    []string{ "R", "Z", "T", "M" },
    []string{ "D", "T", "C", "H", "S", "P", "V" },
    []string{ "G", "P", "T", "L", "D", "Z" },
    []string{ "F", "B", "R", "Z", "J", "Q", "C", "D" },
    []string{ "S", "B", "D", "J", "M", "F", "T", "R" },
    []string{ "L", "H", "R", "B", "T", "V", "M" },
    []string{ "Q", "P", "D", "S", "V" },
  }

  for Stack_Height, Supply_Stack := range Supply_Stacks {
    log.Printf("Stack Row: %v - %v", Stack_Height, Supply_Stack)
  }

  for _, Full_Input_Line := range Full_Input {
    log.Print(Full_Input_Line)
    Line_Values := strings.Split(Full_Input_Line, " ")

    Moves, _ := strconv.Atoi(Line_Values[1])
    From_Row, _ := strconv.Atoi(Line_Values[3])
    To_Row, _ := strconv.Atoi(Line_Values[5])
    log.Printf("Move: %v | %v -> %v", Moves, From_Row - 1, To_Row - 1)

    for Moves != 0 {
      Top_From_Element_Location := len(Supply_Stacks[From_Row - 1]) - 1
      From_Element_Value := Supply_Stacks[From_Row - 1][Top_From_Element_Location]

      // Pop
      Supply_Stacks[From_Row - 1] = Supply_Stacks[From_Row - 1][:Top_From_Element_Location]

      // Push
      Supply_Stacks[To_Row - 1] = append(Supply_Stacks[To_Row - 1], From_Element_Value)

      log.Printf("Moved: %v | %v to %v", From_Element_Value, From_Row - 1, To_Row - 1)

      Moves -= 1
    }

    for Stack_Height, Supply_Stack := range Supply_Stacks {
      log.Printf("Stack Row: %v - %v", Stack_Height, Supply_Stack)
    }
  }

  for Stack_Height, Supply_Stack := range Supply_Stacks {
    log.Printf("Stack Row: %v - %v", Stack_Height, Supply_Stack)
  }
}
