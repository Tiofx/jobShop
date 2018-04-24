package neighborhood

import (
	"jobShop/state"
	"jobShop/tabuSearch/graph_state"
	"jobShop/base"
	"jobShop/util"
	"github.com/philopon/go-toposort"
	"strconv"
	"math"
	"fmt"
)

type ByCriticalPath struct{ byJob }

func NewByCriticalPath(jobState *state.State, graph *graph_state.DisjunctiveGraph) ByCriticalPath {
	return ByCriticalPath{
		byJob: byJob{
			JobState: jobState,
			Graph:    graph,
		},
	}
}

type pathNode struct{ time, maximumEarliestTime, minimumLatestTime int }

func (pn *pathNode) slack() int       { return pn.minimumLatestTime - pn.maximumEarliestTime }
func (pn *pathNode) isCritical() bool { return pn.slack() == 0 }

func (pn *pathNode) reset() {
	pn.maximumEarliestTime = 0
	pn.minimumLatestTime = math.MaxInt64
}

type path []pathNode

func (p path) critical() (criticalNodeIndexes []int) {
	for i, pathNode := range p {
		if pathNode.isCritical() {
			criticalNodeIndexes = append(criticalNodeIndexes, i)
		}
	}
	return
}

func newPath(jobs base.Jobs) path {
	path := make(path, jobs.TotalTaskNumber())

	for index := range path {
		node := &path[index]
		node.reset()
		job, task := toJobAndTask(jobs, index)
		node.time = jobs[job][task].Time
	}

	return path
}

func (p path) updateMET(from, to int) {
	p[to].maximumEarliestTime = util.Max(p[to].maximumEarliestTime, p[from].maximumEarliestTime+p[from].time)
}

func (p path) updateMLT(from, to int) {
	p[from].minimumLatestTime = util.Min(p[from].minimumLatestTime, p[to].minimumLatestTime-p[from].time)
}

func (p path) formTasks(jobs base.Jobs, graph graph_state.DisjunctiveGraph) criticalTasks {
	res := make(criticalTasks)

	for _, index := range p.critical() {
		job, task := toJobAndTask(jobs, index)
		machine := jobs[job][task].Machine
		position := determinatePosition(jobs, graph, job, task)
		res[machine] = append(res[machine], taskPosition(position))
	}

	return res
}

type edge struct{ from, to int }

func (r ByCriticalPath) taskPosition() criticalTasks {
	edgeList := newEdgeList(r.JobState.Jobs, *r.Graph)
	edgeOrder := edgeList.toposort(r.JobState.Jobs, *r.Graph)
	path := newPath(r.JobState.Jobs)

	var reverseOrder []edge

	for _, nodeFrom := range edgeOrder {
		for _, nodeTo := range edgeList[nodeFrom] {
			path.updateMET(nodeFrom, nodeTo)
			reverseOrder = append([]edge{{nodeFrom, nodeTo}}, reverseOrder...)
		}
	}

	path[reverseOrder[0].to].minimumLatestTime = path[reverseOrder[0].to].maximumEarliestTime

	for _, edge := range reverseOrder {
		path.updateMLT(edge.from, edge.to)
	}

	return path.formTasks(r.JobState.Jobs, *r.Graph)
}

func (r ByCriticalPath) Generate() []Move {
	tasks := r.taskPosition()

	return r.generateFor(tasks)
}

// ==================================================================

func determinatePosition(jobs base.Jobs, graph graph_state.DisjunctiveGraph, job, task int) int {
	position := onTheSameMachine(jobs, job, task)
	machine := jobs[job][task].Machine
	indexOnMachine := positionInMachine(graph, int(machine), job, position)
	return indexOnMachine
}

func onTheSameMachine(jobs base.Jobs, job, taskIndex int) int {
	machine := jobs[job][taskIndex].Machine
	numberTaskOnMachine := 0
	for i, task := range jobs[job] {
		if taskIndex == i {
			return numberTaskOnMachine
		}
		if machine == task.Machine {
			numberTaskOnMachine++
		}
	}

	panic("problem with onTheSameMachine, not found task")
}

func positionInMachine(graph graph_state.DisjunctiveGraph, machine, job, position int) int {
	repeating := -1
	for i, currentJob := range graph[machine] {
		if job == int(currentJob) {
			repeating++
		}
		if repeating == position {
			return i
		}
	}

	panic("problem with positionInMachine")
}

// ==================================================================

type edgeList [][]int

func (el edgeList) String() string {
	var res string
	for nodeFrom, toList := range el {
		res += fmt.Sprintf("%v -> %v\n", nodeFrom, toList)
	}
	return res
}

func newEdgeList(jobs base.Jobs, jobsOrderOnMachine graph_state.DisjunctiveGraph) edgeList {
	graph := make(edgeList, jobs.TotalTaskNumber())

	graph.setUp(jobs)
	graph.setUpMore(jobsOrderOnMachine, jobs)

	return graph
}

func (el edgeList) toposort(jobs base.Jobs, jobsOrderOnMachine graph_state.DisjunctiveGraph) (edgeOrder []int) {
	graph := toposort.NewGraph(jobs.TotalTaskNumber())

	for i, job := range jobs {
		for j := range job {
			graph.AddNode(node(jobs, i, j))
		}
	}

	for from, toList := range el {
		for _, to := range toList {
			graph.AddEdge(strconv.Itoa(from), strconv.Itoa(to))
		}
	}

	strings, b := graph.Toposort()

	if !b {
		panic("problem with toposort")
	}

	edgeOrder = make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic("in graph non integer node")
		}

		edgeOrder[i] = num
	}

	return
}

func (g edgeList) addEdge(jobs base.Jobs, jobFrom, taskFrom, jobTo, taskTo int) {
	nodeFrom := indexOf(jobs, jobFrom, taskFrom)
	nodeTo := indexOf(jobs, jobTo, taskTo)
	g[nodeFrom] = append(g[nodeFrom], nodeTo)
}

func (g *edgeList) setUp(jobs base.Jobs) {
	for jobIndex, job := range jobs {
		for taskIndex := range job[1:] {
			g.addEdge(jobs, jobIndex, taskIndex, jobIndex, taskIndex+1)
		}
	}
}

func (g *edgeList) setUpMore(graph graph_state.DisjunctiveGraph, jobs base.Jobs) {
	indexOfNextTask := make([]int, len(jobs))

	for machine, jobSequence := range graph {
		util.FillIntsWith(indexOfNextTask, -1)
		job := jobSequence[0]
		nodeTo, exist := nextTask(int(job), indexOfNextTask[job], machine, jobs)
		if !exist {
			panic("no next node...")
		}
		indexOfNextTask[job] = nodeTo

		for i, job := 0, jobSequence[0]; i < len(jobSequence)-1; i, job = i+1, jobSequence[i+1] {
			nextJob := jobSequence[i+1]
			nodeTo, exist := nextTask(int(nextJob), indexOfNextTask[nextJob], machine, jobs)
			if !exist {
				panic("no next node...")
			}

			indexOfNextTask[nextJob] = nodeTo
			g.addEdge(jobs, int(job), indexOfNextTask[job], int(nextJob), indexOfNextTask[nextJob])
		}
		for _, job := range jobSequence {
			indexOfNextTask[job]++
		}
	}
}

func updateNextTaskFor(indexOfNextTask []int, machine int, jobs base.Jobs) {
	for jobIndex, job := range jobs {
		for taskIndex, task := range job {
			if machine == int(task.Machine) && indexOfNextTask[jobIndex] < taskIndex {
				indexOfNextTask[jobIndex] = taskIndex
				break
			}
		}
	}
}

func nextTask(jobNumber, taskNumber, machine int, jobs base.Jobs) (nextTaskNumber int, exist bool) {
	for taskIndex, task := range jobs[jobNumber] {
		if machine == int(task.Machine) && taskIndex > taskNumber {
			nextTaskNumber = taskIndex
			return nextTaskNumber, true
		}
	}

	return -1, false
}

func node(jobs base.Jobs, job int, task int) string {
	return strconv.Itoa(indexOf(jobs, job, task))
}

func indexOf(jobs base.Jobs, job int, task int) int {
	return jobs.TaskNumber(job) + task
}

func toJobAndTask(jobs base.Jobs, index int) (job, task int) {
	job = numberOfJob(jobs, index)
	task = index - jobs.TaskNumber(job)

	return
}

func numberOfJob(jobs base.Jobs, index int) (job int) {
	for jobIndex := range jobs {
		if jobs.TaskNumber(jobIndex) > index {
			return jobIndex - 1
		}
	}

	return len(jobs) - 1
}
