package main

type Event struct {
    Name string
    Data interface{}
}

type Subscriber struct {
    Name string
    Events chan Event
}

type Publisher struct {
    subscribers []*Subscriber
    events chan Event
}

func (s *Subscriber) Listen() {
    for {
        select {
        case e := <-s.Events:
            fmt.Printf("Subscriber %s received event %s\n", s.Name, e.Name)
        }
    }
}

func (p *Publisher) Publish(e Event) {
    p.events <- e
}

func (p *Publisher) subscribe(s *Subscriber) {
    p.subscribers = append(p.subscribers, s)
}

func (p *Publisher) broadcast() {
    for {
        select {
        case e := <-p.events:
            for _, subscriber := range p.subscribers {
                subscriber.Events <- e
            }
        }
    }
}

func main() {
    publisher := &Publisher{
        events: make(chan Event),
    }

    subscriber1 := &Subscriber{
        Name: "Subscriber 1",
        Events: make(chan Event),
    }

    subscriber2 := &Subscriber{
        Name: "Subscriber 2",
        Events: make(chan Event),
    }

    publisher.subscribe(subscriber1)
    publisher.subscribe(subscriber2)

    go subscriber1.Listen()
    go subscriber2.Listen()
    go publisher.broadcast()

    publisher.Publish(Event{Name: "Event 1", Data: "Hello, World"})
}
