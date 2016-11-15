package context


import (
	"time"
	"net"
	"sync"
)

type ServerContext struct {
	_client_contexts [](*ClientContext)
	_listener        net.Listener
	_mutex           sync.Mutex
}

func NewServerContext() *ServerContext{
	return &ServerContext{_client_contexts:make([](*ClientContext),0,1000)}
}

func (s *ServerContext)SetListener(l net.Listener) {
	s._listener =l
}

func (s *ServerContext)GetListener() net.Listener{
	return s._listener
}

func (s *ServerContext)OnlineClient()  []*ClientContext {
	result :=make( []*ClientContext,0,1024)
	for _,c :=range s._client_contexts {
		if c!=nil&&c.IsClosed==false {
			result=append(result,c)
		}
	}
	return result
}

func (s *ServerContext)FindClientById(id int64) *ClientContext {
	for _,c :=range s._client_contexts {
		if c.Id==id {
			return c
		}
	}
	return nil
}

func (s *ServerContext)FindClientByName(n string) *ClientContext {
	for _,c := range s._client_contexts {
		if c.UserName==n {
			return c
		}
	}
	return nil
}

func  (s *ServerContext)ClientCount() int {
	s._mutex.Lock()
	defer s._mutex.Unlock()

	var count=0

	for _,c := range s._client_contexts {
		if c!=nil {
			count++
		}
	}

	return count
}

func (s *ServerContext)AddClientContext(c * ClientContext) {
	defer s._mutex.Unlock()

	s._mutex.Lock()

	for i,lc :=range s._client_contexts {
		if lc==nil {
			s._client_contexts[i]=c
			return
		}
	}

	s._client_contexts =append(s._client_contexts,c)

}

func (s *ServerContext)ClearClosedClient() {
	for {

		func() {
			s._mutex.Lock()
			defer s._mutex.Unlock()

			for i, c := range s._client_contexts {
				if c != nil&&c.IsClosed {
					s._client_contexts[i] = nil
				}
			}
		}()

		time.Sleep(time.Second * 3)

	}
}




