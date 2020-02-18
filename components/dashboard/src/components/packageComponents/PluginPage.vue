<template>
    <div>
        <pre>{{logs}}</pre>
    </div>
</template>
<script lang="ts">
import Vue from 'vue'
import { getExtLogs } from '@/middlewares/api.mdw'
export default Vue.extend({
    data() {
        return {
            logs : ""
        }
    },
    mounted() {
        let self = this
        getExtLogs("dive", "1").then(function(res:any){
            self.logs = res.content
        })
    }
})
</script>
<style scoped>

</style>