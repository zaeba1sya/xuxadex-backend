package activity

import (
	"context"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/sockets"
)

type ActivityMonitor struct {
	socketPool *sockets.SocketsPool
}

type activity struct {
	ID        string
	StartTime time.Time
}

var (
	activities = []activity{
		{ID: "07374340-bab9-4d4b-8dee-a54f0149757e", StartTime: time.Now().Add(10 * time.Second)},
		{ID: "c39b3801-592c-43a6-a9e8-87db5ba6f760", StartTime: time.Now().Add(20 * time.Second)},
		{ID: "d98ad8ec-2740-4b7b-bac5-8e676ec47998", StartTime: time.Now().Add(30 * time.Second)},
		{ID: "a17e2e97-1f9f-4edc-8dd8-c69bd2f8407a", StartTime: time.Now().Add(40 * time.Second)},
		{ID: "42b7ea08-9988-495a-a656-6fa9e35dc55a", StartTime: time.Now().Add(50 * time.Second)},
	}
)

func NewActivityMonitor(socketPool *sockets.SocketsPool) *ActivityMonitor {
	return &ActivityMonitor{
		socketPool: socketPool,
	}
}

func (m *ActivityMonitor) monitor(activityID string, ch chan<- string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			ch <- activityID
			return
		default:
		}
	}
}

func (m *ActivityMonitor) StartMonitoring() {
	startedActivity := make(chan string, 10)

	for _, activity := range activities {
		go m.monitor(activity.ID, startedActivity, activity.StartTime.Sub(time.Now()))
	}

	for activityID := range startedActivity {
		m.socketPool.SendToAll(activityID, "Started")
	}

}

func (m *ActivityMonitor) Release() {

}
