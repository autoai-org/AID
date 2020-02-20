<template>
  <div>
    <BaseHeading
      class="text-xs-center"
      text="Sign in"
      sub="to continue to Auto AI"
    />

    <VTextField
      v-model="internalValue"
      autofocus
      label="Email or phone"
      name="email"
    />

    <BaseText>Forgot email?</BaseText>

    <div class="my-5">
      <div class="subheading">
        Want to learn more about Auto AI?
      </div>
      <BaseText href="https://hub.autoai.org">
        Learn more
      </BaseText>
    </div>

    <VLayout align-center>
      <VFlex xs8>
        <BaseText to="/auth/signup">
          Create account
        </BaseText>
      </VFlex>
      <VFlex text-xs-right>
        <BaseBtn
          :disabled="!internalValue"
          class="ma-0"
          @click="$emit('next')"
        >
          Next
        </BaseBtn>
      </VFlex>
    </VLayout>
  </div>
</template>

<script>
// Utilities
import {
  mapMutations
} from 'vuex'

export default {
  name: 'Email',

  computed: {
    internalValue: {
      get () {
        return this.$store.state.email
      },
      set (val) {
        this.setEmail(val)
      }
    }
  },

  methods: {
    ...mapMutations(['setEmail'])
  }
}
</script>
