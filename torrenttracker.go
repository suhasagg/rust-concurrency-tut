type Tracker struct {
	peers map[string]Peer
	mu    sync.Mutex
}

type Peer struct {
	IP   string
	Port int
}

func (t *Tracker) AddPeer(peer Peer) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.peers[peer.IP] = peer
}

func (t *Tracker) GetPeers() []Peer {
	t.mu.Lock()
	defer t.mu.Unlock()
	peers := make([]Peer, 0, len(t.peers))
	for _, peer := range t.peers {
		peers = append(peers, peer)
	}
	return peers
}

func (t *Tracker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/announce":
		peer := Peer{
			IP:   r.RemoteAddr,
			Port: int(r.FormValue("port")),
		}
		t.AddPeer(peer)
		w.Write([]byte("OK"))
	case "/peers":
		peers := t.GetPeers()
		data, _ := json.Marshal(peers)
		w.Write(data)
	}
}

func main() {
	tracker := &Tracker{
		peers: make(map[string]Peer),
	}
	http.Handle("/", tracker)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
