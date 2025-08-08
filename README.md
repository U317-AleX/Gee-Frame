gee-frame
一个基于 Trie 树实现的、轻量级的类 Gin Web 框架

gee-frame 是一个从零开始构建的 Go 语言 Web 框架。它的设计灵感来源于著名的 Gin 框架，适用于构建 RESTful API 服务。

✨ 主要特性
动态路由：采用**Trie 树（前缀树）**实现动态路由，支持 :param 参数匹配。

路由分组：提供灵活的**路由分组（Group）**功能，方便管理和组织相关的路由。

中间件（Middleware）：支持为整个框架、路由组注册中间件，可以轻松实现日志记录、权限验证、全局异常捕获等功能。

类 Gin 的 Context：提供一个功能丰富的 Context 对象，封装了请求（Request）和响应（Response）对象，简化了参数获取、数据绑定、JSON/HTML 响应等操作。

丰富的内置中间件：提供开箱即用的中间件，如：

Logger：记录请求的详细信息，包括方法、路径、状态码和耗时。

Recovery：全局异常捕获中间件，在发生 panic 时恢复程序并返回友好的错误信息，避免服务崩溃。

HTML 模板渲染：支持通过 html/template 包进行 HTML 模板渲染，方便构建 Web 页面。
