<template>
    <v-expansion-panels flat>
        <v-expansion-panel flat>
            <v-expansion-panel-header>
                Dive
            </v-expansion-panel-header>
            <v-expansion-panel-content>
                <pre v-html="logs"></pre>
            </v-expansion-panel-content>
        </v-expansion-panel>
    </v-expansion-panels>
</template>
<script lang="ts">
import Vue from 'vue'
import { getExtLogs } from '@/middlewares/api.mdw'
import { ansi2html } from '@/middlewares/logs.mdw'
export default Vue.extend({
    data() {
        return {
            logs : ""
        }
    },
    mounted() {
        let self = this
        getExtLogs("dive", "1").then(function(res:any){
            self.logs = ansi2html(res.content)
        })
    }
})
</script>
<style scoped>

</style>