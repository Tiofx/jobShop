package json_test

import (
	"fmt"
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/test"
	"github.com/Tiofx/jobShop/json"
)

func Example_taskString() {
	baseTask := base.Task{1, 3}
	fmt.Println(json.Task(baseTask))

	//Output: {"machine":1,"time":3}
}

func Example_taskStringToTaskOrPanic() {
	fmt.Println(json.TaskString(`{"machine":1,"time":3}`).ToTaskOrPanic())

	//Output: {"machine":1,"time":3}
}

func Example_jobString() {
	baseJob := test.SimpleTestCase().Jobs[0]
	fmt.Println(json.Job(baseJob))

	//Output:
	//[{"machine":0,"time":3},{"machine":1,"time":2},{"machine":2,"time":2}]
}

func Example_jobsString() {
	baseJobs := test.SimpleTestCase().Jobs
	fmt.Println(json.Jobs(baseJobs))

	//Output:
	//[[{"machine":0,"time":3},{"machine":1,"time":2},{"machine":2,"time":2}],[{"machine":0,"time":2},{"machine":2,"time":1},{"machine":1,"time":4}],[{"machine":1,"time":4},{"machine":2,"time":3}]]
}
