## Ephemeral Pull Request Environments

## GKE Cluster
- Create GKE cluster: https://cloud.google.com/kubernetes-engine/docs/deploy-app-cluster
- Create a Google cloud service account and download its JSON. 
- Install and authenticate into Gcloud cli: https://cloud.google.com/sdk/docs/install

## Argo CD 

Install ArgoCD on the cluster:

```
kubectl create namespace argocd

kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

```
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```

```
argocd admin initial-password -n argocd
```

ArgoCD application set

```
kubectl apply -n argocd -f argocd/application-set.yaml
```

## Crossplane

Install crossplane on the cluster:

```
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

# Install Crossplane into crossplane-system namespace
helm install crossplane \
  crossplane-stable/crossplane \
  --namespace crossplane-system \
  --create-namespace \
  --set "provider.packages[0]=xpkg.upbound.io/upbound/provider-helm:v0.20.3" \
  --set "provider.packages[1]=xpkg.upbound.io/upbound/provider-kubernetes:v0.16.2"
```

```
kubectl get pods -n crossplane-system
kubectl get providers
kubectl -n crossplane-system get sa -o name
```

```
SA=$(kubectl -n crossplane-system get sa -o name | grep provider-helm | sed -e 's|serviceaccount\/|crossplane-system:|g')
kubectl create clusterrolebinding provider-helm-admin-binding --clusterrole cluster-admin --serviceaccount="${SA}"
```

```
kubectl apply -f crossplane-vcluster/helm-provider-config.yaml
kubectl apply -f crossplane-vcluster/composition.yaml
kubectl apply -f crossplane-vcluster/environment-resource-definition.yaml
kubectl apply -f crossplane-vcluster/environment-resource.yaml
```

```
vcluster list
```