package main

import (
  "fmt"
  "math/rand"
)

type Creature struct {
  health int
  strength int
  defence int
  speed int
  luck int
}

func (e Creature) getInfo() {
  fmt.Printf("INFO: %d %d %d %d %d", e.health, e.strength, e.defence, e.speed, e.luck)
}

func randomInt(min int, max int) int {
  return rand.Intn(max - min) + min
}

func main() {
  names := [5]string {"Adalgar", "Berthold", "Friedrich", "Gerhardt", "Reinhold"}
  hero := Creature {
    health: randomInt(70, 100),
    strength: randomInt(70, 80),
    defence: randomInt(45, 55),
    speed: randomInt(40, 50),
    luck: randomInt(10, 30),
  }
  hero.getInfo()

  fmt.Printf("\nFirst beast is %s\n", names[rand.Intn(5)])
  beast := Creature {
    health: randomInt(60, 90),
    strength: randomInt(60, 90),
    defence: randomInt(40, 60),
    speed: randomInt(40, 60),
    luck: randomInt(25, 40),
  }
  beast.getInfo()
  
  var attacker Creature
  if hero.speed > beast.speed {
    attacker = hero
  } else if hero.speed < beast.speed {
    attacker = beast
  } else {
    if hero.luck > beast.luck {
      attacker = hero
    } else {
      attacker = beast
    }
  }
  fmt.Println(attacker)
}

