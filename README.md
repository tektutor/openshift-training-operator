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
```

## Create a training operator
```
operator-sdk init --domain tektutor.org --repo github.com/tektutor/training-operator
operator-sdk create api --group training --version v1 --kind Training --resource --controller
make install

```
