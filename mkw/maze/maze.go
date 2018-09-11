package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	//读一行
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// i 方向是否越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// j 方向是否越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

//上左下右
var dirs = [4]point{
	{-1, 0}, //上
	{0, -1}, //左
	{1, 0},  //下
	{0, 1},  //右
}

func walk(maze [][]int, start, end point) [][]int {
	//步数
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//队列 点
	//刚开始就一个点
	Q := []point{start}
	for len(Q) > 0 {
		//当前点
		cur := Q[0]
		//队列数据
		Q = Q[1:]
		//终点退出
		if cur == end {
			break
		}

		for _, dir := range dirs {
			//下一点
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}
func main() {
	maze := readMaze("mkw/maze/maze.in")

	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
