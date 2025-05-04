package components

import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/zMoooooritz/go-let-observer/pkg/ui/shared"
	"github.com/zMoooooritz/go-let-observer/pkg/util"
)

const (
	NOTIFICATION_DURATION = 3 * time.Second
	NOTIFICATION_COUNT    = 5
)

type Notification struct {
	Message   string
	Timestamp time.Time
}

type NotificationManager struct {
	messages []Notification
	mutex    sync.Mutex
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		messages: []Notification{},
	}
}

func (nm *NotificationManager) Push(message string) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	nm.messages = append(nm.messages, Notification{
		Message:   message,
		Timestamp: time.Now(),
	})

	if len(nm.messages) > NOTIFICATION_COUNT {
		nm.messages = nm.messages[len(nm.messages)-NOTIFICATION_COUNT:]
	}
}

func (nm *NotificationManager) Update() {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	now := time.Now()
	filtered := nm.messages[:0]
	for _, msg := range nm.messages {
		if now.Sub(msg.Timestamp) < NOTIFICATION_DURATION {
			filtered = append(filtered, msg)
		}
	}
	nm.messages = filtered
}

func (nm *NotificationManager) Draw(screen *ebiten.Image) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	if len(nm.messages) == 0 {
		return
	}

	height := 20*len(nm.messages) + 15

	util.DrawScaledRect(screen, 0, 1000-height, 400, height, shared.CLR_OVERLAY)

	x, y := 10, screen.Bounds().Dy()-10
	for i := len(nm.messages) - 1; i >= 0; i-- {
		msg := nm.messages[i]
		util.DrawText(screen, msg.Message, x, y, shared.CLR_WHITE, util.Font.Normal)
		y -= 20
	}
}
