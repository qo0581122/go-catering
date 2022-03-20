<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="代金券名称">
            <el-input
              v-model="listQuery.voucher_name"
              placeholder="代金券名称"
            ></el-input>
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="listQuery.status" placeholder="状态">
              <el-option label="启用" value="1"></el-option>
              <el-option label="禁用" value="2"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleAdd">添加</el-button>
          </el-form-item>
        </el-form>
        <el-table
          :data="data"
          style="width: 100%"
          :loading="loading"
          border
          align="center"
        >
          <el-table-column prop="id" label="ID" align="center">
          </el-table-column>
          <el-table-column prop="uid" width="120px" label="UID" align="center">
          </el-table-column>
          <el-table-column prop="order_type" label="订单类型">
            <template #default="scope">
              <span
                v-if="this.handleOrderTypeFormat(scope.row.order_type) == '-'"
                >-</span
              >
              <el-tag v-else>
                {{ this.handleOrderTypeFormat(scope.row.order_type) }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column label="取餐类型" align="center">
            <template #default="scope">
              <span
                v-if="this.handleOrderGetTypeFormat(scope.row.get_type) == '-'"
                >-</span
              >
              <el-tag v-else>
                {{ this.handleOrderGetTypeFormat(scope.row.get_type) }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column
            prop="get_time"
            label="取餐时间"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column prop="order_number" label="订单号" width="200px">
          </el-table-column>
          <el-table-column prop="order_status" label="订单状态">
            <template #default="scope">
              <span
                v-if="
                  this.handleOrderStatusFormat(scope.row.order_status) == '-'
                "
                >-</span
              >
              <el-tag v-else>
                {{
                  this.handleOrderStatusFormat(scope.row.order_status)
                }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column prop="integration" label="获取积分">
          </el-table-column>
          <el-table-column prop="product_count" label="商品数量" align="center">
          </el-table-column>
          <el-table-column
            prop="remark"
            label="备注"
            width="200px"
            align="center"
          >
          </el-table-column>
          <el-table-column label="订单餐厅ID" width="200px">
            <template #default="scope">
              {{ scope.row.order_shop.id }}
            </template>
          </el-table-column>
          <el-table-column label="订单餐厅名称" width="200px">
            <template #default="scope">
              {{ scope.row.order_shop.shop.shop_name }}
            </template>
          </el-table-column>
          <el-table-column label="订单餐厅号" width="200px">
            <template #default="scope">
              {{ scope.row.order_shop.shop_number }}
            </template>
          </el-table-column>
          <el-table-column label="订单交易号" width="200px">
            <template #default="scope">
              {{ scope.row.order_pay.pay_number }}
            </template>
          </el-table-column>
          <el-table-column label="订单支付类型" width="200px">
            <template #default="scope">
              <span
                v-if="
                  this.handleOrderPayTypeFormat(scope.row.order_pay.pay_type) ==
                  '-'
                "
                >-</span
              >
              <el-tag v-else>
                {{
                  this.handleOrderPayTypeFormat(scope.row.order_pay.pay_type)
                }}</el-tag
              >
            </template>
          </el-table-column>
          <el-table-column label="订单支付金额" width="200px">
            <template #default="scope">
              {{ scope.row.order_pay.pay_money / 100 }}
            </template>
          </el-table-column>
          <el-table-column label="订单折扣金额" width="200px">
            <template #default="scope">
              {{ scope.row.order_pay.discount_money / 100 }}
            </template>
          </el-table-column>
          <el-table-column label="订单总金额" width="200px">
            <template #default="scope">
              {{ scope.row.order_pay.total_money / 100 }}
            </template>
          </el-table-column>
          <el-table-column label="订单支付时间" width="200px">
            <template #default="scope">
              {{ scope.row.order_pay.pay_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="created_time"
            align="center"
            label="创建时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.created_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="updated_time"
            align="center"
            label="更新时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.updated_time }}
            </template>
          </el-table-column>
          <el-table-column
            fixed="right"
            label="操作"
            align="center"
            width="280px"
          >
            <template #default="scope">
              <el-button
                type="success"
                size="small"
                @click="handleShopDetail(scope.row)"
                >查看商铺详情</el-button
              >
              <el-button
                type="success"
                size="small"
                @click="handleProductList(scope.row)"
                >查看订单商品</el-button
              >
            </template>
          </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <div style="text-align: center">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="listQuery.pageNum"
            :page-sizes="pageSize"
            :page-size="listQuery.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
          >
          </el-pagination>
        </div>
      </el-footer>
    </el-container>
    <order-product-detail :visible="showProduct" :productIds="productIds" @close="handleDialogClose" />
  </div>
</template>

<script>
import { fetchOrderPage } from "@/api/order.js";
import global_ from "@/components/global/global.vue";
import OrderProductDetail from "./OrderProductDetail.vue"

const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  voucher_name: null,
  status: null,
};

export default {
  components: { OrderProductDetail },
  name: "OrderDetail",
  data() {
    return {
      pageSize: defaultPageSize,
      listQuery: Object.assign({}, defaultListQuery),
      total: null,
      dialogVisiable: false,
      isEdit: false,
      form: null,
      loading: false,
      data: null,
      orderTypeOptions: global_.OrderTypeOptions,
      orderGetTypeOptions: global_.OrderGetTypeOptions,
      orderStatusOptions: global_.OrderStatusOptions,
      orderPayTypeOptions: global_.OrderPayTypeOptions,
      showProduct: false,
      productIds: [],
    };
  },

  created() {
    this.handleListData();
  },
  methods: {
    handleDialogClose() {
      this.showProduct = false
    },
    handleProductList(data) {
      let ids = []
      data.order_products.forEach((item)=> {
        ids.push(item.product_id)
      })
      this.productIds = ids.join(",")
      this.showProduct = true
    },
    formatForEach(value, list) {
      for (let i = 0; i < list.length; i++) {
        if (value == list[i].value) {
          return list[i].label;
        }
      }
      return "-";
    },
    handleOrderTypeFormat(value) {
      return this.formatForEach(value, this.orderTypeOptions);
    },
    handleOrderGetTypeFormat(value) {
      return this.formatForEach(value, this.orderGetTypeOptions);
    },
    handleOrderStatusFormat(value) {
      return this.formatForEach(value, this.orderStatusOptions);
    },
    handleOrderPayTypeFormat(value) {
      return this.formatForEach(value, this.orderPayTypeOptions);
    },
    handleListData() {
      this.loading = true;
      fetchOrderPage(this.listQuery)
        .then((res) => {
          this.data = res.data.list;
          this.total = res.data.total;
          this.$message({
            message: "查询成功",
            type: "success",
          });
          this.loading = false;
        })
        .catch((res) => {
          console.log(res);
          this.$message({
            message: "查询失败",
            type: "error",
          });
          this.loading = false;
        });
    },
    handleSizeChange(val) {
      this.listQuery.pageSize = val;
      this.listQuery.pageNum = 1;
      this.handleListData();
    },
    handleCurrentChange(val) {
      this.listQuery.pageNum = val;
      this.handleListData();
    },
    handleClear() {
      this.listQuery = Object.assign({}, defaultListQuery);
    },
    handleQuery() {
      this.handleListData();
    },
  },
};
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>