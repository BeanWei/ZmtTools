<template>
  <v-card>
    <v-container id="dropdown-example" grid-list-xl>
      <v-layout row wrap>
        <v-flex xs12 sm2>
          <p>平台</p>
          <v-select
            :items="platform"
            pingtai
            label="全部"
            target="#dropdown-example"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>领域</p>
          <v-select
            :items="domain"
            label="全部"
            lingyu 
            target="#dropdown-example"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>排序方式</p>
          <v-select
            :items="sortway"
            label="阅读量"
            pxfs
            item-value="text"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>阅读量</p>
          <v-select
            :items="totalview"
            label="不限"
            yueduliang
            item-value="text"
          ></v-select>
        </v-flex>
        <v-flex xs12 sm2>
          <p>时间段</p>
          <v-select
            :items="timerange"
            label="24小时"
            shijianduan
            item-value="text"
          ></v-select>
        </v-flex>
        <v-btn color="blue-grey" class="white--text" style="margin-top: 60px;">
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
    >
      <template slot="items" slot-scope="props">
        <td></td>
        <td class="text-xs-right"></td>
        <td class="text-xs-right"></td>
        <td class="text-xs-right"></td>
        <td class="text-xs-right"></td>
        <td class="text-xs-right"></td>
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
      platform: [
        { text: '全部' },
        { text: '微信' },
        { text: '今日头条' },
        { text: '一点资讯' },
        { text: '大鱼号' },
        { text: '百家号' },
        { text: '搜狐号' },
        { text: '网易号' },
      ],
      domain: [
        { text: '全部' },
        { text: '育儿' },
        { text: '教育' },
        { text: '生活' },
        { text: '美食' },
        { text: '情感' },
        { text: '旅行' },
        { text: '职场' },
        { text: '时政' },
        { text: '财经' },
        { text: '军事' },
        { text: '文化' },
        { text: '历史' },
        { text: '科技' },
        { text: '数码' },
        { text: '影视' },
        { text: '星座' },
        { text: '游戏' },
        { text: '动漫' },
        { text: '萌宠' },
        { text: '娱乐' },
        { text: '搞笑' },
        { text: '时尚' },
        { text: '体育' },
        { text: '健康' },
        { text: '房产' },
        { text: '汽车' },
        { text: '传媒' },
        { text: '三农' },
        { text: '综合' }
      ],
      sortway: [
        { text: '发布时间' },
        { text: '阅读量' },
        { text: '评论数' }
      ],
      totalview: [
        { text: '不限' },
        { text: '1万+' },
        { text: '5万+' },
        { text: '10万+' },
        { text: '50万+' }
      ],
      timerange: [
        { text: '不限' },
        { text: '12小时' },
        { text: '24小时' },
        { text: '3天内' },
        { text: '7天内' }
      ],
      search: '',
      headers: [
        { text: '平台', align: 'left', sortable: false, value: '' },
        { text: '标题', align: 'left', sortable: false, value: '' },
        { text: '作者', align: 'left', sortable: false, value: '' },
        { text: '领域', align: 'left', sortable: false, value: '' },
        { text: '时间', align: 'left', sortable: false, value: '' },
        { text: '阅读', align: 'left', sortable: false, value: '' },
        { text: '评论', align: 'left', sortable: false, value: '' }
      ],
      // desserts: [
      //   {
      //     value: false,
      //     name: 'Frozen Yogurt',
      //     calories: 159,
      //     fat: 6.0,
      //     carbs: 24,
      //     protein: 4.0,
      //     iron: '1%'
      //   },
      //   {
      //     value: false,
      //     name: 'Ice cream sandwich',
      //     calories: 237,
      //     fat: 9.0,
      //     carbs: 37,
      //     protein: 4.3,
      //     iron: '1%'
      //   },
      //   {
      //     value: false,
      //     name: 'Eclair',
      //     calories: 262,
      //     fat: 16.0,
      //     carbs: 23,
      //     protein: 6.0,
      //     iron: '7%'
      //   },
      //   {
      //     value: false,
      //     name: 'Cupcake',
      //     calories: 305,
      //     fat: 3.7,
      //     carbs: 67,
      //     protein: 4.3,
      //     iron: '8%'
      //   },
      //   {
      //     value: false,
      //     name: 'Gingerbread',
      //     calories: 356,
      //     fat: 16.0,
      //     carbs: 49,
      //     protein: 3.9,
      //     iron: '16%'
      //   },
      //   {
      //     value: false,
      //     name: 'Jelly bean',
      //     calories: 375,
      //     fat: 0.0,
      //     carbs: 94,
      //     protein: 0.0,
      //     iron: '0%'
      //   },
      //   {
      //     value: false,
      //     name: 'Lollipop',
      //     calories: 392,
      //     fat: 0.2,
      //     carbs: 98,
      //     protein: 0,
      //     iron: '2%'
      //   },
      //   {
      //     value: false,
      //     name: 'Honeycomb',
      //     calories: 408,
      //     fat: 3.2,
      //     carbs: 87,
      //     protein: 6.5,
      //     iron: '45%'
      //   },
      //   {
      //     value: false,
      //     name: 'Donut',
      //     calories: 452,
      //     fat: 25.0,
      //     carbs: 51,
      //     protein: 4.9,
      //     iron: '22%'
      //   },
      //   {
      //     value: false,
      //     name: 'KitKat',
      //     calories: 518,
      //     fat: 26.0,
      //     carbs: 65,
      //     protein: 7,
      //     iron: '6%'
      //   }
      // ]
    })
  }
</script>

<style scoped>

</style>
