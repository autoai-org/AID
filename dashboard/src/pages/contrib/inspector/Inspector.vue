<template>
  <v-container>
    <v-card>
      <v-data-table
        :items="inspectInfo"
        :headers="headers"
        class="elevation-1"
      >
        <template
          slot="items"
          slot-scope="props"
        >
          <td class="text-xs-left">
            {{ props.item.client_ip }}
          </td>
          <td class="text-xs-left">
            {{ props.item.request_url }}
          </td>
          <td class="text-xs-left">
            <v-chip
              v-if="isSuccessfulRequest(props.item.http_status)"
              color="green"
              text-color="white"
            >
              {{ props.item.http_status }}
            </v-chip>
            <v-chip
              v-else
              color="red"
              text-color="white"
            >
              {{ props.item.http_status }}
            </v-chip>
          </td>
          <td class="text-xs-left">
            <time :datetime="props.item.requested_at">{{ formatRequestTime(props.item.requested_at) }}</time>
          </td>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script>
import { systemService } from '@/services/system'
import dayjs from 'dayjs'
export default {
  data () {
    return {
      inspectInfo: [],
      headers: [
        {
          text: this.$t('Inspector.client_ip'),
          align: 'left',
          sortable: true,
          value: 'client_ip'
        },
        {
          text: this.$t('Inspector.target'),
          align: 'left',
          sortable: true,
          value: 'request_url'
        },
        {
          text: this.$t('Inspector.result'),
          align: 'left',
          sortable: false,
          value: 'http_status'
        },
        {
          text: this.$t('Inspector.time'),
          align: 'left',
          sortable: false,
          value: 'requested_at'
        }
      ]
    }
  },
  created () {
    this.fetchInspector()
  },
  methods: {
    fetchInspector () {
      const self = this
      systemService.getInspectorInfo().then(function (res) {
        self.inspectInfo = res.data.requests
        console.log(self.inspectInfo)
      })
    },
    formatRequestTime (requestAt) {
      return dayjs(requestAt).format('YYYY-MM-DD HH:mm:ss A')
    },
    isSuccessfulRequest (requestStatus) {
      console.log(requestStatus)
      return requestStatus.toString().startsWith('2')
    }
  }
}
</script>

<style>
</style>
