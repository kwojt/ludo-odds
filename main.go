package main

type pawn struct {
	field int
}

const (
	blueSpawn   = 0
	greenSpawn  = 10
	yellowSpawn = 20
	redSpawn    = 30
)
const (
	b1 = iota
	b2
	b3
	b4
	g1
	g2
	g3
	g4
	y1
	y2
	y3
	y4
	r1
	r2
	r3
	r4
)

func oddsOfDanger(player int) {
	player = 0 // 0, 3, 7, 11
	pawns := [4][4]int{
		{-1, -1, -1, 0}, //Blue team
		{-1, -1, -1, 0}, //Green team
		{-1, -1 - 1, 0}, //Yellow team
		{-1, -1, -1, 0}} //Red team
	baseOffset := [4]int{0, 10, 20, 30}
	dangerCounter := 0

	for plrPawn := 0; plrPawn < 4; plrPawn++ {
		for enemy := 0; enemy < 4; enemy++ {
			// Do not check you own pawns
			if enemy == player {
				continue
			}
			for enmPawn := 0; enmPawn < 4; enmPawn++ {
				// If enemy's pawn is in the base
				if enmPawn < 0 {
					break
				}
				// If player's pawn is in enemy's pawn range it adds one to danger counter
				if (pawns[player][plrPawn]+baseOffset[player])%40 > (pawns[enemy][enmPawn]+baseOffset[enemy])%40 && (pawns[player][plrPawn]+baseOffset[player])%40 <= (pawns[enemy][enmPawn]+baseOffset[enemy]+6)%40 {
					dangerCounter++
				}
				// If player's pawn is at enemy's spawn, it adds two to danger counter(1 or 6 to go)
				for enmSpawn := 0; enmSpawn < 4; enmSpawn++ {
					if enmSpawn == player {
						continue
					}
					if (pawns[player][plrPawn] + baseOffset[player]) == baseOffset[player] {
						dangerCounter += 2
					}
				}
			}
		}
	}
}

func main() {

}
