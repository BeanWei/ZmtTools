# 领域新闻爬取

## 凤凰娱乐
  * http://ent.ifeng.com/listpage/3/1/list.shtml  

## 一点资讯(暂不考虑)
  * http://www.yidianzixun.com/channel/c3   不能获取最新的，显示时间与发文时间不对应

## 微博


## 其他新闻聚合站点
  * ZAKER新闻
    - http://app.myzaker.com/index.php?app_id= 
      - app_id = 0  热点
      - 1       国内
      - 2       国际
      - 3       军事
      - 4       财经
      - 5       互联网
      - 6       房地产
      - 7       汽车
      - 8       体育
      - 9       娱乐
      - 10      游戏频道
      - 11      教育
      - 12      时尚
      - 13      科技
      - 14      社会
      - 15      影视
      - 959     亲子
      - 981     旅游
      - 1014    星座
      - 1039    科学
      - 1067    奢侈品
      - 10376   游戏
      - 10386   美食
      - 10530   电影
      - 10802   健康
      - 11195   理财

  * 新闻头条
    - http://toutiao.jxnews.com.cn/m/channelnewslist.php?from=null&cate=48&pagesize=18&page=0&lasttime=0
    - cate=48 娱乐频道   pagesize=18 获取数量  ,   get请求 返回值为json
    - 返回值
      ```
        id	2363445
        classname	娱乐
        titleurl	/p/20180701/2363445.html   //目标url：http://toutiao.jxnews.com.cn/p/20180701/2363445.html
        newstime	30分钟前
        timestamp	1530437619
        cateid	48
        title	她比汤唯还敬业拍戏不用替身亲自哺乳却遭到谩骂被逼退出演艺圈
        titlepic	http://ylimg.ufile.ucloud.com.cn/20180701/2018070117331293878907.jpg
        titlepic2	http://ylimg.ufile.ucloud.com.cn/20180701/2018070117331293842840.jpg
        titlepic3	http://ylimg.ufile.ucloud.com.cn/20180701/2018070117331293811072.jpg
        befrom	期货张宁
        smalltext	在娱乐圈中女明星们想要成名可以说比登天还要难，除了自己的先天条件要过硬之外，还要有精湛的演技和敬业的精神，不仅要敬业
      ```



# 热搜实时监控
  * 搜狗: http://top.sogou.com/hot/shishi_1.html



