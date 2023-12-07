package main

import (
  "os"
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

  var Game_Power int
  var Game_Power_Sum int
  var Game_Power_Errors int
  for _, Full_Input_Line := range Full_Input {
    log.Printf(Full_Input_Line)

    Game_Power, err = Parse_Line(Full_Input_Line)
    if err != nil {
      log.Printf("Game had error: %v", err)
      Game_Power_Errors += 1
      continue
    }

    Game_Power_Sum += Game_Power
  }

  log.Printf("Game Power Errors: %v", Game_Power_Errors)
  log.Printf("Game Power Sum: %v", Game_Power_Sum)
}

func Parse_Line(Line string) (int, error) {
  var err error
  var Game_ID int
  var Red_Count int
  var Blue_Count int
  var Green_Count int

  Game_ID_Split := strings.Split(Line, ":")

  Game_ID_Arr := strings.Split(Game_ID_Split[0], " ")
  Game_ID, err = strconv.Atoi(Game_ID_Arr[1])
  if err != nil {
    return 0, errors.New("Game ID conversion failed")
  }

  log.Printf("Game ID: %v", Game_ID)

  Roll_Arr := strings.Split(Game_ID_Split[1], ";")
  for Roll_Count, Roll := range Roll_Arr {
    Roll = strings.TrimSpace(Roll)
    log.Printf("Game %v: %v - %v", Game_ID, Roll_Count, Roll)

    var Cube_Count int
    Cube_Arr := strings.Split(Roll, ", ")
    for Cube_ID, Cube := range Cube_Arr {
      log.Printf("Game %v: Roll %v: %v - %v", Game_ID, Roll_Count, Cube_ID, Cube)
      if strings.Contains(Cube, "red") {
        Cube_Count, err = strconv.Atoi(strings.Replace(Cube, " red", "", -1))
	if err != nil {
          return 0, errors.New("Red cube count failed to convert")
	}

	if Cube_Count > Red_Count { Red_Count = Cube_Count }
      }

      if strings.Contains(Cube, "blue") {
        Cube_Count, err = strconv.Atoi(strings.Replace(Cube, " blue", "", -1))
	if err != nil {
          return 0, errors.New("Blue cube count failed to convert")
	}

	if Cube_Count > Blue_Count { Blue_Count = Cube_Count }
      }

      if strings.Contains(Cube, "green") {
        Cube_Count, err = strconv.Atoi(strings.Replace(Cube, " green", "", -1))
	if err != nil {
          return 0, errors.New("Green cube count failed to convert")
	}

	if Cube_Count > Green_Count { Green_Count = Cube_Count }
      }
    }
  }

  Total_Power := Red_Count * Blue_Count * Green_Count
//  log.Printf("Game %v: Red - %v | Blue - %v | Green - %v", Game_ID, Red_Count, Blue_Count, Green_Count)

  return Total_Power, nil
}
