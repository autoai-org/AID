<template>
  <v-card>
    <v-card-title>
      <h2>Actions</h2>
    </v-card-title>
    <v-card-actions>
      <v-btn color="indigo darken-1" outline @click="triggerDialog('test')">Test</v-btn>
      <v-btn color="indigo darken-1" outline @click="triggerDialog('ticket')">Ticket</v-btn>
    </v-card-actions>
    <v-dialog v-model="enableTest">
      <cvpm-request 
        v-if="enableTest" 
        v-on:closeDialog="triggerDialog('test')"
        :config="config"
        :vendor="vendor"
        :packageName="packageName"
        ></cvpm-request>
    </v-dialog>
    <v-dialog v-model="enableTicket">
      <cvpm-create-ticket
        v-if="enableTicket"
        :vendor="vendor"
        :packageName="packageName"
        v-on:closeDialog="triggerDialog('ticket')"
      ></cvpm-create-ticket>
    </v-dialog>
  </v-card>
</template>

<script>
import cvpmRequest from '@/components/CVPM-Request'
import cvpmCreateTicket from '@/components/plugin/ticket/CVPM-Create-Ticket'
export default {
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
  },
  components: {
    'cvpm-request': cvpmRequest,
    'cvpm-create-ticket': cvpmCreateTicket
  },
  props: ['config', 'vendor', 'packageName']
}
</script>

<style>
</style>
