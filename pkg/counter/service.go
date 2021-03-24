package counter

type Service interface {
    Add(v int) int
}