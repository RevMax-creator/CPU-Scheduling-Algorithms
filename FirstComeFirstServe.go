package main

import (
	"fmt"
)

// Process struct to hold process details
type Process struct {
	ID             int
	ArrivalTime    int
	BurstTime      int
	CompletionTime int
	TurnaroundTime int
	WaitingTime    int
}

// FCFS performs the First-Come, First-Serve Scheduling algorithm
func FCFS(processes []Process) (results []Process, avgTurnaroundTime, avgWaitingTime float64) {
	// Sort processes by ArrivalTime
	for i := 0; i < len(processes); i++ {
		for j := i + 1; j < len(processes); j++ {
			if processes[i].ArrivalTime > processes[j].ArrivalTime {
				processes[i], processes[j] = processes[j], processes[i]
			}
		}
	}

	n := len(processes)
	var totalTurnaroundTime, totalWaitingTime int

	// Calculate completion time, turnaround time, and waiting time
	for i := 0; i < n; i++ {
		if i == 0 {
			processes[i].CompletionTime = processes[i].ArrivalTime + processes[i].BurstTime
		} else {
			if processes[i].ArrivalTime > processes[i-1].CompletionTime {
				processes[i].CompletionTime = processes[i].ArrivalTime + processes[i].BurstTime
			} else {
				processes[i].CompletionTime = processes[i-1].CompletionTime + processes[i].BurstTime
			}
		}
		processes[i].TurnaroundTime = processes[i].CompletionTime - processes[i].ArrivalTime
		processes[i].WaitingTime = processes[i].TurnaroundTime - processes[i].BurstTime

		totalTurnaroundTime += processes[i].TurnaroundTime
		totalWaitingTime += processes[i].WaitingTime
	}

	// Calculate average turnaround and waiting times
	avgTurnaroundTime = float64(totalTurnaroundTime) / float64(n)
	avgWaitingTime = float64(totalWaitingTime) / float64(n)

	return processes, avgTurnaroundTime, avgWaitingTime
}

func main() {
	// Example process list
	processList := []Process{
		{ID: 1, ArrivalTime: 0, BurstTime: 4},
		{ID: 2, ArrivalTime: 1, BurstTime: 3},
		{ID: 3, ArrivalTime: 2, BurstTime: 1},
		{ID: 4, ArrivalTime: 3, BurstTime: 2},
	}

	results, avgTurnaroundTime, avgWaitingTime := FCFS(processList)

	fmt.Println("FCFS Scheduling Results:")
	for _, process := range results {
		fmt.Printf("Process %d: Arrival Time = %d, Burst Time = %d, Completion Time = %d, Turnaround Time = %d, Waiting Time = %d\n",
			process.ID, process.ArrivalTime, process.BurstTime, process.CompletionTime, process.TurnaroundTime, process.WaitingTime)
	}

	fmt.Printf("Average Turnaround Time: %.2f\n", avgTurnaroundTime)
	fmt.Printf("Average Waiting Time: %.2f\n", avgWaitingTime)
}