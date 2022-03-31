<template>
  <div class="cell text-left" :style="styleValue">
    <v-btn class="mx-1" icon>
      <v-avatar size="25px">
        <img alt="Avatar" :src="'http://localhost:8080/media/users/' + user.image">
      </v-avatar>
    </v-btn>
    {{ user.name }}
  </div>
</template>

<script>
import Vue from 'vue'
export default Vue.extend({
  props: ['styleValue', 'user']
})
</script>