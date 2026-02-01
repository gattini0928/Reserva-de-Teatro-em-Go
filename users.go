package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func main() {
	reader:= bufio.NewReader(os.Stdin)
	const playA = 16
	const playB = 18 
	const playC = 12

	plays := map[string][]string {
		"sombras do silêncio": {
			"A-1", "A-2", "A-3", "A-4", "A-5",
			"A-6", "A-7", "A-8", "A-9", "A-10",
			"B-1", "B-2", "B-3", "B-4", "B-5",
			"B-6", "B-7", "B-8", "B-9", "B-10",
			"C-1", "C-2", "C-3", "C-4", "C-5",
			"C-6", "C-7", "C-8", "C-9", "C-10"},
		"o último ato": {
			"A-1", "A-2", "A-3", "A-4", "A-5",
			"A-6", "A-7", "A-8", "A-9", "A-10",
			"B-1", "B-2", "B-3", "B-4", "B-5",
			"B-6", "B-7", "B-8", "B-9", "B-10",
			"C-1", "C-2", "C-3", "C-4", "C-5",
			"C-6", "C-7", "C-8", "C-9", "C-10"},
		"entre cortinas": {
			"A-1", "A-2", "A-3", "A-4", "A-5",
			"A-6", "A-7", "A-8", "A-9", "A-10",
			"B-1", "B-2", "B-3", "B-4", "B-5",
			"B-6", "B-7", "B-8", "B-9", "B-10",
			"C-1", "C-2", "C-3", "C-4", "C-5",
			"C-6", "C-7", "C-8", "C-9", "C-10"},
		}

	for {
		fmt.Println("🎭 - {Bem-vindo ao teatro das máscaras!} ")

		nameValid, age , ok := readSpectator(reader)
		if !ok {
			continue
		}

		menu(nameValid)

		var input string
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Tente Novamente, opção inválida.")
			continue
		}

		var playName string

		if choice == 1 {
			playName = "sombras do silêncio"
			playOk := manageChoosedPlay(reader, age, playA, plays, playName)
			if !playOk {
				continue
			}
			
		} else if choice == 2 {
			playName = "o último ato"
			playOk := manageChoosedPlay(reader, age, playB, plays, playName)
			if !playOk {
				continue
			}
		}else if choice == 3 {
			playName = "entre cortinas" 
			playOk := manageChoosedPlay(reader, age, playC, plays, playName)
			if !playOk {
				continue
			}
		}else if choice == 4 {
			fmt.Printf("🖐️- Volte sempre %s!", nameValid)
			break
		}else {
			fmt.Println("❌- Opção Inválida, Tente Novamente.")
			continue
		}
	}
}

func readSpectator(reader *bufio.Reader) (string, int, bool){
	var name string
	var age string

	fmt.Println("🪪 - Seu Nome: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("🪪 - Sua Idade: ")
	age, _ = reader.ReadString('\n')
	age = strings.TrimSpace(age)
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return "", 0, false
	}
	nameValid, ok := validName(name)
		if !ok {
			return nameValid, ageInt, false
		}

		spectators := make(map[string]string)
		spectators["name"] = nameValid
		spectators["age"] = age

		return nameValid, ageInt, true
}


func manageChoosedPlay(reader *bufio.Reader, age int, agePerm int, plays map[string][]string, playName string) bool{
	if  age < agePerm {
		fmt.Printf("🪪 - Desculpe a peça %s é para maiores de %d anos \n", strings.ToUpper(playName), agePerm)
			return false
		} else {
				if !hasAvaibleSeats(plays, playName) {
					fmt.Printf("😢 os acentos para a peça (%s) foram esgotados. \n", strings.ToUpper(playName))
					return false
				}
				seat := chooseSeat(reader, playName, plays)
				reserveSeat(seat, plays, playName)
			}
	return true
}

func hasAvaibleSeats(plays map[string][]string, playName string) bool {
	return len(plays[playName]) > 0
}

func menu(nameValid string) {
	fmt.Printf("🎭 %s escolha uma das peças abaixo 👇\n", nameValid)
	fmt.Println("👤 1. Sombras do Silêncio (+16)")
	fmt.Println("🔪 2. O Último Ato (+18)")
	fmt.Println("🧕 3. Entre Cortinas (+12)")
	fmt.Println("🏃‍♀️‍➡️ 4. Sair do Teatro 😢")
	fmt.Print("> ")
}

func chooseSeat(reader *bufio.Reader, playName string, plays map[string][]string) string{
	fmt.Printf("🪑- Acentos Disponíveis para a peça %s:\n", playName)
	fmt.Println(plays[playName])
	fmt.Println("🪑- Escolha um acento: ")
	var seat string
	seat, _ = reader.ReadString('\n')
	seat = strings.TrimSpace(seat)
	seat = strings.ToUpper(seat)
	return seat
}

func reserveSeat(seat string,plays map[string][]string, playName string) {
		seats := plays[strings.ToLower(playName)]
		seatIndex := -1

		for i, s := range seats{
			if s == seat {
				seatIndex = i
				break
			}
		}

		// Verificar se o assento existe
		if seatIndex == -1  {
			fmt.Println("❌ Assento indisponível ou não encontrado!")
		} else {
			// Remover o assento
			plays[playName] = append(seats[:seatIndex], seats[seatIndex+1:]...)
			fmt.Printf("🎊 - Assento (%s) escolhido com sucesso para peça (%s), se divirta!\n", seat, strings.ToUpper(playName))
		}
}

func validName(name string) (string, bool){
	if len(name) < 4 {
		fmt.Println("🪪 Digite seu nome completo.")
		return name, false
	}
	return name, true
}


