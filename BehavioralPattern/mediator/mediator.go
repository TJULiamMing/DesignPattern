package mediator

type Mediator interface {
	CanArrive(Train) bool
	NotifyDeparture()
}

type TrainMediator struct {
	IsPlatformFree bool
	TrainQueue     []Train
	Name           string
}

func (t *TrainMediator) CanArrive(train Train) bool {
	if t.IsPlatformFree {
		t.IsPlatformFree = false
		return true
	}

	t.TrainQueue = append(t.TrainQueue, train)
	return false
}

func (t *TrainMediator) NotifyDeparture() {
	if !t.IsPlatformFree {
		t.IsPlatformFree = true
	}

	if len(t.TrainQueue) > 0 {
		firstTrain := t.TrainQueue[0]
		t.TrainQueue = t.TrainQueue[1:]
		firstTrain.PermitArrival()
	}
}

func NewTrainMediator() *TrainMediator {
	return &TrainMediator{
		Name:           "Mediator",
		IsPlatformFree: true,
	}
}
