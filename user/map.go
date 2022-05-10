package user

import (
	"regexp"
)

var DirectionRegexp = regexp.MustCompile(`(up|down|left|right)`)

type Layout [][]TileType

type ChestInfo struct {
	Rare bool `json:"rare"`
	Open bool `json:"open"`
}

type Map struct {
	Layout Layout `json:"layout"`
	Exit Coords `json:"exit"`
	// strings.Join(Coords, ",") : Monsters
	Monsters map[string]Monsters `json:"monsters"`
	// strings.Join(Coords, ",") : completed them
	Rooms map[string]bool `json:"roomsCleared"`
	
	// strings.Join(Coords, ",") : completed them
	ChestsInfo map[string]ChestInfo `json:"chests"`
	
	User Coords `json:"user"`
	Level int `json:"level"`
	NPC map[string]NPCType `json:"npcs"`
}

type NPCType int8

// 0 - y, 1 - x
type Coords [2]int

type TileType int8

const (
	MapWall TileType = iota // MUST BE THERE BY DEFAULT!
	MapEmptySpace
	MapChest
	MapMonster
	MapArena
	MapNPC
	MapExit
)

const (
	NPCShopWeapon NPCType = iota
	NPCShopArmor
	NPCShopPotion
	NPCUpgrade
	NPCLore0
	NPCLore1
	NPCLore2
	NPCLore3
	NPCLore4
	NPCLore5
	NPCLore6
	NPCLore7
	NPCLore8
	NPCLore9
)

type Direction int8

const (
	DirUp Direction = iota
	DirRight
	DirDown
	DirLeft
	// Special direction - same place. Used only in "action" endpoints, where you can litterly be over the action, which is currently only in areanas.
	DirCur
)

var DirectionEnum = map[Direction]Coords{
	DirLeft: {0, -1},
	DirRight: {0, 1},
	DirUp: {-1, 0},
	DirDown: {1, 0},
	DirCur: {0,0},
}

const DIMENSIONS = 26

const MAX_TUNNEL = 80
const MAX_TUNNEL_LEN = 9


const MONSTER_MAX = 50
const MONSTER_SUMMON_CHANCE = 6

const CHEST_MAX = 8
const CHEST_CHANCE = 3
const CHEST_RARE_CHANCE = 10

const ROOM_NUM = 2
const ROOM_CHANCE = 2

const NPC_CHANCE = 5
const NPC_MAX = 3

var NPC_DISTRIBUTION = ChanceDistribution[NPCType]{
	{Max: 4, Value: NPCShopArmor},
	{Max: 8, Value: NPCShopPotion},
	{Max: 12, Value: NPCShopPotion},
	{Max: 20, Value: NPCUpgrade},
	{Max: 28, Value: NPCLore0},
	{Max: 36, Value: NPCLore1},
	{Max: 44, Value: NPCLore2},
	{Max: 52, Value: NPCLore3},
	{Max: 60, Value: NPCLore4},
	{Max: 68, Value: NPCLore5},
	{Max: 76, Value: NPCLore6},
	{Max: 84, Value: NPCLore7},
	{Max: 92, Value: NPCLore8},
	{Max: 100, Value: NPCLore9},
}

var MONSTER_DISTRIBUTION = ChanceDistribution[MonsterType]{
	{Max: 5, Value: MTRange},
	{Max: 15, Value: MTTank},
	{Max: 16, Value: MTGlassSword},
	{Max: 21, Value: MTFire},
	{Max: 26, Value: MTEarth},
	{Max: 31, Value: MTAir},
	{Max: 36, Value: MTWater},
	{Max: 41, Value: MTDark},
	{Max: 46, Value: MTLight},
	{Max: 51, Value: MTSlow},
}