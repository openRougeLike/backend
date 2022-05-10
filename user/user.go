package user

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	
	XP int `json:"xp"`
	Lvl int `json:"level"`
	Money int `json:"money"`
	
	Base Upgradable `json:"baseSkills"`

	// Current HP
	Hp int `json:"currentHP"`
	// Current Mana
	Mana int `json:"currentMP"`

	BaseElements Elements `json:"baseElements"`

	Statistics Statistics `json:"statistics"`

	State UserState `json:"userState"`
	
	Map Map `json:"map"`

	Inventory Inventory `json:"inventory"`

	Fight Fight `json:"fight"`

	Token string `json:"-"`
}

type UserState int8

const (
	STATE_NORMAL UserState = iota
	STATE_FIGHT
)

type StatisticsCurrent struct {
	Kills int
}

type Statistics struct {
	StatisticsCurrent
	CurrentRun StatisticsCurrent
	Deaths int
}

type Elements struct {
	Fire ElementAtcDef `json:"fire"`
	Water ElementAtcDef `json:"water"`
	Air ElementAtcDef `json:"air"`
	Earth ElementAtcDef `json:"earth"`
	Light ElementAtcDef `json:"light"`
	Dark ElementAtcDef `json:"dark"`
}

type ElementAtcDef struct {
	Attack float64 `json:"attack"`
	Defense float64 `json:"defense"`
}

type Upgradable struct {
	// Health amount
	Health float64 `json:"health"`
	// Mana amount
	Mana float64 `json:"mana"`
	// Defense amount
	Defense float64 `json:"defense"`
	// Dodge chance
	Dodge float64 `json:"dodge"`
	// Attack dmg
	Attack float64 `json:"attack"`
	// 'Power' value. Can be used for skills
	Power float64 `json:"power"`
	// Crit chance
	Crit float64 `json:"critical"`
	// Insta kill chance (if crit is achieved, then this chance is applied, so in reality its crit chance * Insta kill chance)
	InstaKill float64 `json:"instaKill"`

	// % of extra money you get
	Greed float64 `json:"greed"`
	// % of extra xp you get
	XpGreed float64 `json:"xpGreed"`
	// chance for a better item from a chest, get item from monster, chance for a better chest type 
	Luck float64 `json:"luck"`
}

type Stats struct {
	Upgradable
	Elements
}

func (user User) FetchStats() Stats {
	stats := Stats{
		Upgradable: user.Base,
		Elements: user.BaseElements,
	}
	stats.AddArmor(user.Inventory.Armor.Boots)
	stats.AddArmor(user.Inventory.Armor.Chestplate)
	stats.AddArmor(user.Inventory.Armor.Helmet)
	stats.AddArmor(user.Inventory.Armor.Necklace)
	stats.AddArmor(user.Inventory.Armor.Pants)
	stats.AddArmor(user.Inventory.Armor.Ring)
	
	return stats
}

func (stats *Stats) AddArmor(armor Armor) {
	if armor.Set == 0 {
		return
	}

	stats.AddEnch(armor.Ability1)
	stats.AddEnch(armor.Ability2)
	stats.AddEnch(armor.Enhancement)
	stats.Health += armor.Health
	stats.Defense += armor.Defense
}

func (stats *Stats) AddEnch(ability Enhancement) {
	var abilityThing *float64

	switch ability.Enhancement {
	case ENCH_UnEnhanced:
		return
	case ENCH_Health:
		abilityThing = &stats.Health
	case ENCH_Mana:
		abilityThing = &stats.Mana
	case ENCH_Defense:
		abilityThing = &stats.Defense
	case ENCH_Dodge:
		abilityThing = &stats.Dodge
	case ENCH_Attack:
		abilityThing = &stats.Attack
	case ENCH_Power:
		abilityThing = &stats.Power
	case ENCH_Crit:
		abilityThing = &stats.Crit
	case ENCH_InstaKill:
		abilityThing = &stats.InstaKill
	case ENCH_Greed:
		abilityThing = &stats.Greed
	case ENCH_XpGreed:
		abilityThing = &stats.XpGreed
	case ENCH_Luck:
		abilityThing = &stats.Luck
	case ENCH_ElemFireAttack:
		abilityThing = &stats.Fire.Attack
	case ENCH_ElemFireDefense:
		abilityThing = &stats.Fire.Defense
	case ENCH_ElemWaterAttack:
		abilityThing = &stats.Water.Attack
	case ENCH_ElemWaterDefense:
		abilityThing = &stats.Water.Defense
	case ENCH_ElemAirAttack:
		abilityThing = &stats.Air.Attack
	case ENCH_ElemAirDefense:
		abilityThing = &stats.Air.Defense
	case ENCH_ElemEarthAttack:
		abilityThing = &stats.Earth.Attack
	case ENCH_ElemEarthDefense:
		abilityThing = &stats.Earth.Defense
	case ENCH_ElemLightAttack:
		abilityThing = &stats.Light.Attack
	case ENCH_ElemLightDefense:
		abilityThing = &stats.Light.Defense
	case ENCH_ElemDarkAttack:
		abilityThing = &stats.Dark.Attack
	case ENCH_ElemDarkDefense:
		abilityThing = &stats.Dark.Defense
	}
	
	*abilityThing += ability.Value
}