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

## Crossplane

Install crossplane on the cluster:

```
helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

# Install Crossplane into crossplane-system namespace
helm install crossplane \
  crossplane-stable/crossplane \
  --namespace crossplane-system \
  --create-namespace
```

I also followed the steps to configure secrets, provider etc. based on this guide (not sure if all steps are actually needed): https://docs.crossplane.io/latest/getting-started/provider-gcp/

Also, install crossplane CLI: https://docs.crossplane.io/latest/cli/

## Vcluster

1. Install Vcluster CLI: https://www.vcluster.com/docs/v0.19/getting-started/setup

2. Sign up for Vcluster Cloud (ie. vcluster Platform): https://www.vcluster.com/install and create a new vcluster Platform instance

3. Create a vcluster access key using Profile > Access key. Keep it handy

4. using the CLI, login to Vcluster platform:

```
vcluster platform login https://ephemeral-pr-vcluster.vcluster.cloud/ --access-key <YOUR_ACCESS_KEY>
```

5. Connect your host cluster (GKE cluster) with the vcluster platform:

```
https://cloud.google.com/sdk/docs/install
```

6. Create a test vcluster using the CLI:

```
vcluster create ephemeral-pr-env-vcluster
```

Open questions:
- I am using vcluster platform because that's what the [Youtube video](https://www.youtube.com/watch?v=j7ZMqzsse9c&ab_channel=PlatformEngineering) also uses. 

