package main

import (
	"flag"

	"github.com/wojiaowenzhong233/graph_order_homework/sorter"
)

func main() {

	filepath := flag.String("input", "", "Graph file to sort.")
	output := flag.String("output", "", "Output file name.")
	sorter_type := flag.String("sorter", "degree", "Graph ordering algorithm, degree or visit")

	flag.Parse()

	if *filepath == "" {
		flag.Usage()
		return
	}

	g_sorter := sorter.CreateSorter(*sorter_type)
	G := sorter.CreateGraph(*filepath)

	sorter.SaveGraph(*output, g_sorter.Sort(G))
}
