package service

import (
	"context"
	"errors"
	"github.com/gorilla/websocket"
	"go-topics/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"sync"
)

func GetPods(namespaceName string) ([]v1.Pod, error) {
	ctx := context.Background()
	clientset, err := client.GetK8sClientSet()
	if err != nil {
		klog.Fatal(err)
		return nil, err
	}
	list, err := clientset.CoreV1().Pods(namespaceName).List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
		return nil, err
	}
	return list.Items, nil
}

type WsMessage struct {
	MessageType int
	Data        []byte
}

type WsConnection struct {
	wsSocket  *websocket.Conn
	inChan    chan *WsMessage
	outChan   chan *WsMessage
	mutex     sync.Mutex
	isClosed  bool
	closeChan chan byte
}

func (wsConn *WsConnection) wsWrite(messageType int, data []byte) (err error) {
	select {
	case wsConn.outChan <- &WsMessage{MessageType: messageType, Data: data}:
		return nil
	case <-wsConn.closeChan:
		err = errors.New("websocket closed")
		return err
	}
}
