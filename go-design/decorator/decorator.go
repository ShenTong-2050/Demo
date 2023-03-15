package decorator

type IDraw interface {
	Draw() string
}

// Square 画个正方形
type Square struct {}

// Draw 正方形 实现 画 的 接口
func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare 给正方形填充颜色
type ColorSquare struct {
	square Square
	color  string
}

// NewColorSquare 获取 ColorSquare 对象
func NewColorSquare(s Square, color string) ColorSquare {
	return ColorSquare{square: s,color: color}
}

// Draw 彩色正方形 实现 画 的接口
func (cs ColorSquare) Draw() string {
	return cs.square.Draw() + " the color is " + cs.color
}
