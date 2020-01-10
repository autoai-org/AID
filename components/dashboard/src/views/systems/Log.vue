<template>
  <v-card width="100%">
    <v-row>
      <v-col>
        <v-select :items="logs.map(a=>a.Source)" label="Sources" v-model="current_source"></v-select>
      </v-col>
      <v-col>
        <v-select
          :items="logs.filter(a => a.Source===current_source).map(a=>a.Title)"
          label="Title"
          v-model="current_title"
        ></v-select>
      </v-col>
    </v-row>
    <comp-log :title="title" :message="message" />
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Log from "@/components/Log.vue";
import { mapState } from "vuex";
import { fetchSystemLog } from "@/middlewares/ws.mdw";
import { fetchAllLogs, fetchLog } from "@/middlewares/api.mdw";
export default Vue.extend({
  data: () => ({
    title: "System Log",
    message: "",
    current_source: "",
    current_title: ""
  }),
  components: {
    "comp-log": Log
  },
  computed: {
    ...mapState({
      logs: "logs"
    })
  },
  methods: {
    renderMsg(text: string, source: string) {
      text = text.replace(/(\r\n|\n|\r)/gm,"")
      text = '[' + text + ']'
      text = text.replace(/}{/g, "},{")
      var content = JSON.parse(text)
      if (source === "Default") {
        var message = ""
        content.forEach((element:any) => {
          message += "[" + element.level+"] " +element.time+" "+ element.msg + "\n"
        });
        return message
      }
    }
  },
  watch: {
    current_title(value) {
      let self = this
      const found = this.logs.find((item:any) => item.Title === value);
      fetchLog(found.ID).then(function(res:any) {
        self.message = self.renderMsg(res.content, found.Source) || ""
      })
    }
  },
  mounted() {
    let self = this;
    fetchAllLogs();
    fetchLog(0).then(function(res: any) {
      self.message = res;
    });
  }
});
</script>

<style scoped>
</style>