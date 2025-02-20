package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const Grid int = 20

type Coor struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type RequestMethod struct {
	Start Coor `json:"start"`
	End   Coor `json:"end"`
}

var directions = []Coor{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func IsValid(x, y int, isVisited [][]bool) bool {
	return x >= 0 && y >= 0 && x < Grid && y < Grid && !isVisited[x][y]
}

func distance(x, y, endX, endY int) int {
	return abs(x-endX) + abs(y-endY)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(x, y, endX, endY int, path *[]Coor, isVisited [][]bool) {
	if x == endX && y == endY {
		return
	}

	isVisited[x][y] = true

	for _, d := range directions {
		newX := x + d.X
		newY := y + d.Y

		if IsValid(newX, newY, isVisited) {
			if distance(x, y, endX, endY) > distance(newX, newY, endX, endY) {

				*path = append(*path, Coor{newX, newY})
				dfs(newX, newY, endX, endY, path, isVisited)
				break
			}
		}
	}

}

func findPath(start, end Coor) []Coor {

	isVisited := make([][]bool, Grid)

	for i := range isVisited {
		isVisited[i] = make([]bool, Grid)
	}

	path := []Coor{start}

	dfs(start.X, start.Y, end.X, end.Y, &path, isVisited)

	return path
}

func handleFindPathDFS(c *gin.Context) {
	var request RequestMethod

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	fmt.Println("start", request.Start, "end", request.End)

	path := findPath(request.Start, request.End)

	c.JSON(http.StatusOK, gin.H{"Path": path})
}

// bfs approach just for checking the paths

func findPathBFS(start, end Coor) []Coor {

	isVisited := make([][]bool, Grid)

	for i := range isVisited {
		isVisited[i] = make([]bool, Grid)
	}

	queue := []struct {
		coord Coor
		path  []Coor
	}{{start, []Coor{start}}}

	isVisited[start.X][start.Y] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.coord.X == end.X && curr.coord.Y == end.Y {
			return curr.path
		}

		for _, d := range directions {
			newX := curr.coord.X + d.X
			newY := curr.coord.Y + d.Y

			if IsValid(newX, newY, isVisited) {
				isVisited[newX][newY] = true

				newPath := append([]Coor{}, curr.path...)
				newPath = append(newPath, Coor{newX, newY})

				queue = append(queue, struct {
					coord Coor
					path  []Coor
				}{coord: Coor{newX, newY}, path: newPath})
			}
		}
	}

	return []Coor{}
}

func handleFindPathBFS(c *gin.Context) {
	var request RequestMethod

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	fmt.Println("start", request.Start, "end", request.End)

	path := findPathBFS(request.Start, request.End)

	c.JSON(http.StatusOK, gin.H{"Path": path})
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Origin"},
	}))

	router.POST("/find-path", handleFindPathDFS)

	router.POST("/find-path-BFS", handleFindPathBFS)

	router.Run(":4000")
}
