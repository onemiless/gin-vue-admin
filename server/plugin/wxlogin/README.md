## GVA 微信公众号登录（绑定）插件

### 附带基础接口和微信回调功能可扩展

#### 开发者：Mr.奇淼

### 插件功能流程（图）

![插件功能流程（图）](https://qmplusimg.henrongyi.top/wxload.png)

### 使用步骤（后端）

#### 1. 前往GVA主程序下的initialize/plugin.go 下注册插件

	PluginInit(PublicGroup, wxlogin.CreateWxLoginPlug(
		AppID,
		Appsecret,
		Token,
		AuthorityID))
        // 前端已有对应调用逻辑，无需任何改动，只要填充配置皆可使用

### 2. 配置说明

#### 2-1 全局配置结构体说明

    type WxGlobanConfig struct {
        AppID       string // 微信AppID
        Appsecret   string // 微信secret
        AccessToken string // 会自动刷新计算 不用管
        Token       string // 自定义的安全码
        AuthorityID string // 微信扫码注册角色的ID（必须是已存在角色）
    }



### 3. 可直接调用的接口

    微信绑定接口（用于微信url token验证）： /wxLogin/pingpang [get] 
    微信事件回调接口（扫码 关注 取消关注等）： /wxLogin/pingpang [post] 
    获取暂时二维码用户扫码登录（引导关注）： /wxLogin/getLoginPic [get] 
    二维码轮询接口，下一步事件前端控制： /wxLogin/loginOrCreate [post] 
    二维码扫码后的登录（自动绑定微信）： /wxLogin/register [post] 
    
### 使用步骤（前端）
    
    前端页面和方法均在web/plugin/wx下 components 为扫码登录绑定按钮 和 解除绑定按钮
    在需要使用的页面引入此组件即可使用扫码登录 扫码绑定  和 取消绑定功能（可根据业务自行将按钮拆分或组合）
    /api 文件夹为和后端的交互
    /utils 文件夹提供了一个随机字符产生方法
    
    需要使用的页面 引入 组件 import wxLogin from '@/plugin/wx/components/bind.vue' 然后页面中 <wxLogin /> 即可
    需要使用的页面 引入 组件 import removeWx from '@/plugin/wx/components/remove.vue' 然后页面中 <removeWx /> 即可


### 微信相关注意事项

    1.需要有已认证微信公众号（或者https://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo 申请测试号）
    2.获取对应的 appID appsecret 填写 Token后 将url指向服务所在的公网可访问的url（线上云服务器或者本地内网穿透）
    全url为 你的url/wxLogin/pingpang  
    如果是指向前端url 为 你的url/api/wxLogin/pingpang 
    3.线上发布请一定记得添加ip白名单，否则你将无法获取accessToken