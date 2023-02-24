package private

type data struct {
	x   int
	Y   int
}

// NewData 将 data 私有结构供包外使用
func NewData() *data {
	return new(data)
}
