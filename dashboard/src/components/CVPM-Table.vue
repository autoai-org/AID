<template>
  <v-card>
    <v-data-table
      :items="items"
      :headers="headers"
      class="elevation-1"
      :loading="loading"
    >
      <v-progress-linear
        slot="progress"
        color="indigo"
        indeterminate
      />
      <template
        slot="items"
        slot-scope="props"
      >
        <td class="text-xs-left">
          {{ props.item.Vendor }}
        </td>
        <td class="text-xs-left">
          {{ props.item.Name }}
        </td>
        <td class="text-xs-left">
          {{ props.item.LocalFolder }}
        </td>
        <td
          v-if="hasActions"
          class="text-xs-left"
        >
          <v-icon
            small
            @click="inspectDetails(props.item)"
          >
            info
          </v-icon>
        </td>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  props: ['items', 'headers', 'loading'],
  data () {
    return {}
  },
  computed: {
    hasActions () {
      if (this.headers.slice(-1)[0].text === 'Actions') {
        return true
      } else {
        return false
      }
    }
  },
  created () {
  },
  methods: {
    inspectDetails (item) {
      this.$router.push({
        name: 'Detail',
        params: { vendor: item.Vendor, name: item.Name }
      })
    }
  }
}
</script>

<style>
</style>
