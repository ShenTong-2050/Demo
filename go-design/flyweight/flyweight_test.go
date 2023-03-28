package flyweight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChessBoard(t *testing.T) {

	chessBoard1 := NewChessBoard()
	// 棋子 1 移动到 2,3 位置
	chessBoard1.Move(1,2,3)

	// 获取 2 3 坐标上的棋子
	for _,chess := range chessBoard1.chessPieces {
		if chess.X == 0 && chess.Y == 0 {
			fmt.Println(chess.Unit)
		}
	}

	chessBoard2 := NewChessBoard()
	// 棋子 2 移动到 3,4 位置
	chessBoard2.Move(2,3,4)

	// 第一个棋盘对象的 第一个棋子 与 第二个棋盘的 第一个棋子是否相同
	assert.Equal(t, chessBoard1.chessPieces[1].Unit,chessBoard2.chessPieces[1].Unit)
	assert.Equal(t, chessBoard1.chessPieces[2].Unit,chessBoard2.chessPieces[2].Unit)
}
