// TIL: This was the first Hard problem which required major recursive brute
// force to solve. The recursive func isn't very clean but it works and took
// several tries to especially have it print well.
package day9

import (
	"fmt"
	"strings"

	"valbaca.com/advent2015/utils"
)

var minDist, maxDist int

func BothParts(in string) string {
	minDist = 100000
	maxDist = 0
	chart := &Chart{}
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		ReadLine(line, chart)
	}
	printAllPaths(chart)
	return fmt.Sprintf("min=%v,max=%v", minDist, maxDist)
}

type Edge struct {
	from string
	to   string
	dist int
}

type Location struct {
	name  string
	edges []Edge
}

type Chart map[string]Location

type Visited []Location

func ReadLine(s string, c *Chart) {
	edge := SplitLine(s)
	c.addBiEdge(edge)
}

func SplitLine(s string) Edge {
	sp := strings.Split(s, " ")
	// London to Dublin = 464
	// 0      1  2      3 4
	if len(sp) < 5 {
		panic("not enough words to SplitLine")
	}
	from, to, dist := sp[0], sp[2], utils.AtoI(sp[4])
	return Edge{from, to, dist}
}

func (c Chart) addBiEdge(e Edge) {
	c.addEdge(e)
	c.addEdge(Edge{e.to, e.from, e.dist})
}

func (c Chart) addEdge(e Edge) {
	from, ok := c[e.from]
	if !ok {
		from = Location{e.from, nil}
	}
	(&from).addEdge(e)
	c[e.from] = from
}

func (loc *Location) addEdge(e Edge) {
	loc.edges = append(loc.edges, e)
}

func (loc *Location) calcDist(to Location) int {
	for _, edge := range loc.edges {
		if edge.to == to.name {
			return edge.dist
		}
	}
	panic("missing path from " + loc.name + " to " + to.name)
}

func printAllPaths(c *Chart) {
	for _, loc := range *c {
		// fmt.Print(buildPaths(loc, Visited{}, c, 0)) // Debug
		buildPaths(loc, Visited{}, c, 0)
	}
}

func buildPaths(loc Location, v Visited, c *Chart, dist int) string {
	visited := append(v[:0:0], v...)
	visited = append(visited, loc)
	toVisit := []Location{}
	for _, to := range *c {
		if !visited.contains(to) {
			toVisit = append(toVisit, to)
		}
	}
	if len(toVisit) == 0 {
		// print all visited
		out := ""
		for i, vl := range visited {
			if i == 0 {
				out += vl.name
			} else {
				out += " -> " + vl.name
			}
		}
		out += fmt.Sprintf(" = %v\n", dist)
		if dist < minDist {
			minDist = dist
		}
		if dist > maxDist {
			maxDist = dist
		}
		return out
	}
	paths := make([]string, len(toVisit))
	for _, to := range toVisit {
		path := buildPaths(to, visited, c, dist+loc.calcDist(to))
		paths = append(paths, path)
	}
	return strings.Join(paths, "")
}

func (v Visited) contains(target Location) bool {
	for _, loc := range v {
		if loc.name == target.name {
			return true
		}
	}
	return false
}
