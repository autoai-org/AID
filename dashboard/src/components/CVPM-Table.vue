<template>
  <v-card>
    <v-data-table :items="items" :headers="headers" class="elevation-1" :loading="loading">
      <v-progress-linear slot="progress" color="indigo" indeterminate></v-progress-linear>
      <template slot="items" slot-scope="props">
        <td class="text-xs-left">{{ props.item.Vendor }}</td>
        <td class="text-xs-left">{{ props.item.Name }}</td>
        <td class="text-xs-left">{{ props.item.LocalFolder }}</td>
        <td class="text-xs-left" v-if="hasActions">
          <v-icon small @click="inspectDetails(props.item)">info</v-icon>
        </td>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  data () {
    return {}
  },
  props: ['items', 'headers', 'loading'],
  computed: {
    hasActions () {
      if (this.headers.slice(-1)[0].text === 'Actions') {
        return true
      } else {
        return false
      }
    }
  },
  methods: {
    inspectDetails (item) {
      this.$router.push({
        name: 'Detail',
        params: { vendor: item.Vendor, name: item.Name }
      })
    }
  },
  created () {
    console.log(this.items)
  }
}
</script>

<style>
</style>
