package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChessBoard(t *testing.T) {

	chessBoard1 := NewChessBoard()
	// 棋子 1 移动到 2,3 位置
	chessBoard1.Move(1,2,3)

	chessBoard2 := NewChessBoard()
	// 棋子 2 移动到 3,4 位置
	chessBoard2.Move(2,3,4)

	// 第一个棋盘的第一个棋子 与 第二个棋盘的第二个棋子 是返回的内容应该是一致的
	assert.Equal(t, chessBoard1.chessPieces[1].Unit,chessBoard2.chessPieces[1].Unit)
	assert.Equal(t, chessBoard1.chessPieces[2].Unit,chessBoard2.chessPieces[2].Unit)
}
