# CPU Scheduling Algorithms

This repository contains implementations of essential **CPU scheduling algorithms**, which are fundamental to operating systems. These algorithms determine the sequence in which processes are executed by the CPU, aiming to optimize performance and ensure efficient resource utilization.

## Implemented Algorithms

1. **First-Come, First-Served (FCFS)**
   - Non-preemptive scheduling based on process arrival time.

2. **Shortest Job Next (SJN)**
   - Non-preemptive scheduling that prioritizes processes with the shortest execution time.

3. **Priority Scheduling**
   - Preemptive and non-preemptive algorithms based on process priority.

4. **Round Robin (RR)**
   - Time-sharing mechanism that allocates CPU time in fixed intervals (time quantum).

5. **Multilevel Queue Scheduling**
   - Segregates processes into priority-based queues for scheduling.

6. **Multilevel Feedback Queue Scheduling**
   - Dynamically adjusts process priorities across multiple queues for optimal scheduling.

## Features

- Written in **[Go Programming Language]**.
- Modular, clean, and well-documented code.
- Example inputs and outputs for testing and understanding algorithms.
- Performance metrics:
  - Average waiting time
  - Turnaround time
  - CPU utilization
- **Optional Visualization**: Gantt chart generation for better understanding.

## Applications

- **Educational Resource**: Ideal for students and developers learning operating system concepts.
- **Practical Implementation**: Suitable for projects requiring efficient process scheduling.
- **Performance Tuning**: Experiment with scheduling strategies to improve system efficiency.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/username/CPU-Scheduling-Algorithms.git
```
2. Navigate to the repository directory:
```bash
    cd CPU-Scheduling-Algorithms
```
Follow the instructions in the README.md to execute individual algorithms.

License

This project is licensed. See the LICENSE file for more details.
