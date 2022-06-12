<template>
  <v-card class="mx-auto">
    <slot name="main"></slot>
    <v-divider></v-divider>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn
        :disabled="loading || !formReady"
        :loading="loading"
        class="white--text"
        color="#295caa"
        depressed
        @click="$emit('send')"
      >
        <slot name="button"></slot>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'login',
  layout: 'login',
  props: {
    formReady: Boolean,
    loading: Boolean,
  },
})
</script>