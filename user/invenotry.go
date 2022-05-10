package user

type Inventory struct {
	Armor ArmorSlots
	Weapon Weapon
	Container Container
	BagLevels BagLevels
}

type BagLevels struct {
	Other,
	Potions,
	Weapons,
	Materials int
}

type Container struct {
	Other map[int]GeneralItem
	Potions map[int]Potion
	Weapons map[int]Weapon
	Materials map[int]Material
	Map map[int]BagType
}

type GeneralItem struct {
	Potion
	Weapon
	Material
	BagType
}

type Potion struct {
	Enhancement
	Permanent bool
}

type Material struct {
	Upgrade1 Enhancement
	Upgrade2 Enhancement
}

type ArmorSlots struct {
	Helmet,
	Chestplate,
	Pants,
	Boots,
	Ring,
	Necklace Armor
}

type ItemBase struct {
	Enhancement Enhancement
	Ability1 Enhancement
	Ability2 Enhancement
	// Set 0 = nothing
	Set int8
	ItemType 

	Quality QualityType
	IsGold bool
}

type Enhancement struct {
	Enhancement EnhancementType
	Value float64
}

type Armor struct {
	ItemBase
	ArmorType ArmorType
	Defense,
	Health float64
}

type WeaponAbility struct {
/* -1 = All Enemies
0 = Allies
1, 2, 3, etc = N number of enemies */
	EffectedN int8
	Type WeaponAbilityType
	Value int8
}

type Weapon struct {
	ItemBase
	Damage int
	WeaponAbility WeaponAbility
}

type ArmorType int8
type EnhancementType int8
type QualityType int8
type WeaponAbilityType int8
type BagType int8
type ItemType int8

const (
	ItemTypeArmor ItemType = iota
	ItemTypeWeapon
	
	ItemTypePotion
	ItemTypeMaterial
)

const (
	Other BagType = iota
	Potions
	Weapons
	Materials
)

const (
	WAbHeal = iota
	WAbManaHeal
	WAbAttack
	WAb
)

const (
	QualityBad QualityType = iota
	QualityGood
	QualityGreat
	QualityLegendary
)

const (
	ENCH_UnEnhanced EnhancementType = iota
	ENCH_Health
	ENCH_Mana
	ENCH_Defense
	ENCH_Dodge
	ENCH_Attack
	ENCH_Power
	ENCH_Crit
	ENCH_InstaKill
	ENCH_Greed
	ENCH_XpGreed
	ENCH_Luck
	ENCH_ElemFireAttack
	ENCH_ElemFireDefense
	ENCH_ElemWaterAttack
	ENCH_ElemWaterDefense
	ENCH_ElemAirAttack
	ENCH_ElemAirDefense
	ENCH_ElemEarthAttack
	ENCH_ElemEarthDefense
	ENCH_ElemLightAttack
	ENCH_ElemLightDefense
	ENCH_ElemDarkAttack
	ENCH_ElemDarkDefense
)

const (
	ArmorHelmet ArmorType = iota
	ArmorChestplate
	ArmorPants
	ArmorBoots
	ArmorRing
	ArmorNecklace
)
