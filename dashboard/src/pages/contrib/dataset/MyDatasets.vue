<template>
  <v-container>
    <v-btn
      outline
      color="indigo"
      @click="uploadNewDatasets()"
    >
      <v-icon
        left
        dark
      >
        fas fa-star
      </v-icon>
      {{ $t(`Datasets.upload`) }}
    </v-btn>
    <v-text-field
      v-model="searchKW"
      :label="$t(`Datasets.search`)"
    />
    <v-card>
      <v-data-table
        :items="datasets"
        :headers="headers"
        class="elevation-1"
      >
        <template
          slot="items"
          slot-scope="props"
        >
          <td class="text-xs-left">
            {{ props.item.Name }}
          </td>
          <td class="text-xs-left">
            {{ (props.item.Size/1024/1024).toFixed(2) }} MB
          </td>
          <v-btn
            outline
            small
            fab
            color="indigo"
          >
            <v-icon>info</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script>
import { systemService } from '@/services/system'
import { searchService } from '@/services/search'
export default {
  data () {
    return {
      datasets: [],
      allDatasets: [],
      searchName: '',
      detailDialog: false,
      searchKW: '',
      datasetSyncDialog: false,
      databaseURL:
        'https://premium.file.cvtron.xyz/cvpm/data/registry/dataset.toml',
      detailInfo: {},
      searchTag: '',
      headers: [
        {
          text: this.$t(`Datasets.name`),
          align: 'left',
          sortable: true,
          value: 'Name'
        },
        {
          text: this.$t(`Datasets.size`),
          align: 'left',
          sortable: true,
          value: 'Size'
        },
        {
          text: this.$t(`Datasets.actions`),
          align: 'left',
          sortable: false,
          value: 'Operation'
        }
      ]
    }
  },
  watch: {
    searchKW (val) {
      let self = this
      if (val === '') {
        self.datasets = self.allDatasets
      }
      searchService.searchItems(val).then(function (res) {
        if (res.length !== 0) {
          self.datasets = res.map(function (each) {
            return self.allDatasets[each]
          })
        }
      })
    }
  },
  created () {
    this.fetchMyDatasets()
  },
  methods: {
    search () {
      this.fetchMyDatasets(this.searchName)
    },
    fetchMyDatasets (name = '') {
      let self = this
      systemService.getMyFiles().then(function (res) {
        self.datasets = res.data.result.filter(function (each) {
          if (each.Name.startsWith('dataset-')) {
            if (each.Name !== '' && each.Name.search(name) !== -1) {
              return each
            }
          }
        })
        console.log(self.datasets)
        self.allDatasets = self.datasets
        for (var index in self.datasets) {
          searchService.addItem(index, self.datasets[index])
        }
      })
    }
  }
}
</script>

<style scoped>
pre {
  white-space: pre-wrap;
  word-wrap: break-word;
}
.keyword-hit {
  color: red;
}
</style>
