# solarlunar-go

Chinese solar lunar calendar converter

# Installation

```sh
go get github.com/jetfueltw/solarlunar-go
```

# Useage

公曆轉農曆
```go
solar := solarlunar.Solar{Year: 2015, Month: 1, Day: 15}
lunar := solarlunar.SolarToLunar(solar)
// Lunar {
//     Year: 2014,
//     Month: 11,
//     Day: 25,
//     IsLeap: false
// }
```

農曆轉公曆
```go
lunar := solarlunar.Lunar{Year: 2014, Month: 11, Day: 25, IsLeap: false}
solar := solarlunar.LunarToSolar(lunar)
// Solar {
//     Year: 2015,
//     Month: 1,
//     Day: 15
// }
```
