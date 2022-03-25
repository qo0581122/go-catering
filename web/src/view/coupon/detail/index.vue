<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="优惠券名称">
            <el-input
              v-model="listQuery.coupon_name"
              placeholder="优惠券名称"
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
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="id" align="center">
            <template #default="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column
            prop="coupon_name"
            width="120px"
            label="优惠券名称"
            align="center"
          >
            <template #default="scope">
              {{ scope.row.coupon_name }}
            </template>
          </el-table-column>
          <el-table-column prop="description" label="描述" width="200px">
            <template #default="scope">
              {{ scope.row.description }}
            </template>
          </el-table-column>
          <el-table-column prop="price" label="优惠后价格" align="center">
            <template #default="scope">
              {{ scope.row.price }}
            </template>
          </el-table-column>
          <el-table-column prop="price" label="最低使用订单价格" align="center">
            <template #default="scope">
              {{ scope.row.least_use_price }}
            </template>
          </el-table-column>
          <el-table-column prop="specis" label="商品名称" align="center">
            <template #default="scope">
              {{ scope.row.product.product_name }}
            </template>
          </el-table-column>
          <el-table-column prop="url" label="图片地址" width="200px">
            <template #default="scope">
              {{ scope.row.pic_url }}
            </template>
          </el-table-column>
          <el-table-column prop="url" label="发放数量" width="200px">
            <template #default="scope">
              {{ scope.row.total_count }}
            </template>
          </el-table-column>
          <el-table-column prop="url" label="剩余数量" width="200px">
            <template #default="scope">
              {{ scope.row.remain_count }}
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
                <el-tag
                  v-if="scope.row.get_type == item.value"
                  type="success"
                  effect="dark"
                  >{{ item.label }}</el-tag
                >
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="url" label="单用户限定获取数量" width="200px">
            <template #default="scope">
              {{ scope.row.get_count }}
            </template>
          </el-table-column>
          <el-table-column
            prop="valid_time_type"
            label="有效时间计算方式"
            width="200px"
            align="center"
          >
            <template #default="scope">
              <span v-for="(item, index) in validTimeFormat" v-bind:key="index">
                <el-tag
                  v-if="scope.row.valid_time_type == item.value"
                  type="success"
                  effect="dark"
                  >{{ item.label }}</el-tag
                >
              </span>
            </template>
          </el-table-column>
          <el-table-column prop="url" label="获取开始时间" width="200px">
            <template #default="scope">
              {{ scope.row.get_begin_time }}
            </template>
          </el-table-column>
          <el-table-column prop="url" label="获取结束时间" width="200px">
            <template #default="scope">
              {{ scope.row.get_end_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="valid_time_type"
            label="使用开始时间"
            width="200px"
          >
            <template #default="scope">
              <div v-if="scope.row.valid_time_type == 2">
                用户领取时间
              </div>
              <div v-else>
                {{ scope.row.use_begin_time }}
              </div>
            </template>
          </el-table-column>
          <el-table-column
            prop="use_end_time"
            label="使用结束时间"
            width="200px"
          >
            <template #default="scope">
              <div v-if="scope.row.valid_time_type == 2">
                用户领取后<span style="color: red">{{
                  scope.row.valid_time_type
                }}</span
                >天过期
              </div>
              <div v-else>
                {{ scope.row.use_end_time }}
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="status" align="center" label="状态">
            <template #default="scope">
              <el-tag
                v-if="scope.row.status == 1"
                type="success"
                effect="dark"
              >
                启用
              </el-tag>
              <el-tag v-else type="danger" effect="dark"> 禁用 </el-tag>
            </template>
          </el-table-column>
          <el-table-column
            prop="created_at"
            align="center"
            label="创建时间"
            width="200px"
          >
            <template #default="scope">
              {{ scope.row.created_time }}
            </template>
          </el-table-column>
          <el-table-column
            prop="updated_at"
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
            width="200px"
          >
            <template #default="scope">
              <el-button
                @click="handleUpdate(scope.row)"
                type="success"
                size="small"
                >编辑</el-button
              >
              <el-button
                type="success"
                size="small"
                @click="handleDelete(scope.row)"
                >删除</el-button
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
    <dform
      :dialogVisiable="dialogVisiable"
      :is-edit="isEdit"
      :form="form"
      @visableChange="handleCloseDialog"
      @submit="handleSubmit"
    ></dform>
  </div>
</template>

<script>
import Dform from "./form.vue";
import {
  fetchCouponPage,
  createCoupon,
  updateCoupon,
  deleteCoupon,
} from "@/api/coupon.js";
import global_ from "@/components/global/global.vue";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  coupon_name: null,
  status: null,
};
const defaultForm = {
  id: undefined,
  coupon_name: null,
  price: 0,
  least_use_price: 0,
  product_id: null,
  pic_url: null,
  total_count: 0,
  remain_count: 0,
  get_type: null,
  valid_time_type: null,
  valid_day: 0,
  use_type: null,
  status: 1,
  get_begin_time: null,
  get_end_time: null,
  use_begin_time: null,
  use_end_time: null,
  get_count: 1,
  description: null,
};
export default {
  components: { Dform },
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
      getTypeFormat: global_.GetTypeFormatOptions,
      validTimeFormat: global_.ValidTimeTypeOptions,
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchCouponPage(this.listQuery)
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
    handleCloseDialog(value) {
      this.dialogVisiable = false;
      this.form = Object.assign({}, defaultForm);
    },
    handleUpdate(data) {
      console.log(data);
      this.dialogVisiable = true;
      this.isEdit = true;
      this.form = {
        id: data.coupon.id,
        coupon_name: data.coupon.coupon_name,
        price: data.coupon.price,
        least_use_price: data.coupon.least_use_price,
        product_id: data.coupon.product_id,
        pic_url: data.coupon.pic_url,
        total_count: data.coupon.total_count,
        remain_count: data.coupon.remain_count,
        get_type: data.coupon.get_type,
        valid_time_type: data.coupon.valid_time_type,
        valid_day: data.coupon.valid_day,
        use_type: data.coupon.use_type,
        status: data.coupon.status,
        get_begin_time: data.coupon.get_begin_time,
        get_end_time: data.coupon.get_end_time,
        use_begin_time: data.coupon.use_begin_time,
        use_end_time: data.coupon.use_end_time,
        get_count: data.coupon.get_count,
        description: data.coupon.description,
      };
    },
    handleAdd() {
      this.dialogVisiable = true;
      this.isEdit = false;
      this.form = Object.assign({}, defaultForm);
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
    handleSubmit(data) {
      let form = Object.assign({}, data);
      console.log(data);
      form.status = form.status == true ? 1 : 2;
      if (this.isEdit) {
        //更新
        updateCoupon(form)
          .then((res) => {
            this.$message({
              message: res.data,
              type: "success",
            });
            this.handleListData();
          })
          .catch((res) => {
            console.log(res);
            this.$message({
              message: res.data,
              type: "error",
            });
          });
      } else {
        //增加
        createCoupon(form)
          .then((res) => {
            this.$message({
              message: res.data,
              type: "success",
            });
            this.handleListData();
          })
          .catch((res) => {
            this.$message({
              message: res.data,
              type: "error",
            });
          });
      }
      this.handleCloseDialog();
    },
    handleDelete(data) {
      this.$confirm("此操作将永久删除该记录, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        center: true,
      })
        .then(() => {
          deleteCoupon(data.coupon.id)
            .then((res) => {
              this.$message({
                message: res.data,
                type: "success",
              });
              this.handleListData();
            })
            .catch((res) => {
              this.$message({
                message: res.data,
                type: "error",
              });
            });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
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