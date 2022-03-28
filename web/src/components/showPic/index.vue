<template>
  <span>
    <el-image
        :src="pic"
        style="width: 100px; height: 100px"
        @click="showPic"
      />
  </span>
  <el-dialog center title="" v-model="showDialog" width="50%" destroy-on-close style="text-align:center" append-to-body="true">
    <div style="text-align:center">
      <el-image
        :src="pic"
      />
    </div>
  </el-dialog>
</template>

<script>
const path = import.meta.env.VITE_BASE_API + '/'

export default {
  name: "ShowPic",
  data() {
    return {
      showDialog: false,
    }
  },
  props:{
    url: String,
  },
  computed: {
    pic: function() {
      if (this.url.length == 0) {
        return ''
      }
      if (this.url !== '' && this.url.slice(0, 4) === 'http') {
        return this.url
      }
      console.log(path + this.url)
      return path + this.url
    }
  },
  methods: {
    showPic() {
      this.showDialog = true
    }
  }
}
</script>

<style></style>