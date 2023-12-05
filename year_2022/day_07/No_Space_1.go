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

//  Directory_Count := 0
  Current_Path := ""
  Directory_Read := false
  var Current_Directory Directory
  var Full_Directory_List []Directory
  for _, Full_Input_Line := range Full_Input {
    if Directory_Read && !strings.Contains(Full_Input_Line, "$") {
      if !strings.Contains(Full_Input_Line, "dir ") {
        File_Info := strings.Split(Full_Input_Line, " ")
        File_Size, _ := strconv.Atoi(File_Info[0])
        Current_Directory.Size += File_Size
        Current_Directory.Files = append(Current_Directory.Files, File{ Name: File_Info[1], Size: File_Size })
      }
      continue
    }

    if Directory_Read == true {
      Directory_Read = false
      Full_Directory_List = append(Full_Directory_List, Current_Directory)
//      log.Printf("Directory List: %v", Directory_List)
//      Directory_Count += 1
//      if Directory_Count > 4 {
//        for _, Final_Directory := range Full_Directory_List {
//          log.Printf("%+v", Final_Directory)
//          _ = Final_Directory
//        }
//        break
//      }
    }


    if strings.Contains(Full_Input_Line, "$ cd") {
      if !strings.Contains(Full_Input_Line[5:], "..") {
        if Full_Input_Line[5:] != "/" {
          Current_Path = Current_Path + Full_Input_Line[5:] + "/"
        } else {
          Current_Path = "/"
        }
      } else {
        Current_Path = strings.Replace(Current_Path, Current_Directory.Name + "/", "", -1)
      }

      Current_Directory = Directory{}
      Current_Directory.Name = Full_Input_Line[5:]
      log.Printf("Current_Path: %v", Current_Path)
    }

    if strings.Contains(Full_Input_Line, "$ ls") {
      Directory_Read = true
      continue
    }
  }
}

type Directory struct {
  Name    string
  Path    string
  Size    int
  Files   []File
}

type File struct {
  Name  string
  Size  int
}
