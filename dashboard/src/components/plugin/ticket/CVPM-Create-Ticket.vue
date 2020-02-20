<template>
  <v-card>
    <v-card-title>
      <h2>Create Ticket / Report Abuse</h2>
      <v-subheader>We strongly recommend you to login first to track your tickets</v-subheader>
    </v-card-title>
    <v-alert
      :value="submitSuccess"
      type="success"
    >
      {{ submitSuccessInfo }}
    </v-alert>
    <v-alert
      :value="submitError"
      type="error"
    >
      {{ submitErrorInfo }}
    </v-alert>
    <v-card-text>
      <v-select
        v-model="selectedType"
        :items="items"
        label="Type"
      />
      <v-text-field
        v-model="email"
        label="Email"
      />
      <v-text-field
        v-model="subject"
        label="Subject"
      />
      <v-textarea
        v-model="description"
        label="Description"
      />
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="indigo darken-1"
        outline
        :loading="requesting"
        @click="submit()"
      >
        Submit
      </v-btn>
      <v-btn
        color="indigo darken-1"
        outline
        @click="closeDialog()"
      >
        Close
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { discovery } from '@/services/discovery'
export default {
  props: {
    vendor: {
      type: String,
      default: ''
    },
    packageName: {
      type: String,
      default: ''
    }
  },
  data: () => ({
    email: '',
    subject: '',
    requesting: false,
    submitSuccess: false,
    submitSuccessInfo: '',
    submitErrorInfo: '',
    submitError: false,
    description: '',
    selectedType: 'Not Working',
    items: ['Not Working', 'Malware', 'Abuse', 'Others']
  }),
  methods: {
    submit () {
      const self = this
      self.requesting = true
      discovery
        .putTicket({
          name: this.email,
          email: this.email,
          subject: '[' + this.selectedType + '] ' + this.subject,
          description: '[' + this.vendor + '/' + this.packageName + '] ' + this.description,
          priority: 1
        })
        .then(function (res) {
          if (res.status === 200 && res.data.code === 200) {
            self.submitSuccess = true
            self.submitError = false
            self.submitSuccessInfo =
              'Your Ticket has been created successfully with track ID: ' +
              res.data.results.id
          }
        })
        .catch(function (err) {
          self.submitError = true
          self.submitSuccess = false
          self.submitErrorInfo = err.response.data.description
        })
        .finally(function () {
          self.requesting = false
        })
    },
    closeDialog () {
      this.$emit('closeDialog', true)
    }
  }
}
</script>

<style>
</style>
