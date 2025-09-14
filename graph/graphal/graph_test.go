package graphal

// func TestNewGraph(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	g.Print()
// }

// func TestAddEdges(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// }

// func TestHasEdge(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	fmt.Println(g.HasEdge(1, 2))
// }

// func TestRemoveEdge(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	g.RemoveEdge(1, 3)
// 	fmt.Println("")
// 	g.Print()
// }

// func TestNeighbors(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	fmt.Println("")
// 	n, err := g.Neighbors(0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(n)
// }

// func TestBFS(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	fmt.Println("")
// 	found, path, err := g.BFS(0, 4)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if found {
// 		fmt.Printf("Path: %v\n", path)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func TestDFS(t *testing.T) {
// 	g, err := NewGraph(5)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	g.AddEdge(0, 1)
// 	g.AddEdge(0, 2)
// 	g.AddEdge(0, 3)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(3, 4)
// 	g.Print()
// 	fmt.Println("")
// 	found, path, err := g.DFS(0, 4)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if found {
// 		fmt.Printf("Path: %v\n", path)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }
