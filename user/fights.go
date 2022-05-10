package user

import (
	"errors"
	"math"
)

type MonsterType int8

const (
	MTRange MonsterType = iota
	MTTank
	MTGlassSword
	MTFire
	MTEarth
	MTAir
	MTWater
	MTDark
	MTLight
	MTSlow
	MonsterType10 //TODO:
	MonsterType11
	MonsterType12
)

type Monsters []MonsterType

type Fight struct {
	MonstersAtStart Monsters `json:"monstersAtStart"`
	Monsters []Monster `json:"currentMonsters"`
	Level int `json:"level"`
	Coords string `json:"location"`
}

type Monster struct {
	Type MonsterType `json:"type"`
	MaxHP int `json:"MaxHealth"`
	HP int `json:"health"`
	Attack int `json:"attack"`
	Charge int `json:"charge"`
	ChargeNeeded int `json:"chargeNeeded"`
	Elements Elements `json:"elements"`
}

var MonsterTypeMap = map[MonsterType]monsterSummonInfo{ //TODO:
	MTRange: {
		Range: summonInfoInfo{
			Base: 40,
			Upgradable: true,
			UpgradeFactor: 2,
		},
		Attack: summonInfoInfo{
			Base: 41,
			Upgradable: true,
			UpgradeFactor: 1.7,
		},
		Health: summonInfoInfo{
			Base: 75,
			Upgradable: true,
			UpgradeFactor: 1.5,
		},
		Charge: summonInfoInfo{
			Base: 1,
			Upgradable: false,
		},
	},
	MTTank: {
		Range: summonInfoInfo{
			Base: 20,
			Upgradable: false,
		},
		Attack: summonInfoInfo{
			Base: 50,
			Upgradable: true,
			UpgradeFactor: 2.3,
		},
		Health: summonInfoInfo{
			Base: 125,
			Upgradable: true,
			UpgradeFactor: 2.1,
		},
		Charge: summonInfoInfo{
			Base: 1,
			Upgradable: false,
		},
	},
	MTGlassSword: {
		Range: summonInfoInfo{
			Base: 0,
			Upgradable: false,
		},
		Attack: summonInfoInfo{
			Base: 100,
			Upgradable: true,
			UpgradeFactor: 2.3,
		},
		Health: summonInfoInfo{
			Base: 1,
			Upgradable: false,
		},
		Charge: summonInfoInfo{
			Base: 1,
			Upgradable: false,
		},
	},
	MTFire: {
		Health: summonInfoInfo{
			Base: 1,
			Upgradable: false,
		},
		Elements: summonInfoElements{
			Fire: summonInfoElementAtt{
				Attack: summonInfoInfo{
					Base: 30,
					Upgradable: true,
					UpgradeFactor: 2.5,
				},
				Defense: summonInfoInfo{
					Base: -1,
					Upgradable: false,
				},
			},
			Water: summonInfoElementAtt{
				Defense: summonInfoInfo{
					Base: 0,
					Upgradable: false,
				},
			},
			Air: summonInfoElementAtt{
				Defense: summonInfoInfo{
					Base: 20,
					Upgradable: true,
					UpgradeFactor: 3.5,
				},
			},
			Dark: summonInfoElementAtt{
				Defense: summonInfoInfo{
					Base: 20,
					Upgradable: true,
					UpgradeFactor: 3.5,
				},
			},
			Light: summonInfoElementAtt{
				Defense: summonInfoInfo{
					Base: 20,
					Upgradable: true,
					UpgradeFactor: 3.5,
				},
			},
			Earth: summonInfoElementAtt{
				Defense: summonInfoInfo{
					Base: 20,
					Upgradable: true,
					UpgradeFactor: 3.5,
				},
			},
		},
	},
	MTEarth: {
		// TODO:
	},
	MTSlow: {
		Range: summonInfoInfo{
			Base: 0,
			Upgradable: false,
		},
		Attack: summonInfoInfo{
			Base: 160,
			Upgradable: true,
			UpgradeFactor: 2.3,
		},
		Health: summonInfoInfo{
			Base: 75,
			Upgradable: true,
			UpgradeFactor: 2.75,
		},
		Charge: summonInfoInfo{
			Base: 3,
			Upgradable: false,
		},
	},
}

type monsterSummonInfo struct {
	Range summonInfoInfo
	Health summonInfoInfo
	Charge summonInfoInfo
	Attack summonInfoInfo
	Elements summonInfoElements
}

type summonInfoElements struct {
	Fire,
	Water,
	Air,
	Earth,
	Light,
	Dark summonInfoElementAtt
}

type summonInfoElementAtt struct {
	Attack,
	Defense summonInfoInfo
}

type summonInfoInfo struct {
	Upgradable bool
	Base float64
	UpgradeFactor float64
}

func (info summonInfoInfo) Value(Level int) int {
	if info.Upgradable {
		return int(math.Floor(info.Base * info.UpgradeFactor * float64(Level)))
	} else {
		return int(math.Floor(info.Base))
	}
}

func (info summonInfoInfo) ValueFloat(Level int) float64 {
	if info.Upgradable {
		return info.Base * info.UpgradeFactor * float64(Level)
	} else {
		return info.Base
	}
}

func (monsters Monsters) Generate(Level int) Fight {
	fight := Fight{
		MonstersAtStart: monsters,
		Level: Level,
	}

	for _, monster := range monsters {
		summon := MonsterTypeMap[monster]
		
		curMonster := Monster{
			Type: monster,
			MaxHP: summon.Health.Value(Level),
			HP: summon.Health.Value(Level),
			Attack: summon.Attack.Value(Level),
			Charge: 0,
			ChargeNeeded: summon.Charge.Value(Level),
			Elements: Elements{
				Fire: ElementAtcDef{
					Attack: summon.Elements.Fire.Attack.ValueFloat(Level),
					Defense: summon.Elements.Fire.Defense.ValueFloat(Level),
				},
				Water: ElementAtcDef{
					Attack: summon.Elements.Water.Attack.ValueFloat(Level),
					Defense: summon.Elements.Water.Defense.ValueFloat(Level),
				},
				Air: ElementAtcDef{
					Attack: summon.Elements.Air.Attack.ValueFloat(Level),
					Defense: summon.Elements.Air.Defense.ValueFloat(Level),
				},
				Earth: ElementAtcDef{
					Attack: summon.Elements.Earth.Attack.ValueFloat(Level),
					Defense: summon.Elements.Earth.Defense.ValueFloat(Level),
				},
				Light: ElementAtcDef{
					Attack: summon.Elements.Light.Attack.ValueFloat(Level),
					Defense: summon.Elements.Light.Defense.ValueFloat(Level),
				},
				Dark: ElementAtcDef{
					Attack: summon.Elements.Dark.Attack.ValueFloat(Level),
					Defense: summon.Elements.Dark.Defense.ValueFloat(Level),
				},
			},
		}

		fight.Monsters = append(fight.Monsters, curMonster)
	}

	return fight
}

var ErrNotInFight = errors.New("NotInFight")

func (user *User) FightEnd() error {
	if user.State != STATE_FIGHT {
		return ErrNotInFight
	}
	
	fight := user.Fight
	newMonsters := []MonsterType{}

	for _, monster := range fight.Monsters {
		if monster.HP == 0 {
			continue
		}

		newMonsters = append(newMonsters, monster.Type)
	}

	user.Map.Monsters[fight.Coords] = newMonsters

	user.Fight = Fight{}
	user.State = STATE_NORMAL

	return nil
}

func (user *User) FightStart(Coordinates Coords) error {
	// handle locations w/ coords.String()
	// errors
	// TODO:

	return nil
}