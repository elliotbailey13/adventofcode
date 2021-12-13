package main

import (
  "os"
  "log"
  "bufio"
  "sort"
  "strconv"
  "strings"
)

func main() {
  Raw_Input, err := os.Open("./input.txt")
  if err != nil { log.Fatal(err) }

  var Raw_String_Array []string
  scanner := bufio.NewScanner(Raw_Input)
  for scanner.Scan() {
    Raw_String_Array = append(Raw_String_Array, scanner.Text())
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }

  Raw_Input.Close()

  var First_Signal_Display [][]string
  var Output_Signal_Display [][]string
  for _, Raw_String_Row := range Raw_String_Array {
    Row_Split := strings.Split(Raw_String_Row, " | ")
    Output_Signal_Display = append(Output_Signal_Display, strings.Split(Row_Split[1], " "))
    First_Signal_Display = append(First_Signal_Display, strings.Split(Row_Split[0], " "))
  }

  _ = First_Signal_Display

  Total_Value := 0
  for Index, Output_Signal_Display_Row := range Output_Signal_Display {
    Decoder_Map := Decoder(Output_Signal_Display_Row, First_Signal_Display[Index])
//    log.Printf("Decoder Map: %+v", Decoder_Map)
    Value := Decode(Decoder_Map, Output_Signal_Display_Row)
    Total_Value += Value
  }

  log.Printf("Total Value of all Ooutputs: %v", Total_Value)
}

func Decoder(Output_Values, First_Values []string) map[string]string {
  var Five_Codes, Six_Codes []string
  var One_Code, Four_Code, Seven_Code, Eight_Code string
  Full_Values := append(Output_Values, First_Values...)
  for _, Full_Value := range Full_Values {
    switch len(Full_Value) {
      case 2:
        if One_Code == "" { One_Code = Full_Value }

      case 3:
        if Seven_Code == "" { Seven_Code = Full_Value }

      case 4:
        if Four_Code == "" { Four_Code = Full_Value }

      case 5:
        Five_Codes = append(Five_Codes, Full_Value)

      case 6:
        Six_Codes = append(Six_Codes, Full_Value)

      case 7:
        if Eight_Code == "" { Eight_Code = Full_Value }
    }
  }

//  log.Printf("One Code: %v - Four Code: %v - Seven Code: %v - Eight Code: %v", One_Code, Four_Code, Seven_Code, Eight_Code)

  Decoder_Map := map[string]string{ "a": "z", "b": "z", "c": "z", "d": "z", "e": "z", "f": "z", "g": "z" }

  for _, Code := range strings.Split(Seven_Code, "") {
    Found := false
    for _, Compare_Code := range strings.Split(One_Code, "") {
      if Compare_Code == Code { Found = true }
    }
    if Found == false { Decoder_Map["a"] = Code }
  }

  bd_String := ""
  for _, Code := range strings.Split(Four_Code, "") {
    Found := false
    for _, Compare_Code := range strings.Split(One_Code, "") {
      if Compare_Code == Code { Found = true }
    }
    if Found == false { bd_String += Code }
  }


  eg_String := ""
  for _, Code := range strings.Split(Eight_Code, "") {
    Found := false
    for _, Compare_Code := range strings.Split(Four_Code, "") {
      if Compare_Code == Code { Found = true }
    }
    if Code == Decoder_Map["a"] { Found = true }

    if Found == false { eg_String += Code }
  }

  Five_Six_Codes := append(Five_Codes, Six_Codes...)
  for _, Five_Six_Code := range Five_Six_Codes {
    cf_Temp := ""
    bd_Temp := ""
    eg_Temp := ""
    cf_Count := 0
    bd_Count := 0
    eg_Count := 0
    for _, Code := range strings.Split(Five_Six_Code, "") {
      for _, Compare_Code := range strings.Split(One_Code, "") {
        if Compare_Code == Code {
          cf_Temp = Code
          cf_Count++
        }
      }
      for _, Compare_Code := range strings.Split(bd_String, "") {
        if Compare_Code == Code {
          bd_Temp = Code
          bd_Count++
        }
      }
      for _, Compare_Code := range strings.Split(eg_String, "") {
        if Compare_Code == Code {
          eg_Temp = Code
          eg_Count++
        }
      }
    }

    if cf_Count == 2 && bd_Count == 1 && eg_Count == 2 {
      // 0
      Decoder_Map["b"] = bd_Temp
    }

    if cf_Count == 1 && bd_Count == 1 && eg_Count == 2 {
      // 2
      Decoder_Map["c"] = cf_Temp
      Decoder_Map["d"] = bd_Temp
    }

    if cf_Count == 2 && bd_Count == 1 && eg_Count == 1 {
      // 3
      Decoder_Map["d"] = bd_Temp
      Decoder_Map["g"] = eg_Temp
    }

    if cf_Count == 1 && bd_Count == 2 && eg_Count == 1 {
      // 5
      Decoder_Map["f"] = cf_Temp
      Decoder_Map["g"] = eg_Temp
    }

    if cf_Count == 1 && bd_Count == 2 && eg_Count == 2 {
      // 6
      Decoder_Map["f"] = cf_Temp
    }
    if cf_Count == 2 && bd_Count == 2 && eg_Count == 1 {
      // 9
      Decoder_Map["g"] = eg_Temp
    }
  }

  Codes := []string{"a" , "b", "c", "d", "e", "f", "g"}
  for Index, Code := range Codes {
    for _, Value := range Decoder_Map {
      if Value == Code { Codes[Index] = "" }
    }
  }

  for _, Code := range Codes {
    if Code != "" { Decoder_Map["e"] = Code }
  }

  return Decoder_Map
}

func Decode(Decoder_Map map[string]string, Output_Values []string) int {
  Real_Values_Map := map[string]string{
    "0": "abcefg",
    "1": "cf",
    "2": "acdeg",
    "3": "acdfg",
    "4": "bcdf",
    "5": "abdfg",
    "6": "abdefg",
    "7": "acf",
    "8": "abcdefg",
    "9": "abcdfg"}

  Real_Value := ""
  for _, Output_Value := range Output_Values {
    Decoded_Value_Array := []string{}
    for Real_Code, Map_Code := range Decoder_Map {
      for _, Value := range strings.Split(Output_Value, "") {
        if Value == Map_Code { Decoded_Value_Array = append(Decoded_Value_Array, Real_Code) }
      }
    }

    sort.Strings(Decoded_Value_Array)
    Decoded_Value := ""
    for _, Value := range Decoded_Value_Array { Decoded_Value += Value }

    for Value, Code := range Real_Values_Map {
      if Code == Decoded_Value { Real_Value += Value }
    }
  }

  log.Printf("Real Value: %v", Real_Value)

  Real_Value_Int, _ := strconv.Atoi(Real_Value)

  return Real_Value_Int
}
