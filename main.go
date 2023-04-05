package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var hero Creature
var beast Creature
var run bool = true

type Creature struct {
  name string
  role string
  health int
  strength int
  defence int
  speed int
  luck int
  rapidStrike int
  magicShield int
}

func (e Creature) getInfo() {
  fmt.Printf("\n%s:\nhealth: %d\nstrength: %d\ndefence: %d\nspeed: %d\nluck: %d\n", strings.ToUpper(e.name), e.health, e.strength, e.defence, e.speed, e.luck)
}

func attack(nr int) {
  // choose an attacker & defender
  var attacker *Creature
  var defender *Creature
  if nr == 0 {
    if hero.speed > beast.speed {
      attacker, defender = &hero, &beast
    } else if hero.speed < beast.speed {
      attacker, defender = &beast, &hero
    } else {
      if hero.luck > beast.luck {
        attacker, defender = &hero, &beast
      } else { 
        attacker, defender = &beast, &hero
      }
    }
  } else {
    if attacker == &hero {
      attacker, defender = &beast, &hero
    } else {
      attacker, defender = &hero, &beast
    }
  }

  var attackKoef = 1


  fmt.Printf("\n---> Attack %d\n* Attacker: %s\n* Defender: %s - health %d", nr + 1, attacker.name, defender.name, defender.health)
  if !isLucky(defender.luck) {
    attackKoef = 0
    fmt.Printf("\n- %s is lucky this time!", defender.name)
  }
  if attacker == &hero && isLucky(hero.rapidStrike) {
    attackKoef = 2
    fmt.Printf("\n- %s uses skill 'Rapid Strike'", attacker.name)
  }
  if defender == &hero && isLucky(hero.magicShield) {
    attackKoef = 0
    fmt.Printf("\n- %s uses skill 'Magic Shield'", defender.name)
  }
  
  damage := (attacker.strength - defender.defence) * attackKoef
  defender.health -= damage
  if defender.health <= 0 {
    fmt.Printf("\n%s health - 0, damage - %d\n", defender.name, damage + defender.health)
    fmt.Printf("\nTHE WINNER: %s!", attacker.name)
    run = false
  } else {
    fmt.Printf("\n%s health - %d, damage - %d\n", defender.name, defender.health, damage)
  }
  
}


func randomInt(min int, max int) int {
  return rand.Intn(max - min) + min
}

func isLucky(percent int) bool {
  return rand.Intn(100) < percent
}

func main() {
  hero = Creature {
    name: "Orderus",
    health: randomInt(70, 100),
    strength: randomInt(70, 80),
    defence: randomInt(45, 55),
    speed: randomInt(40, 50),
    luck: randomInt(10, 30),
    rapidStrike: 10,
    magicShield: 20,
  }
  hero.getInfo()

  names := [5]string {"Adalgar", "Berthold", "Friedrich", "Gerhardt", "Reinhold"}

  beast = Creature {
    name: names[rand.Intn(5)],
    health: randomInt(60, 90),
    strength: randomInt(60, 90),
    defence: randomInt(40, 60),
    speed: randomInt(40, 60),
    luck: randomInt(25, 40),
  }
  beast.getInfo()

  done := make(chan bool)

    go func() {
      var i = 0

      for run {
        time.Sleep(5 * time.Second)
        attack(i)
        if i == 20 {
          fmt.Print("DRAW!")
          done <- true
        }
        i++
      }

      done <- true
    }()

  <- done
}



