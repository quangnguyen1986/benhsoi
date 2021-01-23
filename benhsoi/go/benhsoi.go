package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/flogging"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// SmartContract Define the Smart Contract structure
type SmartContract struct {
}

// Benhsoi :  Define the benhsoi structure, with 4 properties.  Structure tags are used by encoding/json library
type Benhsoi struct {
	Mabenhsoi   string `json:"mabenhsoi"`
	Ngaynhap  string `json:"ngaynhap"`
	Thongtinchitiet string `json:"thongtinchitiet"`
	Nhanviencdc  string `json:"nhanviencdc"`
}

type benhsoiPrivateDetails struct {
	Nhanviencdc string `json:"nhanviencdc"`
	Thongtinrieng string `json:"thongtinrieng"`
}

// Init ;  Method for initializing smart contract
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

var logger = flogging.MustGetLogger("benhsoi_cc")

// Invoke :  Method for INVOKING smart contract
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	logger.Infof("Function name is:  %d", function)
	logger.Infof("Args length is : %d", len(args))

	switch function {
	case "queryBenhsoi":
		return s.queryBenhsoi(APIstub, args)
	case "initLedger":
		return s.initLedger(APIstub)
	case "createBenhsoi":
		return s.createBenhsoi(APIstub, args)
	case "queryAllBenhsois":
		return s.queryAllBenhsois(APIstub)
	case "changeBenhsoiNhanviencdc":
		return s.changeBenhsoiNhanviencdc(APIstub, args)
	case "getHistoryForAsset":
		return s.getHistoryForAsset(APIstub, args)
	case "queryBenhsoisByNhanviencdc":
		return s.queryBenhsoisByNhanviencdc(APIstub, args)
	case "restictedMethod":
		return s.restictedMethod(APIstub, args)
	case "test":
		return s.test(APIstub, args)
	case "createPrivateBenhsoi":
		return s.createPrivateBenhsoi(APIstub, args)
	case "readPrivateBenhsoi":
		return s.readPrivateBenhsoi(APIstub, args)
	case "updatePrivateData":
		return s.updatePrivateData(APIstub, args)
	case "readBenhsoiPrivateDetails":
		return s.readBenhsoiPrivateDetails(APIstub, args)
	case "createPrivateBenhsoiImplicitForOrg1":
		return s.createPrivateBenhsoiImplicitForOrg1(APIstub, args)
	case "createPrivateBenhsoiImplicitForOrg2":
		return s.createPrivateBenhsoiImplicitForOrg2(APIstub, args)
	case "queryPrivateDataHash":
		return s.queryPrivateDataHash(APIstub, args)
	default:
		return shim.Error("Invalid Smart Contract function name.")
	}

	// return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryBenhsoi(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	benhsoiAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) readPrivateBenhsoi(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	// collectionBenhsois, collectionBenhsoiPrivateDetails, _implicit_org_Org1MSP, _implicit_org_Org2MSP
	benhsoiAsBytes, err := APIstub.GetPrivateData(args[0], args[1])
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get private details for " + args[1] + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if benhsoiAsBytes == nil {
		jsonResp := "{\"Error\":\"Benhsoi private details does not exist: " + args[1] + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) readPrivateBenhsoiIMpleciteForOrg1(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	benhsoiAsBytes, _ := APIstub.GetPrivateData("_implicit_org_Org1MSP", args[0])
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) readBenhsoiPrivateDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	benhsoiAsBytes, err := APIstub.GetPrivateData("collectionBenhsoiPrivateDetails", args[0])

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get private details for " + args[0] + ": " + err.Error() + "\"}"
		return shim.Error(jsonResp)
	} else if benhsoiAsBytes == nil {
		jsonResp := "{\"Error\":\"Marble private details does not exist: " + args[0] + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) test(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	benhsoiAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	benhsois := []Benhsoi{
		Benhsoi{Mabenhsoi: "Toyota", Ngaynhap: "Prius", Thongtinchitiet: "blue", Nhanviencdc: "Tomoko"},
		Benhsoi{Mabenhsoi: "Ford", Ngaynhap: "Mustang", Thongtinchitiet: "red", Nhanviencdc: "Brad"},
		Benhsoi{Mabenhsoi: "Hyundai", Ngaynhap: "Tucson", Thongtinchitiet: "green", Nhanviencdc: "Jin Soo"},
		Benhsoi{Mabenhsoi: "Volkswagen", Ngaynhap: "Passat", Thongtinchitiet: "yellow", Nhanviencdc: "Max"},
		Benhsoi{Mabenhsoi: "Tesla", Ngaynhap: "S", Thongtinchitiet: "black", Nhanviencdc: "Adriana"},
		Benhsoi{Mabenhsoi: "Peugeot", Ngaynhap: "205", Thongtinchitiet: "purple", Nhanviencdc: "Michel"},
		Benhsoi{Mabenhsoi: "Chery", Ngaynhap: "S22L", Thongtinchitiet: "white", Nhanviencdc: "Aarav"},
		Benhsoi{Mabenhsoi: "Fiat", Ngaynhap: "Punto", Thongtinchitiet: "violet", Nhanviencdc: "Pari"},
		Benhsoi{Mabenhsoi: "Tata", Ngaynhap: "Nano", Thongtinchitiet: "indigo", Nhanviencdc: "Valeria"},
		Benhsoi{Mabenhsoi: "Holden", Ngaynhap: "Barina", Thongtinchitiet: "brown", Nhanviencdc: "Shotaro"},
	}

	i := 0
	for i < len(benhsois) {
		benhsoiAsBytes, _ := json.Marshal(benhsois[i])
		APIstub.PutState("CAR"+strconv.Itoa(i), benhsoiAsBytes)
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createPrivateBenhsoi(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	type benhsoiTransientInput struct {
		Mabenhsoi  string `json:"mabenhsoi"` //the fieldtags are needed to keep case from bouncing around
		Ngaynhap string `json:"ngaynhap"`
		Thongtinchitiet string `json:"thongtinchitiet"`
		Nhanviencdc string `json:"nhanviencdc"`
		Thongtinrieng string `json:"thongtinrieng"`
		Key   string `json:"key"`
	}
	if len(args) != 0 {
		return shim.Error("1111111----Incorrect number of arguments. Private marble data must be passed in transient map.")
	}

	logger.Infof("11111111111111111111111111")

	transMap, err := APIstub.GetTransient()
	if err != nil {
		return shim.Error("222222 -Error getting transient: " + err.Error())
	}

	benhsoiDataAsBytes, ok := transMap["benhsoi"]
	if !ok {
		return shim.Error("benhsoi must be a key in the transient map")
	}
	logger.Infof("********************8   " + string(benhsoiDataAsBytes))

	if len(benhsoiDataAsBytes) == 0 {
		return shim.Error("333333 -marble value in the transient map must be a non-empty JSON string")
	}

	logger.Infof("2222222")

	var benhsoiInput benhsoiTransientInput
	err = json.Unmarshal(benhsoiDataAsBytes, &benhsoiInput)
	if err != nil {
		return shim.Error("44444 -Failed to decode JSON of: " + string(benhsoiDataAsBytes) + "Error is : " + err.Error())
	}

	logger.Infof("3333")

	if len(benhsoiInput.Key) == 0 {
		return shim.Error("name field must be a non-empty string")
	}
	if len(benhsoiInput.Mabenhsoi) == 0 {
		return shim.Error("thongtinchitiet field must be a non-empty string")
	}
	if len(benhsoiInput.Ngaynhap) == 0 {
		return shim.Error("ngaynhap field must be a non-empty string")
	}
	if len(benhsoiInput.Thongtinchitiet) == 0 {
		return shim.Error("thongtinchitiet field must be a non-empty string")
	}
	if len(benhsoiInput.Nhanviencdc) == 0 {
		return shim.Error("nhanviencdc field must be a non-empty string")
	}
	if len(benhsoiInput.Thongtinrieng) == 0 {
		return shim.Error("thongtinrieng field must be a non-empty string")
	}

	logger.Infof("444444")

	// ==== Check if benhsoi already exists ====
	benhsoiAsBytes, err := APIstub.GetPrivateData("collectionBenhsois", benhsoiInput.Key)
	if err != nil {
		return shim.Error("Failed to get marble: " + err.Error())
	} else if benhsoiAsBytes != nil {
		fmt.Println("This benhsoi already exists: " + benhsoiInput.Key)
		return shim.Error("This benhsoi already exists: " + benhsoiInput.Key)
	}

	logger.Infof("55555")

	var benhsoi = Benhsoi{Mabenhsoi: benhsoiInput.Mabenhsoi, Ngaynhap: benhsoiInput.Ngaynhap, Thongtinchitiet: benhsoiInput.Thongtinchitiet, Nhanviencdc: benhsoiInput.Nhanviencdc}

	benhsoiAsBytes, err = json.Marshal(benhsoi)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = APIstub.PutPrivateData("collectionBenhsois", benhsoiInput.Key, benhsoiAsBytes)
	if err != nil {
		logger.Infof("6666666")
		return shim.Error(err.Error())
	}

	benhsoiPrivateDetails := &benhsoiPrivateDetails{Nhanviencdc: benhsoiInput.Nhanviencdc, Thongtinrieng: benhsoiInput.Thongtinrieng}

	benhsoiPrivateDetailsAsBytes, err := json.Marshal(benhsoiPrivateDetails)
	if err != nil {
		logger.Infof("77777")
		return shim.Error(err.Error())
	}

	err = APIstub.PutPrivateData("collectionBenhsoiPrivateDetails", benhsoiInput.Key, benhsoiPrivateDetailsAsBytes)
	if err != nil {
		logger.Infof("888888")
		return shim.Error(err.Error())
	}

	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) updatePrivateData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	type benhsoiTransientInput struct {
		Nhanviencdc string `json:"nhanviencdc"`
		Thongtinrieng string `json:"thongtinrieng"`
		Key   string `json:"key"`
	}
	if len(args) != 0 {
		return shim.Error("1111111----Incorrect number of arguments. Private marble data must be passed in transient map.")
	}

	logger.Infof("11111111111111111111111111")

	transMap, err := APIstub.GetTransient()
	if err != nil {
		return shim.Error("222222 -Error getting transient: " + err.Error())
	}

	benhsoiDataAsBytes, ok := transMap["benhsoi"]
	if !ok {
		return shim.Error("benhsoi must be a key in the transient map")
	}
	logger.Infof("********************8   " + string(benhsoiDataAsBytes))

	if len(benhsoiDataAsBytes) == 0 {
		return shim.Error("333333 -marble value in the transient map must be a non-empty JSON string")
	}

	logger.Infof("2222222")

	var benhsoiInput benhsoiTransientInput
	err = json.Unmarshal(benhsoiDataAsBytes, &benhsoiInput)
	if err != nil {
		return shim.Error("44444 -Failed to decode JSON of: " + string(benhsoiDataAsBytes) + "Error is : " + err.Error())
	}

	benhsoiPrivateDetails := &benhsoiPrivateDetails{Nhanviencdc: benhsoiInput.Nhanviencdc, Thongtinrieng: benhsoiInput.Thongtinrieng}

	benhsoiPrivateDetailsAsBytes, err := json.Marshal(benhsoiPrivateDetails)
	if err != nil {
		logger.Infof("77777")
		return shim.Error(err.Error())
	}

	err = APIstub.PutPrivateData("collectionBenhsoiPrivateDetails", benhsoiInput.Key, benhsoiPrivateDetailsAsBytes)
	if err != nil {
		logger.Infof("888888")
		return shim.Error(err.Error())
	}

	return shim.Success(benhsoiPrivateDetailsAsBytes)

}

func (s *SmartContract) createBenhsoi(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var benhsoi = Benhsoi{Mabenhsoi: args[1], Ngaynhap: args[2], Thongtinchitiet: args[3], Nhanviencdc: args[4]}

	benhsoiAsBytes, _ := json.Marshal(benhsoi)
	APIstub.PutState(args[0], benhsoiAsBytes)

	indexName := "nhanviencdc~key"
	thongtinchitietNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{benhsoi.Nhanviencdc, args[0]})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	APIstub.PutState(thongtinchitietNameIndexKey, value)

	return shim.Success(benhsoiAsBytes)
}

func (S *SmartContract) queryBenhsoisByNhanviencdc(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}
	nhanviencdc := args[0]

	nhanviencdcAndIdResultIterator, err := APIstub.GetStateByPartialCompositeKey("nhanviencdc~key", []string{nhanviencdc})
	if err != nil {
		return shim.Error(err.Error())
	}

	defer nhanviencdcAndIdResultIterator.Close()

	var i int
	var id string

	var benhsois []byte
	bArrayMemberAlreadyWritten := false

	benhsois = append([]byte("["))

	for i = 0; nhanviencdcAndIdResultIterator.HasNext(); i++ {
		responseRange, err := nhanviencdcAndIdResultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		objectType, compositeKeyParts, err := APIstub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}

		id = compositeKeyParts[1]
		assetAsBytes, err := APIstub.GetState(id)

		if bArrayMemberAlreadyWritten == true {
			newBytes := append([]byte(","), assetAsBytes...)
			benhsois = append(benhsois, newBytes...)

		} else {
			// newBytes := append([]byte(","), benhsoisAsBytes...)
			benhsois = append(benhsois, assetAsBytes...)
		}

		fmt.Printf("Found a asset for index : %s asset id : ", objectType, compositeKeyParts[0], compositeKeyParts[1])
		bArrayMemberAlreadyWritten = true

	}

	benhsois = append(benhsois, []byte("]")...)

	return shim.Success(benhsois)
}

func (s *SmartContract) queryAllBenhsois(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAR0"
	endKey := "CAR999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllBenhsois:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) restictedMethod(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// get an ID for the client which is guaranteed to be unique within the MSP
	//id, err := cid.GetID(APIstub) -

	// get the MSP ID of the client's identity
	//mspid, err := cid.GetMSPID(APIstub) -

	// get the value of the attribute
	//val, ok, err := cid.GetAttributeValue(APIstub, "attr1") -

	// get the X509 certificate of the client, or nil if the client's identity was not based on an X509 certificate
	//cert, err := cid.GetX509Certificate(APIstub) -

	val, ok, err := cid.GetAttributeValue(APIstub, "role")
	if err != nil {
		// There was an error trying to retrieve the attribute
		shim.Error("Error while retriving attributes")
	}
	if !ok {
		// The client identity does not possess the attribute
		shim.Error("Client identity doesnot posses the attribute")
	}
	// Do something with the value of 'val'
	if val != "approver" {
		fmt.Println("Attribute role: " + val)
		return shim.Error("Only user with role as APPROVER have access this method!")
	} else {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1")
		}

		benhsoiAsBytes, _ := APIstub.GetState(args[0])
		return shim.Success(benhsoiAsBytes)
	}

}

func (s *SmartContract) changeBenhsoiNhanviencdc(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	benhsoiAsBytes, _ := APIstub.GetState(args[0])
	benhsoi := Benhsoi{}

	json.Unmarshal(benhsoiAsBytes, &benhsoi)
	benhsoi.Nhanviencdc = args[1]

	benhsoiAsBytes, _ = json.Marshal(benhsoi)
	APIstub.PutState(args[0], benhsoiAsBytes)

	return shim.Success(benhsoiAsBytes)
}

func (t *SmartContract) getHistoryForAsset(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	benhsoiName := args[0]

	resultsIterator, err := stub.GetHistoryForKey(benhsoiName)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForAsset returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) createPrivateBenhsoiImplicitForOrg1(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect arguments. Expecting 5 arguments")
	}

	var benhsoi = Benhsoi{Mabenhsoi: args[1], Ngaynhap: args[2], Thongtinchitiet: args[3], Nhanviencdc: args[4]}

	benhsoiAsBytes, _ := json.Marshal(benhsoi)
	// APIstub.PutState(args[0], benhsoiAsBytes)

	err := APIstub.PutPrivateData("_implicit_org_Org1MSP", args[0], benhsoiAsBytes)
	if err != nil {
		return shim.Error("Failed to add asset: " + args[0])
	}
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) createPrivateBenhsoiImplicitForOrg2(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect arguments. Expecting 5 arguments")
	}

	var benhsoi = Benhsoi{Mabenhsoi: args[1], Ngaynhap: args[2], Thongtinchitiet: args[3], Nhanviencdc: args[4]}

	benhsoiAsBytes, _ := json.Marshal(benhsoi)
	APIstub.PutState(args[0], benhsoiAsBytes)

	err := APIstub.PutPrivateData("_implicit_org_Org2MSP", args[0], benhsoiAsBytes)
	if err != nil {
		return shim.Error("Failed to add asset: " + args[0])
	}
	return shim.Success(benhsoiAsBytes)
}

func (s *SmartContract) queryPrivateDataHash(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	benhsoiAsBytes, _ := APIstub.GetPrivateDataHash(args[0], args[1])
	return shim.Success(benhsoiAsBytes)
}

// func (s *SmartContract) CreateBenhsoiAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	var benhsoi Benhsoi
// 	err := json.Unmarshal([]byte(args[0]), &benhsoi)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	benhsoiAsBytes, err := json.Marshal(benhsoi)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	err = APIstub.PutState(benhsoi.ID, benhsoiAsBytes)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return shim.Success(nil)
// }

// func (s *SmartContract) addBulkAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
// 	logger.Infof("Function addBulkAsset called and length of arguments is:  %d", len(args))
// 	if len(args) >= 500 {
// 		logger.Errorf("Incorrect number of arguments in function CreateAsset, expecting less than 500, but got: %b", len(args))
// 		return shim.Error("Incorrect number of arguments, expecting 2")
// 	}

// 	var eventKeyValue []string

// 	for i, s := range args {

// 		key :=s[0];
// 		var benhsoi = Benhsoi{Mabenhsoi: s[1], Ngaynhap: s[2], Thongtinchitiet: s[3], Nhanviencdc: s[4]}

// 		eventKeyValue = strings.SplitN(s, "#", 3)
// 		if len(eventKeyValue) != 3 {
// 			logger.Errorf("Error occured, Please mabenhsoi sure that you have provided the array of strings and each string should be  in \"EventType#Key#Value\" format")
// 			return shim.Error("Error occured, Please mabenhsoi sure that you have provided the array of strings and each string should be  in \"EventType#Key#Value\" format")
// 		}

// 		assetAsBytes := []byte(eventKeyValue[2])
// 		err := APIstub.PutState(eventKeyValue[1], assetAsBytes)
// 		if err != nil {
// 			logger.Errorf("Error coocured while putting state for asset %s in APIStub, error: %s", eventKeyValue[1], err.Error())
// 			return shim.Error(err.Error())
// 		}
// 		// logger.infof("Adding value for ")
// 		fmt.Println(i, s)

// 		indexName := "Event~Id"
// 		eventAndIDIndexKey, err2 := APIstub.CreateCompositeKey(indexName, []string{eventKeyValue[0], eventKeyValue[1]})

// 		if err2 != nil {
// 			logger.Errorf("Error coocured while putting state in APIStub, error: %s", err.Error())
// 			return shim.Error(err2.Error())
// 		}

// 		value := []byte{0x00}
// 		err = APIstub.PutState(eventAndIDIndexKey, value)
// 		if err != nil {
// 			logger.Errorf("Error coocured while putting state in APIStub, error: %s", err.Error())
// 			return shim.Error(err.Error())
// 		}
// 		// logger.Infof("Created Composite key : %s", eventAndIDIndexKey)

// 	}

// 	return shim.Success(nil)
// }

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
