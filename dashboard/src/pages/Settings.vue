<template>
  <v-layout
    row
    wrap
  >
    <v-flex
      v-if="info"
      xs12
    >
      <v-alert
        outline
        class="cvpm-setting-alert"
        :value="info"
        type="info"
      >
        {{ info }}
      </v-alert>
    </v-flex>
    <v-flex
      md8="md8"
      xs12
    >
      <v-form>
        <v-switch
          v-model="config.developerMode"
          label="Developer Mode"
        />
        <v-switch
          v-model="config.hintMode"
          label="Enable Hint"
        />
        <v-text-field
          v-model="config.endpoint"
          label="Endpoint"
          required
        />
      </v-form>
    </v-flex>
    <v-flex md4="md4" />
    <v-btn
      primary
      @click="writeConfig()"
    >
      Save
    </v-btn>
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
  mounted () {
    this.loadConfig()
  },
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
  }
}
</script>

<style scoped>
.tooltip {
  flex: 1 1 auto;
}
</style>
