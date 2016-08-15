package context


import (
	"log"
	"time"
	"net"
	"sync"
)

var ServerContext =&struct {
	clients  [](*ClientContext)
	Listener net.Listener
	mutex    sync.Mutex
}{clients:make([](*ClientContext),0,1000)}


func OnlineClient()  []*ClientContext {
	result :=make( []*ClientContext,0,1024)
	for _,c :=range ServerContext.clients {
		if c!=nil&&c.IsClosed==false {
			result=append(result,c)
		}
	}
	return result
}

func FindClientById(id int64) *ClientContext {
	for _,c :=range ServerContext.clients {
		if c.Id==id {
			return c
		}
	}
	return nil
}

func FindClientByName(n string) *ClientContext {
	for _,c := range ServerContext.clients {
		if c.UserName==n {
			return c
		}
	}
	return nil
}



func  ClientCount() int {
	ServerContext.mutex.Lock()
	defer ServerContext.mutex.Unlock()

	var count=0

	for _,c := range ServerContext.clients {
		if c!=nil {
			count++
		}
	}

	return count
}


func ServerContextAddClientContext(c * ClientContext) {
	defer ServerContext.mutex.Unlock()


	log.Println("Try to add a new client Context:",c)

	ServerContext.mutex.Lock()

	for i,lc :=range ServerContext.clients {
		if lc==nil {
			log.Println("Reuse the recoveried client context",i)
			ServerContext.clients[i]=c
			return
		}
	}

	log.Println("Add a new client Context:",c)

	ServerContext.clients =append(ServerContext.clients,c)

}

func ServerContextClearClosedClient() {
	for {
		func (){
			ServerContext.mutex.Lock()
			defer ServerContext.mutex.Unlock()

			for i,c:=range ServerContext.clients {
				if c!=nil&&c.IsClosed {
					log.Println(`Find a closed connect to clear:`,c)
					ServerContext.clients[i]=nil
				}
			}
		}()

		time.Sleep(time.Second*3)
	}
}




