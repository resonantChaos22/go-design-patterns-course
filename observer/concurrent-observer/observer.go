package concurrent

import (
	"fmt"
	"sync"
	"time"
)

type Observer interface {
	Notify(data any)
	GetID() string
}

type Observable struct {
	subs map[string]chan any
	mu   sync.RWMutex
}

func (o *Observable) Subscribe(x Observer) chan any {
	o.mu.Lock()
	defer o.mu.Unlock()
	ch := make(chan any, 2)
	o.subs[x.GetID()] = ch
	go func(observer Observer, dataChan <-chan any) {
		for data := range dataChan {
			observer.Notify(data)
		}
	}(x, ch)

	return ch
}

func (o *Observable) Unsubscribe(x Observer) {
	o.mu.Lock()
	defer o.mu.Unlock()
	close(o.subs[x.GetID()])
	delete(o.subs, x.GetID())
}

func (o *Observable) Fire(data any) {
	o.mu.RLock()
	subChans := make([]chan any, 0)
	for _, ch := range o.subs {
		subChans = append(subChans, ch)
	}
	o.mu.RUnlock()

	for _, subChan := range subChans {
		//	select is here so that the loop is not blocked if any channel cant take the data
		select {
		case subChan <- data:

		default:
			fmt.Printf("Warning: Subscriber channel blocked for data: %v\n", data)
		}
	}
}
func (o *Observable) UnsubscribeAll() {
	o.mu.Lock()
	defer o.mu.Unlock()
	for key, ch := range o.subs {
		close(ch)
		delete(o.subs, key)
	}
}

type Person struct {
	Observable
	Name string
}

func (p *Person) CatchACold() {
	p.Fire(p.Name)
}

func NewPerson(name string) *Person {
	return &Person{
		Observable: Observable{
			subs: make(map[string]chan any),
		},
		Name: name,
	}
}

type DoctorService struct {
	Name string
}

func (d *DoctorService) Notify(data any) {
	fmt.Printf("Doctor %s has been notified about patient %s\n", d.Name, data.(string))
}
func (d *DoctorService) GetID() string {
	return d.Name
}

func TestConcurrentObserver() {
	p := NewPerson("Damru")
	defer p.UnsubscribeAll()
	d1 := DoctorService{
		Name: "Rajesh",
	}
	d2 := DoctorService{
		Name: "Saurabh",
	}
	p.Subscribe(&d1)
	// time.Sleep(100 * time.Millisecond)
	p.Subscribe(&d2)
	// time.Sleep(100 * time.Millisecond)

	p.CatchACold()
	// time.Sleep(100 * time.Millisecond)
	p.Unsubscribe(&d1)
	p.CatchACold()
	time.Sleep(100 * time.Millisecond)
}
