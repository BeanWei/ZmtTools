<template>
  <v-card>
    <div id="toolbar"></div>
    <div id="content">
      <p>请输入内容</p>
    </div>
    <!-- <div id="editorElem" style="text-align:left">
      <p>欢迎使用富文本编辑器</p>
    </div> -->
    <v-speed-dial
      v-model="fab"
      bottom
      right
      fixed
      :direction="direction"
      :open-on-hover="hover"
      :transition="transition"
    >
      <v-btn
        slot="activator"
        v-model="fab"
        color="teal"
        dark
        fab
      >
        <v-icon>widgets</v-icon>
        <v-icon>close</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        small
        color="green"
      >
        <v-icon>edit</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        small
        color="indigo"
      >
        <v-icon>add</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        small
        color="red"
      >
        <v-icon>delete</v-icon>
      </v-btn>
    </v-speed-dial>
  </v-card>
</template>

<script>
import WangEditor from 'wangeditor'
export default {
  data() {
    return{
      fab: false,
      editorContent: '',
    }
  },
  mounted() {
    var editor = new WangEditor('#toolbar', '#content')
    editor.customConfig.onchange = (html) => {
      this.editorContent = html
      editor.customConfig.uploadImgServer = '' 
      editor.customConfig.uploadFileName = '' 
      editor.customConfig.uploadImgHeaders = { 
        'Accept': '*/*', 'Authorization':'Bearer ' + 'token'  //头部token 
      }
      editor.customConfig.menus = [ //菜单配置 
        'head', 
        'list', // 列表 
        'justify', // 对齐方式 
        'bold', 
        'fontSize', // 字号 
        'italic', 
        'underline', 
        'image', // 插入图片 
        'foreColor', // 文字颜色 
        'undo', // 撤销 
        'redo', // 重复 
      ]

      editor.customConfig.uploadImgHooks = {
        before: function (xhr, editor, files) {

        },
        success: function (xhr, editor, result) {
          this.imgUrl = Object.values(result.data).toString()
        },
        fail: function (xhr, editor, result) {

        },
        error: function (xhr, editor) {

        },
        customInsert: function (insertImg, result, editor) {
          let url = Object.values(result.data)
          JSON.stringify(url)
          insertImg(url)
        }
      }
    }
    editor.create()
  },
  methods: {
    // getEditorContent() {
    //   this.editorContent = this.editor.$txt.html()
    // }
  }
}
</script>

<style scoped>
  #toolbar {
    border: 1px solid #ccc;
  }
  #content {
    border: 1px solid #ccc;
    height: 600px;
    min-height: 450px;
  }
</style>
