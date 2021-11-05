package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/controller"
	"kubernetesa-admin/global/views"
	"net/http"
)

func JSONAppErrorReporter() gin.HandlerFunc {
	return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)
		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			var parsedError *views.ApiError
			switch err.(type) {
			//如果产生的error是自定义的结构体,转换error,返回自定义的code和msg
			case *views.ApiError:
				parsedError = err.(*views.ApiError)
			default:
				parsedError = &views.ApiError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
			}
			c.IndentedJSON(parsedError.Code, parsedError)
			return
		}

	}
}
func Route(server *gin.Engine)  {
	server.Use(JSONAppErrorReporter())
	server.Use(cors.Default())
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK,"pong")
	})
	api := server.Group("/api/v1")

	api.GET("/nodes",controller.GetNodeList)
	api.GET("/namespaces",controller.GetNamespaceList)
	api.GET("/deployments",controller.GetDeploymentList)
	api.GET("/statefulset",controller.GetStatefulSetList)
	api.GET("/jobs",controller.GetJobList)
	api.GET("/crontabjobs",controller.GetCrontabJobList)
	api.GET("/pods",controller.GetPodList)
	api.GET("/services",controller.GetServicesList)
	api.GET("/ingress",controller.GetIngressList)
	api.GET("/configmaps",controller.GetConfigMapList)
	api.GET("/secrets",controller.GetSecretList)
	api.GET("/pvcs",controller.GetPersistentVolumeClaimList)
	api.GET("/pvs",controller.GetPersistentVolumeList)
	api.GET("/StorageClass",controller.GetStorageClassList)
	api.GET("/node/exec",controller.SSH)
	api.GET("/getcluster",controller.GetCluster)
	api.GET("/getpodinfo",controller.GetPodInfo)
	api.GET("/getpodlogs",controller.GetPodlog)
	api.GET("/podexec",controller.PodExec)
	api.POST("/upload",controller.Upload)
}
