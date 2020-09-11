package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type File_onChain struct{
	Filename       string `json:"Filename"`
	File_branch    string `json:"File_branch"`
	Access_time    string `json:"Access_time"`
	Visitor        string `json:"Visitor"`
	Visitor_branch string `json:"Visitor_branch"`
	Visitor_role   string `json:"Visitor_role"`
}


func (t *File_onChain) Init (stub shim.ChaincodeStubInterface) pb.Response{
	return shim.Success(nil)
}

func (t *File_onChain) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	funcName,args := stub.GetFunctionAndParameters()
	if(funcName=="save"){
		return t.saveBasic(stub,args)
	}else if(funcName=="query"){
		return t.queryBasic(stub,args)
	}else{
		return shim.Error("no such function")
	}
}

func (t *File_onChain) saveBasic(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if(len(args)!=5){
		return shim.Error("except five args")
	}else{
		err:=stub.PutState(args[0],[]byte(args[1]))
		if(err!=nil) {
			return shim.Error(err.Error())
		}
		return shim.Success(nil)
	}

}

func (t *File_onChain) queryBasic(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	if(len(args)!=1){
		return shim.Error("except one arg")
	}else{
	  value,err :=stub.GetState(args[0])
	  if(err!=nil){
	  	shim.Error("no data found")
	  }
	  return shim.Success(value)
	}
}

func main(){
	err:=shim.Start(new(File_onChain))
	if(err!=nil){
		fmt.Println("emr File_onChain chaincode start error")
	}
}
