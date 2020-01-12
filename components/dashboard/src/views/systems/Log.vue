<template>
  <v-card width="100%">
    <v-container fluid>
      <v-row align="center">
        <v-col class="d-flex" cols="12" sm="6">
          <v-select :items="logs.map(a=>a.Source)" label="Sources" v-model="current_source"></v-select>
        </v-col>
        <v-col class="d-flex" cols="12" sm="6">
          <v-select
            :items="logs.filter(a => a.Source===current_source).map(a=>processLogItem(a))"
            label="Title"
            v-model="current_title"
          ></v-select>
        </v-col>
      </v-row>
    </v-container>
    <comp-log :title="title" :messages="messages" :logid="id" />
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
    id:"",
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
    processLogItem(item:any) {
      return {
        text: item.CreatedAt + " " + item.Title,
        value: item.ID
      }
    },
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
          let msgContent = JSON.parse(element.msg).stream;
          if (msgContent !== "\n" && msgContent) {
            msgContent = msgContent.replace(/(\r\n|\n|\r)/gm, "");
            msgContent = msgContent.trim();
            messages.push({
              level: element.level,
              time: element.time,
              msg: msgContent
            });
          }
        });
      }
      return messages;
    },
    getAndWatchLog(logid: string, source: string) {
      let self = this;
      fetchLog(logid).then(function(res: any) {
        self.id = logid
        self.messages = self.renderMsg(res.content, source);
        watchLog(logid, (wsres: any) => {
          let lines = wsres.split("\n");
          lines.splice(0, 1);
          lines = lines.join("\n");
          self.messages = self.renderMsg(lines, source);
        });
      });
    }
  },
  watch: {
    current_title(value) {
      console.log(value)
      let self = this;
      const found = this.logs.find((item: any) => item.ID === value);
      if (found) {
        this.id = found.ID
        this.title = found.CreatedAt + " " + found.Title;
        this.getAndWatchLog(found.ID, found.Source);
      }
    }
  },
  mounted() {
    let logid = this.$route.query.logid || "";
    let self = this;
    fetchAllObjects("logs").then(function(res: any) {
      self.current_source = res[0].Source;
      self.current_title = res[0].ID;
      self.id = res[0].ID
    });
    if (logid) {
      const found = this.logs.find((item: any) => item.ID === logid);
      if (found) {
        this.getAndWatchLog(logid.toString(), found.source);
      }
    }
  }
});
</script>

<style scoped>
</style>