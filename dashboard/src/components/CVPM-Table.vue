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
        <td class="text-xs-left">
          <v-icon
            color="indigo"
            small
            @click="inspectDetails(props.item)"
          >
            info
          </v-icon>
          <v-icon
            color="indigo"
            small
            @click="showTrainHistory()(props.item)"
          >
            fas fa-book-reader
          </v-icon>
        </td>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  props: {
    config: {
      type: Object,
      default: function () {
        return {}
      }
    },
    headers: {
      type: Array,
      default: function () {
        return []
      }
    },
    loading: {
      type: Boolean,
      default: false
    },
    items: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  data () {
    return {}
  },
  created () {},
  methods: {
    inspectDetails (item) {
      this.$router.push({
        name: 'Detail',
        params: { vendor: item.Vendor, name: item.Name }
      })
    },
    showTrainHistory (item) {
      console.log(item)
    }
  }
}
</script>

<style>
</style>
