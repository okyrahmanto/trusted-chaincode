/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

/*
// SmartContract provides functions for managing a car
type SmartContract struct {
	contractapi.Contract
}

// Car describes basic details of what makes up a car
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	cars := []Car{
		Car{Make: "Toyota", Model: "Prius", Colour: "blue", Owner: "Tomoko"},
		Car{Make: "Ford", Model: "Mustang", Colour: "red", Owner: "Brad"},
		Car{Make: "Hyundai", Model: "Tucson", Colour: "green", Owner: "Jin Soo"},
		Car{Make: "Volkswagen", Model: "Passat", Colour: "yellow", Owner: "Max"},
		Car{Make: "Tesla", Model: "S", Colour: "black", Owner: "Adriana"},
		Car{Make: "Peugeot", Model: "205", Colour: "purple", Owner: "Michel"},
		Car{Make: "Chery", Model: "S22L", Colour: "white", Owner: "Aarav"},
		Car{Make: "Fiat", Model: "Punto", Colour: "violet", Owner: "Pari"},
		Car{Make: "Tata", Model: "Nano", Colour: "indigo", Owner: "Valeria"},
		Car{Make: "Holden", Model: "Barina", Colour: "brown", Owner: "Shotaro"},
	}

	for i, car := range cars {
		carAsBytes, _ := json.Marshal(car)
		err := ctx.GetStub().PutState("CAR"+strconv.Itoa(i), carAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateCar adds a new car to the world state with given details
func (s *SmartContract) CreateCar(ctx contractapi.TransactionContextInterface, carNumber string, make string, model string, colour string, owner string) error {
	car := Car{
		Make:   make,
		Model:  model,
		Colour: colour,
		Owner:  owner,
	}

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryCar(ctx contractapi.TransactionContextInterface, carNumber string) (*Car, error) {
	carAsBytes, err := ctx.GetStub().GetState(carNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if carAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", carNumber)
	}

	car := new(Car)
	_ = json.Unmarshal(carAsBytes, car)

	return car, nil
}

// QueryAllCars returns all cars found in world state
func (s *SmartContract) QueryAllCars(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {
	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		car := new(Car)
		_ = json.Unmarshal(queryResponse.Value, car)

		queryResult := QueryResult{Key: queryResponse.Key, Record: car}
		results = append(results, queryResult)
	}

	return results, nil
}

// ChangeCarOwner updates the owner field of car with given id in world state
func (s *SmartContract) ChangeCarOwner(ctx contractapi.TransactionContextInterface, carNumber string, newOwner string) error {
	car, err := s.QueryCar(ctx, carNumber)

	if err != nil {
		return err
	}

	car.Owner = newOwner

	carAsBytes, _ := json.Marshal(car)

	return ctx.GetStub().PutState(carNumber, carAsBytes)
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
*/

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	AgentID        string `json:"AgentID"`
	DeviceID       string `json:"DeviceID"`
	SubcribePath   string `json:"SubcribePath"`
	TrustValue     string `json:"TrustValue"`
	AppraisedValue string `json:"AppraisedValue"`
}

type Agent struct {
	AgentID      string             `json:"AgentID"`
	DeviceID     string             `json:"DeviceID"`
	SubcribePath string             `json:"SubcribePath"`
	TrustValue   string             `json:"TrustValue"`
	Tolerance    string             `json:"Tolerance"`
	Transaction  []TransactionAgent `json:"TransactionAgent"`
	Ratings      []Rating           `json:"Rating"`
}

type TransactionAgent struct {
	TransactionID string `json:"TransactionID"`
	Timestamp     string `json:"Timestamp"`
	Source        string `json:"Source"`
	Destination   string `json:"Destination"`
	MessageType   string `json:"MessageType"` // messagetype : request, response. forward
	Data          string `json:"Data"`
}

type Rating struct {
	AgentID    string `json:"AgentID"`
	Value      string `json:"Value"`
	Reputation string `json:"Reputation"` // rep : good | bad
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAgent(ctx contractapi.TransactionContextInterface, agentID string, deviceID string, subcribePath string, trustValue string, tolerance string) error {
	exists, err := s.AgentExists(ctx, agentID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the agent %s already exists", agentID)
	}

	tAgent := []TransactionAgent{
		{TransactionID: "nil",
			Timestamp:   "",
			Source:      "",
			Destination: "",
			MessageType: "",
			Data:        ""},
	}

	rAgent := []Rating{

		{AgentID: "nil",
			Value:      "",
			Reputation: ""},
	}

	asset := Agent{
		AgentID:      agentID,
		DeviceID:     deviceID,
		SubcribePath: subcribePath,
		TrustValue:   trustValue,
		Tolerance:    tolerance,
		Transaction:  tAgent,
		Ratings:      rAgent,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(agentID, assetJSON)
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) UpdateAgent(ctx contractapi.TransactionContextInterface, agentID string, deviceID string, subcribePath string, trustValue string, tolerance string) error {
	exists, err := s.AgentExists(ctx, agentID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the agent %s not exists", agentID)
	}

	agent, err := s.GetAgent(ctx, agentID)
	if err != nil {
		return err
	}

	agent.SubcribePath = subcribePath
	agent.TrustValue = trustValue
	agent.Tolerance = tolerance
	agentJSON, err := json.Marshal(agent)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(agentID, agentJSON)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AgentExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) GetAgent(ctx contractapi.TransactionContextInterface, id string) (*Agent, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var agent Agent
	err = json.Unmarshal(assetJSON, &agent)
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

func (s *SmartContract) AddTransactionAgent(ctx contractapi.TransactionContextInterface, id string, timestamp string, source string, destination string, messageType string, data string) error {
	/*assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	var agent Agent
	err = json.Unmarshal(assetJSON, &agent)
	if err != nil {
		return err
	}*/

	agent, err := s.GetAgent(ctx, id)
	if err != nil {
		return err
	}

	/*assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	//var agent Agent
	err = json.Unmarshal(assetJSON, &agent)
	if err != nil {
		return err
	}*/

	//return &agent, nil

	//create transactionAgent
	//time := createTimestamp()

	tAgent := TransactionAgent{
		TransactionID: id + "." + timestamp,
		Timestamp:     timestamp,
		Source:        source,
		Destination:   destination,
		Data:          "",
	}

	//var transactions []TransactionAgent
	transactions := agent.Transaction
	//agent.Transaction[0] = tAgent
	transactions = append(transactions, tAgent)
	agent.Transaction = transactions

	agentJSON, err := json.Marshal(agent)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, agentJSON)

}

func (s *SmartContract) AddRatingAgent(ctx contractapi.TransactionContextInterface, id string, trusteeAgentID string, value string, reputation string) error {

	agent, err := s.GetAgent(ctx, id)
	if err != nil {
		return err
	}

	rAgent := Rating{
		AgentID:    trusteeAgentID,
		Value:      value,
		Reputation: reputation,
	}

	//var transactions []TransactionAgent
	ratings := agent.Ratings
	ratings = append(ratings, rAgent)
	agent.Ratings = ratings

	agentJSON, err := json.Marshal(agent)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, agentJSON)

}

func (s *SmartContract) UpdateRatingAgent(ctx contractapi.TransactionContextInterface, id string, trusteeAgentID string, value string, reputation string) error {

	agent, err := s.GetAgent(ctx, id)
	if err != nil {
		return err
	}

	//var transactions []TransactionAgent
	ratings := agent.Ratings

	for i := range ratings {
		if ratings[i].AgentID == trusteeAgentID {
			// Found!
			ratings[i].Value = value
			ratings[i].Reputation = reputation
			break
		}
	}

	agent.Ratings = ratings

	agentJSON, err := json.Marshal(agent)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, agentJSON)

}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAgent(ctx contractapi.TransactionContextInterface) ([]*Agent, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var agents []*Agent
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var agent Agent
		err = json.Unmarshal(queryResponse.Value, &agent)
		if err != nil {
			return nil, err
		}
		agents = append(agents, &agent)
	}

	return agents, nil
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedgerMod(ctx contractapi.TransactionContextInterface) error {
	transactionsAgent := []TransactionAgent{
		{TransactionID: "111", Timestamp: "3424", Source: "", Destination: "", Data: ""},
	}

	assets := []Agent{
		{AgentID: "asset1", DeviceID: "blue", SubcribePath: "5", TrustValue: "Tomoko", Tolerance: "300", Transaction: transactionsAgent},
		{AgentID: "asset2", DeviceID: "red", SubcribePath: "5", TrustValue: "Brad", Tolerance: "400", Transaction: transactionsAgent},
		{AgentID: "asset3", DeviceID: "green", SubcribePath: "10", TrustValue: "Jin Soo", Tolerance: "500", Transaction: transactionsAgent},
		{AgentID: "asset4", DeviceID: "yellow", SubcribePath: "10", TrustValue: "Max", Tolerance: "600", Transaction: transactionsAgent},
		{AgentID: "asset5", DeviceID: "black", SubcribePath: "15", TrustValue: "Adriana", Tolerance: "700", Transaction: transactionsAgent},
		{AgentID: "asset6", DeviceID: "white", SubcribePath: "15", TrustValue: "Michel", Tolerance: "800", Transaction: transactionsAgent},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.AgentID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAssetMod(ctx contractapi.TransactionContextInterface, agentID string, deviceID string, subcribePath string, trustValue string, appraisedValue string) error {
	exists, err := s.AssetExists(ctx, agentID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", agentID)
	}

	transactionsAgent := []TransactionAgent{
		{TransactionID: "114", Timestamp: "3424", Source: "", Destination: "", Data: ""},
	}

	asset := Agent{
		AgentID:      agentID,
		DeviceID:     deviceID,
		SubcribePath: subcribePath,
		TrustValue:   trustValue,
		Tolerance:    appraisedValue,
		Transaction:  transactionsAgent,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(agentID, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAssetMod(ctx contractapi.TransactionContextInterface, id string) (*Agent, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Agent
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{AgentID: "asset1", DeviceID: "blue", SubcribePath: "5", TrustValue: "Tomoko", AppraisedValue: "300"},
		{AgentID: "asset2", DeviceID: "red", SubcribePath: "5", TrustValue: "Brad", AppraisedValue: "400"},
		{AgentID: "asset3", DeviceID: "green", SubcribePath: "10", TrustValue: "Jin Soo", AppraisedValue: "500"},
		{AgentID: "asset4", DeviceID: "yellow", SubcribePath: "10", TrustValue: "Max", AppraisedValue: "600"},
		{AgentID: "asset5", DeviceID: "black", SubcribePath: "15", TrustValue: "Adriana", AppraisedValue: "700"},
		{AgentID: "asset6", DeviceID: "white", SubcribePath: "15", TrustValue: "Michel", AppraisedValue: "800"},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.AgentID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, agentID string, deviceID string, subcribePath string, trustValue string, appraisedValue string) error {
	exists, err := s.AssetExists(ctx, agentID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", agentID)
	}

	asset := Asset{
		AgentID:        agentID,
		DeviceID:       deviceID,
		SubcribePath:   subcribePath,
		TrustValue:     trustValue,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(agentID, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, agentID string, deviceID string, subcribePath string, trustValue string, appraisedValue string) error {
	exists, err := s.AssetExists(ctx, agentID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", agentID)
	}

	// overwriting original asset with new asset
	asset := Asset{
		AgentID:        agentID,
		DeviceID:       deviceID,
		SubcribePath:   subcribePath,
		TrustValue:     trustValue,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(agentID, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the owner field of asset with given id in world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newTrustValue string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.TrustValue = newTrustValue
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
