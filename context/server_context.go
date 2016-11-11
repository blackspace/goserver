package context


import (
	"log"
	"time"
	"net"
	"sync"
)

type ServerContext struct {
	_clients [](*ClientContext)
	Listener net.Listener
	_mutex   sync.Mutex
}


func NewServerContext() *ServerContext{
	return &ServerContext{_clients:make([](*ClientContext),0,1000)}
}


func (s *ServerContext)OnlineClient()  []*ClientContext {
	result :=make( []*ClientContext,0,1024)
	for _,c :=range s._clients {
		if c!=nil&&c.IsClosed==false {
			result=append(result,c)
		}
	}
	return result
}

func (s *ServerContext)FindClientById(id int64) *ClientContext {
	for _,c :=range s._clients {
		if c.Id==id {
			return c
		}
	}
	return nil
}

func (s *ServerContext)FindClientByName(n string) *ClientContext {
	for _,c := range s._clients {
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

	for _,c := range s._clients {
		if c!=nil {
			count++
		}
	}

	return count
}


func (s *ServerContext)AddClientContext(c * ClientContext) {
	defer s._mutex.Unlock()


	log.Println("Try to add a new client Context:",c)

	s._mutex.Lock()

	for i,lc :=range s._clients {
		if lc==nil {
			log.Println("Reuse the recoveried client context",i)
			s._clients[i]=c
			return
		}
	}

	log.Println("Add a new client Context:",c)

	s._clients =append(s._clients,c)

}

func (s *ServerContext)ClearClosedClient() {
	for {
		func (){
			s._mutex.Lock()
			defer s._mutex.Unlock()

			for i,c:=range s._clients {
				if c!=nil&&c.IsClosed {
					log.Println(`Find a closed connect to clear:`,c)
					s._clients[i]=nil
				}
			}
		}()

		time.Sleep(time.Second*3)
	}
}




