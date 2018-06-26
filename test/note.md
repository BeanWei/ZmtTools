# 领域新闻爬取

## 凤凰娱乐
  * http://ent.ifeng.com/listpage/3/1/list.shtml  

## 一点资讯(暂不考虑)
  * http://www.yidianzixun.com/channel/c3   不能获取最新的，显示时间与发文时间不对应

## 微博


## 其他娱乐站点




# 热搜实时监控
  * 搜狗: http://top.sogou.com/hot/shishi_1.html



# 乐观号
  * http://www.myleguan.com/lgEditor/lgEditor.html#/ESArticle

  ## 十万爆文
  * url: http://www.myleguan.com/lg_res/focus/flr/la
  * method: POST
  * POSTDATA: 
              lgCustomerId	1000186598
              lptime	1529904816
              page	1
              read	0
              readOrder	DESC
              reqTime	1529991215035
              size	10
              status	1
      - lgCustomerId获取:  http://www.myleguan.com/lgh_uc/user/base/grant_visitor
      - method: POST
      - POSTDATA: reqTime	1529991210584
      - 响应: {"code":1,"num":1,"msg":"OK","reObj":{"userId":1000186598}}
  * 响应:
        {"restNum":null,"code":1,"num":1,"msg":"成功返回文章列表","reObj":"{\"content\":[{\"id\":\"5b317eaae1382338df484eee\",\"title\":\"美媒指责：马哈蒂尔“甚至没有站在美国军方自由航行的立场”\",\"url\":\"http://toutiao.com/a6570973196156142093\",\"platform\":\"今日头条\",\"account\":\"4541998456\",\"domain\":\"生活\",\"status\":1,\"read\":2185186,\"like\":13,\"comment\":0,\"favor\":0,\"origin\":1,\"mediaType\":2,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 18:51:58\",\"createTime\":\"2018-06-26 07:45:44\",\"updateTime\":\"2018-06-26 07:45:44\",\"lptime\":1529923918,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a6a10263817832ca4d6d178\",\"app\":0,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p3.pstatp.com/medium/97d001136e21733e22d\",\"nickName\":\"火星方阵\"},{\"id\":\"5b30f6a7e1382377b71b775d\",\"title\":\"美团赴港上市 招股书披露\",\"url\":\"http://www.sohu.com/a/237589672_115565\",\"platform\":\"搜狐号\",\"account\":\"a2VqaXF1YW5qaWFud2VuQHNvaHUuY29t=\",\"domain\":\"科技\",\"status\":1,\"read\":2162565,\"like\":0,\"comment\":3,\"favor\":0,\"origin\":0,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 16:53:47\",\"createTime\":\"2018-06-25 22:04:16\",\"updateTime\":\"2018-06-25 22:05:28\",\"lptime\":1529916827,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a8e58a63817832c30965280\",\"app\":1,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://sucimg.itc.cn/avatarimg/98548f5ec9444377986a3281d6737476_1486691657080\",\"nickName\":\"搜狐科技视界\"},{\"id\":\"5b31864f6d03b66b4bfa9a48\",\"title\":\"湖北监利通报：电力施工人员触电1死1伤，原因正调查\",\"url\":\"http://www.sohu.com/a/237752844_260616\",\"platform\":\"搜狐号\",\"account\":\"c29odXptdGZtNmkzaWlAc29odS5jb20=\",\"domain\":\"教育\",\"status\":1,\"read\":2062085,\"like\":0,\"comment\":0,\"favor\":0,\"origin\":0,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 23:52:38\",\"createTime\":\"2018-06-26 08:17:45\",\"updateTime\":\"2018-06-26 08:21:05\",\"lptime\":1529941958,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a8e58a03817832c309651b7\",\"app\":1,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://sucimg.itc.cn/avatarimg/ece34a261c7147b0864ef095bf1c02d8_1506051877694\",\"nickName\":\"澎湃新闻\"},{\"id\":\"5b30ed6fe138236fac4ffb14\",\"title\":\"19岁少女跳楼自杀，每个起哄者都是凶手\",\"url\":\"http://toutiao.com/a6571001589849391623\",\"platform\":\"今日头条\",\"account\":\"5140856177\",\"domain\":\"职场\",\"status\":1,\"read\":1825443,\"like\":4546,\"comment\":49727,\"favor\":0,\"origin\":1,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 20:42:09\",\"createTime\":\"2018-06-25 21:26:05\",\"updateTime\":\"2018-06-25 21:26:05\",\"lptime\":1529930529,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a8a972438178303fcec17f5\",\"app\":0,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p3.pstatp.com/medium/1480/7186611868\",\"nickName\":\"中国新闻周刊\"},{\"id\":\"5b3166d2e138236ac5695609\",\"title\":\"胡春华任总指挥的指挥部，组成人员公布\",\"url\":\"http://toutiao.com/a6570965400845025795\",\"platform\":\"今日头条\",\"account\":\"4644596221\",\"domain\":\"时政\",\"status\":1,\"read\":1705782,\"like\":563,\"comment\":2,\"favor\":0,\"origin\":1,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 18:21:43\",\"createTime\":\"2018-06-26 06:04:00\",\"updateTime\":\"2018-06-26 06:04:00\",\"lptime\":1529922103,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a6a10263817832ca4d6ca4b\",\"app\":0,\"appOnline\":1,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p9.pstatp.com/medium/5677/4633951464\",\"nickName\":\"新京报政事儿\"},{\"id\":\"5b312e40e138232624286ab5\",\"title\":\"20多辆“宾利”多辆“劳斯莱斯”10多辆“哈雷”的豪华婚车队被查\",\"url\":\"http://toutiao.com/a6570912671896437255\",\"platform\":\"今日头条\",\"account\":\"5800672047\",\"domain\":\"生活\",\"status\":1,\"read\":1675489,\"like\":4661,\"comment\":14755,\"favor\":0,\"origin\":1,\"mediaType\":2,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 14:57:06\",\"createTime\":\"2018-06-26 02:02:38\",\"updateTime\":\"2018-06-26 02:02:38\",\"lptime\":1529909826,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a8a96c538178303fcec08ad\",\"app\":0,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p1.pstatp.com/medium/888f0000eb4b5b7ac157\",\"nickName\":\"新闻大连\"},{\"id\":\"5b31318f6d03b64aa7e66dcc\",\"title\":\"中国男篮红队确定打NBA夏联 赛程未定期待惊喜\",\"url\":\"http://www.sohu.com/a/237657224_458722\",\"platform\":\"搜狐号\",\"account\":\"cHBhZzU4NjJkMmIzNGIzZUBzb2h1LmNvbQ=\",\"domain\":\"体育\",\"status\":1,\"read\":1563242,\"like\":0,\"comment\":30,\"favor\":0,\"origin\":0,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 15:01:51\",\"createTime\":\"2018-06-26 02:16:15\",\"updateTime\":\"2018-06-26 02:17:53\",\"lptime\":1529910111,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a8e593c3817832c30966508\",\"app\":1,\"appOnline\":1,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://sucimg.itc.cn/avatarimg/130c0965ce51493ca5108d2271b652e4_1517370199108\",\"nickName\":\"NBA广角\"},{\"id\":\"5b3116b2e13823668d49e96c\",\"title\":\"中纪委原副书记出任“扫黑钦差”\",\"url\":\"http://toutiao.com/a6570915903968379399\",\"platform\":\"今日头条\",\"account\":\"4327876576\",\"domain\":\"文化\",\"status\":1,\"read\":1492584,\"like\":2118,\"comment\":0,\"favor\":0,\"origin\":1,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 15:09:39\",\"createTime\":\"2018-06-26 00:22:08\",\"updateTime\":\"2018-06-26 00:22:08\",\"lptime\":1529910579,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a6a10263817832ca4d6d242\",\"app\":0,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p1.pstatp.com/medium/3538/9145332\",\"nickName\":\"长安街知事\"},{\"id\":\"5b3116b2e13823668d49e963\",\"title\":\"意味深长！美军四星上将访华前，中国最强导弹来了一轮齐射\",\"url\":\"http://toutiao.com/a6570974399447433736\",\"platform\":\"今日头条\",\"account\":\"4327876576\",\"domain\":\"文化\",\"status\":1,\"read\":1490868,\"like\":3851,\"comment\":0,\"favor\":0,\"origin\":1,\"mediaType\":1,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 18:56:38\",\"createTime\":\"2018-06-26 00:22:08\",\"updateTime\":\"2018-06-26 00:22:08\",\"lptime\":1529924198,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a6a10263817832ca4d6d242\",\"app\":0,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p1.pstatp.com/medium/3538/9145332\",\"nickName\":\"长安街知事\"},{\"id\":\"5b30f58e6d03b617e58b0819\",\"title\":\"占豪：美国终于说实话，策略是让中国更痛苦，中国却有三招制衡！\",\"url\":\"http://toutiao.com/a6570947828107969028\",\"platform\":\"今日头条\",\"account\":\"54653259147\",\"domain\":\"时政\",\"status\":1,\"read\":1487687,\"like\":3148,\"comment\":0,\"favor\":0,\"origin\":1,\"mediaType\":2,\"coverPic\":\"\",\"sn\":\"\",\"publicTime\":\"2018-06-25 17:13:32\",\"createTime\":\"2018-06-25 22:00:43\",\"updateTime\":\"2018-06-25 22:00:43\",\"lptime\":1529918012,\"isCrawl\":0,\"crawl\":1,\"area\":\"\",\"province\":\"\",\"childDomain\":\"\",\"userId\":\"5a8a97e538178303fcec351a\",\"app\":0,\"appOnline\":0,\"priority\":0,\"social\":0,\"haveHtml\":0,\"es\":1,\"logo\":\"http://p1.pstatp.com/medium/150c0008678d394b93ab\",\"nickName\":\"自媒体联播\"}],\"totalElements\":44821,\"totalPages\":4483,\"numberOfElements\":10,\"number\":2}","columnTotle":null}



