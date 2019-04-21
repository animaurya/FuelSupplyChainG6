/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright Ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Vehicle management Use Case - WORK IN  PROGRESS
 */

package main


import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}


// Define the petrol inventory
type PetrolInv struct {
	Owner string `json:"owner"`
	Volume int `json:"volume"`
	Temperature int `json:"temperature"`
	Density int `json:"density"`
	// weight float `json:"weight"`
	// location string `json:"location"`
	// pressure string `json:"pressure"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {

	fmt.Println("Initiating Storage1 inventory ")

	empty_Inv := PetrolInv{Owner: "Storage1", Volume: 0, Temperature: 30, Density: 750}


	InvBytes, err := json.Marshal(empty_Inv)

    if err != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }
	fmt.Println("empty_Inv ", empty_Inv, InvBytes)

    err = APIstub.PutState("Storage1",InvBytes)

    if err != nil {
    	return shim.Error("failed to add Inventory for this Owner")
    }

	fmt.Println("Initiated Storage1 inventory ")

	fmt.Println("Initiating Transport1 inventory ")

	empty_Inv = PetrolInv{Owner: "Transport1", Volume: 0, Temperature: 30, Density: 750}
	InvBytes, err = json.Marshal(empty_Inv)

    if err != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err = APIstub.PutState("Transport1",InvBytes)

    if err != nil {
    	return shim.Error("failed to add Inventory for this Owner")
    }

	fmt.Println("Initiated Transport1 inventory ")

	fmt.Println("Initiating Station1 inventory ")

	empty_Inv = PetrolInv{Owner: "Station1", Volume: 0, Temperature: 30, Density: 750}
	InvBytes, err = json.Marshal(empty_Inv)

    if err != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err = APIstub.PutState("Station1",InvBytes)

    if err != nil {
    	return shim.Error("failed to add Inventory for this Owner")
    }

	fmt.Println("Initiated Station1 inventory ")

	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addFuel" {
		return s.addFuel(APIstub, args)
	} else if function == "move" {
		return s.move(APIstub, args)
	} else if function == "update" {
		return s.update(APIstub, args)
	} else if function == "transfer" {
		return s.transfer(APIstub, args)
	} else if function == "sell" {
		return s.sell(APIstub, args)
	} else if function == "viewStatus" {
		return s.viewStatus(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// This function is only used by initial storage
func (s *SmartContract) addFuel(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Println("Updating ", args[0], "inventory ")

	CurrInvBytes,err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	var CurrInv PetrolInv
	err = json.Unmarshal(CurrInvBytes, &CurrInv)

	fmt.Println("Current invetory shows" , CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}

	addvol,e1 := strconv.Atoi(args[1])
	addtemp,e2 := strconv.Atoi(args[2])
	adddens,e3 := strconv.Atoi(args[3])

	fmt.Println("features of petrol being added are %d %d %d " , addvol, addtemp, adddens)

	if e1 != nil {
		return shim.Error("Issue with string conversion")
	}

        if e2 != nil {
                return shim.Error("Issue with string conversion")
        }

        if e3 != nil {
                return shim.Error("Issue with string conversion")
        }

	Inv := PetrolInv{Owner: args[0], Volume: addvol+CurrInv.Volume, Temperature: (addvol*addtemp + CurrInv.Volume*CurrInv.Temperature)/(addvol+CurrInv.Volume), Density: (addvol*adddens + CurrInv.Volume*CurrInv.Density)/(addvol+CurrInv.Volume)}
	InvBytes, err2 := json.Marshal(Inv)

	fmt.Println("new inventory shows ",Inv,InvBytes)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err2 = APIstub.PutState(args[0],InvBytes)

    if err2 != nil {
    	return shim.Error("failed to update Inventory for this Owner")
    }
	CurrInvBytes,err = APIstub.GetState(args[0])
        if err != nil {
                return shim.Error(" failed to retrieve inventory with this Owner ")
        }
        err = json.Unmarshal(CurrInvBytes, &CurrInv)

        fmt.Println("Current invetory shows" , CurrInv)
	fmt.Println("Updated ", args[0], "inventory")

	return shim.Success(nil)
}

// This function is only used by initial storage
func (s *SmartContract) move(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	fmt.Println("Updating ", args[0], "inventory ")

	CurrInvBytes,err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	var CurrInv PetrolInv
	err = json.Unmarshal(CurrInvBytes, &CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}

	temp := CurrInv.Temperature
	dens := CurrInv.Density

	subvol,e := strconv.Atoi(args[1])


        if e != nil {
                return shim.Error("Issue with string conversion")
        }
	var Inv PetrolInv

	if (subvol > CurrInv.Volume) {
		return shim.Error("Removing more Volume than present, exiting with error")		
	} else if (subvol == CurrInv.Volume) {
		Inv = PetrolInv{Owner: args[0], Volume: 0.0, Temperature: 30.0, Density:750.0 }
	} else {
		Inv = PetrolInv{Owner: args[0], Volume: CurrInv.Volume - subvol, Temperature: CurrInv.Temperature, Density:CurrInv.Density }	
	}

	InvBytes, err2 := json.Marshal(Inv)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err2 = APIstub.PutState(args[0],InvBytes)

    if err2 != nil {
    	return shim.Error("failed to update Inventory for this Owner")
    }

	fmt.Println("Updated ", args[0], "inventory")

	fmt.Println("Updating ", args[2], "inventory ")

	CurrInvBytes,err = APIstub.GetState(args[2])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	err = json.Unmarshal(CurrInvBytes, &CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}

	addvol,e := strconv.Atoi(args[1])

	Inv = PetrolInv{Owner: args[2], Volume: addvol+CurrInv.Volume, Temperature: (addvol*temp + CurrInv.Volume*CurrInv.Temperature)/(addvol+CurrInv.Volume), Density: (addvol*dens + CurrInv.Volume*CurrInv.Density)/(addvol+CurrInv.Volume)}
	InvBytes, err2 = json.Marshal(Inv)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err2 = APIstub.PutState(args[2],InvBytes)

    if err2 != nil {
    	return shim.Error("failed to update Inventory for this Owner")
    }

	fmt.Println("Updated ", args[2], "inventory")

	return shim.Success(nil)
}

// This function is used by all parties in timely intervals by IOTs
func (s *SmartContract) update(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	InvBytes,err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	var CurrInv PetrolInv
	err = json.Unmarshal(InvBytes, &CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}
	// should check if the caller is the current Owner
	CurrInv.Volume,err = strconv.Atoi(args[1])
	CurrInv.Temperature,err = strconv.Atoi(args[2])
	CurrInv.Density,err = strconv.Atoi(args[3])

	fmt.Println("updated values are", CurrInv)

	Bytes2, err2 := json.Marshal(CurrInv)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err = APIstub.PutState(args[0],Bytes2)

    if err != nil {
    	return shim.Error(" failed to update details of the inventory with this Owner ")
    }

	fmt.Println("update details Requested -> ", args[0])

	return shim.Success(nil)
}

// This function is only used by Transport 
func (s *SmartContract) transfer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	fmt.Println("Updating ", args[0], "inventory ")

	CurrInvBytes,err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	var CurrInv PetrolInv
	err = json.Unmarshal(CurrInvBytes, &CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}

	temp := CurrInv.Temperature
	dens := CurrInv.Density
	var Inv PetrolInv
	subvol,e := strconv.Atoi(args[1])
	

        if e != nil {
                return shim.Error("Issue with string conversion")
        }
	if (subvol > CurrInv.Volume) {
		return shim.Error("Removing more Volume than present, exiting with error")		
	} else if (subvol == CurrInv.Volume) {
		Inv = PetrolInv{Owner: args[0], Volume: 0.0, Temperature: 30.0, Density:750.0 }
	} else {
		Inv = PetrolInv{Owner: args[0], Volume: CurrInv.Volume - subvol, Temperature: CurrInv.Temperature, Density:CurrInv.Density }	
	}

	InvBytes, err2 := json.Marshal(Inv)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err2 = APIstub.PutState(args[0],InvBytes)

    if err2 != nil {
    	return shim.Error("failed to update Inventory for this Owner")
    }

	fmt.Println("Updated ", args[0], "inventory")

	fmt.Println("Updating ", args[2], "inventory ")

	CurrInvBytes,err = APIstub.GetState(args[2])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	err = json.Unmarshal(CurrInvBytes, &CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}

	addvol,e1 := strconv.Atoi(args[1])
	addtemp :=  temp
	adddens := dens


        if e1 != nil {
                return shim.Error("Issue with string conversion")
        }
	Inv = PetrolInv{Owner: args[0], Volume: addvol+CurrInv.Volume, Temperature: (addvol*addtemp + CurrInv.Volume*CurrInv.Temperature)/(addvol+CurrInv.Volume), Density: (addvol*adddens + CurrInv.Volume*CurrInv.Density)/(addvol+CurrInv.Volume)}
	InvBytes, err2 = json.Marshal(Inv)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err2 = APIstub.PutState(args[2],InvBytes)

    if err2 != nil {
    	return shim.Error("failed to update Inventory for this Owner")
    }

	fmt.Println("Updated ", args[2], "inventory")

	return shim.Success(nil)	
}

// This function is only used by Fuel Station
func (s *SmartContract) sell(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	fmt.Println("Updating ", args[0], "inventory ")

	CurrInvBytes,err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}
	var CurrInv PetrolInv
	err = json.Unmarshal(CurrInvBytes, &CurrInv)

	if err != nil {
		return shim.Error("Issue with Inventory json unmarshaling")
	}

	subvol,e := strconv.Atoi(args[1])
	

        if e != nil {
                return shim.Error("Issue with string conversion")
        }
	var Inv PetrolInv
	if (subvol > CurrInv.Volume) {
		return shim.Error("Removing more Volume than present, exiting with error")		
	} else if (subvol == CurrInv.Volume) {
		Inv = PetrolInv{Owner: args[0], Volume: 0.0, Temperature: 30.0, Density:750.0 }
	} else {
		Inv = PetrolInv{Owner: args[0], Volume: CurrInv.Volume - subvol, Temperature: CurrInv.Temperature, Density:CurrInv.Density }	
	}

	InvBytes, err2 := json.Marshal(Inv)

    if err2 != nil {
    	return shim.Error("Issue with Inventory json marshaling")
    }

    err2 = APIstub.PutState(args[0],InvBytes)

    if err2 != nil {
    	return shim.Error("failed to update Inventory for this Owner")
    }

	fmt.Println("Updated ", args[0], "inventory")
	return shim.Success(nil)
}

// This function is only used by all parties
func (s *SmartContract) viewStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	Owner := args[0];
	
	InvBytes,err := APIstub.GetState(Owner)
	if err != nil {
		return shim.Error(" failed to retrieve inventory with this Owner ")
	}

	fmt.Println("Status view Requested for -> ", Owner,string(InvBytes))
	return shim.Success(InvBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

