<template>
  <div class="terminal">
    <div class="header">
      <h4>{{title}}</h4>
        <ul class="shell-dots">
          <v-icon color="blue darken-2" @click="navHelp()">help_outline</v-icon>
        </ul>
    </div>
    <div
      style="position:absolute;top:0;left:0;right:0;overflow:auto;z-index:1;margin-top:30px;max-height:500px"
      ref="terminalWindow"
    >
      <div class="terminal-window" id="terminalWindow">
        <p v-for="(item, index) in parsedMessageList" :key="index">
          <span v-if="item.type" :class="item.type">{{item.type}}</span>
          <span class="cmd">{{item.text}}</span>
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
    }
  },
  computed: {
    parsedMessageList: function () {
      return this.messageList.map(function (each) {
        if (each.indexOf('error') > -1) {
          return {'type': 'error', 'text': each}
        } else {
          return {'text': each}
        }
      })
    }
  },
  props: ['title', 'messageList'],
  methods: {
    navHelp () {
      window.open('https://cvpm.autoai.org')
    }
  }
}
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
    .info {
      padding: 2px 3px;
      background: #2980b9;
    }
    .warning {
      padding: 2px 3px;
      background: #f39c12; // https://github.com/Mayccoll/Gogh/blob/master/content/themes.md #Flat
    }
    .success {
      padding: 2px 3px;
      background: #27ae60;
    }
    .error {
      padding: 2px 3px;
      background: #c0392b;
    }
    .system {
      padding: 2px 3px;
      background: #bdc3c7;
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