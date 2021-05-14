# k8s-Management
# k8s管理平台
# kubeadmin是后端接口接口代码，kube-front是前端代码
使用方式
git clone xxxxxx
go build k8s-Management/kubeadmin/src/kubeadmin/main/main.go
后端默认端口：8080，可修改main文件重新编译，config默认从/root/.kube/下读取
前端：
npm build 
然后把dist目录部署到NGINX下即可
