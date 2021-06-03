# k8s-Management
# k8s管理平台
# kubeadmin是后端接口接口代码，kube-front是前端代码  
kubeadmin 的 prod下面init读取生产的config路径，controller的init文件读取的路径是测试或者研发环境  

## 使用方式  
## go 1.11+ 如下
git clone xxxxxx  
export GO111MODULE="on"  
cp -fr  k8s-Management/kubeadmin/src/kubeadmin $GOROOT/src/  
cd $GOROOT/src/
go mod init  
上述命令执行后可能会后后续提示操作如下：  
Example usage:  
	'go mod init example.com/m' to initialize a v0 or v1 module  
	'go mod init example.com/m/v2' to initialize a v2 module  
执行：  
go mod init example.com/m/v2  
go mod tidy  
然后执行run  
cd $GOROOT/src/kubeadmin/main/  
go build -o xxxx main.go  
如果出现依赖包没有,根据提示go get xxxx下载后重新执行go build -o xxxx main.go即可  
config.yaml为配置文件模板
执行下面启动服务：  
./xxx -f ./config.yaml    
后端默认端口：8080，  
nginx代理配置：   
location / {    
        root   /usr/share/nginx/html/;    
        index  index.html index.htm;    
        try_files $uri $uri/ /index.html;    
    }    
    location /api {    
        proxy_read_timeout 300s;    
                                     
      proxy_set_header Host $host;  
      proxy_set_header X-Real-IP $remote_addr;  
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;  
                                       
      proxy_http_version 1.1;  
      proxy_pass http://127.0.0.1:8080;  
    }  
    location /api/getpodlog {  
        proxy_read_timeout 300s;  
                                      
      proxy_set_header Host $host;  
      proxy_set_header X-Real-IP $remote_addr;  
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;  
      proxy_http_version 1.1;  
      proxy_set_header Upgrade $http_upgrade;  
      proxy_set_header Connection $connection_upgrade;  
        proxy_pass http://127.0.0.1:8080;  
    }  
前端:  
需要先安装npm这里请自行百度安装  
安装npm之后为了下载依赖包方便安装cnpm
npm install -g cnpm --registry=https://registry.npm.taobao.org
安装cnpm之后执行下面命令，下载需要的依赖包
cnpm install
这里在编译前端项目之前请根据情况自行修改下面文件后进行编译  
store.js还有views/Log.vue,这两个文件里面配置了连接后端的ip和端口，这里根据情况自行修改，用ngixn代理后端的修改为NGINX地址
npm build   
然后把dist目录部署到NGINX下即可  
