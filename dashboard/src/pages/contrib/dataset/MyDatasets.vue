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
            {{ props.item.Desc }}
          </td>
          <td class="text-xs-left">
            <v-chip
              v-for="(item, index) in props.item.Tags"
              :key="index"
              color="primary"
              text-color="white"
              @click="setSearchTag(item);search()"
            >
              <B
                v-if="item===searchTag"
                style="color:yellow"
              >
                <I>{{ item }}</I>
              </B>
              <span v-else>{{ item }}</span>
            </v-chip>
          </td>
          <v-btn
            outline
            small
            fab
            color="indigo"
            @click="viewDetail(props.item)"
          >
            <v-icon>info</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>
    <v-dialog v-model="detailDialog">
      <v-card>
        <v-card-title
          class="headline grey lighten-2"
          primary-title
        >
          {{ detailInfo.Name }}
        </v-card-title>

        <v-card-text>
          <pre>{{ detailInfo.FullDesc }}</pre>
        </v-card-text>

        <v-divider />
        <v-card-text>
          External Link:
          <a :href="detailInfo.Link">{{ detailInfo.Link }}</a>
        </v-card-text>
        <v-card-text>
          Files:
          {{ detailInfo.Files }}
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            color="primary"
            flat
            @click="detailDialog = false"
          >
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="datasetSyncDialog">
      <v-card>
        <v-card-title>
          <span class="headline">Sync Database</span>
        </v-card-title>
        <v-card-text>
          <v-text-field
            v-model="databaseURL"
            label="Datasets URL*"
            required
            hint="e.g: https://premium.file.cvtron.xyz/cvpm/data/registry/dataset.toml"
          />
        </v-card-text>
        <v-card-actions>
          <v-btn
            color="indigo darken-1"
            outline
            @click="sync()"
          >
            Sync
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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
          text: this.$t(`Datasets.desc`),
          align: 'left',
          sortable: true,
          value: 'Desc'
        },
        {
          text: this.$t(`Datasets.tags`),
          align: 'left',
          sortable: false,
          value: 'Tags'
        },
        {
          text: this.$t(`Datasets.actions`),
          align: 'left',
          sortable: false,
          value: 'Tags'
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
    this.fetchAllDatasets()
  },
  methods: {
    alert (message) {
      alert(message)
    },
    trigger_sync () {
      this.datasetSyncDialog = true
    },
    sync () {
      systemService.SyncDatabase(this.databaseURL).then(function (res) {
        location.reload()
      })
    },
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
