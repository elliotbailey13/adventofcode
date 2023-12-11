package main

import (
  "os"
  "log"
  "bufio"
  "slices"
  "errors"
  "strconv"
  "strings"
)

var (
  Seeds []int

  // 0 - Reset | 1 - Seed_to_Soil | 2 - Soil_to_Fertilizer | 3 - Fertilizer_to_Water
  // 4 - Water_to_Light | 5 - Light_to_Temperature | 6 - Temperature_to_Humidity | 7 - Humidity_to_Location
  Mode = 0

  Seed_Soil []Seed_Map
  Soil_Fertilizer []Seed_Map
  Fertilizer_Water []Seed_Map
  Water_Light []Seed_Map
  Light_Temperature []Seed_Map
  Temperature_Humidity []Seed_Map
  Humidity_Location []Seed_Map
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

  for Index, Full_Input_Line := range Full_Input {
    log.Printf(Full_Input_Line)
    err = Load_Data(Full_Input_Line)
    if err != nil {
      log.Printf("File load failed on line %v\n%v", Index, Full_Input_Line)
      log.Fatal(err)
    }
  }

  log.Printf("Finished loading data. Total seeds %v.", len(Seeds))

  var Locations []int
  for _, Seed := range Seeds {
    Soil_Location, _ := Find_Next_Location(Seed, Seed_Soil)
    Fertilizer_Location, _ := Find_Next_Location(Soil_Location, Soil_Fertilizer)
    Water_Location, _ := Find_Next_Location(Fertilizer_Location, Fertilizer_Water)
    Light_Location, _ := Find_Next_Location(Water_Location, Water_Light)
    Temperature_Location, _ := Find_Next_Location(Light_Location, Light_Temperature)
    Humidity_Location_2, _ := Find_Next_Location(Temperature_Location, Temperature_Humidity)
    Location, _ := Find_Next_Location(Humidity_Location_2, Humidity_Location)

    Locations = append(Locations, Location)
  }

  log.Printf("Finished loading locations. Total locations %v.", len(Locations))

  slices.Sort(Locations)

  log.Printf("Lowest location: %v", Locations[0])
}

func Load_Data(Line string) error {
  var err error

  if Line == "" {
    Mode = 0

    return nil
  }

  if strings.Contains(Line, "seeds:") {
    var Seed_ID int

    var Start int
    var Length bool
    Seed_Split := strings.Split(Line, " ")
    for Index, Seed_ID_Str := range Seed_Split {
      if Index == 0 { continue }

      Seed_ID, err = strconv.Atoi(Seed_ID_Str)
      if err != nil { return err }

      if Length == false {
        Start = Seed_ID
        Length = true

        continue
      }

      Org_Start := Start
      Seeds = append(Seeds, Start)
      for i := Seed_ID; i != 1; i-- {
        Start += 1
	Seeds = append(Seeds, Start)
      }

      Length = false

      log.Printf("Finished seed %v", Org_Start)
    }

    return nil
  }

  if Mode == 0 {
    switch {
      case strings.Contains(Line, "seed-to-soil"):
        Mode = 1
      case strings.Contains(Line, "soil-to-fertilizer"):
        Mode = 2
      case strings.Contains(Line, "fertilizer-to-water"):
        Mode = 3
      case strings.Contains(Line, "water-to-light"):
        Mode = 4
      case strings.Contains(Line, "light-to-temperature"):
        Mode = 5
      case strings.Contains(Line, "temperature-to-humidity"):
        Mode = 6
      case strings.Contains(Line, "humidity-to-location"):
        Mode = 7
    }

    return nil
  }

  Data_Split := strings.Split(Line, " ")

  var Seed_Info_Ptr *[]Seed_Map
  switch Mode {
    case 1:
      Seed_Info_Ptr = &Seed_Soil

    case 2:
      Seed_Info_Ptr = &Soil_Fertilizer

    case 3:
      Seed_Info_Ptr = &Fertilizer_Water

    case 4:
      Seed_Info_Ptr = &Water_Light

    case 5:
      Seed_Info_Ptr = &Light_Temperature

    case 6:
      Seed_Info_Ptr = &Temperature_Humidity

    case 7:
      Seed_Info_Ptr = &Humidity_Location

    default:
      return errors.New("Seed mode can't be found")
  }

  var Destination int
  var Source int
  var Range int

  Destination, err = strconv.Atoi(Data_Split[0])
  if err != nil { return err }

  Source, err = strconv.Atoi(Data_Split[1])
  if err != nil { return err }

  Range, err = strconv.Atoi(Data_Split[2])
  if err != nil { return err }

  *Seed_Info_Ptr = append(*Seed_Info_Ptr, Seed_Map{
       Destination: Destination,
       Source: Source,
       Range: Range,
     })

  return nil
}

func Find_Next_Location(Source int, Seed_Map_Arr []Seed_Map) (int, error) {
  for _, Seed_Map_Item := range Seed_Map_Arr {
    if Source >= Seed_Map_Item.Source && Source <= ( Seed_Map_Item.Source + ( Seed_Map_Item.Range - 1 ) ) {

      return ( (Source - Seed_Map_Item.Source) + Seed_Map_Item.Destination ), nil
    }
  }

  return Source, nil
}

type Seed_Map struct {
  Destination int
  Source int
  Range int
}
