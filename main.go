package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	all = flag.Bool("all", false, "output all metrics")

	do_sum = flag.Bool("sum", false, "output sum of values")
	do_avg = flag.Bool("avg", false, "output avg of values")
	do_max = flag.Bool("max", false, "output max value")
	do_min = flag.Bool("min", false, "output min value")

	delimiter = flag.String("d", ",", "delimiter")
	field     = flag.Int("f", 1, "field to parse")

	sum, max, min float64

	min_set, max_set bool

	count uint64
)

func init() {
	flag.Parse()
	if !*all && !*do_sum && !*do_avg && !*do_max && !*do_min {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {

	cr := csv.NewReader(os.Stdin)
	cr.Comma = rune((*delimiter)[0])
	f := *field - 1

	for {
		line, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if len(line) < *field {
			continue
		}
		process(line[f])
	}

	if *all || *do_sum {
		fmt.Printf("Sum: %g\n", sum)
	}

	if *all || *do_avg {
		fmt.Printf("Avg: %g\n", sum/float64(count))
	}

	if *all || *do_max {
		fmt.Printf("Max: %g\n", max)
	}

	if *all || *do_min {
		fmt.Printf("Min: %g\n", min)
	}

}

func process(val string) {

	v, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	count++

	sum += v

	if v > max || !max_set {
		max = v
		max_set = true
	}

	if v < min || !min_set {
		min = v
		min_set = true
	}

}
