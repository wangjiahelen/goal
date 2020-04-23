package leetcode

//岛屿数量，只能横向和竖向，1为岛，0为水
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func numIslands(grid [][]byte) int {
	//可以使用visited数组用来保存已经访问过的位置，本题可以直接将其置为0表示已经访问
	nr := len(grid)
	if nr == 0 {
		return 0
	}
	nc := len(grid[0])
	var count int
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(grid, r, c)
				//bfs(grid, r, c)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r, c int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] != '1' {
		return
	}
	grid[r][c] = '0'
	for i := 0; i < 4; i++ {
		dfs(grid, r+dx[i], c+dy[i])
	}
}

func bfs(grid [][]byte, r, c int) {
	queue := make([]int, 0)
	queue = append(queue, r, c)
	grid[r][c] = '0'
	for len(queue) != 0 {
		i, j := queue[0], queue[1]
		queue = queue[2:]
		for k := 0; k < 4; k++ {
			tmp_x, tmp_y := i+dx[k], j+dy[k]
			if tmp_x >= 0 && tmp_x < len(grid) && tmp_y >= 0 && tmp_y < len(grid[0]) && grid[tmp_x][tmp_y] == '1' {
				bfs(grid, tmp_x, tmp_y)
			}
		}
	}
}
