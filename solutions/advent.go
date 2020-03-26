package solutions

type ProblemID int
type Solution interface {
	Solve(user string)
}

var solutionsMap map[ProblemID]Solution

func Register(id ProblemID, solution Solution) {
	solutionsMap[id] = solution
}

func Find(id ProblemID) (solution Solution, ok bool) {
	solution, ok = solutionsMap[id]
	return
}
