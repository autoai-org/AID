<template>
  <v-card width="100%">
    <v-container fluid>
      <v-row align="center">
        <v-col class="d-flex" cols="12" sm="6">
          <v-select :items="logs.map(a=>a.Source)" label="Sources" v-model="current_source"></v-select>
        </v-col>
        <v-col class="d-flex" cols="12" sm="6">
          <v-select
            :items="logs.filter(a => a.Source===current_source).map(a=>a.Title)"
            label="Title"
            v-model="current_title"
          ></v-select>
        </v-col>
      </v-row>
    </v-container>
    <comp-log :title="title" :messages="messages" />
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Log from "@/components/Log.vue";
import { LogContent } from "@/entities";
import { mapState } from "vuex";
import { watchLog } from "@/middlewares/ws.mdw";
import { fetchAllObjects, fetchLog } from "@/middlewares/api.mdw";
export default Vue.extend({
  data: () => ({
    title: "system",
    messages: [],
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
      text = text.replace(/(\r\n|\n|\r)/gm, "");
      text = "[" + text + "]";
      text = text.replace(/}{/g, "},{");
      let content = JSON.parse(text);
      let messages = [] as any;
      if (source === "Default") {
        content.forEach((element: any) => {
          messages.push({
            level: element.level,
            time: element.time,
            msg: element.msg
          });
        });
      } else if (source === "docker-build") {
        content.forEach((element: any) => {
          let msgContent = JSON.parse(element.msg).stream
          if (msgContent !== "\n" && msgContent) {
            msgContent = msgContent.replace(/(\r\n|\n|\r)/gm, "")
            msgContent = msgContent.trim()
            messages.push({
              level: element.level,
              time: element.time,
              msg: msgContent
            });
          }
        });
      }
      return messages;
    }
  },
  watch: {
    current_title(value) {
      this.title = value
      let self = this;
      const found = this.logs.find((item: any) => item.Title === value);
      fetchLog(found.ID).then(function(res: any) {
        self.messages = self.renderMsg(res.content, found.Source);
      });
    }
  },
  mounted() {
    let self = this;
    fetchAllObjects("logs");
    fetchLog(0).then(function(res: any) {
      self.messages = self.renderMsg(res.content, "Default");
    });
  }
});
</script>

<style scoped>
</style>