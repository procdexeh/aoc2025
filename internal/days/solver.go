package days

type Solver interface {
	Part1(input string) any
	Part2(input string) any
}

func GetSolver(day int) Solver {
	solvers := map[int]Solver{
		1: &Day01{},
		2: &Day02{},
		3: &Day03{},
	}
	return solvers[day]
}
