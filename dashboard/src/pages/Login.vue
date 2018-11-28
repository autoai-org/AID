<template>
<div>
    <v-dialog :value="true" persistent="">
        <v-card hover="" style="background:white;">
            <v-card-row class="deep-purple darken-1">
                <v-card-title class="white--text">
                    <div class="text-xs-center"> {{$t("Login")}}</div>
                </v-card-title>
            </v-card-row>
            <v-card-row>
                <v-card-text class="pt-4">
                    <v-form v-model="model" action="login" :fields="fields" @success="onSuccess" submitButtonText="Login">
                        <div class="flex pb-2"><small>{{$t("* Indicates required field")}}</small></div>
                    </v-form>
                </v-card-text>
            </v-card-row>
        </v-card>
    </v-dialog>
</div>
</template>

<script>
export default {
  data () {
    return {
      model: {
        username: 'admin',
        password: '123456'
      },
      fields: {
        username: { label: 'Username' },
        password: { label: 'Password', type: 'password' }
      },
      show: true
    }
  },
  methods: {
    onSuccess (data) {
      this.$store.commit('setAuth', data)
      this.$router.replace('/')
    }
  },
  mounted () {
  }
}
</script>

<style scoped>
  body{
    background: #666 !important;
  }
</style>