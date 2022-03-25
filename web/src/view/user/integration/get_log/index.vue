<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="Uid">
            <el-input v-model="listQuery.uid" placeholder="Uid"></el-input>
          </el-form-item>
          <el-form-item label="积分获取方式">
            <el-select v-model="listQuery.change_type" placeholder="请选择">
            <el-option
              v-for="item in changeTypeOptions"
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
          <el-table-column prop="id" label="id" align="center">
          </el-table-column>
          <el-table-column prop="uid" width="120px" label="Uid" align="center">
          </el-table-column>
          <el-table-column
            prop="level"
            label="等级"
            align="center"
            width="200px"
          >
            <template #default="scope"> Lv{{ scope.row.level }} </template>
          </el-table-column>
          <el-table-column
            prop="level_name"
            label="等级称号"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column label="变更类型" align="center" width="100">
            <template #default="scope">
              <span
                v-for="(item, index) in changeTypeOptions"
                v-bind:key="index"
              >
                <el-tag
                  type="success"
                  v-if="scope.row.change_type == item.value"
                >
                  {{ item.label }}
                </el-tag>
              </span>
            </template>
          </el-table-column>
          <el-table-column
            prop="change_value"
            label="变更值"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="before_value"
            label="变更前积分值"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="after_value"
            label="变更后积分值"
            align="center"
            width="200px"
          >
          </el-table-column>

          <el-table-column
            prop="change_time"
            align="center"
            label="更新时间"
            width="200px"
          >
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
import { fetchUserIntegrationLogList } from "@/api/user_integration.js";
import global_ from "@/components/global/global.vue";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  uid: null,
  change_type: null,
};
export default {
  name: "Voucher",
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
      changeTypeOptions: global_.IntegrationChangeOptions,
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchUserIntegrationLogList(this.listQuery)
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