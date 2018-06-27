<template>
  <v-card>
    <v-container id="baowen" grid-list-xl>
      <v-layout row wrap>
        <v-flex xs12 sm2>
          <p>平台</p>
          <v-select
            :items="platformlist"
            v-model="platform" 
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>领域</p>
          <v-select
            :items="domainlist"
            v-model="domain"  
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>排序方式</p>
          <v-select
            :items="sortwaylist"
            v-model="sortway" 
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>阅读量</p>
          <v-select
            :items="totalviewlist"
            v-model="totalview"  
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>时间段</p>
          <v-select
            :items="timerangelist"
            v-model="timerange" 
          ></v-select>
        </v-flex>
        <v-btn color="blue-grey" class="white--text" style="margin-top: 60px;" @click="submit">
          提交
          <v-icon right dark>cloud_download</v-icon>
        </v-btn>
      </v-layout>
    </v-container>
    <v-card-title>
      *一点资讯、大鱼号和网易号平台不提供阅读量数据，请悉知。
      <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="search"
        label="『搜索』请输入文章标题"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :search="search"
      :items="respdata"
    >
      <template slot="items" slot-scope="props">
        <td class="text-xs-right">{{props.item.platform}}</td>
        <td class="text-xs-right">{{props.item.title}}</td>
        <td class="text-xs-right">{{props.item.nickName}}</td>
        <td class="text-xs-right">{{props.item.domain}}</td>
        <td class="text-xs-right">{{props.item.publicTime}}</td>
        <td class="text-xs-right">{{props.item.read}}</td>
        <td class="text-xs-right">{{props.item.comment}}</td>
      </template>
      <v-alert slot="no-results" :value="true" color="error" icon="warning">
        Your search for "{{ search }}" found no results.
      </v-alert>
    </v-data-table>
  </v-card>
</template>

<script>
  export default {
    data: () => ({
      platformlist: [
        '全部' ,
        '微信' ,
        '今日头条' ,
        '一点资讯' ,
        '大鱼号' ,
        '百家号' ,
        '搜狐号' ,
        '网易号' ,
      ],
      domainlist: [
        '全部' ,
        '育儿' ,
        '教育' ,
        '生活' ,
        '美食' ,
        '情感' ,
        '旅行' ,
        '职场' ,
        '时政' ,
        '财经' ,
        '军事' ,
        '文化' ,
        '历史' ,
        '科技' ,
        '数码' ,
        '影视' ,
        '星座' ,
        '游戏' ,
        '动漫' ,
        '萌宠' ,
        '娱乐' ,
        '搞笑' ,
        '时尚' ,
        '体育' ,
        '健康' ,
        '房产' ,
        '汽车' ,
        '传媒' ,
        '三农' ,
        '综合' 
      ],
      sortwaylist: [
        '默认' ,
        '发布时间' ,
        '阅读量' ,
        '评论数' 
      ],
      totalviewlist: [
        '不限' ,
        '1万+' ,
        '5万+' ,
        '10万+' ,
        '50万+' 
      ],
      timerangelist: [
        '不限' ,
        '12小时' ,
        '24小时' ,
        '3天内' ,
        '7天内' 
      ],
      search: '',
      headers: [
        { text: '平台', align: 'left', sortable: false, value: 'platform' },
        { text: '标题', align: 'left', sortable: false, value: 'title' },
        { text: '作者', align: 'left', sortable: false, value: 'nickName' },
        { text: '领域', align: 'left', sortable: false, value: 'domain' },
        { text: '时间', align: 'left', sortable: false, value: 'publicTime' },
        { text: '阅读', align: 'left', sortable: false, value: 'read' },
        { text: '评论', align: 'left', sortable: false, value: 'comment' }
      ],
      respdata: [],
      platform: "全部",
      domain: "全部",
      sortway: "发布时间",
      totalview: "不限",
      page:	1,
      timerange: "24小时"

    }),
    methods: {
      submit() {
        let postData = {
          lgCustomerId:	"1000186598",
          page:	this.page,
          reqTime:	new Date().getTime().toString(),
          size:	10,
          status:	1
        }
        if (this.platform !== "全部") {
          postData.platform = this.platform
        }
        if (this.domain !== "全部") {
          postData.domain = this.domain
        }
        if (this.sortway !== "默认") {
          if (this.sortway === "发布时间") {
            postData.lptimeOrder = "DESC"
          } else if (this.sortway === "阅读量") {
            postData.readOrder = "DESC"
          } else if (this.sortway === "评论数") {
            postData.commentOrder =	"DESC"
          }
        }
        if (this.totalview === "不限") {
           postData.read = 0
        } else if (this.totalview === "1万+") {
          postData.read = 10000
        } else if (this.totalview === "5万+") {
          postData.read = 50000
        } else if (this.totalview === "10万+") {
          postData.read = 100000
        } else if (this.totalview === "50万+") {
          postData.read = 500000
        } 
        if (this.timerange === "不限") {
          postData.lptime = "0"
        } else if (this.timerange === "12小时") {
          postData.lptime = Math.round((new Date().getTime()/1000)-3600*12).toString()
        } else if (this.timerange === "24小时") {
          postData.lptime = Math.round((new Date().getTime()/1000)-3600*24).toString()
        } else if (this.timerange === "3天内") {
          postData.lptime = Math.round((new Date().getTime()/1000)-3600*24*3).toString()
        } else if (this.timerange === "7天内") {
          postData.lptime = Math.round((new Date().getTime()/1000)-3600*24*7).toString()
        } 
        this.$axios.post(
          "http://www.myleguan.com/lg_res/focus/flr/la",
          this.$qs.stringify(postData)
        ).then(response => {
          console.log(response.data)
          for (var item in JSON.parse(response.data.reObj).content) {
            const data = {}
            data.platform = item.platform
            data.title = item.title
            data.nickName = item.nickName
            data.domain = item.domain
            data.publicTime = item.publicTime
            data.read = item.read
            data.comment = item.comment
            this.respdata.push(data)
          }
          console.log(this.respdata)
        }, response => {
          alert("出错了！")
        })
      }
    }
  }
</script>

<style scoped>

</style>
