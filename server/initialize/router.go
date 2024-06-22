package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/organization"
	swaggerFiles "github.com/swaggo/files"
	"net/http"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	InstallPlugin(Router)
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	//Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.GVA_CONFIG.Local.StorePath)})

	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)

	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)
		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)

	}
	PluginInit(PublicGroup, organization.CreateOrganizationPlug())
	{
		alphaRouter := router.RouterGroupApp.Alpha
		alphaRouter.InitMdUnitMeasureRouter(PrivateGroup)
		alphaRouter.InitMdClientRouter(PrivateGroup)
		alphaRouter.InitCMSMERouter(PrivateGroup)
		alphaRouter.InitCMSXCRouter(PrivateGroup)
		alphaRouter.InitItemTypeRouter(PrivateGroup)
		alphaRouter.InitMutiSelectRouter(PrivateGroup)
		alphaRouter.InitMdFirstLevelRouter(PrivateGroup)
		alphaRouter.InitMdSecondLevelRouter(PrivateGroup)
		alphaRouter.InitMdThirdLevelRouter(PrivateGroup)
		alphaRouter.InitTecBaseInfoRouter(PrivateGroup)

		alphaRouter.InitTecBaseInfoExtRouter(PrivateGroup)
		alphaRouter.InitQualityBaseInfoRouter(PrivateGroup)
		alphaRouter.InitCostCollectionRouter(PrivateGroup)
		alphaRouter.InitProofingInformationRouter(PrivateGroup)
		//不鉴权
		alphaRouter.InitEntryNumberRouter(PublicGroup)
		alphaRouter.InitTecSNRuleRouter(PublicGroup)
		alphaRouter.InitProcessFileInformationRouter(PrivateGroup, PublicGroup)
		alphaRouter.InitTestFileAndImgRouter(PrivateGroup, PublicGroup)
		alphaRouter.
			InitMassProductionTransferRouter(PrivateGroup, PublicGroup)
		alphaRouter.InitEngineeringChangeRouter(PrivateGroup, PublicGroup)
		alphaRouter.InitPackageRequirementRouter(PrivateGroup, PublicGroup)
		alphaRouter.InitTecBaseProcessRouter(PrivateGroup, PublicGroup)

	}

	global.GVA_LOG.Info("router register success")
	return Router
}
