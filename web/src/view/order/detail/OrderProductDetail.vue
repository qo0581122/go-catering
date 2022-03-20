<template>
  <el-dialog
    v-model="visible"
    title="订单商品"
    width="60%"
    :before-close="handleClose"
  >
    <el-table :data="data" style="width: 100%" :loading="loading" border>
      <el-table-column
        prop="product_name"
        width="120px"
        label="商品名称"
        align="center"
      >
        <template #default="scope">
          {{ scope.row.product_name }}
        </template>
      </el-table-column>
      <el-table-column prop="description" label="描述" width="200px">
        <template #default="scope">
          {{ scope.row.description }}
        </template>
      </el-table-column>
      <el-table-column prop="category_name" label="商品分类" align="center">
        <template #default="scope">
          {{ scope.row.Category.category_name }}
        </template>
      </el-table-column>
      <el-table-column prop="specis" label="商品规格" align="center">
        <template #default="scope">
          <el-tag
            v-if="scope.row.specis == 1"
            type="success"
            effect="dark"
          >
            单品
          </el-tag>
          <el-tag v-else type="success" effect="dark"> 套餐 </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="url" label="图片地址" width="200px">
        <template #default="scope">
          {{ scope.row.url }}
        </template>
      </el-table-column>
      <el-table-column  label="原价（元）"  align="center">
        <template #default="scope">
          {{ scope.row.original_price / 100 }}
        </template>
      </el-table-column>
      <el-table-column  label="售价（元）"  align="center">
        <template #default="scope">
          {{ scope.row.pay_price / 100 }}
        </template>
      </el-table-column>
      <el-table-column  label="折扣"   align="center">
        <template #default="scope">
          {{ scope.row.discount / 10 }}
        </template>
      </el-table-column>
    </el-table>
  </el-dialog>
</template>

<script>
import { fetchOrderProduct } from "@/api/order.js";
export default {
  name: "OrderProductDetail",
  components: {},
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    productIds: {
      type: String,
    },
  },
  data() {
    return {
      data: [],
    };
  },
  watch: {
    productIds(newValue) {
      fetchOrderProduct({product_ids: newValue})
        .then((res) => {
          this.data = res.data
        })
        .catch((err) => {});
    },
    visible(newValue) {
    },
  },

  created() {},
  methods: {
    handleClose() {
      this.$emit("close")
    },
  },
};
</script>

<style>
</style>