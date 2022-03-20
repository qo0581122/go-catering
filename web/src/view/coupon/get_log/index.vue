<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="UID">
            <el-input v-model="listQuery.uid" placeholder="UID"></el-input>
          </el-form-item>
          <el-form-item label="优惠券ID">
            <el-input
              v-model="listQuery.coupon_id"
              placeholder="优惠券ID"
            ></el-input>
          </el-form-item>
          <el-form-item label="获取方式">
            <el-select v-model="listQuery.get_type" placeholder="请选择">
              <el-option
                v-for="item in getTypeOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="使用状态">
            <el-select v-model="listQuery.use_status" placeholder="请选择">
              <el-option
                v-for="item in useStatusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="ID" align="center">
            <template #default="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column prop="uid" width="120px" label="UID" align="center">
            <template #default="scope">
              {{ scope.row.uid }}
            </template>
          </el-table-column>
          <el-table-column prop="coupon_id" label="优惠券ID" align="center">
            <template #default="scope">
              {{ scope.row.coupon_id }}
            </template>
          </el-table-column>
          <el-table-column
            prop="coupon_name"
            label="优惠券名称"
            align="center"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.coupon.coupon_name }}
            </template>
          </el-table-column>
          <el-table-column
            prop="get_time"
            align="center"
            label="获取时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.get_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="use_begin_time"
            align="center"
            label="使用开始时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.use_begin_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="created_at"
            align="center"
            label="使用结束时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.use_end_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="get_type"
            label="获取方式"
            width="200px"
            align="center"
          >
          <template #default="scope">
            <span v-for="(item, index) in getTypeFormat" v-bind:key="index">
                <el-tag v-if="scope.row.get_type == item.value" type="success" effect="dark">{{item.label}}</el-tag>
            </span>
          </template>
          </el-table-column>
          <el-table-column
            prop="use_status"
            label="使用状态"
            width="200px"
            align="center"
          >
          <template #default="scope">
            <span v-for="(item, index) in useStatusOptions" v-bind:key="index">
                <el-tag v-if="scope.row.use_status == item.value" type="success" effect="dark">{{item.label}}</el-tag>
            </span>
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
  </div>
</template>

<script>
import { fetchCouponGetLogPage } from "@/api/coupon_get_log.js";
import global_ from "@/components/global/global.vue";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  uid: null,
  coupon_id: null,
  use_status: null,
  get_type: null,
};
export default {
  name: "Coupon",
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
      getTypeOptions: global_.GetTypeOptions,
      getTypeFormat: global_.GetTypeFormatOptions,
      useStatusOptions: global_.UseStatusOptions,
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchCouponGetLogPage(this.listQuery)
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