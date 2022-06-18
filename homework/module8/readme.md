1. 镜像使用阿里云私有仓库，镜像secert 文件：pullImageSecret.yaml
  pullImageSecret.yaml 中密文为假值
2. deploy 包括 readinessProbe，startupProbe 探针，startupProbe 使用了读取 configMap 中配置的环境变量
3. configMap 测试成功环境变量和读取挂载的配置文件(/etc/cm01/db.config)