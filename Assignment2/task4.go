package main

import (
    "fmt"
    "errors"
)

type Weapon interface {
    SetName(name string) error
    SetPower(power int) error
    GetName() string
    GetPower() int
}

type Gun struct {
    name  string
    power int
}

func (g *Gun) SetName(name string) error {
    if name == "" {
        return errors.New("weapon name cant be empty")
    }
    g.name = name
    return nil
}

func (g *Gun) GetName() string {
    return g.name
}

func (g *Gun) SetPower(power int) error {
    if power < 0 {
        return errors.New("weapon power cant be negative")
    }
    g.power = power
    return nil
}

func (g *Gun) GetPower() int {
    return g.power
}

type AK47 struct {
    Gun
}

func NewAK47() Weapon {
    return &AK47{
        Gun: Gun{
            name:  "AK-47",
            power: 4,
        },
    }
}

type Musket struct {
    Gun
}

func NewMusket() Weapon {
    return &Musket{
        Gun: Gun{
            name:  "Musket",
            power: 1,
        },
    }
}

func GetWeapon(weaponType string) (Weapon, error) {
    switch weaponType {
    case "ak47":
        return NewAK47(), nil
    case "musket":
        return NewMusket(), nil
    default:
        return nil, errors.New("unsupported weapon type")
    }
}

func main() {
    ak47, err := GetWeapon("ak47")
    if err != nil {
        fmt.Println(err)
        return
    }

    musket, err := GetWeapon("musket")
    if err != nil {
        fmt.Println(err)
        return
    }

    PrintWeaponDetails(ak47)
    PrintWeaponDetails(musket)
}

func PrintWeaponDetails(weapon Weapon) {
    fmt.Printf("Weapon: %s\n", weapon.GetName())
    fmt.Printf("Power: %d\n", weapon.GetPower())
}
