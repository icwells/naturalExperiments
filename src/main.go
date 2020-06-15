// Identifies species that have recently diverged and have different cancer rates.

package main

import (
	"fmt"
	"github.com/icwells/go-tools/dataframe"
	"gopkg.in/alecthomas/kingpin.v2"
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	app       = kingpin.New("naturalExperiments", "Indentifies species that have recently diverged and have different cancer rates.")
	infile    = kingpin.Flag("infile", "Path to input cancer rates file.").Short('i').Required().String()
	malignant = kingpin.Flag("malignant", "Examine malignancy rates (examines neoplasia rate by default).").Default("false").Bool()
	max       = kingpin.Flag("max", "The maximum divergeance allowed to compare species.").Default("10.0").Float()
	min       = kingpin.Flag("min", "The minimum difference in cancer rates to report results.").Default("0.2").Float()
	outfile   = kingpin.Flag("outfile", "Name of output file.").Short('o').Default("nil").String()
	treefile  = kingpin.Flag("treefile", "Path to newick tree file.").Short('t').Required().String()
)

type cancerRate struct {
	name string
	rate float64
}

func newCancerRate(name string, rate float64) *cancerRate {
	// Returns filled struct
	var r cancerRate
	r.name = strings.Replace(name, " ", "_", 1)
	r.rate = rate
	return &r
}

type identifier struct {
	max     float64
	min     float64
	rates   []*cancerRate
	results *dataframe.Dataframe
	tree    *NewickTree
}

func newIdentifier() *identifier {
	// Returns initialized identifier struct
	id := new(identifier)
	id.max = *max
	id.min = *min
	id.results, _ = dataframe.NewDataFrame(-1)
	id.results.SetHeader([]string{"SpeciesA", "RateA", "SpeciesB", "RateB", "Difference", "Divergence(MYA)"})
	fmt.Println("\n\tReading tree file...")
	id.tree = FromFile(*treefile)
	id.setRates(*infile, *malignant)
	return id
}

func (id *identifier) setRates(infile string, mal bool) {
	// Reads cancer rates from file
	r := "NeoplasiaRate"
	if mal {
		r = "MalignancyRate"
	}
	fmt.Println("\tReading cancer rate file...")
	df, err := dataframe.FromFile(infile, -1)
	if err == nil {
		for idx := range df.Rows {
			if species, er := df.GetCell(idx, "Species"); er == nil {
				if rate, e := df.GetCellFloat(idx, r); e == nil {
					id.rates = append(id.rates, newCancerRate(species, rate))
				}
			}
		}
	}
}

func (id *identifier) greater(a, b *cancerRate) bool {
	// Returns true if difference between a and b cancer rates are greater than min
	if math.Abs(a.rate-b.rate) >= id.min {
		return true
	}
	return false
}

func (id *identifier) formatFloat(f float64) string {
	// Returns float formatted to string
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func (id *identifier) checkDistance(a, b *cancerRate) {
	// Stores results if distance is less than max
	d := id.tree.Divergence(a.name, b.name)
	if d > 0.0 && d <= id.max {
		row := []string{a.name, id.formatFloat(a.rate), b.name, id.formatFloat(b.rate), id.formatFloat(math.Abs(a.rate - b.rate)), id.formatFloat(d)}
		id.results.AddRow(row)
	}
}

func (id *identifier) identify() {
	// Compares cancer rates and determines distance between possible hits
	fmt.Println("\tIndentifying natural experiments...")
	for idx, i := range id.rates[:len(id.rates)-1] {
		for _, j := range id.rates[idx:] {
			if id.greater(i, j) {
				go id.checkDistance(i, j)
			}
		}
	}
}

func (id *identifier) writeResults() {

}

func main() {
	start := time.Now()
	kingpin.Parse()
	id := newIdentifier()
	id.identify()
	id.results.ToCSV(*outfile)
	fmt.Printf("\tFinished. Runtime: %s\n\n", time.Since(start))
}
