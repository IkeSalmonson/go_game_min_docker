package piece

import (
	"fmt" // Mantido para fmt.Sprintf, que é necessário para formatar o símbolo
)

// Piece representa uma peça genérica no jogo.
// Com as regras simplificadas, todas as peças têm o mesmo tipo de movimento.
type Piece struct {
	Index  int    // Identificador numérico único para a peça (0-base)
	Color  string // Cor da peça ("black" ou "white")
	Row    int    // Posição atual da peça (linha, 0-base)
	Col    int    // Posição atual da peça (coluna, 0-base)
	Symbol string // Símbolo para representação no terminal (ex: "B1", "W2")
}

// NewPiece cria e retorna uma nova instância de Piece.
// O símbolo da peça é determinado pela sua cor e índice numérico.
func NewPiece(index int, color string, row, col int) *Piece {
	var symbol string
	if color == "black" {
		symbol = fmt.Sprintf("B%d ", index+1) // Peça Preta, com índice 1-base (ex: B1, B2)
	} else {
		symbol = fmt.Sprintf("W%d ", index+1) // Peça Branca, com índice 1-base (ex: W1, W2)
	}

	return &Piece{
		Index:  index,
		Color:  color,
		Row:    row,
		Col:    col,
		Symbol: symbol,
	}
}

// UpdatePosition atualiza a linha e coluna da peça.
func (p *Piece) UpdatePosition(newRow, newCol int) {
	p.Row = newRow
	p.Col = newCol
}
