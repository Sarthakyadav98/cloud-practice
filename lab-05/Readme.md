minikube start
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl get pods
minikube service my-service