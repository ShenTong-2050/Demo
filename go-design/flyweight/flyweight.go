package flyweight

// units 所有棋子
var units = map[int]*ChessPieceUnits{
	1:{
		ID: 1,
		Name: "車",
		Color: "red",
	},
	2:{
		ID: 2,
		Name: "马",
		Color: "red",
	},
	3:{
		ID: 3,
		Name: "相",
		Color: "red",
	},
}

// ChessPieceUnits 单个执行操作的象棋棋子【棋子享元】
type ChessPieceUnits struct {
	ID 			int
	Name,Color 	string
}

// NewChessPieceUnit 棋子工厂
func NewChessPieceUnit(id int) *ChessPieceUnits {
	return units[id]
}

// ChessPiece 棋子位置
type ChessPiece struct {
	Unit 	*ChessPieceUnits
	X,Y 	int
}

// ChessBoard 棋盘
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// NewChessBoard 初始化棋盘 起始位置
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{Unit: NewChessPieceUnit(id),X: 0,Y: 0}
	}
	return board
}

// Move 移动棋子
func (c ChessBoard) Move(id ,x,y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}