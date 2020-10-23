package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	Status  string
	Error   string
	Payload interface{}
}

func retryHTTPCommunication(op string, protocol string, method string, path string, targetNode *Node, payload interface{}, logMsg string, seconds int, retryCnt int, retrylimitation int) (interface{}, error) {
	if logMsg != "" {
		log.Println(logMsg)
	}
	time.Sleep(time.Second * time.Duration(seconds))
	return HTTPCommunicate(op, protocol, method, path, targetNode, payload, retryCnt+1, retrylimitation)
}

// HTTPCommunicate access target node, if retry >= 0, then enable retry mode, if retry < 0, then disable retry mode
func HTTPCommunicate(operationName string, protocol string, method string, path string, targetNode *Node, payload interface{}, retryCnt int, retryLimitation int) (interface{}, error) {
	if targetNode == nil || (strings.ToLower(protocol) != "http" && strings.ToLower(protocol) != "https") {
		msg := "ERROR: invalid target node for " + operationName
		log.Println(msg)
		return nil, errors.New(msg)
	}
	if retryCnt > retryLimitation {
		msg := "Retried maximum times: " + strconv.Itoa(retryCnt) + ", will no longer retry " + operationName
		log.Println(msg)
		return nil, errors.New(msg)
	}

	tailstr := ""
	if retryCnt > 0 {
		tailstr = ", retry count: " + strconv.Itoa(retryCnt)
	}
	log.Println("[comm] "+operationName+" started toward target node:", targetNode.Key(), tailstr)

	reqbody, err := json.Marshal(payload)
	if err != nil {
		msg := "ERROR during encoding payload: " + err.Error() + ", will retry in 30 seconds"
		return retryHTTPCommunication(operationName, protocol, method, path, targetNode, payload, msg, 30, retryCnt+1, retryLimitation)
	}
	targetURL := protocol + "://" + targetNode.Key() + path
	// Send the register information to upper node
	req, err := http.NewRequest(strings.ToUpper(method), targetURL, bytes.NewBuffer(reqbody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		msg := "Error when sending http request, will retry after 30 seconds: " + err.Error()
		return retryHTTPCommunication(operationName, protocol, method, path, targetNode, payload, msg, 30, retryCnt+1, retryLimitation)
	}

	if res == nil || res.Body == nil {
		msg := "Error of the communication for the response is nil, will retry after 30 seconds"
		return retryHTTPCommunication(operationName, protocol, method, path, targetNode, payload, msg, 30, retryCnt+1, retryLimitation)
	}

	// parse the response message of upper node for registering
	resMsg := Response{}
	json.NewDecoder(res.Body).Decode(&resMsg)
	if resMsg.Error != "" {
		msg := "target node responded ERROR message: " + resMsg.Error + ", will retry " + operationName + " registering in 30 seconds"
		return retryHTTPCommunication(operationName, protocol, method, path, targetNode, payload, msg, 30, retryCnt+1, retryLimitation)
	} else {
		if resMsg.Status == "OK" {
			log.Println("[comm] "+operationName+" complete with target node:", targetNode.Key())
			return resMsg.Payload, nil
		} else {
			msg := "target node responded abnormal status: " + resMsg.Status + ", will retry " + operationName + " in 30 seconds"
			return retryHTTPCommunication(operationName, protocol, method, path, targetNode, payload, msg, 30, retryCnt+1, retryLimitation)
		}
	}
}

func DecodeRequestBody(r *rest.Request, v interface{}) ([]byte, error) {
	content, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		return nil, err
	}
	if len(content) == 0 {
		return nil, errors.New("JSON payload is empty")
	}
	if v != nil {
		err = json.Unmarshal(content, v)
		if err != nil {
			return nil, err
		}
	}
	return content, nil
}
