### 练习云原生devops

1. minikube start
2. helm install jenkins jenkins/jenkins
3. kubectl exec -ti po/jenkins-0 -c jenkins
4. kubectl --namespace default port-forward svc/jenkins 8080:8080
5. make install
6. make run
7. kubectl apply -f config/samples/devops_v1alpha1_pipeline.yaml
8. kubectl patch pipeline/test-pipeline -n test --patch '{"spec": "action": "run"}'