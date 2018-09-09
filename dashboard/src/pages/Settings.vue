<template>
  <v-layout row wrap>
    <v-flex xs12 v-if="info">
    <v-alert outline class="cvpm-setting-alert" :value="info" type="info">{{ info }}</v-alert>
    </v-flex>
    <v-flex md8="md8" xs12>
      <v-form>
        <v-switch label="Developer Mode" v-model="config.developerMode"></v-switch>
        <v-switch label="Enable Hint" v-model="config.hintMode"></v-switch>
        <v-text-field v-model="config.endpoint" label="Endpoint" required></v-text-field>
      </v-form>
    </v-flex>
    <v-flex md4="md4"></v-flex>
    <v-btn primary @click="writeConfig()">Save</v-btn>
  </v-layout>
</template>

<script>
import { configService } from '@/services/config'
export default {
  data: () => ({
    info: '',
    isTest: true,
    endpoint: '',
    config: configService
  }),
  methods: {
    onSubmit () {},
    onSuccess (data) {},
    loadConfig () {
      configService.read()
    },
    writeConfig () {
      configService.write()
      this.info = 'The page needs to be refreshed to apply changes'
    }
  },
  mounted () {
    this.loadConfig()
  }
}
</script>

<style scoped>
.tooltip {
  flex: 1 1 auto;
}
</style>
