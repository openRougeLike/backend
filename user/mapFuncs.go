package user

import (
	"fmt"
)

func NewMap(level int) (userMap Map) {
	userMap.Level = level

	curRow := RandInt(0, DIMENSIONS-1)
	curCol := RandInt(0, DIMENSIONS-1)

	userMap.ChestsInfo = map[string]ChestInfo{}
	userMap.Monsters = map[string]Monsters{}
	userMap.NPC = map[string]NPCType{}
	userMap.Rooms = map[string]bool{}

	userMap.Layout = make(Layout, DIMENSIONS)

	for i := range userMap.Layout {
		userMap.Layout[i] = make([]TileType, DIMENSIONS)
	}

	// lastDir := DirDown
	randDir := Direction(RandInt(0, len(DirectionEnum) - 1))

	empties := []Coords{}
	emptyHashMap := map[string]bool{}

	tunnelLeft := MAX_TUNNEL

	for tunnelLeft != 0 {
		rand := RandInt(0, 1)
		if randDir == DirDown || randDir == DirUp {
			if rand == 0 {
				randDir = DirLeft
			} else {
				randDir = DirRight	
			}
		} else {
			if rand == 0 {
				randDir = DirUp
			} else {
				randDir = DirDown
			}
		}

		randLen := RandInt(0, MAX_TUNNEL_LEN)
		curLen := 0 // the 

		// panic(randLen)
		
		dirParsed := DirectionEnum[randDir]

		for 
		(curLen < randLen) && !(
		(curRow == 0 && dirParsed[0] == -1) ||
		(curRow >= DIMENSIONS - 1 && dirParsed[0] == 1) ||
		(curCol <= 0 && dirParsed[1] == -1) ||
		(curCol >= DIMENSIONS - 1 && dirParsed[1] == 1)) {

			userMap.Layout[curRow][curCol] = MapEmptySpace
			cordsToAdd := Coords{
				curRow,
				curCol,
			}
			if !emptyHashMap[cordsToAdd.String()] {
				emptyHashMap[cordsToAdd.String()] = true
				empties = append(empties, Coords{curRow, curCol})
			}
			
			curRow += dirParsed[0]
			curCol += dirParsed[1]

			curLen++
		}

		if curLen != 0 {
			// lastDir = randDir
			tunnelLeft--
		}
	}

	chestsMade := 0
	npcMade := 0
	monstersMade := 0
	roomsMade := 0
	taken := map[string]bool{}

	nonWalkableLeft := []Coords{}
	walkableLeft := []Coords{}


	for _, coords := range empties {
		if taken[coords.String()] {continue}

		if userMap.Layout.canPlaceNonWalkableTile(coords[1], coords[0]) {
			if Chance(CHEST_CHANCE) && chestsMade <= CHEST_MAX {
				userMap.ChestsInfo[coords.String()] = ChestInfo{
					Rare: Chance(CHEST_RARE_CHANCE),
				}
				chestsMade++
				userMap.Layout[coords[0]][coords[1]] = MapChest
				around := userMap.summonAroundCoords(coords)

				monstersMade += len(around)

				for _, coords2 := range around {
					taken[coords2.String()] = true
				}

				continue
			} else if Chance(NPC_CHANCE) && npcMade <= NPC_MAX {
				userMap.NPC[coords.String()] = NPC_DISTRIBUTION.Fetch()
				userMap.Layout[coords[0]][coords[1]] = MapNPC
				npcMade++
				continue
			}
		}
		if Chance(MONSTER_SUMMON_CHANCE) && monstersMade <= MONSTER_MAX {
			monstersMade++
			userMap.summonMonster(coords, "regular")
			continue
		} else if Chance(ROOM_CHANCE) && roomsMade == ROOM_NUM {
			roomsMade++
			userMap.Rooms[coords.String()] = false
			userMap.Layout[coords[0]][coords[1]] = MapArena
			continue
		}

		if userMap.Layout.canPlaceNonWalkableTile(coords[1], coords[0]) {
			nonWalkableLeft = append(nonWalkableLeft, coords)
		} else {
			walkableLeft = append(walkableLeft, coords)
		}
	}

	if len(nonWalkableLeft) < 10 {
		// TODO: Setup propper logs
		fmt.Println("LOG: TOO DENSE")
		userMap = NewMap(level)
		return 
	}

	fails := 0

	for {
		if fails > 10 {
			userMap = NewMap(level)
			return 
		}

		exitLoc := RandInt(0, len(nonWalkableLeft)-1)
		
		exitCords := nonWalkableLeft[exitLoc]
		
		if taken[exitCords.String()] {
			fails++
			continue
		}

		userMap.Exit = exitCords
		userMap.Layout[exitCords[0]][exitCords[1]] = MapExit
		// c := userMap.summonAroundCoords(exitCords)
		// for _, coordsE := range c {
		// 	taken[coordsE.String()] = true
		// }
		fails = 0
		taken[exitCords.String()] = true
		break
	}

	for {
		if fails > 10 {
			userMap = NewMap(level)
			return 
		}

		uLoc := RandInt(0, len(walkableLeft)+len(nonWalkableLeft)-2)
		uCoords := Coords{}

		if uLoc > len(walkableLeft) - 1 {
			uCoords = nonWalkableLeft[uLoc-len(walkableLeft)]
		} else {
			uCoords = walkableLeft[uLoc]
		}

		if taken[uCoords.String()] {
			fails++
			continue
		}

		userMap.User = uCoords
		break
	}
	
	return
}

func (m *Map) summonAroundCoords(coords Coords) []Coords {
	summonCoodrds := []Coords{}

	for i := 0; i < 4; i++ {
		curCoords := Coords{
			coords[0],
			coords[1],
		}
		switch i {
		case 0:
			if coords[0] == 0 {continue}
			curCoords[0] -= 1
		case 1:
			if coords[0] == DIMENSIONS-1 {continue}
			curCoords[0]+=1
		case 2:
			if coords[1] == 0 {continue}
			curCoords[1] -= 1
		case 3:
			if coords[1] == DIMENSIONS-1 {continue}
			curCoords[1] += 1
		}

		if m.Layout[curCoords[0]][curCoords[1]] == MapEmptySpace && !curCoords.isEqual(m.Exit) && !curCoords.isEqual(m.User) {
			if Chance(96) {
				m.summonMonster(curCoords, "coords")
				summonCoodrds = append(summonCoodrds, curCoords)
			}
		}
	}

	return summonCoodrds
}

func (m *Map) summonMonster(coords Coords, r string) {
	monst := Monsters{
		MONSTER_DISTRIBUTION.Fetch(),
	}

	chance := MONSTER_SUMMON_CHANCE*3

	summoned := 0
	
	for Chance(chance) && summoned <= 1 {
		monst = append(monst, MONSTER_DISTRIBUTION.Fetch())
		chance /= 3
		summoned++
	}

	m.Monsters[coords.String()] = monst
	m.Layout[coords[0]][coords[1]] = MapMonster
}

// Tests if you can place a non walkable tile here. If returns as true, means you can.
func (m Layout) canPlaceNonWalkableTile(x int, y int) bool {
	// nums := 0

	if x == 0 || x == DIMENSIONS - 1 || y == 0 || y == DIMENSIONS -1 {
		return false
	}

	dirUp := m[y-1][x].IsBuildWalkable()
	dirDown := m[y+1][x].IsBuildWalkable()
	dirLeft := m[y][x-1].IsBuildWalkable()
	dirRight := m[y][x+1].IsBuildWalkable()
	dirUpLeft := m[y-1][x-1].IsBuildWalkable()
	dirUpRight := m[y-1][x+1].IsBuildWalkable()
	dirDownLeft := m[y+1][x-1].IsBuildWalkable()
	dirDownRight := m[y+1][x+1].IsBuildWalkable()
	
	if dirUp && dirRight && dirDown && dirLeft {
		if !(dirUpLeft && dirUpRight && dirDownLeft && dirDownRight) {
			return false
		}
	}

	if dirUp && dirLeft && dirDown {
		if !(dirUpLeft && dirDownLeft) {
			return false
		}
	}

	if dirUp && dirRight && dirDown {
		if !(dirUpRight && dirDownRight) {
			return false
		}
	}

	if dirLeft && dirUp && dirRight {
		if !(dirUpLeft && dirUpRight) {
			return false
		}
	}

	if dirLeft && dirDown && dirRight {
		if !(dirDownLeft && dirDownRight) {
			return false
		}
	}

	if dirUp && dirLeft && !dirRight {
		if !dirUpLeft {
			return false
		}
	}

	if dirDown && dirLeft {
		if !dirDownLeft {
			return false
		}
	}

	if dirUp && dirRight {
		if !dirUpRight {
			return false
		}
	}

	if dirDown && dirRight {
		if !dirDownRight {
			return false
		}
	}

	if dirDown && dirUp {
		if !(dirLeft || dirRight) {
			return false
		}
	}

	if dirRight && dirLeft {
		if !(dirUp || dirDown) {
			return false
		}
	}

	return true
}

// Move a user in `dir` direction, by `amount`
func (u *User) Move(dir Direction, amount int) error {
	endY := u.Map.User[0] + DirectionEnum[dir][0]*amount
	endX := u.Map.User[1] + DirectionEnum[dir][1]*amount
	if endY < 0 || DIMENSIONS <= endY || endX < 0 || DIMENSIONS <= endX  {
		return ErrOutOfBoundry
	}
	if amount <= 0 {
		return ErrOutOfBoundry
	}

	for i := 1; i < amount+1; i++ {
		if !u.Map.Layout[u.Map.User[0] + DirectionEnum[dir][0]*i][u.Map.User[1] + DirectionEnum[dir][1]*i].IsWalkWalkable()  {
			return ErrWall
		}
	}

	u.Map.User = Coords{endY, endX}
	return nil
}

// go up a level in the map
// Caller is responsible for checking if its right for user to move up
func (u *User) Exit() {
	u.Map = NewMap(u.Map.Level + 1)
}
