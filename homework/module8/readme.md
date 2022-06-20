1. 镜像使用阿里云私有仓库，镜像secert 文件：pullImageSecret.yaml
  pullImageSecret.yaml 中密文为假值
2. deploy 包括 readinessProbe，startupProbe 探针，startupProbe 使用了读取 configMap 中配置的环境变量
3. configMap 测试成功环境变量和读取挂载的配置文件(/etc/cm01/db.config)

#### 创建证书


+ key csr  101 文档中生成crt 命令提示错误 `unknown option -addext` ，该用 cfssl 

``` shell

# 生成 key 和 csr
./cfssl genkey cfssl.json | ./cfssljson -bare server

kubectl create -f ./csr.yaml

``` 

+ 签发证书

``` shell

# 获取crt 
kubectl get csr httpsvc.default -o jsonpath='{.spec.request}' | base64 --decode > tls.crt 

kubectl create secret generic cncamp-tls --from-file=tls.crt --from-file=server-key.pem  

```

#### 证书绑定ingress

``` shell

kubectl create -f ingress.yaml

```