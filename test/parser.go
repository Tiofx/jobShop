package test

import (
	"github.com/Tiofx/jobShop/base"
	"os"
	"bufio"
	"log"
	"strconv"
	"strings"
)

type testParser struct {
	words []string
	*bufio.Scanner
}

func (tp *testParser) next() string {
	tp.Scan()
	tp.words = strings.Fields(tp.Text())
	return tp.Text()
}

func (tp *testParser) nextInt() (uint64, error) {
	if len(tp.words) == 0 {
		tp.next()
	}
	word := tp.words[0]
	tp.words = tp.words[1:]

	i, e := strconv.Atoi(word)
	return uint64(i), e
}

func newTestParser(filename string) *testParser {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	skipLine(scanner)

	return &testParser{Scanner: scanner}
}

func skipLine(scanner *bufio.Scanner) {
	scanner.Scan()
}

func (scanner *testParser) parseAllData() (jobsNumber, taskNumbers, optimum uint64, jobs base.Jobs) {
	jobsNumber, taskNumbers, optimum = scanner.parseMainData()
	jobs = scanner.parseJobs(jobsNumber, taskNumbers)

	return
}
func (scanner *testParser) parseMainData() (jobsNumber, tasksNumber, optimum uint64) {
	var mainValues []uint64

	for i := uint64(0); i < 3; i++ {
		if value, err := scanner.nextInt(); err == nil {
			mainValues = append(mainValues, value)
		} else {
			log.Panic(err)
		}
	}

	return mainValues[0], mainValues[1], mainValues[2]
}

func (scanner *testParser) parseJobs(jobNumber, taskNumber uint64) base.Jobs {
	jobs := make([]base.Job, jobNumber)

	for i := uint64(0); i < jobNumber; i++ {
		jobs[i] = scanner.parseJob(taskNumber)
	}

	return jobs
}

func (scanner *testParser) parseJob(taskNumber uint64) base.Job {
	job := make([]base.Task, taskNumber)

	for i := uint64(0); i < taskNumber; i++ {
		job[i] = scanner.parseTask()
	}

	return job
}

func (scanner *testParser) parseTask() base.Task {
	machine, err1 := scanner.nextInt()
	time, err2 := scanner.nextInt()

	if err1 != nil || err2 != nil {
		log.Panic(err1, err2)
	}

	return base.Task{
		Machine: base.Machine(machine),
		Time:    time,
	}
}
