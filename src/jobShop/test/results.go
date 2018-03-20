package test

import "time"
import (
	"github.com/montanaflynn/stats"
	"log"
	"fmt"
)

type Results []Result

func (r Results) String() string {
	outStr := ""
	outStr += fmt.Sprintf("== Total test number: %v\n", len(r))
	for i, result := range r {
		outStr += "==============================\n"
		outStr += fmt.Sprintf("========== Test №%v ===========\n", i+1)
		outStr += "==============================\n"
		outStr += result.String()
	}

	outStr += "==============================\n"
	outStr += "======== Statistics ==========\n"
	outStr += "==============================\n"
	outStr += r.resultErrorStrings()
	outStr += r.timesStrings()

	return outStr
}

func (r Results) resultErrorStrings() string {
	var outStr string
	mean, min, max := r.ResultError()

	outStr += "-- Result error statistics:\n"
	outStr += fmt.Sprintf("avg.:  %.2f%%\n", mean*100)
	outStr += fmt.Sprintf("best:  %.2f%%\n", min*100)
	outStr += fmt.Sprintf("worst: %.2f%%\n", max*100)
	outStr += "---------------------------\n"

	return outStr
}

func (r Results) timesStrings() string {
	var outStr string
	total, mean, min, max := r.Time()

	outStr += "-- Time statistics:\n"
	outStr += fmt.Sprintf("total: %v\n", total)
	outStr += fmt.Sprintf("mean:  %v\n", mean)
	outStr += fmt.Sprintf("min:   %v\n", min)
	outStr += fmt.Sprintf("max:   %v\n", max)
	outStr += "---------------------------\n"

	return outStr
}

func (r Results) Time() (total, mean, min, max time.Duration) {
	times := make([]time.Duration, 0, len(r))
	for _, res := range r {
		times = append(times, res.Elapsed)
	}
	d := stats.LoadRawData(times)

	rawSum, err1 := d.Sum()
	rawMean, err2 := d.Mean()
	rawMin, err3 := d.Min()
	rawMax, err4 := d.Max()
	logErrors(err1, err2, err3, err4)

	total = time.Duration(rawSum)
	mean = time.Duration(rawMean)
	min = time.Duration(rawMin)
	max = time.Duration(rawMax)
	return
}

func logErrors(errors ...error) {
	for _, e := range errors {
		if e != nil {
			log.Println(e)
		}
	}
}

func (r Results) ResultError() (mean, min, max float64) {
	resultErrors := make([]float64, 0, len(r))
	for _, res := range r {
		resultErrors = append(resultErrors, res.HowMuchWorse())
	}

	mean, err1 := stats.Mean(resultErrors)
	min, err2 := stats.Min(resultErrors)
	max, err3 := stats.Max(resultErrors)
	logErrors(err1, err2, err3)

	return
}
