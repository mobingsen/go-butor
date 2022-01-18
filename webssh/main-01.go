package main

// go get k8s.io/client-go
// go mod tidy  命令会把所有依赖都补上
//func main() {
//	var kubeConfig *string
//	ctx := context.Background()
//	if home := homedir2.HomeDir(); home != "" {
//		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to kubeConfig file")
//	} else {
//		kubeConfig = flag.String("kubeConfig", "", "absolute path to kubeConfig file")
//	}
//	flag.Parse()
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
//	if err != nil {
//		klog.Fatal(err)
//		return
//	}
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		klog.Fatal(err)
//		return
//	}
//	namespaceList, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
//	if err != nil {
//		klog.Fatal(err)
//		return
//	}
//	namespaces := namespaceList.Items
//	for _, namespace := range namespaces {
//		fmt.Println("name: " + namespace.Name + ", status: " + string(namespace.Status.Phase))
//	}
//	//podList, err := clientset.CoreV1().Pods("default").List()
//}
