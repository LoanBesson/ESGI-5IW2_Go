package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type game struct {
	isEnd     bool
	nbPlayers int
	players   []player
}

type player struct {
	firtsname string
	lastname  string
	health    int8
	score     int8
}

func main() {

	var game = game{false, 0, nil}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Création d'un nouveau joueur...")
	fmt.Println("Saisissez un prénom :")
	fmt.Print("> ")

	firstname, _ := reader.ReadString('\n')
	firstname = strings.Replace(firstname, "\n", "", -1)

	fmt.Println("Saisissez un nom :")
	fmt.Print("> ")

	lastname, _ := reader.ReadString('\n')
	lastname = strings.Replace(lastname, "\n", "", -1)

	game.addPlayer(player{firtsname: firstname, lastname: lastname, health: 100, score: 0})

	fmt.Println("Joueur ajouté !")

	fmt.Println("Informations du joueur :")
	fmt.Println("Prénom : ", game.players[0].firtsname)
	fmt.Println("Nom : ", game.players[0].lastname)
	fmt.Println("Santé : ", game.players[0].health)
	fmt.Println("Score : ", game.players[0].score)

}

func (myGame *game) addPlayer(player player) {
	myGame.players = append(myGame.players, player)
	myGame.nbPlayers++
}
