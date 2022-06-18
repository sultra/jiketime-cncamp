## statefulset 部署etcd

`sts.yaml` 中 `volumeClaimTemplates` 部分是通过模版自动创建 pvc，但是前提是有 StorageClass
 https://kubernetes.io/zh/docs/concepts/storage/storage-classes/

 ### 使用本地存储 StorageClass 

+ `sc.yaml` 创建本地 StorageClass

配置中 etcd 是3节点，使用本地存储需要手动创建3个pv和3个pvc

+ 创建3个pv

先在各主机节点创建 `/mnt/data` 目录

创建一个后再修改 `pv.yaml` 中

``` yaml
metadata:
  # pv 的名字
  name: pv-loc-sc-2 

nodeAffinity:
    ……
          # K8s节点名称 （hostname）
          values:
          - k8s3
```

+ 创建3个pvc 

创建一个后再修改 `pvc.yaml`

``` yaml
metadata:
  # pvc 的名字
  name: datadir-etcd-2
  labels:
    kubeless: etcd
spec:
  # pv 名称
  volumeName: pv-loc-sc-2
```

pvc 名称需要根据 etcd `sts.yaml` 中 name 组合，规则是 `(volumeClaimTemplates.[metadata].name)-(metadata.name)-编号`

做好以上步骤可以在正常部署 sts.yaml , 自动创建pvc 需要其他统一存储的 StorageClass 插件

由于 pv 配置中设定了不同节点，所以 etcd 启动后每个节点会被调度到各K8s对应节点上