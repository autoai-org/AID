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
          :label="$t(`Settings.developer_mode`)"
        />
        <v-switch
          v-model="config.hintMode"
          :label="$t(`Settings.enable_hint`)"
        />
        <v-text-field
          v-model="config.endpoint"
          :label="$t(`Settings.endpoint`)"
          required
        />
      </v-form>
    </v-flex>
    <v-flex md4="md4" />
    <v-btn
      primary
      @click="writeConfig()"
    >
      {{$t(`Settings.save`)}}
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
