package stats_test

import (
	"fmt"
	"github.com/Tiofx/jobShop/stats"
	"github.com/Tiofx/jobShop/test"
	"time"
)

func Example_parallelGettingImprovementsStats() {
	solver := stats.NewSolverWithStatistic(test.GeneralTestCase().Jobs, 1000, 222, 11)

	go func() {
		for newImprovement := range solver.GetImprovementStatsChannel() {
			fmt.Println(newImprovement)
		}
	}()

	solver.FindSolution()

	// Example of output:
	//On iteration number [0] scheduler's makespan was improved to [1846]. From the start passed [58.83ms]
	//On iteration number [1] scheduler's makespan was improved to [1753]. From the start passed [61.033ms]
	//On iteration number [2] scheduler's makespan was improved to [1690]. From the start passed [62.603ms]
	//On iteration number [3] scheduler's makespan was improved to [1648]. From the start passed [64.222ms]
	//On iteration number [4] scheduler's makespan was improved to [1630]. From the start passed [65.988ms]
	//On iteration number [5] scheduler's makespan was improved to [1571]. From the start passed [67.807ms]
	//On iteration number [6] scheduler's makespan was improved to [1566]. From the start passed [69.171ms]
	//On iteration number [7] scheduler's makespan was improved to [1543]. From the start passed [70.615ms]
	//On iteration number [8] scheduler's makespan was improved to [1512]. From the start passed [71.985ms]
	//On iteration number [9] scheduler's makespan was improved to [1507]. From the start passed [74.061ms]
	//On iteration number [10] scheduler's makespan was improved to [1486]. From the start passed [75.364ms]
	//On iteration number [11] scheduler's makespan was improved to [1471]. From the start passed [76.672ms]
	//On iteration number [12] scheduler's makespan was improved to [1439]. From the start passed [78.567ms]
	//On iteration number [13] scheduler's makespan was improved to [1424]. From the start passed [80.331ms]
	//On iteration number [14] scheduler's makespan was improved to [1421]. From the start passed [81.942ms]
	//On iteration number [15] scheduler's makespan was improved to [1365]. From the start passed [83.311ms]
	//On iteration number [26] scheduler's makespan was improved to [1363]. From the start passed [92.81ms]
	//On iteration number [27] scheduler's makespan was improved to [1362]. From the start passed [94.733ms]
	//On iteration number [28] scheduler's makespan was improved to [1341]. From the start passed [96.128ms]
	//On iteration number [29] scheduler's makespan was improved to [1319]. From the start passed [97.419ms]
	//On iteration number [30] scheduler's makespan was improved to [1290]. From the start passed [98.701ms]
}

func Example_waitingForFinishProcessingOnStatsStream() {
	solver := stats.NewSolverWithStatistic(test.GeneralTestCase().Jobs, 1000, 222, 11)
	statsProcessing := func(stats.Snapshot) { time.Sleep(1 * time.Second) }

	go func() {
		defer solver.FinishProcessingStats()

		for newImprovement := range solver.GetImprovementStatsChannel() {
			statsProcessing(newImprovement)
			fmt.Println(newImprovement)
		}
	}()

	solver.FindSolution()
	solver.WaitForProcessingStats()

	// Example of output:
	//On iteration number [0] scheduler's makespan was improved to [1846]. From the start passed [58.83ms]
	//On iteration number [1] scheduler's makespan was improved to [1753]. From the start passed [61.033ms]
	//On iteration number [2] scheduler's makespan was improved to [1690]. From the start passed [62.603ms]
	//On iteration number [3] scheduler's makespan was improved to [1648]. From the start passed [64.222ms]
	//On iteration number [4] scheduler's makespan was improved to [1630]. From the start passed [65.988ms]
	//On iteration number [5] scheduler's makespan was improved to [1571]. From the start passed [67.807ms]
	//On iteration number [6] scheduler's makespan was improved to [1566]. From the start passed [69.171ms]
	//On iteration number [7] scheduler's makespan was improved to [1543]. From the start passed [70.615ms]
	//On iteration number [8] scheduler's makespan was improved to [1512]. From the start passed [71.985ms]
	//On iteration number [9] scheduler's makespan was improved to [1507]. From the start passed [74.061ms]
	//On iteration number [10] scheduler's makespan was improved to [1486]. From the start passed [75.364ms]
	//On iteration number [11] scheduler's makespan was improved to [1471]. From the start passed [76.672ms]
	//On iteration number [12] scheduler's makespan was improved to [1439]. From the start passed [78.567ms]
	//On iteration number [13] scheduler's makespan was improved to [1424]. From the start passed [80.331ms]
	//On iteration number [14] scheduler's makespan was improved to [1421]. From the start passed [81.942ms]
	//On iteration number [15] scheduler's makespan was improved to [1365]. From the start passed [83.311ms]
	//On iteration number [26] scheduler's makespan was improved to [1363]. From the start passed [92.81ms]
	//On iteration number [27] scheduler's makespan was improved to [1362]. From the start passed [94.733ms]
	//On iteration number [28] scheduler's makespan was improved to [1341]. From the start passed [96.128ms]
	//On iteration number [29] scheduler's makespan was improved to [1319]. From the start passed [97.419ms]
	//On iteration number [30] scheduler's makespan was improved to [1290]. From the start passed [98.701ms]
}
