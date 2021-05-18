# k8s-Management
# k8s管理平台
# kubeadmin是后端接口接口代码，kube-front是前端代码
使用方式  
## go 1.11+ 如下
git clone xxxxxx  
export GO111MODULE="on"  
cp -fr  k8s-Management/kubeadmin/src/kubeadmin $GOROOT/src/  
go mod init  
上述命令执行后可能会后后续提示操作如下：  
Example usage:  
	'go mod init example.com/m' to initialize a v0 or v1 module  
	'go mod init example.com/m/v2' to initialize a v2 module  
执行：  
go mod init example.com/m/v2  
go mod tidy  
然后执行run  
go run $GOROOT/src/kubeadmin/main/main.go
如果出现依赖包没有根据提示go get xxxx下载即可  
后端默认端口：8080，可修改main文件重新编译，config默认从/root/.kube/下读取  
前端:  
npm build   
然后把dist目录部署到NGINX下即可  
