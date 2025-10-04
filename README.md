# openshift-training-operator

## Install golang
```
wget https://go.dev/dl/go1.21.2.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

## Install operator SDK
```
export RELEASE_VERSION=v1.34.1
curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk_linux_amd64
sudo install -m 755 operator-sdk_linux_amd64 /usr/local/bin/operator-sdk
sudo apt install -y make
```

## Create a training operator
```
operator-sdk init --domain tektutor.org --repo github.com/tektutor/training-operator
operator-sdk create api --group training --version v1 --kind Training --resource --controller

export PATH=$PATH:$(go env GOPATH)/bin
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
controller-gen --version

go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.7.0
mkdir -p /root/openshift-training-operator/bin
ln -s $(which controller-gen) /root/openshift-training-operator/bin/controller-gen
# Install latest stable Kustomize (v5.x is fine)
go install sigs.k8s.io/kustomize/kustomize/v4@v4.5.7

# Verify installation
kustomize version

# Create bin directory inside your project if not exists
mkdir -p /root/openshift-training-operator/bin

# Symlink your global kustomize to the project bin
ln -s $(which kustomize) /root/openshift-training-operator/bin/kustomize


make install
kubectl get crd | grep training
make deploy
kubectl get pods -n openshift-custom-operator-system
kubectl logs -n openshift-custom-operator-system deploy/openshift-custom-operator-controller-manager -c manager
kubectl apply -f config/samples/training_v1_training.yaml
kubectl get trainings.training.tektutor.org
kubectl describe training go-operator-training
kubectl logs -n openshift-custom-operator deploy/openshift-custom-operator-controller-manager
```
