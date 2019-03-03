<template>
  <v-container>
    <v-card>
      <v-data-table
        :items="datasets"
        :headers="headers"
        class="elevation-1"
      >
        <template
          slot="items"
          slot-scope="props"
        >
          <td class="text-xs-left">
            {{ props.item.Name }}
          </td>
          <td class="text-xs-left">
            {{ props.item.Desc }}
          </td>
          <td class="text-xs-left">
            <v-chip
              v-for="(item, index) in props.item.Tags"
              :key="index"
              color="primary"
              text-color="white"
              @click="setSearchTag(item);search()"
            >
              <B
                v-if="item===searchTag"
                style="color:yellow"
              >
                <I>{{ item }}</I>
              </B>
              <span v-else>{{ item }}</span>
            </v-chip>
          </td>
          <v-btn
            outline
            small
            fab
            color="indigo"
            @click="viewDetail(props.item)"
          >
            <v-icon>info</v-icon>
          </v-btn>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script>
import { systemService } from '@/services/system'

export default {
  data () {
    return {
      inspectInfo: []
    }
  },
  created () {
    this.fetchInspector()
  },
  methods: {
    fetchInspector () {
      let self = this
      systemService.getInspectorInfo().then(function (res) {
        self.inspectInfo = res
        console.log(res)
      })
    }
  }
}
</script>

<style>
</style>
