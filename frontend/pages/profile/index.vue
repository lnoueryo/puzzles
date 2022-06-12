<template>
  <div>
    <card-profile :selectedUser="organization" :disabled="false"></card-profile>
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