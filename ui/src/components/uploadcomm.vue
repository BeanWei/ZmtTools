<template>
  <v-card>
    <v-container fluid grid-list-md>
      <v-flex xs12 sm12>
        <v-card>
          <v-layout row wrap>
            <div class="upload">
              <div class="upload_warp">
                <div class="upload_warp_left" @click="fileClick">
                  <v-icon large>add</v-icon>
                </div>
                <div class="upload_warp_right" @drop="drop($event)" @dragenter="dragenter($event)" @dragover="dragover($event)">
                  或者将文件拖到此处
                </div>
              </div>
              <div class="upload_warp_text">
                选中{{videoList.length}}个文件，共{{bytesToSize(this.size)}}
              </div>
              <input @change="fileChange($event)" type="file" id="upload_file" multiple style="display: none"/>
                <!-- <div class="upload_warp_img" v-show="imgList.length!=0">
                  <div class="upload_warp_img_div" v-for="(item,index) of imgList">
                    <div class="upload_warp_img_div_top">
                      <div class="upload_warp_img_div_text">
                        {{item.file.name}}
                      </div>
                      <div class="upload_warp_img_div_del" @click="fileDel(index)"></div>
                    </div>
                    <img :src="item.file.src">
                  </div>
                </div> -->
            </div>
          </v-layout>
        </v-card>
      </v-flex>
      <v-form ref="videos" :model="videos" v-show="videoList.length!=0">
        <v-flex xs12 sm12>
          <v-card v-for="(item,index) of videoList" style="margin-top: 10px;">
            <v-layout row wrap>
              <v-card-text>
                <p class="text-sm-left">附件名：{{item.file.name}}</p>
              </v-card-text>
            </v-layout>
            <v-layout row wrap style="margin-top: -28px;">
              <v-flex xs12 sm4 style="margin-right: 7px; margin-left: 7px;">
                  <v-text-field
                  name="title"
                  label="标题"
                  v-model="videos[index].title"
                  ></v-text-field>
              </v-flex>
              <v-flex xs12 sm3 style="margin-right: 7px;">
                  <v-text-field
                  name="describe"
                  label="简介"
                  v-model="videos[index].describe"
                  ></v-text-field>
              </v-flex>
              <v-flex xs12 sm3 style="margin-right: 7px;">
                  <v-text-field
                  name="tags"
                  label="标签(逗号隔开)"
                  v-model="videos[index].tags"
                  ></v-text-field>
              </v-flex>
              <v-flex xs12 sm1>
                  <v-select :items="items" v-model="videos[index].field" label="领域"></v-select>
              </v-flex>
            </v-layout>
          </v-card>
        </v-flex>
      </v-form>
    </v-container>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      items:[
        { text: "科技" },
        { text: "娱乐" },
        { text: "搞笑" },
        { text: "财经" },
        { text: "游戏" },
      ],
      videos: [],
      videoList: [],
      size: 0,
    }
  },
  watch: {
    videoList: function() {
      for (var i = this.videos.length + 1; i <= this.videoList.length; i ++) {
        var obj = {
          filename: '',
          title: '',
          describe: '',
          tags: '',
          field: '',
        };
        obj.filename = this.videoList[i-1].file.name
        this.videos.push(obj)
      } 
      this.$store.dispatch("setVideos", this.videos)
    }
  },
  methods: {
    //设置
    // limitClick(state) {
    //   this.imgList = [];
    //   if (state)
    //     this.limit = 2;
    //   else
    //     this.limit = undefined;
    // },
    fileClick() {
      document.getElementById('upload_file').click()
    },
    fileChange(el) {
      if (!el.target.files[0].size) return;
      this.fileList(el.target);
      el.target.value = ''
    },
    fileList(fileList) {
      let files = fileList.files;
      for (let i = 0; i < files.length; i++) {
        //判断是否为文件夹
        if (files[i].type != '') {
          this.fileAdd(files[i]);
        } else {
          //文件夹处理
          this.folders(fileList.items[i]);
        }
      }
    },
    //文件夹处理
    folders(files) {
      let _this = this;
      //判断是否为原生file
      if (files.kind) {
        files = files.webkitGetAsEntry();
      }
      files.createReader().readEntries(function (file) {
        for (let i = 0; i < file.length; i++) {
          if (file[i].isFile) {
            _this.foldersAdd(file[i]);
          } else {
            _this.folders(file[i]);
          }
        }
      })
    },
    foldersAdd(entry) {
      let _this = this;
      entry.file(function (file) {
        _this.fileAdd(file)
      })
    },
    fileAdd(file) {
      if (this.limit !== undefined) this.limit--;
      if (this.limit !== undefined && this.limit < 0) return;
      //总大小
      this.size = this.size + file.size;
      //判断是否为视频文件
      if (file.type.indexOf('image') != -1) {
        this.videoList.push({
          file
        });
      } else {
        alert("不支持非视频格式文件上传")
      }
    },
    fileDel(index) {
      this.size = this.size - this.List[index].file.size;//总大小
      this.videoList.splice(index, 1);
      if (this.limit !== undefined) this.limit = this.videoList.length;
    },
    bytesToSize(bytes) {
      if (bytes === 0) return '0 B';
      let k = 1000, // or 1024
        sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
        i = Math.floor(Math.log(bytes) / Math.log(k));
      return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
    },
    dragenter(el) {
      el.stopPropagation();
      el.preventDefault();
    },
    dragover(el) {
      el.stopPropagation();
      el.preventDefault();
    },
    drop(el) {
      el.stopPropagation();
      el.preventDefault();
      this.fileList(el.dataTransfer);
    }
  }
}
</script>
<style scoped>
  .upload_warp_img_div_del {
    position: absolute;
    top: 6px;
    width: 16px;
    right: 4px;
  }
  .upload_warp_img_div_top {
    position: absolute;
    top: 0;
    width: 100%;
    height: 30px;
    background-color: rgba(0, 0, 0, 0.4);
    line-height: 30px;
    text-align: left;
    color: #fff;
    font-size: 12px;
    text-indent: 4px;
  }
  .upload_warp_img_div_text {
    white-space: nowrap;
    width: 80%;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  .upload_warp_img_div img {
    max-width: 100%;
    max-height: 100%;
    vertical-align: middle;
  }
  .upload_warp_img_div {
    position: relative;
    height: 100px;
    width: 120px;
    border: 1px solid #ccc;
    margin: 0px 30px 10px 0px;
    float: left;
    line-height: 100px;
    display: table-cell;
    text-align: center;
    background-color: #eee;
    cursor: pointer;
  }
  .upload_warp_img {
    border-top: 1px solid #D2D2D2;
    padding: 14px 0 0 14px;
    overflow: hidden
  }
  .upload_warp_text {
    text-align: left;
    margin-bottom: 10px;
    padding-top: 10px;
    text-indent: 14px;
    border-top: 1px solid #ccc;
    font-size: 14px;
  }
  .upload_warp_right {
    float: left;
    width: 57%;
    margin-left: 2%;
    height: 100%;
    border: 1px dashed #999;
    border-radius: 4px;
    line-height: 130px;
    color: #999;
  }
  .upload_warp_left img {
    margin-top: 32px;
  }
  .upload_warp_left {
    float: left;
    width: 40%;
    height: 100%;
    border: 1px dashed #999;
    border-radius: 4px;
    cursor: pointer;
  }
  .upload_warp {
    margin: 14px;
    height: 130px;
  }
  .upload {
    border: 1px solid #ccc;
    background-color: #fff;
    width: 100%;
    box-shadow: 0px 1px 0px #ccc;
    border-radius: 4px;
  }
</style>