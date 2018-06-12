# parcel-vue
零配置VUE开发模板，基于Parcel全新构建

[传送门：Parcel-VUE官网文档](http://parcel.xserver.top)

> 使用说明

```bash
请使用Node版本 >= 9.10
```

```bash
$ npm install xserver-cli -g

$ x-cli parcel-vue myproject
```

> 开发调试
```bash
$ npm run dev
```

> 正式构建
```bash
$ npm run build
```

>目录结构
```js
├── dist                构建结果
├── index.html          首页文件
├── node_modules
├── package.json
├── src                 开发源码
│   ├── App.vue
│   ├── assets
│   ├── components
│   ├── main.js
│   └── router
└── static              静态文件
    ├── css
    ├── img
    └── js
```

更新日志
>
	2018.01.28:更新iview至2.9.0版本
	2018.03.31:更新pacel至1.7.0版本，原生支持VUE构建，进一步精简配置，完全零配置

帮助联系
>
	作者:cheneyxu
	邮箱:457299596@qq.com
	QQ:457299596
