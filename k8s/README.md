setup
```
k context get-contexts
# switch context for local
k context use-context docker-desktop

# install ingress to kubenets
# https://kubernetes.github.io/ingress-nginx/deploy/
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.1/deploy/static/provider/cloud/deploy.yaml
```

deployment
```
# confirm
k get all --namespace default

k apply -f file
# delete
k delete -f file
```
