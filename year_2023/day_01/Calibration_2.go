package main

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "errors"
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

  var Calibration_Sum, Calibration, Errors int
  for _, Full_Input_Line := range Full_Input {
//    log.Printf(Full_Input_Line)
    Calibration, err = Find_Number(Full_Input_Line)
    if err != nil {
      Errors += 1
      log.Printf("Finding number failed. (%v)", err)
      continue
    }

    Calibration_Sum += Calibration
  }

  log.Printf("Number of errors: %v", Errors)
  log.Printf("Calibration sum: %v", Calibration_Sum)
}

func Find_Number(Line string) (int, error) {
  Split_Line := strings.Split(Line, "")

  var Int int
  var err error
  var Word []string
  var Numbers []int
  var Full_String string
  for _, Character := range Split_Line {
    Int, err = strconv.Atoi(Character)

    if err != nil {
//      log.Printf("# %v # - Not a number", Character)
      Word = append(Word, Character)

      Full_String = strings.Join(Word, "")
      if strings.Contains(Full_String, "one") {
        Numbers = append(Numbers, 1)
        Word = []string{"n", "e"}
      }
      if strings.Contains(Full_String, "two") {
        Numbers = append(Numbers, 2)
        Word = []string{"w", "o"}
      }
      if strings.Contains(Full_String, "three") {
        Numbers = append(Numbers, 3)
        Word = []string{"h", "r", "e", "e"}
      }
      if strings.Contains(Full_String, "four") {
        Numbers = append(Numbers, 4)
        Word = []string{"o", "u", "r"}
      }
      if strings.Contains(Full_String, "five") {
        Numbers = append(Numbers, 5)
        Word = []string{"i", "v", "e"}
      }
      if strings.Contains(Full_String, "six") {
        Numbers = append(Numbers, 6)
        Word = []string{"i", "x"}
      }
      if strings.Contains(Full_String, "seven") {
        Numbers = append(Numbers, 7)
        Word = []string{"e", "v", "e", "n"}
      }
      if strings.Contains(Full_String, "eight") {
        Numbers = append(Numbers, 8)
        Word = []string{"i", "g", "h", "t"}
      }
      if strings.Contains(Full_String, "nine") {
        Numbers = append(Numbers, 9)
        Word = []string{"i", "n", "e"}
      }

      continue
    }

    Word = []string{}
    Numbers = append(Numbers, Int)
  }

  if len(Numbers) == 0 {
    return 0, errors.New("No numbers in line")
  }

  var Converted_Number int
  Converted_Number, err = strconv.Atoi(fmt.Sprintf("%v%v", Numbers[0], Numbers[len(Numbers) - 1]))

  if err != nil {
    return 0, errors.New("Converted number failed")
  }

  return Converted_Number, nil
}
