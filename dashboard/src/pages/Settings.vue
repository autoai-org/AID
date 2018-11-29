<template>
  <v-layout>
    <v-flex md8="md8">
      <v-form>
        <v-text-field v-model="name" :counter="10" label="Name" required></v-text-field>
        <v-text-field v-model="email" label="E-mail" required></v-text-field>
      </v-form>
    </v-flex>
    <v-flex md4="md4">
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  data: () => ({
    name: '',
    email: ''
  }),
  methods: {
    onSubmit () {},
    onSuccess (data) {}
  }
}
</script>

<style>
</style>
