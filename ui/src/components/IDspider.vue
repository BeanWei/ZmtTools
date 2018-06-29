<template>
  <v-card>
    <v-card-text>
      <v-container fluid>
        <v-layout row wrap style="margin-top: -40px;">
          <v-flex xs12 style="height: 50px;">
            <v-radio-group v-model="sites" row>
              <v-radio
                v-for="(site, i) in ['大鱼号', '百家号', '企鹅号', '头条号']"
                :key="i"
                :value="site"
                :label="site"
                :color="success"
              ></v-radio>
            </v-radio-group>
          </v-flex>
          <v-btn
            color="teal"
            class="white--text"
            style="margin-top: 17px; margin-right: 26px;"
            @click.native="loader = 'loading3'"
          >
            解析ID
            <v-icon right dark>adb</v-icon>
          </v-btn>
          <v-flex xs12 sm4>
            <v-text-field
              v-model="authorID"
              :rules="IDRules"
              label="ID"
              required
            ></v-text-field>
          </v-flex>
          <v-flex xs12 sm6 md3>
            <v-dialog
              ref="dialog"
              v-model="modal"
              :return-value.sync="date"
              persistent
              lazy
              full-width
              width="290px"
            >
              <v-text-field
                slot="activator"
                v-model="datefrom"
                label="选择起始时间段"
                prepend-icon="event"
                readonly
              ></v-text-field>
              <v-date-picker v-model="datefrom" scrollable>
                <v-spacer></v-spacer>
                <v-btn flat color="primary" @click="modal = false">Cancel</v-btn>
                <v-btn flat color="primary" @click="$refs.dialog.save(datefrom)">OK</v-btn>
              </v-date-picker>
            </v-dialog>
          </v-flex>
          <v-flex xs12 sm6 md3>
            <v-dialog
              ref="dialog"
              v-model="modal"
              :return-value.sync="date"
              persistent
              lazy
              full-width
              width="290px"
            >
              <v-text-field
                slot="activator"
                v-model="dateto"
                label="选择终止时间段"
                prepend-icon="event"
                readonly
              ></v-text-field>
              <v-date-picker v-model="dateto" scrollable>
                <v-spacer></v-spacer>
                <v-btn flat color="primary" @click="modal = false">Cancel</v-btn>
                <v-btn flat color="primary" @click="$refs.dialog.save(dateto)">OK</v-btn>
              </v-date-picker>
            </v-dialog>
          </v-flex>
        </v-layout>
      </v-container>
      <v-btn block color="primary" dark style="margin-top: -10px;">开始抓取</v-btn>
    </v-card-text>
    <v-data-table
      :headers="headers"
      :items="desserts"
      :loading="true"
      class="elevation-1"
    >
      <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
      <template slot="items" slot-scope="props">
        <td>{{ props.item.name }}</td>
        <td class="text-xs-right">{{ props.item.calories }}</td>
        <td class="text-xs-right">{{ props.item.fat }}</td>
        <td class="text-xs-right">{{ props.item.carbs }}</td>
        <td class="text-xs-right">{{ props.item.protein }}</td>
        <td class="text-xs-right">{{ props.item.iron }}</td>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
  export default {
    name: 'IDspider',
    data () {
      return {
        valid: false,
        authorID: '',
        IDRules: [
          v => !!v || '请输入作者ID',
          v => v.length <= 10 || 'Name must be less than 10 characters'
        ],
        datefrom: null,
        dateto: null,
        modal: false,
        headers: [
          {
            text: 'Dessert (100g serving)',
            align: 'left',
            sortable: false,
            value: 'name'
          },
          { text: 'Calories', value: 'calories' },
          { text: 'Fat (g)', value: 'fat' },
          { text: 'Carbs (g)', value: 'carbs' },
          { text: 'Protein (g)', value: 'protein' },
          { text: 'Iron (%)', value: 'iron' }
        ],
        desserts: [
          {
            value: false,
            name: 'Frozen Yogurt',
            calories: 159,
            fat: 6.0,
            carbs: 24,
            protein: 4.0,
            iron: '1%'
          },
          {
            value: false,
            name: 'Ice cream sandwich',
            calories: 237,
            fat: 9.0,
            carbs: 37,
            protein: 4.3,
            iron: '1%'
          },
          {
            value: false,
            name: 'Eclair',
            calories: 262,
            fat: 16.0,
            carbs: 23,
            protein: 6.0,
            iron: '7%'
          },
          {
            value: false,
            name: 'Cupcake',
            calories: 305,
            fat: 3.7,
            carbs: 67,
            protein: 4.3,
            iron: '8%'
          },
          {
            value: false,
            name: 'Gingerbread',
            calories: 356,
            fat: 16.0,
            carbs: 49,
            protein: 3.9,
            iron: '16%'
          },
          {
            value: false,
            name: 'Jelly bean',
            calories: 375,
            fat: 0.0,
            carbs: 94,
            protein: 0.0,
            iron: '0%'
          },
          {
            value: false,
            name: 'Lollipop',
            calories: 392,
            fat: 0.2,
            carbs: 98,
            protein: 0,
            iron: '2%'
          },
          {
            value: false,
            name: 'Honeycomb',
            calories: 408,
            fat: 3.2,
            carbs: 87,
            protein: 6.5,
            iron: '45%'
          },
          {
            value: false,
            name: 'Donut',
            calories: 452,
            fat: 25.0,
            carbs: 51,
            protein: 4.9,
            iron: '22%'
          },
          {
            value: false,
            name: 'KitKat',
            calories: 518,
            fat: 26.0,
            carbs: 65,
            protein: 7,
            iron: '6%'
          }
        ]
      }
    }
  }
</script>

<style scoped>

</style>

