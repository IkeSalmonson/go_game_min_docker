package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	board "go_game_min_docker/cmd/board"
	piece "go_game_min_docker/cmd/piece"
)

// gameLoop gerencia o fluxo principal do jogo, incluindo turnos e entrada do jogador.
func gameLoop(gameBoard *board.Board) {
	reader := bufio.NewReader(os.Stdin)
	currentPlayer := "white" // Peças brancas iniciam o jogo.
	turnCount := 0           // Contador de turnos para a condição de 10 turnos.

	for { // Loop principal do jogo, continua até que uma condição de vitória seja atingida.
		turnCount++            // Incrementa o contador de turnos a cada nova rodada.
		gameBoard.PrintBoard() // Exibe o estado atual do tabuleiro.

		fmt.Printf("Turno %d. É a vez das peças %s.\n", turnCount, currentPlayer)

		// 1. Obtém e exibe as peças disponíveis para o jogador atual.
		availablePieces := []*piece.Piece{}
		for _, p := range gameBoard.Pieces {
			if p.Color == currentPlayer {
				availablePieces = append(availablePieces, p)
			}
		}

		// Verifica se o jogador atual tem peças. Se não, o jogo pode ter terminado (por eliminação).
		// A condição de vitória principal será verificada após cada movimento.
		if len(availablePieces) == 0 {
			// Isso pode acontecer se o último movimento do oponente capturou todas as peças.
			// A vitória já teria sido declarada pelo PerformMove, mas é uma segurança.
			fmt.Printf("Não há mais peças %s no tabuleiro. Fim de jogo!\n", currentPlayer)
			break
		}

		fmt.Println("Suas peças disponíveis:")
		for i, p := range availablePieces {
			// Exibe as coordenadas 1-base para o usuário.
			fmt.Printf("%d: %s em (%d, %d)\n", i, p.Symbol, p.Row+1, p.Col+1)
		}

		// 2. Solicita ao jogador que escolha uma peça para mover.
		fmt.Print("Digite o número da peça que você deseja mover: ")
		pieceChoiceInput, _ := reader.ReadString('\n')
		pieceChoiceInput = strings.TrimSpace(pieceChoiceInput)
		pieceChoice, err := strconv.Atoi(pieceChoiceInput)

		if err != nil || pieceChoice < 0 || pieceChoice >= len(availablePieces) {
			fmt.Println("Escolha de peça inválida. Por favor, digite um número válido da lista.")
			continue // Volta para o início do loop para nova entrada.
		}

		selectedPiece := availablePieces[pieceChoice]
		fmt.Printf("Você selecionou a peça %s %s em (%d, %d).\n", selectedPiece.Color, selectedPiece.Symbol, selectedPiece.Row+1, selectedPiece.Col+1) // +1 para exibição ao usuário

		// 3. Obtém e exibe os movimentos válidos para a peça selecionada.
		validMoves := gameBoard.GetValidMoves(selectedPiece)

		if len(validMoves) == 0 {
			fmt.Println("Esta peça não tem movimentos válidos. Por favor, selecione outra peça.")
			continue // Volta para o início do loop para nova seleção de peça.
		}

		fmt.Println("Movimentos válidos para esta peça:")
		for i, move := range validMoves {
			// Exibe as coordenadas 1-base para o usuário.
			fmt.Printf("%d: (%d, %d)\n", i, move[0]+1, move[1]+1)
		}

		// 4. Solicita ao jogador que escolha um destino para o movimento.
		fmt.Print("Digite o número correspondente ao movimento desejado: ")
		moveChoiceInput, _ := reader.ReadString('\n')
		moveChoiceInput = strings.TrimSpace(moveChoiceInput)
		moveChoice, err := strconv.Atoi(moveChoiceInput)

		if err != nil || moveChoice < 0 || moveChoice >= len(validMoves) {
			fmt.Println("Escolha de movimento inválida. Por favor, digite um número válido da lista.")
			continue // Volta para o início do loop para nova entrada.
		}

		chosenMove := validMoves[moveChoice]
		newRow, newCol := chosenMove[0], chosenMove[1]

		// --- Executa o movimento usando o método PerformMove do tabuleiro ---
		// PerformMove lida com a atualização da posição da peça, grade e capturas.
		moveSuccessful, opponentEliminated := gameBoard.PerformMove(selectedPiece, newRow, newCol)

		if !moveSuccessful {
			fmt.Println("Erro ao executar o movimento. Por favor, tente novamente.") // Não deve acontecer se validMoves for correto.
			continue
		}

		fmt.Printf("Moveu a peça %s %s para (%d, %d).\n", selectedPiece.Color, selectedPiece.Symbol, newRow+1, newCol+1) // +1 para exibição ao usuário

		// Verifica a condição de vitória após o movimento e possível captura.
		// Primeiro, a condição de eliminação de peças.
		if opponentEliminated {
			fmt.Printf("\n--- FIM DE JOGO ---\n")
			fmt.Printf("Parabéns, jogador das peças %s! Você eliminou todas as peças do oponente e venceu o jogo!\n", currentPlayer)
			break // Sai do loop do jogo.
		}

		// Segundo, a condição de 10 turnos.
		isGameOver, result := gameBoard.CheckWinCondition(turnCount)
		if isGameOver {
			fmt.Printf("\n--- FIM DE JOGO ---\n")
			if result == "draw" {
				fmt.Println("O jogo terminou em empate após 10 turnos!")
			} else {
				fmt.Printf("O jogo terminou após 10 turnos. O jogador das peças %s venceu por ter mais peças!\n", result)
			}
			break // Sai do loop do jogo.
		}

		// Troca o turno para o próximo jogador.
		if currentPlayer == "white" {
			currentPlayer = "black"
		} else {
			currentPlayer = "white"
		}
	}
	fmt.Println("Obrigado por jogar o Go Game!")
}

func main() {
	fmt.Println("Bem-vindo ao Go Game!")
	fmt.Println("Por favor, digite as dimensões do tabuleiro.")

	// Cria o tabuleiro chamando a função de board.go.
	gameBoard := board.CreateBoard()

	// Inicia o loop principal do jogo.
	gameLoop(gameBoard)
}
