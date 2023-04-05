Once upon a time there was a great hero, called Orderus, with some strengs and weaknesses, as all heroes have.
After battling all kinds of monsters for more than a hundred years, Orderus now has the following stats:

* HEALTH: 70 - 100
* STRENGTH: 70 - 80
* DEFENCE: 45 - 55
* SPEED: 40 - 50
* LUCK: 10% - 30%

He possesses two skills:

**Rapid strike**: Strike twice while it's his turn to attack, there's a 10% he'll use this skill every time he attacks

**Magic shield**: takes only half of the usual damage when an enemy attacks, there's a 20% change he'll use this skill every time he defends.

As Orderus walks through a magical ever-green forest, he encounters some wild beasts, with the following properties:
* HEALTH: 60 - 90 
* STRENGTH: 60 - 90 
* DEFENCE: 40- 60 
* SPEED: 40 - 60 
* LUCK: 25% - 40%

The first attack is done by the player with the higher speed. If both players have the same speed, then the attack is carried on by the player with the highest luck. After an attack, the players switch roles: the attacker now defends and the defender now attacks.
The damage done by the attacker is calculated with the following formula:

`Damage = Attacker strength - Defender defence`

The game ends when one of the players remain without health or the number of turns reaches 20.
The application must output the results each turn: what happens , which skills were used ( if any ), the damage done, defender's health left.
If we have a winner before the maximum number of rounds is reached, he must be declared.
