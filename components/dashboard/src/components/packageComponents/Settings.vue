<template>
  <v-expansion-panels flat>
    <v-expansion-panel flat>
      <v-expansion-panel-header>Environment Variables</v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-card flat color="transparent">
          <v-card-text>
            <v-list-item v-for="(item, index) in variables" :key="index" width="80%">
              <v-list-item-content>
                <v-row class="mb-12">
                  <v-col cols="4">
                    <v-text-field v-model="variables[index]['key']" label="Key" required></v-text-field>
                  </v-col>
                  <v-col cols="6">
                    <v-text-field
                      @click:append="showValue[index] = !showValue[index]; update()"
                      :type="showValue[index] ? 'text' : 'password'"
                      :append-icon="showValue[index] ? 'mdi-eye' : 'mdi-eye-off'"
                      v-model="variables[index]['value']"
                      label="Value"
                      required
                    ></v-text-field>
                  </v-col>
                  <v-col cols="2">
                    <v-btn color="red" icon @click="remove(index)">
                      <v-icon>mdi-delete-outline</v-icon>
                    </v-btn>
                  </v-col>
                </v-row>
              </v-list-item-content>
            </v-list-item>
          </v-card-text>
          <v-card-actions>
            <v-btn color="secondary" @click="addNewEnvironmentVariable">Add</v-btn>
            <v-btn color="primary" @click="save">Save</v-btn>
          </v-card-actions>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      variables: [] as object[],
      showValue: [] as Boolean[],
    };
  },
  methods: {
    addNewEnvironmentVariable() {
      this.variables.push({
        key: "",
        value: ""
      });
      this.showValue.push(false);
    },
    save() {},
    remove(index: number) {
      this.variables.splice(index, 1);
    },
    update() {
      this.$forceUpdate()
    }
  },
});
</script>
<style scoped>
</style>