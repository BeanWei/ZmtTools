<template>
  <v-card>
    <v-container id="baowen" grid-list-xl>
      <v-layout row wrap>
        <v-flex xs12 sm2>
          <p style="margin-top: -20px;">平台</p>
          <v-select
            :items="platformlist"
            v-model="platform" 
            style="margin-top: -25px;"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p style="margin-top: -20px;">领域</p>
          <v-select
            :items="domainlist"
            v-model="domain"  
            style="margin-top: -25px;"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p style="margin-top: -20px;">排序方式</p>
          <v-select
            :items="sortwaylist"
            v-model="sortway" 
            style="margin-top: -25px;"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p style="margin-top: -20px;">阅读量</p>
          <v-select
            :items="totalviewlist"
            v-model="totalview"  
            style="margin-top: -25px;"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p style="margin-top: -20px;">时间段</p>
          <v-select
            :items="timerangelist"
            v-model="timerange" 
            style="margin-top: -25px;"
          ></v-select>
        </v-flex>
        <v-btn color="blue-grey" class="white--text" style="margin-top: 12px;" @click="submit">
          提交
          <v-icon right dark>cloud_download</v-icon>
        </v-btn>
      </v-layout>
    </v-container>
    <v-card-title style="margin-top: -70px;">
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
      :loading = "loading"
      class="elevation-3"
      hide-actions
    >
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">{{props.item.platform}}</td>
        <td class="text-xs-left">{{props.item.title}}</td>
        <td class="text-xs-left">{{props.item.nickName}}</td>
        <td class="text-xs-left">{{props.item.domain}}</td>
        <td class="text-xs-left">{{props.item.publicTime}}</td>
        <td class="text-xs-left">{{props.item.read}}</td>
        <td class="text-xs-left">{{props.item.comment}}</td>
      </template>
      <template slot="no-data">
        <v-alert :value="true" outline color="error" icon="warning">
          Sorry, nothing to display here :(
        </v-alert>
      </template>
      <v-alert slot="no-results" :value="true" color="error" icon="warning">
        Your search for "{{ search }}" found no results.
      </v-alert>
    </v-data-table>
    <v-layout row justify-space-around>
      <v-flex xs12 sm1>
        <v-btn flat icon large color="blue-grey darken-2" @click="previous">
          <v-icon>keyboard_arrow_left</v-icon>
          上一页
        </v-btn>
      </v-flex>
      <v-flex xs12 sm1>
        <v-subheader>跳转至:</v-subheader>
      </v-flex>
      <v-flex xs12 sm4>
        <v-text-field
          v-model="pagenum"
          label="输入页数"
          single-line
          style="margin-top: -12px;"
          @keyup.enter="pagejump"
        ></v-text-field>
      </v-flex>
      <v-flex xs12 sm2>
        <v-subheader>页 | 共{{totalpages}}页 | 当前第-{{page}}-页</v-subheader>
      </v-flex>
      <v-flex xs12 sm1>
        <v-btn flat icon large color="blue-grey darken-2" @click="next">
          下一页
          <v-icon>keyboard_arrow_right</v-icon>
        </v-btn>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
  export default {
    name: 'Baowen',
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
      timerange: "24小时",
      loading: false,
      pagenum: '',
      totalpages: 1,
      pagechage: false
    }),
    methods: {
      submit() {
        this.loading = true
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
        if (this.pagechage === false) {
          this.page = 1
          this.pagenum = ''
        }
        this.$axios.post(
          "http://www.myleguan.com/lg_res/focus/flr/la",
          this.$qs.stringify(postData)
        ).then(response => {
          var obj = JSON.parse(response.data.reObj)
          this.totalpages = obj.totalPages
          var objlist = obj.content
          this.respdata = []
          for (var item in objlist) {
            const data = {}
            data.platform = objlist[item].platform
            data.title = objlist[item].title
            data.nickName = objlist[item].nickName
            data.domain = objlist[item].domain
            data.publicTime = objlist[item].publicTime
            data.read = objlist[item].read
            data.comment = objlist[item].comment
            this.respdata.push(data)
          }
          console.log(this.respdata)
        }, response => {
          alert("出错了！")
        })
        setTimeout(() => {
          this.loading = false
        }, 1000) 
        this.pagechage = false
      },
      pagejump() {
        if (this.pagenum < 1 || this.pagenum > this.totalpages) {
          alert("请输入正确的页数")
        } else {
          this.page = this.pagenum
          this.pagechage = true
          this.submit()
        }   
      },
      previous() {
        if (this.page === 1) {
          alert("没有上一页了")
        } else {
          this.page--
          this.pagechage = true
          this.submit()
        }
      },
      next() {
        if (this.page === this.totalpages) {
          alert("没有下一页了")
        } else {
          this.page++
          this.pagechage = true
          this.submit()
        }
      }
    },
    // created() { 
    //   this.submit()
    // }
  }
</script>

<style scoped>

</style>
