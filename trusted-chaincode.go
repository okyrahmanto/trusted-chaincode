/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const index = "agentID~TransactionID"

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

type TransactionAgent struct {
	TransactionID    string `json:"TransactionID"` //uuid
	AgentID          string `json:"AgentID"`
	Timestamp        string `json:"Timestamp"`
	RefTransactionID string `json:"RefTransactionID"` //uuid
	Source           string `json:"Source"`
	Destination      string `json:"Destination"`
	MessageType      string `json:"MessageType"` // messagetype : request, response. forward
	Data             string `json:"Data"`
}

type Agent struct {
	AgentID      string `json:"AgentID"`
	DeviceID     string `json:"DeviceID"`
	SubcribePath string `json:"SubcribePath"`
	TrustValue   string `json:"TrustValue"`
	Tolerance    string `json:"Tolerance"`
}

type Rating struct {
	RatingID       string `json:"RatingID"` //uuid
	AgentID        string `json:"AgentID"`
	AgentIDTrustee string `json:"AgentIDTrustee"`
	Value          string `json:"Value"`
	Reputation     string `json:"Reputation"` // rep : good | bad
}

type EvaluationParam struct {
	EvaluationID  string `json:"EvaluationID"` //uuid
	TransactionID string `json:"TransactionID"`
	AgentID       string `json:"AgentID"`
	TrusterID     string `json:"TrusterID"`
	TrusteeID     string `json:"TrusteeID"`
	Timestamp     string `json:"Timestamp"`
	ResponseTime  string `json:"ResponseTime"`
	Validity      string `json:"Validity"`
	Correctness   string `json:"Correctness"`
	Cooperation   string `json:"Cooperation"`
	Qos           string `json:"Qos"`
	Availability  string `json:"Availability"`
	Confidence    string `json:"Confidence"`
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

// QueryResult structure used for handling result of query
type QueryResultAgent struct {
	Key    string `json:"Key"`
	Record *Agent
}

// QueryResult structure used for handling result of query
type QueryResultTransaction struct {
	Key    string `json:"Key"`
	Record *TransactionAgent
}

// QueryResult structure used for handling result of query
type QueryResultRatings struct {
	Key    string `json:"Key"`
	Record *Rating
}

// QueryResult structure used for handling result of query
type QueryResultEvaluation struct {
	Key    string `json:"Key"`
	Record *EvaluationParam
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

// InitLedger adds a base set of cars to the ledger
func (s *SmartContract) InitLedgerTrans(ctx contractapi.TransactionContextInterface) error {
	//use concat
	tAgents := []TransactionAgent{
		{TransactionID: "tr01",
			AgentID:     "az01",
			Timestamp:   "001",
			Source:      "11",
			Destination: "32",
			MessageType: "21",
			Data:        "123"},
		{TransactionID: "tr02",
			AgentID:     "az02",
			Timestamp:   "32",
			Source:      "321",
			Destination: "3213",
			MessageType: "31",
			Data:        "3123"},
		{TransactionID: "tr03",
			AgentID:     "az01",
			Timestamp:   "001",
			Source:      "11",
			Destination: "32",
			MessageType: "21",
			Data:        "123"},
		{TransactionID: "tr04",
			AgentID:     "az02",
			Timestamp:   "32",
			Source:      "321",
			Destination: "3213",
			MessageType: "31",
			Data:        "3123"},
	}

	for i, tAgent := range tAgents {
		tAgentAsBytes, _ := json.Marshal(tAgent)
		err := ctx.GetStub().PutState("tAgent"+strconv.Itoa(i), tAgentAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// TRANSACTION AGENT

// Query returns the single transaction Agent stored in the world state with given id
func (s *SmartContract) QueryTransactionAgent(ctx contractapi.TransactionContextInterface, transactionID string) (*TransactionAgent, error) {
	tAgentAsBytes, err := ctx.GetStub().GetState(transactionID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if tAgentAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", transactionID)
	}

	tAgent := new(TransactionAgent)
	_ = json.Unmarshal(tAgentAsBytes, tAgent)

	return tAgent, nil
}

// Create Transsaction With Reference of Agent
func (s *SmartContract) CreateTransactionAgent(ctx contractapi.TransactionContextInterface, TransactionID string, AgentID string, Timestamp string, RefTransactionID string, Source string, Destination string, MessageType string, Data string) error {

	tAgent := TransactionAgent{
		TransactionID:    TransactionID,
		AgentID:          AgentID,
		Timestamp:        Timestamp,
		RefTransactionID: RefTransactionID,
		Source:           Source,
		Destination:      Destination,
		MessageType:      MessageType,
		Data:             Data}

	tAgentAsBytes, _ := json.Marshal(tAgent)
	err := ctx.GetStub().PutState(TransactionID, tAgentAsBytes)
	if err != nil {
		return err
	}
	//create composite key by type asset i.e trxAgent-TransactionID
	idTrxAgentKey, err := ctx.GetStub().CreateCompositeKey(index, []string{tAgent.AgentID, tAgent.TransactionID})
	if err != nil {
		return err
	}
	value := []byte{0x00}
	return ctx.GetStub().PutState(idTrxAgentKey, value)

}

// returned list of transaction
func (t *SmartContract) GetTransactionByAgent(ctx contractapi.TransactionContextInterface, AgentID string) ([]QueryResultTransaction, error) {
	// Execute a key range query on all keys starting with 'color'
	//coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{color})
	coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{AgentID})
	if err != nil {
		return nil, err
	}
	defer coloredAssetResultsIterator.Close()

	results := []QueryResultTransaction{}

	for coloredAssetResultsIterator.HasNext() {
		responseRange, err := coloredAssetResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedAssetID := compositeKeyParts[1]
			asset, err := t.QueryTransactionAgent(ctx, returnedAssetID)
			if err != nil {
				return nil, err
			}

			QueryResultTransaction := QueryResultTransaction{Key: responseRange.Key, Record: asset}
			results = append(results, QueryResultTransaction)
			/*asset.Owner = newOwner
			assetBytes, err := json.Marshal(asset)
			if err != nil {
				return err
			}
			err = ctx.GetStub().PutState(returnedAssetID, assetBytes)
			if err != nil {
				return fmt.Errorf("transfer failed for asset %s: %v", returnedAssetID, err)
			}
			*/
		}

	}

	return results, nil
}

// AGENT

func (s *SmartContract) CreateAgent(ctx contractapi.TransactionContextInterface, AgentID string, DeviceID string, SubcribePath string, TrustValue string, Tolerance string) error {

	agent := Agent{
		AgentID:      AgentID,
		DeviceID:     DeviceID,
		SubcribePath: SubcribePath,
		TrustValue:   TrustValue,
		Tolerance:    Tolerance,
	}

	agentAsBytes, _ := json.Marshal(agent)

	err := ctx.GetStub().PutState(AgentID, agentAsBytes)
	if err != nil {
		return err
	}
	//create composite key by type asset i.e deviceID-agentID
	idDeviceAgentKey, err := ctx.GetStub().CreateCompositeKey(index, []string{agent.DeviceID, agent.AgentID})
	if err != nil {
		return err
	}
	value := []byte{0x00}
	return ctx.GetStub().PutState(idDeviceAgentKey, value)
}

// QueryCar returns the car stored in the world state with given id
func (s *SmartContract) QueryAgent(ctx contractapi.TransactionContextInterface, AgentID string) (*Agent, error) {
	agentAsBytes, err := ctx.GetStub().GetState(AgentID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if agentAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", AgentID)
	}

	agent := new(Agent)
	_ = json.Unmarshal(agentAsBytes, agent)

	return agent, nil
}

func (s *SmartContract) UpdateAgent(ctx contractapi.TransactionContextInterface, AgentID string, DeviceID string, SubcribePath string, TrustValue string, Tolerance string) error {
	agent, err := s.QueryAgent(ctx, AgentID)

	if err != nil {
		return err
	}

	agent.SubcribePath = SubcribePath
	agent.Tolerance = Tolerance
	agent.TrustValue = TrustValue

	agentAsBytes, _ := json.Marshal(agent)

	return ctx.GetStub().PutState(AgentID, agentAsBytes)
}

func (t *SmartContract) GetAgentByDevice(ctx contractapi.TransactionContextInterface, DeviceID string) ([]QueryResultAgent, error) {

	coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{DeviceID})
	if err != nil {
		return nil, err
	}
	defer coloredAssetResultsIterator.Close()

	results := []QueryResultAgent{}

	for coloredAssetResultsIterator.HasNext() {
		responseRange, err := coloredAssetResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedAssetID := compositeKeyParts[1]
			asset, err := t.QueryAgent(ctx, returnedAssetID)
			if err != nil {
				return nil, err
			}

			QueryResultAgent := QueryResultAgent{Key: responseRange.Key, Record: asset}
			results = append(results, QueryResultAgent)
		}

	}

	return results, nil
}

// RATINGS

// Create Transsaction With Reference of Agent
func (s *SmartContract) CreateRatingAgent(ctx contractapi.TransactionContextInterface, RatingID string, AgentID string, AgentIDTrustee string, Value string, Reputation string) error {

	rAgent := Rating{
		RatingID:       RatingID,
		AgentID:        AgentID,
		AgentIDTrustee: AgentIDTrustee,
		Value:          Value,
		Reputation:     Reputation,
	}

	rAgentAsBytes, _ := json.Marshal(rAgent)
	err := ctx.GetStub().PutState(RatingID, rAgentAsBytes)
	if err != nil {
		return err
	}
	//create composite key by type asset i.e trxAgent-TransactionID
	idRtxAgentKey, err := ctx.GetStub().CreateCompositeKey(index, []string{rAgent.AgentID, rAgent.RatingID})
	if err != nil {
		return err
	}
	value := []byte{0x00}
	return ctx.GetStub().PutState(idRtxAgentKey, value)

}

// Query returns the single Rating Agent stored in the world state with given id
func (s *SmartContract) QueryRatingAgent(ctx contractapi.TransactionContextInterface, RatingID string) (*Rating, error) {
	rAgentAsBytes, err := ctx.GetStub().GetState(RatingID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if rAgentAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", RatingID)
	}

	rAgent := new(Rating)
	_ = json.Unmarshal(rAgentAsBytes, rAgent)

	return rAgent, nil
}

// returned list of transaction
func (t *SmartContract) GetRatingByAgent(ctx contractapi.TransactionContextInterface, AgentID string) ([]QueryResultRatings, error) {

	coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{AgentID})
	if err != nil {
		return nil, err
	}
	defer coloredAssetResultsIterator.Close()

	results := []QueryResultRatings{}

	for coloredAssetResultsIterator.HasNext() {
		responseRange, err := coloredAssetResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedAssetID := compositeKeyParts[1]
			asset, err := t.QueryRatingAgent(ctx, returnedAssetID)
			if err != nil {
				return nil, err
			}

			QueryResultRatings := QueryResultRatings{Key: responseRange.Key, Record: asset}
			results = append(results, QueryResultRatings)
			/*asset.Owner = newOwner
			assetBytes, err := json.Marshal(asset)
			if err != nil {
				return err
			}
			err = ctx.GetStub().PutState(returnedAssetID, assetBytes)
			if err != nil {
				return fmt.Errorf("transfer failed for asset %s: %v", returnedAssetID, err)
			}
			*/
		}

	}

	return results, nil
}

// EVALUATION PARAMETER

// Create Transsaction With Reference of Agent
func (s *SmartContract) CreateEvaluationAgent(ctx contractapi.TransactionContextInterface, EvaluationID string, TransactionID string, AgentID string, Timestamp string, TrusterID string, TrusteeID string, ResponseTime string, Validity string, Correctness string, Cooperation string, Qos string, Availability string, Confidence string) error {

	evaluationParam := EvaluationParam{
		EvaluationID:  EvaluationID, //uuid
		TransactionID: TransactionID,
		AgentID:       AgentID,
		TrusterID:     TrusterID,
		TrusteeID:     TrusteeID,
		Timestamp:     Timestamp,
		ResponseTime:  ResponseTime,
		Validity:      Validity,
		Correctness:   Correctness,
		Cooperation:   Cooperation,
		Qos:           Qos,
		Availability:  Availability,
		Confidence:    Confidence,
	}

	evaluationParamAsBytes, _ := json.Marshal(evaluationParam)
	err := ctx.GetStub().PutState(EvaluationID, evaluationParamAsBytes)
	if err != nil {
		return err
	}
	//create composite key by type asset i.e trxAgent-EvaluationID
	idRtxAgentKey, err := ctx.GetStub().CreateCompositeKey(index, []string{evaluationParam.AgentID, evaluationParam.EvaluationID})
	if err != nil {
		return err
	}
	value := []byte{0x00}
	err = ctx.GetStub().PutState(idRtxAgentKey, value)
	if err != nil {
		return err
	}
	idEvaTransactionKey, err := ctx.GetStub().CreateCompositeKey(index, []string{evaluationParam.TransactionID, evaluationParam.EvaluationID})
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(idEvaTransactionKey, value)
	if err != nil {
		return err
	}
	idTrustorTransactionKey, err := ctx.GetStub().CreateCompositeKey(index, []string{evaluationParam.TrusterID, evaluationParam.EvaluationID})
	if err != nil {
		return err
	}
	//value := []byte{0x00}
	err = ctx.GetStub().PutState(idTrustorTransactionKey, value)
	if err != nil {
		return err
	}
	idTrusteeTransactionKey, err := ctx.GetStub().CreateCompositeKey(index, []string{evaluationParam.TrusteeID, evaluationParam.EvaluationID})
	if err != nil {
		return err
	}
	//value := []byte{0x00}
	return ctx.GetStub().PutState(idTrusteeTransactionKey, value)

}

// Query returns the single Rating Agent stored in the world state with given id
func (s *SmartContract) QueryEvaluationAgent(ctx contractapi.TransactionContextInterface, EvaluationID string) (*EvaluationParam, error) {
	evaluationParamAsBytes, err := ctx.GetStub().GetState(EvaluationID)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if evaluationParamAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", EvaluationID)
	}

	evaluationParam := new(EvaluationParam)
	_ = json.Unmarshal(evaluationParamAsBytes, evaluationParam)

	return evaluationParam, nil
}

// returned list of transaction by agent
func (t *SmartContract) GetEvaluationByAgent(ctx contractapi.TransactionContextInterface, AgentID string) ([]QueryResultEvaluation, error) {

	coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{AgentID})
	if err != nil {
		return nil, err
	}
	defer coloredAssetResultsIterator.Close()

	results := []QueryResultEvaluation{}

	for coloredAssetResultsIterator.HasNext() {
		responseRange, err := coloredAssetResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedAssetID := compositeKeyParts[1]
			asset, err := t.QueryEvaluationAgent(ctx, returnedAssetID)
			if err != nil {
				return nil, err
			}

			QueryResultEvaluation := QueryResultEvaluation{Key: responseRange.Key, Record: asset}
			results = append(results, QueryResultEvaluation)
		}

	}

	return results, nil
}

// returned list of transaction by transaction
func (t *SmartContract) GetEvaluationByTransaction(ctx contractapi.TransactionContextInterface, TransactionID string) ([]QueryResultEvaluation, error) {

	coloredAssetResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(index, []string{TransactionID})
	if err != nil {
		return nil, err
	}
	defer coloredAssetResultsIterator.Close()

	results := []QueryResultEvaluation{}

	for coloredAssetResultsIterator.HasNext() {
		responseRange, err := coloredAssetResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedAssetID := compositeKeyParts[1]
			asset, err := t.QueryEvaluationAgent(ctx, returnedAssetID)
			if err != nil {
				return nil, err
			}

			QueryResultEvaluation := QueryResultEvaluation{Key: responseRange.Key, Record: asset}
			results = append(results, QueryResultEvaluation)
		}

	}

	return results, nil
}

// ##########################################################

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
