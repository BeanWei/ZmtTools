<template>
  <v-jumbotron
    :gradient="gradient"
    dark
    :src="aboutCover"
    style="margin-top: -9px;height: 650px;"
  >
    <v-container grid-list-xl>
      <v-layout align-center style="margin: 16px 0px 26px 0;">
        <v-flex text-xs-center>
          <h1 class="display-4">关于·『ZmtTools』</h1>
        </v-flex>
      </v-layout>
      <v-layout align-center>
        <v-flex text-xs-center style="margin: 16px 0px 26px 0;">
          <h3 class="display-3">作者：Bean.Wei</h3>
        </v-flex>
      </v-layout>
      <v-layout align-center>
        <v-flex text-xs-center style="margin: 16px 0px 26px 0;">
          <h3 class="display-3">版本：内测版</h3>
        </v-flex>
      </v-layout>
      <v-layout align-center>
        <v-flex text-xs-center style="margin: 16px 0px 26px 0;">
          <h3 class="display-3">版本号：V20180707-1.0</h3>
        </v-flex>
      </v-layout>
    </v-container>
  </v-jumbotron>
</template>

<script>
import aboutCover from "../assets/material.jpg"
export default {
  data() {
    return {
      aboutCover
    }
  }
}
</script>