package main

import (
  "os"
  "log"
  "bufio"
  "strings"
)

var (
  Directions []string
  Nodes map[string]Node
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

  Nodes = make(map[string]Node)
  for _, Full_Input_Line := range Full_Input {
    Load_Data(Full_Input_Line)
  }

  Steps := Find_End()

  log.Printf("Steps to end: %v", Steps)
}

func Load_Data(Line string) {
  if Line == "" { return }
  if !strings.Contains(Line, "=") {
    Directions = strings.Split(Line, "")

    return
  }

  Line = strings.Replace(Line, " = (", " ", -1)
  Line = strings.Replace(Line, ", ", " ", -1)
  Line = strings.Replace(Line, ")", " ", -1)

  Node_IDs := strings.Split(Line, " ")
  Nodes[Node_IDs[0]] = Node{ Left: Node_IDs[1], Right: Node_IDs[2] }
}

func Find_End() int {
  Steps := 0
  var Current_Node Node
  Current_Location := "AAA"
  for true {
    for _, Direction := range Directions {
      if Current_Location == "ZZZ" { return Steps }

      Current_Node = Nodes[Current_Location]
      if Direction == "L" { Current_Location = Current_Node.Left }
      if Direction == "R" { Current_Location = Current_Node.Right }

      Steps++
    }
  }

  return 0
}

type Node struct {
  Left string
  Right string
}
