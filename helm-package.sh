helm package charts/go-app
rm go-app-helm-chart/*.tgz
mv go-app-helm-chart-1.0.0.tgz go-app-helm-chart/
helm repo index go-app-helm-chart --url https://go-app-helm-chart.storage.googleapis.com