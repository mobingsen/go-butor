package Informer

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"time"
)

/**
kubectl get --raw /
kubectl get --raw /apis/apps/v1
kubectl proxy
kubectl get statefulsets.apps
*/
func main() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
		return
	}
	clientSet, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic(err)
		return
	}
	informerFactory := informers.NewSharedInformerFactory(clientSet, 30*time.Second)
	deploymentInformer := informerFactory.Apps().V1().Deployments()
	informer := deploymentInformer.Informer()
	deploymentLister := deploymentInformer.Lister()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})
	stopper := make(chan struct{})
	defer close(stopper)
	// 启动Informer list & watch
	informerFactory.Start(stopper)
	// 等待所有的Informer缓存同步
	informerFactory.WaitForCacheSync(stopper)
	deployments, err := deploymentLister.Deployments("default").List(labels.Everything())
	// 编辑deploy列表
	for index, deploy := range deployments {
		fmt.Printf("%d -> %s\n", index, deploy.Name)
	}
	<-stopper
}

func onAdd(obj interface{}) {
	deploy := obj.(*v1.Deployment)
	klog.Infoln("add a deploy: ", deploy.Name)
}

func onUpdate(old, new interface{}) {
	oldDeploy := old.(*v1.Deployment)
	newDeploy := new.(*v1.Deployment)
	klog.Infoln("update deploy: ", oldDeploy.Status.Replicas, newDeploy.Status.Replicas)
}

func onDelete(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	klog.Infoln("delete a deploy: ", deployment.Name)
}
