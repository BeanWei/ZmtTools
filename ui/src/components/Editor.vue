<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm12 style="margin-top: -20px;">
        <v-text-field
          v-model="title"
          :rules="[titlerules.required, titlerules.counter]"
          label="文章标题"
          counter
          maxlength="50"
        ></v-text-field>
      </v-flex>
    </v-layout>
    <v-layout row>
      <v-flex xs12 sm12>
        <quill-editor v-model="content"
          ref="textEditor"
          :options="editorOption"
          @blur="onEditorBlur($event)"
          @focus="onEditorFocus($event)"
          @ready="onEditorReady($event)">
        </quill-editor>
        <v-card style="margin-top: 12px;">
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
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import { quillEditor } from 'vue-quill-editor'

export default {
  components: {
    quillEditor
  },
  data() {
    return{
      fab: false,
      title: '',
      editorContent: '',
      titlerules: {
        required: value => !!value || '标题不能为空.',
        counter: value => value.length <= 50 || '标题限制50字',
      },
      editorOption: {
        //theme: 'bubble',
        placeholder: "*欢迎使用*",
        modules: {
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'],
            ['blockquote', 'code-block'],
            [{ 'header': 1 }, { 'header': 2 }],
            [{ 'list': 'ordered' }, { 'list': 'bullet' }],
            [{ 'script': 'sub' }, { 'script': 'super' }],
            [{ 'indent': '-1' }, { 'indent': '+1' }],
            [{ 'direction': 'rtl' }],
            [{ 'size': ['small', false, 'large', 'huge'] }],
            [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
            [{ 'font': [] }],
            [{ 'color': [] }, { 'background': [] }],
            [{ 'align': [] }],
            ['clean'],
            ['link', 'image', 'video']
          ],
        }
      }
    }
  },
  methods: {
    onEditorBlur(editor) {
      console.log('editor blur!', editor)
    },
    onEditorFocus(editor) {
      console.log('editor focus!', editor)
    },
    onEditorReady(editor) {
      console.log('editor ready!', editor)
    },
    onEditorChange({ editor, html, text }) {
      console.log('editor change!', editor, html, text)
      this.content = html
    }
  },
  computed: {
    editor() {
      return this.$refs.textEditor.quill
    }
  },
}
</script>

<style scoped>
.quill-editor {
  height: 450px;
  margin-top: 12px;
}
</style>
