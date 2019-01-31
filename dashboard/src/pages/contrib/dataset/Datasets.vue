<template>
  <v-container>
    <v-card>
      <v-data-table :items="datasets" :headers="headers" class="elevation-1">
        <template slot="items" slot-scope="props">
          <td class="text-xs-left">{{ props.item.Name }}</td>
          <td class="text-xs-left">{{ props.item.Desc }}</td>
          <td class="text-xs-left">
            <v-chip
              v-for="(item, index) in props.item.Tags"
              :key="index"
              color="primary"
              text-color="white"
              @click="setSearchTag(item);search()"
            >
              <B style="color:yellow" v-if="item===searchTag">
                <I>{{ item }}</I>
              </B>
              <span v-else>{{ item }}</span>
            </v-chip>
          </td>
          <v-btn outline small fab color="indigo" @click="viewDetail(props.item)">
            <v-icon>info</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>
    <v-dialog v-model="detailDialog">
      <v-card>
        <v-card-title class="headline grey lighten-2" primary-title>{{detailInfo.Name}}</v-card-title>

        <v-card-text>{{detailInfo.FullDesc}}</v-card-text>

        <v-divider></v-divider>
        <v-card-text>
          External Link:
          <a :href="detailInfo.Link">{{detailInfo.Link}}</a>
        </v-card-text>
        <v-card-text>
          Files:
          {{detailInfo.Files}}
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" flat @click="detailDialog = false">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { systemService } from '@/services/system'
export default {
  data () {
    return {
      datasets: [],
      searchName: '',
      detailDialog: false,
      detailInfo: {},
      searchTag: '',
      headers: [
        {
          text: 'Name',
          align: 'left',
          sortable: true,
          value: 'Name'
        },
        {
          text: 'Desc',
          align: 'left',
          sortable: true,
          value: 'Desc'
        },
        {
          text: 'Tags',
          align: 'left',
          sortable: false,
          value: 'Tags'
        },
        {
          text: 'Actions',
          align: 'left',
          sortable: false,
          value: 'Tags'
        }
      ]
    }
  },
  methods: {
    filterTags (tag) {
      if (tag !== '') {
        return tag
      }
    },
    setSearchTag (item) {
      if (item === this.searchTag) {
        this.searchTag = ''
      } else {
        this.searchTag = item
      }
    },
    filterDatasets (item) {
      if (typeof item !== 'undefined') {
        return item
      }
    },
    viewDetail (item) {
      this.detailDialog = true
      this.detailInfo = item
    },
    search () {
      this.fetchAllDatasets(this.searchName, this.searchTag)
    },
    fetchAllDatasets (name = '', tag = '') {
      let self = this
      systemService.getAllDatasets().then(function (res) {
        self.datasets = res.data.map(function (each) {
          if (each.Tags.search(tag) !== -1) {
            if (each.Name !== '' && each.Name.search(name) !== -1) {
              each.FullDesc = each.Desc
              each.Desc = each.Desc.slice(0, 140) + '...'
              each.Tags = each.Tags.replace(/,/g, ' ')
                .split(' ')
                .filter(self.filterTags)
              return each
            }
          }
        })
        self.datasets = self.datasets.filter(self.filterDatasets)
      })
    }
  },
  created () {
    this.fetchAllDatasets()
  }
}
</script>

<style>
</style>
