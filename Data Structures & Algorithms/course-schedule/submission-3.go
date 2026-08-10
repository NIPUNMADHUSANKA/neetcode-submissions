func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)

	for _, p := range prerequisites{
		course := p[0]
		prerequest := p[1]
		g[prerequest] = append(g[prerequest], course)
	}

	state := make([]int, numCourses)

	var dfs func(course int)bool

	dfs = func(course int)bool{
		if state[course] == 1{
			return false
		}
		if state[course] == 2{
			return true
		}
		state[course] = 1

		for _, neightbour := range g[course]{
			if !dfs(neightbour){
				return false
			}
		}
		state[course] = 2
		return true
	}


	for i := 0; i < numCourses; i++ {
		if !dfs(i){
			return false
		}
	}
	return true
}