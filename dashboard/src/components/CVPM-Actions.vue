<template>
  <v-card>
    <v-card-title>
      <h2>{{ $t('Packages_detail.actions') }}</h2>
    </v-card-title>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="indigo darken-1"
        outline
        @click="triggerDialog('test')"
      >
        {{ $t('Packages_detail.test') }}
      </v-btn>
      <v-btn
        color="indigo darken-1"
        outline
        @click="triggerDialog('ticket')"
      >
        {{ $t('Packages_detail.ticket') }}
      </v-btn>
    </v-card-actions>
    <v-dialog v-model="enableTest">
      <cvpm-request
        v-if="enableTest"
        :config="config"
        :vendor="vendor"
        :package-name="packageName"
        @closeDialog="triggerDialog('test')"
      />
    </v-dialog>
    <v-dialog v-model="enableTicket">
      <cvpm-create-ticket
        v-if="enableTicket"
        :vendor="vendor"
        :package-name="packageName"
        @closeDialog="triggerDialog('ticket')"
      />
    </v-dialog>
  </v-card>
</template>

<script>
import cvpmRequest from '@/components/CVPM-Request'
import cvpmCreateTicket from '@/components/plugin/ticket/CVPM-Create-Ticket'
export default {
  components: {
    'cvpm-request': cvpmRequest,
    'cvpm-create-ticket': cvpmCreateTicket
  },
  props: {
    config: {
      type: Object,
      default: function () {
        return {}
      }
    },
    vendor: {
      type: Array,
      default: function () {
        return []
      }
    },
    packageName: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  data () {
    return {
      enableTest: false,
      enableTicket: false
    }
  },
  methods: {
    triggerDialog (name) {
      if (name === 'test') {
        this.enableTest = !this.enableTest
      } else if (name === 'ticket') {
        this.enableTicket = !this.enableTicket
      }
    }
  }
}
</script>

<style>
</style>
