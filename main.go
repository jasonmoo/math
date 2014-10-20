package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
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

	m, s float64 // stddev

	min_set, max_set bool

	count float64
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

	scanner := bufio.NewScanner(os.Stdin)

	d := rune((*delimiter)[0])
	dfunc := func(c rune) bool { return c == d }
	f := *field - 1

	for scanner.Scan() {
		line := strings.FieldsFunc(scanner.Text(), dfunc)
		if len(line) < *field {
			continue
		}
		process(line[f])
	}

	if *all || *do_sum {
		fmt.Println("Sum:", strconv.FormatFloat(sum, 'f', -1, 64))
	}

	if *all || *do_avg {
		fmt.Println("Avg:", strconv.FormatFloat(sum/float64(count), 'f', -1, 64))
		fmt.Println("Stddev:", strconv.FormatFloat(math.Sqrt(s/count-1), 'f', -1, 64))
	}

	if *all || *do_max {
		fmt.Println("Max:", strconv.FormatFloat(max, 'f', -1, 64))
	}

	if *all || *do_min {
		fmt.Println("Min:", strconv.FormatFloat(min, 'f', -1, 64))
	}

}

func process(val string) {

	v, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	count++

	tm := m
	m = (v - tm) / count
	s = (v - tm) * (v - m)

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
