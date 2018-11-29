<template>
  <v-layout>
    <v-flex md8="md8">
      <v-form>
        <v-switch label="Developer Mode" v-model="config.developerMode">
        </v-switch>
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
    }
  },
  mounted () {
    this.loadConfig()
  }
}
</script>

<style scoped>
.tooltip {
    flex: 1 1 auto
}
</style>
