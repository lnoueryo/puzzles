<template>
  <div>
    <profile :selectedUser="organization" :disabled="false"/>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
export default Vue.extend({
  computed: {
    ...mapGetters([
      'organization'
    ])
  }
})
</script>