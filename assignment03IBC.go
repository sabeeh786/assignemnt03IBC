package assignment03
import (
	"fmt"
	"log"
	"net"
	"encoding/gob" 
	"os"
	"strconv"
a1 "github.com/sabeeh786/assignment01IBC"
)


type User struct {
	Name string
	TotalBalance int
	PortNumber    string
}

func Mining(chainHead *a1.Block,receiver *User,sender *User,miner *User,amount int) *a1.Block{
	if amount<=sender.TotalBalance{
		sender.TotalBalance-=amount
		receiver.TotalBalance+=amount
		miner.TotalBalance+=100
		s:=sender.Name + " sent " +string(amount) + " HAFAZ coins to " + receiver.Name
		chainHead=a1.InsertBlock(s,chainHead)
		fmt.Println(s)
	} else {
		fmt.Println("Transaction is Invalid")
	}
	return chainHead
}


func server(chainHead *a1.Block, Array []user, NumberOfUsers int){


		fmt.Println("Server")
		ln, err := net.Listen("tcp", ":6000")	
		if err != nil {
			log.Fatal(err)
		}
		chainHead=Mining(chainHead,&Array[0],&Array[0],&Array[0],0)
		for i:=0; i<NumberOfUsers; i++  {
			chainHead=Mining(chainHead,&Array[0],&Array[0],&Array[0],0)
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			var new User
			decoder := gob.NewDecoder(conn)
			err = decoder.Decode(&new)
			if err != nil {
				//handle error
			}
			Array=append(Array,new)
			fmt.Println(Array[0].TotalBalance)
		}
		
	fmt.Println(Array)
	for i:=0; i<len(Array); i++{
	
	conn, err := net.Dial("tcp", "localhost:"+Array[i].PortNumber)
	if err != nil {
		//handle error
		}
		
	//Broadcasting Blockchain	
	Encoder := gob.NewEncoder(conn)
	err1 := Encoder.Encode(&chainHead)
	if err1 != nil {
		//handle error
		}
	}
	for i:=0; i<len(Array); i++{
	conn, err := net.Dial("tcp", "localhost:"+Array[i].PortNumber)
	if err != nil {
		//handle error
		}
	//Broadcasting Users
	Encoder1 := gob.NewEncoder(conn)
	err2 := Encoder1.Encode(&Array)
	if err2 != nil {
		//handle error
		}		
	}

}


func client(chainHead *a1.Block, Array []user){

	fmt.Println("client")
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		//handle error
	}
	p:=User{os.Args[1],0,os.Args[3]}
	
	//var recvdBlock a1.Block
	Encoder := gob.NewEncoder(conn)
	err2 := Encoder.Encode(&p)
	if err2 != nil {
		//handle error
	}
		
	
	ln, err := net.Listen("tcp", ":"+os.Args[3])	
		if err != nil {
			log.Fatal(err)
		}	
	conn, err = ln.Accept()
	if err != nil {
		log.Println(err)
	}
	var receivedHead *a1.Block
	decoder := gob.NewDecoder(conn)
		err3 := decoder.Decode(&receivedHead)
		if err3 != nil {
			//handle error
		}
	conn, err = ln.Accept()
	if err != nil {
		log.Println(err)
		
	}
	decoder1 := gob.NewDecoder(conn)
		err4 := decoder1.Decode(&Array)
		if err4 != nil {
			//handle error
		}
	chainHead=receivedHead	
	fmt.Println(Array)

}
