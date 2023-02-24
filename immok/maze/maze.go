package main

import (
	"fmt"
	"os"
)

type point struct {
	i,j int	// i,j 分别代表纵坐标与横坐标
}

var dirs = [4]point{
	{-1,0},{0,-1},{1,0},{0,1},
}

// 探索下一个坐标算法
func (p point) add(n point) point {
	return point{p.i+n.i,p.j+n.j}	
}

// 确定探索的条件
func (p point) at(grid [][]int) (int,bool) {

	// 超出上下边界
	if p.i < 0 || p.i >= len(grid) {
		return 0,false
	}

	// 超出左右边界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0,false
	}

	return grid[p.i][p.j],true
}

func readMaze(filename string) [][]int {
	file,err := os.Open(filename)
	if err != nil {
		return nil
	}

	var row,col int

	fmt.Fscanf(file,"%d %d",&row,&col)

	//fmt.Println(row,col)

	maze := make([][]int,row)

	for i := range maze {
		maze[i] = make([]int,col)
		for j := range maze[i] {
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}

func walk(maze [][]int, start,end point) [][]int {
	// 构建步骤二维切片
	steps := make([][]int,len(maze))
	for i := range maze {
		steps[i] = make([]int,len(maze[i]))
	}

	Q := []point{start}	// 创建一个队列来确定探索的边界条件
	
	for len(Q) > 0 {
		
		cur := Q[0]
		Q = Q[1:]
		
		// 结束条件
		if cur == end {
			break
		}

		for _,dir := range dirs {
			
			next := cur.add(dir)
			
			val, ok := next.at(maze)
			if !ok || val == 1 {	// 如果遇墙或超出边界则探索下一个
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue	// 走过则跳出
			}
		
			if next == start {
				continue	// 回到起始点则跳出
			}

			curSteps,_ := cur.at(steps)

			steps[next.i][next.j] = curSteps + 1
			
			Q = append(Q,next)
			
		}
	}
	return steps
}

func main() {

	filename := "maze.in"

	maze := readMaze(filename)

	/*for _,row := range maze {
		for _,val := range row {
			fmt.Printf("%d ",val)
		}
		fmt.Println()
	}*/

	steps := walk(maze,point{0,0},point{len(maze)-1,len(maze[0])-1})
	
	for _,row := range steps {
		for _,val := range row {
			fmt.Printf("%3d",val)
		}
		fmt.Println()
	}
	
}
