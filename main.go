// به نام خدا
package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

// 3 tover game

// ever tover have hp and Attack every round
// player in the him round have 2 chose Attack or Def
// tover's have 3 mode :
// 		overflow [+attckTime , +Attack , -less hp every Attack , if miss come out] ,
// 		normal [+Def , +hpreg],
// 		inFire [-get dps , -less Attack , +have dps]

// Attack , Attack time , miss , dps , hp , hpreg

// TODO: create Attack and hp | create Attack time | create modes for player | create modes for tavers
// TODO: add menu and mode of easy , normal , hard and milad

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}
type NumberU interface {
	uint | uint8 | uint16 | uint32 | uint64
}
type NumberInt interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Taver struct {
	hp         int8
	Attack     int8
	Def        int8
	AttackTime int8
	Time       int8
	MaxTime    int8
	Died       bool
}

type Player struct {
	hp         int16
	Attack     int8
	Def        int8
	AttackTime int8
}

// porints
var terun uint16 = 0

func main() {

	// clear cmd
	clear()

	// mune show
	fmt.Print(
		`
____________     ____   ___  ___    ________  ______
\___________\   / __ \  \  \ \  \  /   ____/ / __   \
    \    \     / /__\ \  \  \/  / /   /___  / |__|  /
     \    \   /  ____  \  \    / /   ____/ /  __   /
      \____\ /__/    \__\  \__/ /____\___ /__/ 	\__\`,
	)

	// list of mune
	fmt.Println("\n \n #1. Play	#2. Modes	#3. Help	#4. Quit \n \n ")

	var code int
	fmt.Scan(&code)
	switch code {
	case 1:
		Game()
	case 2:

	case 3:
		Help()
	case 4:
		break
	default:
		fmt.Println("your code is wrong!")
	}

}

func Game() {
	// set tavers value
	var tavers [3]Taver = [3]Taver{
		{hp: 99, Def: 3, AttackTime: 3, MaxTime: 7, Attack: 10},
		{hp: 99, Def: 3, AttackTime: 3, MaxTime: 7, Attack: 10},
		{hp: 99, Def: 3, AttackTime: 3, MaxTime: 7, Attack: 10},
	}

	// mack time to max time
	for i := 0; i < len(tavers); i++ {
		tavers[i].Time = tavers[i].MaxTime
	}

	var player Player = Player{hp: 200, Attack: 50, Def: 3, AttackTime: 2}

	var msg string

mainLoop:
	for {

		// add points
		terun++

		// check game is win or loss or continue
		switch {
		case gameOver(&player):
			whenGameOver()
			break mainLoop

		case tavers[0].Died && tavers[1].Died && tavers[2].Died:
			whenGameWin()
			break mainLoop
		}

		// clear the cmd
		clear()

		// send massage
		if msg != "" {
			fmt.Printf("⚠️ %v", msg)
			msg = ""
		}
		// leet tavers attack
		for i := 0; i < len(tavers); i++ {
			tavers[i].taversLive()
			tavers[i].AttackingTaverToPlayer(&player)
		}

		// show tavers and player
		if err := showTaver(tavers); err != nil {
			log.Fatal("\n you have local err : ", err.Error())
		}
		player.show()

		// print how you want
		fmt.Println("\n say what you want to Attack:")

		// get code from player
		var code int
		if _, err := fmt.Scan(&code); err != nil {
			log.Fatal("\n you have binery err : ", err.Error())
		}

		// set code jobs
		switch code {
		case 0:
			break mainLoop
		case 1, 2, 3:
			player.AttackTaver(&tavers[code-1])
		case 5:
			main()
			break mainLoop
		default:
			msg = "your code is worng!!"
		}

	}

}

// help to how play page
func Help() {

	// clear the cmd
	clear()

	// show helps
	fmt.Print(
		`
in Game:

 > attack the taver : 1 , 2 , 3
 > quit fast : 0
 > back to menu : 5

in Mune:

 > play : 1
 > modes : 2
 > help : 3
 > quit : 4
 `,
	)

	// show buttons
	fmt.Println("\n \n #1. Back	#2. Quit")

	// get code
	var code int
	fmt.Scan(&code)

	// set codes
	switch code {
	case 1:
		main()
	case 2:
		break
	default:
		fmt.Println("your code is wrong")
	}

}

// when game is ended
func whenGameOver() {
	fmt.Println(" \n \n .......  you lose !!! ........ \n ")
	fmt.Println(" \n #1. Back   #2. Play agine   #3. Quit \n ")

	// get code
	var code int
	fmt.Scan(&code)

	// set code jobs
	switch code {
	case 1:
		main()
	case 2:
		Game()
	case 3:
		break
	default:
		fmt.Println("your code is wrong!!")
	}
}

// when game is win
func whenGameWin() {
	fmt.Println(" \n \n .......  you win (O_ o) ........ \n ")
	fmt.Println(" \n #1. Back   #2. Play agine   #3. Quit \n ")

	// get code
	var code int
	fmt.Scan(&code)

	// set code jobs
	switch code {
	case 1:
		main()
	case 2:
		Game()
	case 3:
		break
	default:
		fmt.Println("your code is wrong!!")
	}
}

func gameOver(player *Player) bool {
	return player.hp <= 0
}

func showTaver(tavers [3]Taver) error {
	for _, elemnet := range tavers {
		if elemnet.hp > 99 {
			return errors.New("your taver hp too big")
		}
	}

	// log.Fatal("your hp Number is to big");

	// show tavers down or not
	var topOfTaver [3]string = [3]string{" __ ___ __ ", " __ ___ __ ", " __ ___ __ "}

	for i := 0; i < len(tavers); i++ {
		if tavers[i].Died {
			topOfTaver[i] = "           "
		}
	}

	fmt.Printf(` 
%v %v %v
|  |   |  | |  |   |  | |  |   |  |
 \%v/   \%v/   \%v/ 
 |      |    |      |    |      |  
 |   1  |    |   2  |    |   3  | 

  ..%v..      ..%v..      ..%v..
	`, topOfTaver[0], topOfTaver[1], topOfTaver[2],
		reCreate(tavers[0].Time, tavers[0].MaxTime, "."), reCreate(tavers[1].Time, tavers[1].MaxTime, "."), reCreate(tavers[2].Time, tavers[2].MaxTime, "."),
		tavers[0].hp, tavers[1].hp, tavers[2].hp)

	return nil
}

func (taver *Taver) AttackingTaverToPlayer(player *Player) {
	if !taver.Died {
		taver.Time -= ((rand8(taver.AttackTime-1) + 1) - (rand8(player.AttackTime / 2)))

		// Time to attack
		if taver.Time <= 0 {

			player.hp -= int16(max(rand8(taver.Attack)-rand8(player.Def), 0))

			taver.Time = taver.MaxTime
		}
	} else {
		taver.Time = 0
	}
}

// check tavers live or not
func (taver *Taver) taversLive() {
	taver.Died = taver.hp <= 0
}

func reCreate[T Number](num T, maxNum T, text string) string {
	var resulte string
	var i T
	for ; i < num; i++ {
		resulte += text
	}
	for i = 0; i < maxNum-num; i++ {
		resulte += " "
	}
	return resulte
}

func (p Player) AttackTaver(taver *Taver) {
	taver.hp -= max(rand8(p.Attack)-rand8(taver.Def), 0)
	taver.hp = max(taver.hp, 0)
}

func (player Player) show() {

	fmt.Printf(" \n \nplayer : ❤️  %v  💪 %v  🏹  %v", player.hp, player.Attack, player.AttackTime)
}

// my needed funcs in this case

func clear() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func rand8(a int8) int8 {
	return int8(rand.Intn(int(a)))
}

func max[T Number](a T, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

// الهم عجل لویک الفرج
