<template>
  <v-container>
    <v-btn
      outline
      color="indigo"
      @click="newTraining()"
    >
      <v-icon
        left
        dark
      >
        fas fa-star
      </v-icon>
      {{ $t(`Train.new`) }}
    </v-btn>
    <v-text-field
      v-model="searchKW"
      :label="$t(`Datasets.search`)"
    />
    <v-card>
      <v-data-table
        :items="trains"
        :headers="headers"
        class="elevation-1"
      >
        <template
          slot="items"
          slot-scope="props"
        >
          <td class="text-xs-left">
            {{ props.item.rayId }}
          </td>
          <td class="text-xs-left">
            {{ props.item.vendor }}
          </td>
          <td class="text-xs-left">
            {{ props.item.package }}
          </td>
          <td class="text-xs-left">
            {{ props.item.solver }}
          </td>
          <td class="text-xs-left">
            {{ props.item.status }}
          </td>
          <td class="text-xs-left">
            {{ formatRequestTime(props.item.createdAt) }}
          </td>
          <v-btn
            outline
            small
            fab
            color="indigo"
            @click="uncompressFile(props.item.objectId)"
          >
            <v-icon>info</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>
    <v-dialog v-model="enableAnnotMetrics">
      <anno-metrics :annodata="annots" />
    </v-dialog>
  </v-container>
</template>

<script>
import AnnotationsMetrics from '@/components/metrics/AnnotationMetrics.vue'
import { systemService } from '@/services/system'
import { searchService } from '@/services/search'
import dayjs from 'dayjs'

export default {
  components: {
    'anno-metrics': AnnotationsMetrics
  },
  data () {
    return {
      trains: [],
      allTrains: [],
      searchName: '',
      enableAnnotMetrics: false,
      searchKW: '',
      searchTag: '',
      headers: [
        {
          text: 'RayID',
          align: 'left',
          sortable: true,
          value: 'Name'
        },
        {
          text: this.$t(`Packages_detail.vendor`),
          align: 'left',
          sortable: true,
          value: 'Name'
        },
        {
          text: this.$t(`Packages_detail.package`),
          align: 'left',
          sortable: true,
          value: 'Size'
        },
        {
          text: this.$t(`Packages_detail.solver`),
          align: 'left',
          sortable: true,
          value: 'Size'
        },
        {
          text: this.$t(`Train.status`),
          align: 'left',
          sortable: true,
          value: 'Size'
        },
        {
          text: this.$t(`Datasets.uploadedAt`),
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
      const self = this
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
    this.fetchTrainList()
  },
  methods: {
    formatRequestTime (requestAt) {
      return dayjs(requestAt).format('YYYY-MM-DD HH:mm:ss A')
    },
    search () {
      this.fetchMyDatasets(this.searchName)
    },
    newTraining () {
      this.$router.push('/train/new')
    },
    uploadFile (e) {
      const files = e.target.files
      if (files.length) {
        console.log(files[0])
        systemService.uploadFile(files[0], 'dataset').then(function (res) {
          location.reload()
        })
      } else {
        console.error('[err] no file selected!')
      }
    },
    fetchTrainList (name = '') {
      const self = this
      systemService.getTrainList().then(function (res) {
        self.trains = res.data.result
        self.allTrains = self.datasets
        for (var index in self.datasets) {
          searchService.addItem(index, self.datasets[index])
        }
      })
    },
    uncompressFile (id) {
      console.log(id)
    },
    showAnnotationsMetrics (objectId) {
      const self = this
      systemService.getAnnotationInfo(objectId).then(function (res) {
        self.annots = JSON.parse(res.data.annotation)
        self.enableAnnotMetrics = true
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
input[type="file"] {
  position: absolute;
  clip: rect(0, 0, 0, 0);
}
</style>
