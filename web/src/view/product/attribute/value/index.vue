<template>
  <div>
      <el-container style="margin: 15px 20px 20px 20px">
        <el-main>
            <el-form :inline="true" >
                <el-form-item label="属性值">
                    <el-input v-model="listQuery.attribute_value" placeholder="属性值"></el-input>
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
            style="width: 100%" :loading="loading" border >
                <el-table-column
                    prop="id"
                    label="id"
                    >
                </el-table-column>
                <el-table-column
                    prop="attribute_name"
                    label="属性名称"
                    >
                    <template #default="scope">
                        {{ scope.row.attribute.attribute_name }}
                    </template>
                </el-table-column>
                <el-table-column
                    prop="attribute_value"
                    label="属性值"
                    >
                </el-table-column>
                <el-table-column
                    prop="status"
                    label="状态">
                    <template #default="scope">
                        <el-tag
                            v-if="scope.row.status == 1"
                            type="success"
                            effect="dark">
                           启用
                        </el-tag>
                        <el-tag
                            v-else
                            type="danger"
                            effect="dark">
                           禁用
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column
                    prop="created_time"
                    label="创建时间" width="200px">
                </el-table-column>
                <el-table-column
                    prop="updated_time"
                    label="更新时间" width="200px">
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="200px">
                    <template #default="scope">
                        <el-button @click="handleUpdate(scope.row)" type="success" size="small">编辑</el-button>
                        <el-button type="success" size="small" @click="handleDelete(scope.row)">删除</el-button>
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
                    :total="total">
                </el-pagination>
            </div>
        </el-footer>
    </el-container>
    <dform :dialogVisiable="dialogVisiable" :is-edit="isEdit" :form="form" @visableChange="handleCloseDialog" @submit="handleSubmit"></dform>
  </div>
</template>

<script>
import Dform from './form.vue'
import {fetchProductAttributeValueList, createProductAttributeValue, updateProductAttributeValue, deleteProductAttributeValue} from '@/api/product_attribute_value.js'
const defaultPageSize = [10, 20, 50, 100, 200]
const defaultListQuery = {
    pageSize: 20,
    pageNum: 1,
    attribute_value: '',
    status: null
}
const defaultForm = {
    attribute_id: null,
    attribute_value: '',
    status: true
}
export default {
    components: { Dform },
    name: 'Provinces',
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
        }
    },
    created() {
        this.handleListData()
    },
    methods: {
        handleListData() {
            this.loading = true
            fetchProductAttributeValueList(this.listQuery).then(res => {
                this.data = res.data.list
                this.total = res.data.total
                this.$message({
                    message: '查询成功',
                    type: 'success'
                });
                this.loading = false
            }).catch(res => {
                console.log(res)
                this.$message({
                    message: '查询失败',
                    type: 'error'
                });
                this.loading = false
            })
        },
        handleCloseDialog(value) {
            this.dialogVisiable = false
            this.form = Object.assign({}, defaultForm)
        },
        handleUpdate(data) {
            this.dialogVisiable = true
            this.isEdit = true
            this.form = data
        },
        handleAdd() {
            this.dialogVisiable = true
            this.isEdit = false
            this.form = Object.assign({}, defaultForm)
        },
        handleSizeChange(val) {
            this.listQuery.pageSize = val
            this.listQuery.pageNum = 1
            this.handleListData()
        },
        handleCurrentChange(val) {
            this.listQuery.pageNum = val
            this.handleListData()
        },
        handleClear() {
            this.listQuery = Object.assign({}, defaultListQuery)
        },
        handleQuery() {
            this.handleListData()
        },
        handleSubmit(data) {
            let form = Object.assign({}, data)
            console.log(data)
            form.status = form.status == true ? 1 : 2
            if (this.isEdit) {
                //更新
                updateProductAttributeValue(form).then(res => {
                    this.$message({
                        message: res.message,
                        type: 'success'
                    });
                    this.handleListData()
                }).catch(res => {
                    console.log(res)
                    this.$message({
                        message: res.message,
                        type: 'error'
                    });
                })
            }else {
                //增加
                createProductAttributeValue(form).then( res=> {
                    this.$message({
                        message: res.message,
                        type: 'success'
                    });
                     this.handleListData()
                }).catch(res => {
                    this.$message({
                        message: res.message,
                        type: 'error'
                    });
                })
            }
            this.handleCloseDialog()
        },
        handleDelete(data) {
            this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
                center: true,
            })
            .then(() => {
                deleteProductAttributeValue(data.id)
                    .then((res) => {
                    this.$message({
                        message: res.message,
                        type: 'success',
                    })
                    this.handleListData()
                    })
                    .catch((res) => {
                    this.$message({
                        message: res.message,
                        type: 'error',
                    })
                    })
                })
            .catch(() => {
                this.$message({
                    type: 'info',
                    message: '已取消删除',
                })
            })
        }
    }
}
</script>

<style>
  .el-table .warning-row {
    background: oldlace;
  }

  .el-table .success-row {
    background: #f0f9eb;
  }
</style>