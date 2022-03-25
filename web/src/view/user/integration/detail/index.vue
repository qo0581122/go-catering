<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="Uid">
            <el-input v-model="listQuery.uid" placeholder="Uid"></el-input>
          </el-form-item>
          <el-form-item label="等级id">
            <el-input
              v-model="listQuery.level_id"
              placeholder="等级id"
            ></el-input>
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
            prop="level_id"
            label="等级id"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="level"
            label="等级"
            align="center"
            width="200px"
          >
            <template #default="scope">
              Lv{{scope.row.level}}
            </template>
          </el-table-column>
          <el-table-column
            prop="level_name"
            label="等级称号"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="current_integration"
            label="当前积分值"
            align="center"
          ></el-table-column>
          <el-table-column
            prop="history_integration"
            label="历史最高积分值"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="consume_integration"
            label="已消费积分值"
            align="center"
            width="200px"
          >
          </el-table-column>

          <el-table-column
            prop="updated_time"
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
import { fetchUserIntegrationList } from "@/api/user_integration.js";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  uid: null,
  level_id: null,
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
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchUserIntegrationList(this.listQuery)
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