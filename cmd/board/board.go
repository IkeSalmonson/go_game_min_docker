package board

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	piece "go_game_min_docker/cmd/piece" // Importa o pacote piece
)

// Board representa o tabuleiro do jogo.
type Board struct {
	Grid   [][]string     // A grade 2D do tabuleiro, com símbolos das peças ou espaços vazios.
	Width  int            // Largura do tabuleiro.
	Height int            // Altura do tabuleiro.
	Pieces []*piece.Piece // Slice de todas as peças ativas no tabuleiro.
}

// CreateBoard solicita ao usuário as dimensões do tabuleiro, valida-as
// e retorna uma nova instância de Board com as peças posicionadas.
func CreateBoard() *Board {
	reader := bufio.NewReader(os.Stdin)
	var width, height int
	var err error

	// Loop para obter e validar a largura do tabuleiro (entre 5 e 8)
	for {
		fmt.Print("Digite a largura do tabuleiro (entre 5 e 8): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		width, err = strconv.Atoi(input)

		if err != nil || width < 5 || width > 8 {
			fmt.Println("Entrada inválida. A largura deve ser um número entre 5 e 8.")
			continue
		}
		break
	}

	// Loop para obter e validar a altura do tabuleiro (entre 5 e 8)
	for {
		fmt.Print("Digite a altura do tabuleiro (entre 5 e 8): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		height, err = strconv.Atoi(input)

		if err != nil || height < 5 || height > 8 {
			fmt.Println("Entrada inválida. A altura deve ser um número entre 5 e 8.")
			continue
		}
		break
	}

	// Inicializa a grade do tabuleiro com espaços vazios
	grid := make([][]string, height)
	for r := range grid {
		grid[r] = make([]string, width)
		for c := range grid[r] {
			grid[r][c] = "   " // Representa uma casa vazia com 3 espaços para alinhamento
		}
	}

	// Inicializa a slice de peças
	pieces := make([]*piece.Piece, 0)

	// --- Lógica de Posicionamento Inicial das Peças ---
	// Peças Pretas: Na esquerda e topo
	// Posicionamos 3 peças pretas nas primeiras duas linhas/colunas
	pieceID := 0
	for r := 0; r < 1; r++ {
		for c := 0; c < 3; c++ {
			if bpo := piece.NewPiece(pieceID, "black", r, c); bpo != nil {
				pieces = append(pieces, bpo)
				grid[bpo.Row][bpo.Col] = bpo.Symbol
				pieceID++
			}
		}
	}

	// Peças Brancas: Na direita e fundo
	// Posicionamos 3 peças brancas nas últimas duas linhas/colunas
	pieceID = 0 // Reset ID for white pieces
	for r := height - 1; r < height; r++ {
		for c := width - 3; c < width; c++ {
			if wpo := piece.NewPiece(pieceID, "white", r, c); wpo != nil {
				pieces = append(pieces, wpo)
				grid[wpo.Row][wpo.Col] = wpo.Symbol
				pieceID++
			}
		}
	}

	fmt.Printf("Tabuleiro criado com dimensões %dx%d e peças posicionadas.\n", width, height)

	return &Board{
		Grid:   grid,
		Width:  width,
		Height: height,
		Pieces: pieces, // Armazena as peças criadas no tabuleiro
	}
}

// PrintBoard exibe o estado atual do tabuleiro no terminal.
func (b *Board) PrintBoard() {
	fmt.Println("\n----- Tabuleiro -----")
	// Imprime os números das colunas (normalizados para 1-base para o usuário)
	fmt.Print("     ") // Espaço para alinhamento com os números das linhas
	for c := 0; c < b.Width; c++ {
		fmt.Printf(" %-4d", c+1) // Imprime números de coluna 1-base
	}
	fmt.Println()

	// Imprime a linha superior da "moldura"
	fmt.Print("   +")
	for c := 0; c < b.Width; c++ {
		fmt.Print("-----")
	}
	fmt.Println("")

	for r := 0; r < b.Height; r++ {
		fmt.Printf("%2d |", r+1) // Imprime números de linha 1-base e o separador
		for c := 0; c < b.Width; c++ {
			// Usa o símbolo da peça ou espaço vazio
			fmt.Printf(" %s|", b.Grid[r][c]) // Espaço antes do símbolo e separador
		}
		fmt.Println()
		// Imprime o separador entre as linhas
		fmt.Print("   +")
		for c := 0; c < b.Width; c++ {
			fmt.Print("-----")
		}
		fmt.Println("")
	}
	fmt.Println("-------------------\n")
}

// IsValidCoord verifica se as coordenadas fornecidas estão dentro dos limites do tabuleiro.
func (b *Board) IsValidCoord(row, col int) bool {
	return row >= 0 && row < b.Height && col >= 0 && col < b.Width
}

// GetPieceAt retorna a peça nas coordenadas fornecidas, ou nil se nenhuma peça for encontrada.
func (b *Board) GetPieceAt(row, col int) *piece.Piece {
	if !b.IsValidCoord(row, col) {
		return nil
	}
	for _, p := range b.Pieces {
		if p.Row == row && p.Col == col {
			return p
		}
	}
	return nil
}

// IsSquareOccupied verifica se uma casa nas coordenadas fornecidas está ocupada por alguma peça.
func (b *Board) IsSquareOccupied(row, col int) bool {
	return b.GetPieceAt(row, col) != nil
}

// IsOccupiedByFriendly verifica se uma casa nas coordenadas fornecidas está ocupada por uma peça da cor especificada.
func (b *Board) IsOccupiedByFriendly(row, col int, color string) bool {
	p := b.GetPieceAt(row, col)
	return p != nil && p.Color == color
}

// IsOccupiedByOpponent verifica se uma casa nas coordenadas fornecidas está ocupada por uma peça da cor oposta.
func (b *Board) IsOccupiedByOpponent(row, col int, color string) bool {
	p := b.GetPieceAt(row, col)
	return p != nil && p.Color != color
}

// GetValidMoves calcula e retorna uma slice de coordenadas válidas (linha, coluna)
// onde a peça dada pode se mover, considerando o estado atual do tabuleiro e as regras de captura.
func (b *Board) GetValidMoves(p *piece.Piece) [][]int {
	validMoves := make([][]int, 0)

	// As peças podem se movimentar para um espaço vazio adjacente ou capturar uma peça oponente pulando o espaço dela.
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 { // Não pode mover para a própria casa
				continue
			}

			// --- Movimento para espaço vazio adjacente (1 passo) ---
			newRow1, newCol1 := p.Row+dr, p.Col+dc
			if b.IsValidCoord(newRow1, newCol1) && !b.IsSquareOccupied(newRow1, newCol1) {
				validMoves = append(validMoves, []int{newRow1, newCol1})
			}

			// --- Captura por salto (2 passos) ---
			// A casa adjacente (1 passo) deve conter uma peça oponente
			if b.IsValidCoord(newRow1, newCol1) && b.IsOccupiedByOpponent(newRow1, newCol1, p.Color) {
				// A casa de destino (2 passos) deve estar vazia
				newRow2, newCol2 := p.Row+(dr*2), p.Col+(dc*2)
				if b.IsValidCoord(newRow2, newCol2) && !b.IsSquareOccupied(newRow2, newCol2) {
					validMoves = append(validMoves, []int{newRow2, newCol2})
				}
			}
		}
	}
	return validMoves
}

// PerformMove atualiza o estado do tabuleiro movendo uma peça e lidando com capturas.
// Retorna (true, true) se o movimento foi bem-sucedido e o jogo terminou (todas as peças do oponente eliminadas).
// Retorna (true, false) se o movimento foi bem-sucedido, mas o jogo não terminou.
// Retorna (false, false) se o movimento foi inválido (o que não deve acontecer se GetValidMoves for correto).
func (b *Board) PerformMove(p *piece.Piece, newRow, newCol int) (bool, bool) {
	if !b.IsValidCoord(newRow, newCol) {
		return false, false // Coordenadas de destino inválidas
	}

	// Determina se uma captura ocorre
	var capturedPiece *piece.Piece = nil
	dr := newRow - p.Row
	dc := newCol - p.Col

	// Calcula a distância do movimento (máximo deslocamento absoluto)
	distance := int(math.Max(math.Abs(float64(dr)), math.Abs(float64(dc))))

	if distance == 2 { // Se a distância é 2, é um movimento de salto, potencialmente uma captura
		// Normaliza dr, dc para -1, 0 ou 1 para a direção do salto
		drUnit := dr / 2
		dcUnit := dc / 2

		intermediateRow, intermediateCol := p.Row+drUnit, p.Col+dcUnit
		if b.IsValidCoord(intermediateRow, intermediateCol) && b.IsOccupiedByOpponent(intermediateRow, intermediateCol, p.Color) {
			capturedPiece = b.GetPieceAt(intermediateRow, intermediateCol)
		}
	}

	// Remove a peça capturada, se houver
	if capturedPiece != nil {
		fmt.Printf("Peça %s %d capturada em (%d, %d)!\n", capturedPiece.Color, capturedPiece.Index+1, capturedPiece.Row+1, capturedPiece.Col+1) // +1 para exibição ao usuário
		// Remove da slice de Pieces
		for i, cp := range b.Pieces {
			if cp == capturedPiece {
				b.Pieces = append(b.Pieces[:i], b.Pieces[i+1:]...)
				break
			}
		}
		// Limpa seu símbolo da grade
		b.Grid[capturedPiece.Row][capturedPiece.Col] = "   "
	}

	// Limpa a posição antiga da peça que está se movendo na grade
	b.Grid[p.Row][p.Col] = "   "
	// Atualiza a posição interna da peça que está se movendo
	p.UpdatePosition(newRow, newCol)
	// Coloca o símbolo da peça que está se movendo na nova posição na grade
	b.Grid[newRow][newCol] = p.Symbol

	// Verifica a condição de vitória após o movimento e possível captura
	isGameOver, _ := b.CheckWinCondition(0) // turnCount não é relevante aqui para verificar a vitória por eliminação

	return true, isGameOver
}

// CheckWinCondition avalia se o jogo terminou e quem venceu ou se é um empate.
// turnCount é o número de turnos já jogados.
func (b *Board) CheckWinCondition(turnCount int) (bool, string) {
	blackPiecesCount := 0
	whitePiecesCount := 0

	for _, p := range b.Pieces {
		if p.Color == "black" {
			blackPiecesCount++
		} else {
			whitePiecesCount++
		}
	}

	// Condição de vitória 1: Eliminar todas as peças do oponente
	if blackPiecesCount == 0 {
		return true, "white" // Brancas venceram
	}
	if whitePiecesCount == 0 {
		return true, "black" // Pretas venceram
	}

	// Condição de vitória 2: Mais peças ao final de 10 turnos
	if turnCount >= 10 {
		if blackPiecesCount > whitePiecesCount {
			return true, "black" // Pretas venceram por ter mais peças
		} else if whitePiecesCount > blackPiecesCount {
			return true, "white" // Brancas venceram por ter mais peças
		} else {
			return true, "draw" // Empate
		}
	}

	return false, "" // Jogo não terminou
}
