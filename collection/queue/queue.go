package queue

type Queue struct {
	data []interface{}
}

func NewQueue(size int) *Queue {
	queue := new(Queue)
	queue.data = make([]interface{}, 0, size)
	return queue
}

func (q *Queue) Empty() bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}

func (q *Queue) Add(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Peek() interface{} {
	if q.Empty() {
		return nil
	}
	return q.data[0]
}

func (q *Queue) Poll() interface{} {
	if q.Empty() {
		return nil
	}
	result := q.data[0]
	q.data = q.data[1:]
	return result
}
