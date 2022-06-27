<template>
  <div>
    <card-profile :selectedUser="organizationAuthority" :editable="true" />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapGetters } from 'vuex'
export default Vue.extend({
  computed: {
    ...mapGetters([
      'organizationAuthority'
    ])
  }
})
</script>