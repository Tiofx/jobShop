package neighborhood

import (
	"github.com/Tiofx/jobShop/state"
	"github.com/Tiofx/jobShop/tabuSearch/graph_state"
	"github.com/Tiofx/jobShop/base"
	"github.com/Tiofx/jobShop/util"
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

type pathNode struct{ time, maximumEarliestTime, minimumLatestTime int64 }

func (pn *pathNode) slack() int64       { return pn.minimumLatestTime - pn.maximumEarliestTime }
func (pn *pathNode) isCritical() bool { return pn.slack() == 0 }

func (pn *pathNode) reset() {
	pn.maximumEarliestTime = 0
	pn.minimumLatestTime = math.MaxInt64
}

type path []pathNode

func (p path) critical() (criticalNodeIndexes []int64) {
	for i, pathNode := range p {
		if pathNode.isCritical() {
			criticalNodeIndexes = append(criticalNodeIndexes, int64(i))
		}
	}
	return
}

func newPath(jobs base.Jobs) path {
	path := make(path, jobs.TotalTaskNumber())

	for index := range path {
		node := &path[index]
		node.reset()
		job, task := toJobAndTask(jobs, int64(index))
		node.time = int64(jobs[job][task].Time)
	}

	return path
}

func (p path) updateMET(from, to int64) {
	p[to].maximumEarliestTime = int64(util.MaxInt64(p[to].maximumEarliestTime, p[from].maximumEarliestTime+p[from].time))
}

func (p path) updateMLT(from, to int64) {
	p[from].minimumLatestTime = int64(util.MinInt64(p[from].minimumLatestTime, p[to].minimumLatestTime-p[from].time))
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

type edge struct{ from, to int64 }

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

func determinatePosition(jobs base.Jobs, graph graph_state.DisjunctiveGraph, job, task int64) int64 {
	position := onTheSameMachine(jobs, job, task)
	machine := jobs[job][task].Machine
	indexOnMachine := positionInMachine(graph, int64(machine), job, position)
	return indexOnMachine
}

func onTheSameMachine(jobs base.Jobs, job, taskIndex int64) int64 {
	machine := jobs[job][taskIndex].Machine
	numberTaskOnMachine := int64(0)
	for i, task := range jobs[job] {
		if taskIndex == int64(i) {
			return numberTaskOnMachine
		}
		if machine == task.Machine {
			numberTaskOnMachine++
		}
	}

	panic("problem with onTheSameMachine, not found task")
}

func positionInMachine(graph graph_state.DisjunctiveGraph, machine, job, position int64) int64 {
	repeating := -1
	for i, currentJob := range graph[machine] {
		if job == int64(currentJob) {
			repeating++
		}
		if int64(repeating) == position {
			return int64(i)
		}
	}

	panic("problem with positionInMachine")
}

// ==================================================================

type edgeList [][]int64

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

func (el edgeList) toposort(jobs base.Jobs, jobsOrderOnMachine graph_state.DisjunctiveGraph) (edgeOrder []int64) {
	graph := toposort.NewGraph(int(jobs.TotalTaskNumber()))

	for i, job := range jobs {
		for j := range job {
			graph.AddNode(node(jobs, int64(i), int64(j)))
		}
	}

	for from, toList := range el {
		for _, to := range toList {
			graph.AddEdge(strconv.Itoa(from), strconv.Itoa(int(to)))
		}
	}

	strings, b := graph.Toposort()

	if !b {
		panic("problem with toposort")
	}

	edgeOrder = make([]int64, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic("in graph non integer node")
		}

		edgeOrder[i] = int64(num)
	}

	return
}

func (g edgeList) getIndexes(jobs base.Jobs, jobFrom, taskFrom, jobTo, taskTo int64) (nodeFrom, nodeTo int64) {
	nodeFrom, nodeTo = indexOf(jobs, jobFrom, taskFrom), indexOf(jobs, jobTo, taskTo)
	return
}

func (g edgeList) addEdge(jobs base.Jobs, jobFrom, taskFrom, jobTo, taskTo int64) {
	nodeFrom, nodeTo := g.getIndexes(jobs, jobFrom, taskFrom, jobTo, taskTo)
	g[nodeFrom] = append(g[nodeFrom], nodeTo)
}

func (g *edgeList) setUp(jobs base.Jobs) {
	for jobIndex, job := range jobs {
		for taskIndex := range job[1:] {
			g.addEdge(jobs, int64(jobIndex), int64(taskIndex), int64(jobIndex), int64(taskIndex)+1)
		}
	}
}

func (g *edgeList) setUpMore(graph graph_state.DisjunctiveGraph, jobs base.Jobs) {
	indexOfNextTask := make([]int64, len(jobs))
	for machine, jobSequence := range graph {
		util.FillIntsWith(indexOfNextTask, -1)
		job := jobSequence[0]
		nodeTo, exist := nextTask(int64(job), int64(indexOfNextTask[job]), int64(machine), jobs)
		if !exist {
			panic("no next node...")
		}
		indexOfNextTask[job] = int64(nodeTo)

		for i, job := 0, jobSequence[0]; i < len(jobSequence)-1; i, job = i+1, jobSequence[i+1] {
			nextJob := jobSequence[i+1]
			nodeTo, exist := nextTask(int64(nextJob), int64(indexOfNextTask[nextJob]), int64(machine), jobs)
			if !exist {
				panic("no next node...")
			}

			indexOfNextTask[nextJob] = int64(nodeTo)
			if nodeFrom, nodeTo := g.getIndexes(jobs, int64(job), int64(indexOfNextTask[job]), int64(nextJob), int64(indexOfNextTask[nextJob]));
				nodeFrom == nodeTo {
				continue
			}
			g.addEdge(jobs, int64(job), int64(indexOfNextTask[job]), int64(nextJob), int64(indexOfNextTask[nextJob]))
		}
		for _, job := range jobSequence {
			indexOfNextTask[job]++
		}
	}
}

func updateNextTaskFor(indexOfNextTask []int64, machine int64, jobs base.Jobs) {
	for jobIndex, job := range jobs {
		for taskIndex, task := range job {
			if machine == int64(task.Machine) && indexOfNextTask[jobIndex] < int64(taskIndex) {
				indexOfNextTask[jobIndex] = int64(taskIndex)
				break
			}
		}
	}
}

func nextTask(jobNumber, taskNumber, machine int64, jobs base.Jobs) (nextTaskNumber int64, exist bool) {
	for taskIndex, task := range jobs[jobNumber] {
		if machine == int64(task.Machine) && int64(taskIndex) > taskNumber {
			nextTaskNumber = int64(taskIndex)
			return nextTaskNumber, true
		}
	}

	return math.MaxInt64, false
}

func node(jobs base.Jobs, job int64, task int64) string {
	return strconv.Itoa(int(indexOf(jobs, job, task)))
}

func indexOf(jobs base.Jobs, job int64, task int64) int64 {
	return int64(jobs.TaskNumber(uint64(job))) + task
}

func toJobAndTask(jobs base.Jobs, index int64) (job, task int64) {
	job = numberOfJob(jobs, index)
	task = index - int64(jobs.TaskNumber(uint64(job)))

	return
}

func numberOfJob(jobs base.Jobs, index int64) (job int64) {
	for jobIndex := range jobs {
		if jobs.TaskNumber(uint64(jobIndex)) > uint64(index) {
			return int64(jobIndex - 1)
		}
	}

	return int64(len(jobs) - 1)
}
