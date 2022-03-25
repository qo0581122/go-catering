<template>
    <div>
    <el-dialog
        :title="dialogTitle"
        center
        v-model="dialogVisiable"
        width="500px"
        :before-close="handleClose">
        <el-form ref="dialogForm" :model="dialogForm" label-width="80px" label-position="right" :rules="formRules">
            <div v-if="isEdit">
                <el-form-item label="id">
                    <el-input v-model="dialogForm.id"  :disabled="isEdit"></el-input>
                </el-form-item>
            </div>
            <el-form-item label="省份" prop="province_id">
                <el-select v-model="dialogForm.province_id" placeholder="请选择" @change="handleGetCitys">
                    <el-option
                    v-for="item in provinces"
                    :key="item.id"
                    :label="item.province_name"
                    :value="item.id">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="城市" prop="city_id">
                <el-select v-model="dialogForm.city_id" placeholder="请选择">
                    <el-option
                    v-for="item in citys"
                    :key="item.id"
                    :label="item.city_name"
                    :value="item.id">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="区域名称" prop="district_name">
                <el-input v-model="dialogForm.district_name"></el-input>
            </el-form-item>
            <el-form-item label="状态">
                  <el-switch v-model="dialogForm.status"></el-switch>
            </el-form-item>
        </el-form>
        <template #footer>
            <span  class="dialog-footer">
                <el-button type="primary" @click="handleSubmit('dialogForm')">提交</el-button>
                <el-button  @click="handleClose">取消</el-button>
            </span>
        </template>
    </el-dialog>
    </div>
</template>

<script>
import {fetchProvinceList} from '@/api/province.js'
import {fetchCityList} from '@/api/city.js'
const defaultForm = {
    id: "",
    district_name: "",
    status: true,
    province_id: null,
    city_id: null,
}
export default {
    name: 'Dform',
    props: {
        dialogVisiable: {
            type: Boolean,
            required: true,
        },
        isEdit: {
            type: Boolean,
            required: true,
            default: false,
        },
        form: {
            type: Object,
            default: Object.assign({}, defaultForm)
        }
    },
    computed: {
        dialogTitle: function() {
            return this.isEdit ? this.defaultDialogTitles[1] : this.defaultDialogTitles[0]
        },
    },
    watch: {
        form() {
            this.dialogForm = Object.assign({}, this.form)
            this.dialogForm.status = this.dialogForm.status == 1 ? true : false
            if (this.dialogForm.province_id != null) {
                fetchCityList({province_id:this.dialogForm.province_id}).then(res => {
                this.citys = res.data.list
            })
            }
        },
    },
    created() {
        fetchProvinceList().then(res => {
            this.provinces = res.data.list
        })
        if (this.dialogForm.province_id != null) {
            fetchCityList({province_id:this.dialogForm.province_id}).then(res => {
                this.citys = res.data.list
            })
        }
    },
    data() {
        return {
            defaultDialogTitles: ["创建", "编辑"],
            formRules: {
                district_name: [
                    {required: true, message: '请输入区域名称', trigger: 'blur'}
                ],
                province_id: [
                    {required: true, type: 'number', message: '请选择省份', trigger: 'blur'}
                ],
                city_id: [
                    {required: true, type: 'number', message: '请选择城市', trigger: 'blur'}
                ]
            },
            dialogForm: Object.assign({}, defaultForm),
            provinces: [],
            citys: [],
        }
    },
    methods: {
        handleClose() {
            this.dialogForm = Object.assign({}, defaultForm)
            this.$emit("visableChange", false)
         },
        handleSubmit(formName) {
            this.$refs[formName].validate((valid) => {
            if (valid) {
                this.$emit("submit", this.dialogForm)
            } else {
                console.log('error submit!!');
                return false;
            }
            });
        },
        handleGetCitys() {
            fetchCityList({province_id:this.dialogForm.province_id}).then(res => {
                this.citys = res.data.list
            })
        },
    }
}
</script>

<style>

</style>