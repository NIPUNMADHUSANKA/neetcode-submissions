func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)

	for _, p := range prerequisites {
		course := p[0]
		prereq := p[1]
		graph[prereq] = append(graph[prereq], course)
	}

	state := make([]int, numCourses)

	var dfs func(int) bool

	dfs = func(course int) bool {
		if state[course] == 1 {
			return false // cycle found
		}

		if state[course] == 2 {
			return true // already checked
		}

		state[course] = 1

		for _, next := range graph[course] {
			if !dfs(next) {
				return false
			}
		}

		state[course] = 2
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}