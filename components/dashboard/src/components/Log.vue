<template>
  <div class="terminal">
    <confirm-dialog 
      :show="showDeleteDialog"
      @closed="showDeleteDialog=false"
      @cancelled="showDeleteDialog=false"
      @confirmed="performDeleteLog()"
      :info="deleteInfo"
      :title="'Are you sure?'" 
      :message="'You are going to delete log ' + title +' (' +logid+')?'">
    </confirm-dialog>
    <div class="header">
      <h4>{{ title }}</h4>
      <ul class="shell-dots">
        <v-icon color="blue darken-2" @click="navHelp()">mdi-help</v-icon>
        <v-icon color="blue darken-2" @click="showTime = !showTime">mdi-clock-outline</v-icon>
        <v-icon color="blue darken-2" @click="openDelete()">mdi-delete-outline</v-icon>
      </ul>
    </div>
    <div
      ref="terminalWindow"
      style="position:absolute;top:0;left:0;right:0;overflow:auto;z-index:1;margin-top:30px;max-height:500px"
    >
      <div id="terminalWindow" class="terminal-window">
        <pre v-for="item in messages" :key="item.title" v-html="renderMsg(item)" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { LogContent } from "@/entities";
import { deleteLog } from "@/middlewares/api.mdw"
import ConfirmDialog from "@/components/dialogs/ConfirmDialog.vue"
export default Vue.extend({
  props: {
    title: {
      type: String,
      default: ""
    },
    logid: {
      type: String,
      default: ""
    },
    messages: {
      type: Array,
      default: function() {
        return [];
      }
    }
  },
  components: {
    ConfirmDialog
  },
  data() {
    return {
      showTime: false,
      showDeleteDialog: false,
      deleteInfo: ""
    };
  },
  methods: {
    openDelete() {
      this.deleteInfo=''
      this.showDeleteDialog = true
    },
    navHelp() {
      window.open("https://aid.autoai.org");
    },
    performDeleteLog() {
      let self = this
      deleteLog(this.logid).then(function(res:any) {
        if(res.code===200) {
          self.deleteInfo = "Successfully deleted the log file"      
        }
      })
    },
    renderMsg(msgItem: LogContent) {
      let msg = "";
      msg =
        msg +
        "<span class=" +
        msgItem.level +
        ">[" +
        msgItem.level +
        "]</span>";
      if (this.showTime) {
        msg = msg + "<span class=success>[" + msgItem.time + "]</span>";
      }
      msg = msg + " " + msgItem.msg;
      msg = msg + "\n";
      msg = "<p>" + msg + "</p>";
      return msg;
    }
  }
});
</script>

<style lang="scss" scoped>
.terminal {
  position: relative;
  width: 100%;
  border-radius: 4px;
  color: white;
  margin-bottom: 10px;
  max-height: 580px;
}
.terminal .terminal-window {
  padding-top: 50px;
  background-color: rgb(3, 9, 36);
  min-height: 140px;
  padding: 20px;
  font-weight: normal;
  font-family: Monaco, Menlo, Consolas, monospace;
  color: #fff;
  pre {
    font-family: Monaco, Menlo, Consolas, monospace;
    white-space: pre-wrap;
  }
  p {
    overflow-wrap: break-word;
    word-break: break-all;
    font-size: 13px;
    .cmd {
      line-height: 24px;
    }
  }
  pre {
    display: inline;
  }
}
.terminal .header ul.shell-dots li {
  display: inline-block;
  width: 12px;
  height: 12px;
  border-radius: 6px;
  background-color: rgb(3, 9, 36);
  margin-left: 6px;
}
.terminal .header ul.shell-dots li.red {
  background-color: rgb(200, 48, 48);
}
.terminal .header ul.shell-dots li.yellow {
  background-color: rgb(247, 219, 96);
}
.terminal .header ul.shell-dots li.green {
  background-color: rgb(46, 201, 113);
}
.terminal .header {
  position: absolute;
  z-index: 2;
  top: 0;
  right: 0;
  left: 0;
  background-color: rgb(149, 149, 152);
  text-align: center;
  padding: 2px;
  border-top-left-radius: 4px;
  border-top-right-radius: 4px;
}
.terminal .header h4 {
  font-size: 14px;
  margin: 5px;
  letter-spacing: 1px;
}
.terminal .header ul.shell-dots {
  position: absolute;
  top: 5px;
  left: 8px;
  padding-left: 0;
  margin: 0;
}
.terminal .terminal-window .prompt::before {
  content: "$";
  margin-right: 10px;
}
.terminal .terminal-window .cursor {
  margin: 0;
  background-color: white;
  animation: blink 1s step-end infinite;
  -webkit-animation: blink 1s step-end infinite;
  margin-left: -5px;
}
@keyframes blink {
  50% {
    visibility: hidden;
  }
}
@-webkit-keyframes blink {
  50% {
    visibility: hidden;
  }
}
.terminal .terminal-window .loading {
  display: inline-block;
  width: 0;
  overflow: hidden;
  animation: load 1.2s step-end infinite;
  -webkit-animation: load 1.2s step-end infinite;
}
@keyframes load {
  0% {
    width: 0px;
  }
  20% {
    width: 5px;
  }
  40% {
    width: 10px;
  }
  60% {
    width: 15px;
  }
  80% {
    width: 20px;
  }
}
@-webkit-keyframes load {
  0% {
    width: 0px;
  }
  20% {
    width: 5px;
  }
  40% {
    width: 10px;
  }
  60% {
    width: 15px;
  }
  80% {
    width: 20px;
  }
}
.terminal-last-line {
  font-size: 0;
  word-spacing: 0;
  letter-spacing: 0;
}
.input-box {
  position: relative;
  background: rgb(3, 9, 36);
  border: none;
  width: 1px;
  opacity: 0;
  cursor: default;
}
.input-box:focus {
  outline: none;
  border: none;
}
</style>