package context


import (
	"log"
	"time"
	"net"
	"sync"
)

type ServerContext struct {
	clients  [](*ClientContext)
	Listener net.Listener
	mutex    sync.Mutex
}


func NewServerContext() *ServerContext{
	return &ServerContext{clients:make([](*ClientContext),0,1000)}
}


func (s *ServerContext)OnlineClient()  []*ClientContext {
	result :=make( []*ClientContext,0,1024)
	for _,c :=range s.clients {
		if c!=nil&&c.IsClosed==false {
			result=append(result,c)
		}
	}
	return result
}

func (s *ServerContext)FindClientById(id int64) *ClientContext {
	for _,c :=range s.clients {
		if c.Id==id {
			return c
		}
	}
	return nil
}

func (s *ServerContext)FindClientByName(n string) *ClientContext {
	for _,c := range s.clients {
		if c.UserName==n {
			return c
		}
	}
	return nil
}



func  (s *ServerContext)ClientCount() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var count=0

	for _,c := range s.clients {
		if c!=nil {
			count++
		}
	}

	return count
}


func (s *ServerContext)AddClientContext(c * ClientContext) {
	defer s.mutex.Unlock()


	log.Println("Try to add a new client Context:",c)

	s.mutex.Lock()

	for i,lc :=range s.clients {
		if lc==nil {
			log.Println("Reuse the recoveried client context",i)
			s.clients[i]=c
			return
		}
	}

	log.Println("Add a new client Context:",c)

	s.clients =append(s.clients,c)

}

func (s *ServerContext)ClearClosedClient() {
	for {
		func (){
			s.mutex.Lock()
			defer s.mutex.Unlock()

			for i,c:=range s.clients {
				if c!=nil&&c.IsClosed {
					log.Println(`Find a closed connect to clear:`,c)
					s.clients[i]=nil
				}
			}
		}()

		time.Sleep(time.Second*3)
	}
}




