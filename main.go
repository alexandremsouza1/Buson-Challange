package main;

import (
	"fmt"
	"buson.challange.com/src"
)


func main() {
	var diameter, height, width, length int

	fmt.Print("Digite o diâmetro da bola: ")
	fmt.Scan(&diameter)
	ball, err := src.NewBall(diameter)
	if err != nil {
		fmt.Println("Erro ao criar a bola:", err)
		return
	}

	fmt.Print("Digite as dimensões da caixa com virgulas (altura, largura, comprimento): ")
	_, err = fmt.Scanf("%d,%d,%d", &height, &width, &length)
	if err != nil {
		fmt.Println("Erro ao ler as dimensões da caixa:", err)
		return
	}

	box, err := src.NewBox(height, width, length)
	if err != nil {
		fmt.Println("Erro ao criar a caixa:", err)
		return
	}

	if src.CheckHasFit(*ball,*box) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}


