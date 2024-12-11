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
  All_Paths []Node
  Next_Paths []Node
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

  Find_Starting_Nodes()

  for _, Path := range All_Paths {
    log.Printf("Paths: %+v", Path)
  }

  log.Printf("Total paths: %v", len(All_Paths))
  Steps := Find_End()

  for _, Path := range All_Paths {
    log.Printf("Paths: %+v", Path)
  }
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
  Strings := strings.Split(Node_IDs[0], "")
  Nodes[Node_IDs[0]] = Node{
    Index: Node_IDs[0],
    Left: Node_IDs[1],
    Right: Node_IDs[2],
    Last_Index: Strings[2],
  }
}

func Find_End() int {
  Steps := 0
  var Ends int
  var New_Path string
  for true {
    for _, Direction := range Directions {
      Ends = 0
      Next_Paths = []Node{}

      for _, Path := range All_Paths {
        if Direction == "L" { New_Path = Path.Left }
        if Direction == "R" { New_Path = Path.Right }

        if Nodes[New_Path].Last_Index == "Z" { Ends++ }
        Next_Paths = append(Next_Paths, Nodes[New_Path])
      }

      if Ends > 2 {
        log.Printf("Ends: %v Steps: %v", Ends, Steps)
      }

      Steps++

      if Ends == len(All_Paths) { return Steps }
      All_Paths = Next_Paths
    }
  }

  return 0
}

func Find_Starting_Nodes() {
  for Index, Node := range Nodes {
    if Node.Last_Index == "A" {
      log.Printf("Found start: %v - %+v", Index, Node)
      All_Paths = append(All_Paths, Node)
    }
  }
}

type Node struct {
  Index string
  Left string
  Right string
  Last_Index string
}
