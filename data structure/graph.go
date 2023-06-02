package main
import (
	"fmt"
)
type Graph struct {
	Vertices []*Vertix
}
type Vertix struct {
	Key int
	Adjecent []*Vertix
}
func (g Graph) Println(){
	for _,v := range g.Vertices {
		fmt.Print(*v)
	}
}
func (g *Graph) AppendVertix(key int){
	g.Vertices = append(g.Vertices,&Vertix{
		Key:key,
	})
}
func main(){
	var graph Graph
	graph.AppendVertix(4)
	graph.AppendVertix(5)
	graph.Println()
}