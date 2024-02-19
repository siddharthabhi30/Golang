package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

//
// Common RPC request/reply definitions
//

type PutArgs struct {
	Key   string
	Value string
}

type PutReply struct {
	Value string
}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Value string
}

//
// Client
//

func connect() *rpc.Client {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func get(key string) string {
	client := connect()
	args := GetArgs{"subject"}
	reply := GetReply{}
	err := client.Call("KV.Get", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	client.Close()
	return reply.Value
}

func put(key string, val string) {
	client := connect()
	args := PutArgs{key, val}
	reply := PutReply{}
	err := client.Call("KV.Put", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Printf("received value of put call from server %s\n", reply.Value)
	client.Close()
}

func append(key string, val string) {
	client := connect()
	args := PutArgs{key, val}
	reply := PutReply{}
	err := client.Call("KV.Append", &args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Printf("received value of append call from server %s\n", reply.Value)
	client.Close()
}

//
// Server
//

type KV struct {
	mu   sync.Mutex
	data map[string]string
}

func server() {
	kv := new(KV)
	kv.data = map[string]string{}
	rpcs := rpc.NewServer()
	rpcs.Register(kv)
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				fmt.Println("reached at break")
				break
			}
		}
		l.Close()
	}()

	fmt.Println("this ended")
}

func (kv *KV) Get(args *GetArgs, reply *GetReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	reply.Value = kv.data[args.Key]

	return nil
}

func (kv *KV) Put(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if oldVal, ok := kv.data[args.Key]; ok {
		reply.Value = oldVal
	} else {
		reply.Value = ""
	}
	kv.data[args.Key] = args.Value
	fmt.Printf("received put call from client %s\n", args.Value)
	return nil
}

func (kv *KV) Append(args *PutArgs, reply *PutReply) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	if oldVal, ok := kv.data[args.Key]; ok {
		reply.Value = oldVal
	} else {
		reply.Value = ""
	}
	fmt.Printf("received append call from client %s\n", args.Value)
	kv.data[args.Key] = reply.Value + args.Value
	return nil
}

//
// main
//

func main() {
	server()
	append("subject", "init-")
	put("subject", "6.824")
	put("subject", "6.824-new")
	append("subject", "course")
	fmt.Printf("Put(subject, 6.824) done\n")
	fmt.Printf("get(subject) -> %s\n", get("subject"))
}
