# k8s test

This is my tiny docker webserver container that serves me to test kubernetes.

* It's built using scratch
* The web server listen on env variable K8STEST_PORT
* It shows container/host name
* It shows network interfaces and the associated IP addresses

You can use the container at `pmercado/k8stest`
You can also use the pod and service at the kubernetes folder
