<template>
  <v-card
    class="mx-auto"
  >
    <slot></slot>
    <v-divider></v-divider>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn
        :loading="isLoading"
        class="white--text"
        color="amber darken-3"
        depressed
        @click="$emit('send')"
      >
        Submit
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from 'vue'
  export default Vue.extend({
    name: 'login',
    layout: 'login',
    data: () => ({
      isLoading: false,
    }),
  })
</script>