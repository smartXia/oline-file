<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                            
        <el-form-item label="图片来源哪个任务ID">
          <el-input placeholder="搜索条件" v-model="searchInfo.taskId"></el-input>
        </el-form-item>      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增images表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="中台标识" prop="appkey" width="120"></el-table-column> 
    
    <el-table-column label="渠道标识" prop="channel" width="120"></el-table-column> 
    
    <el-table-column label="账户id" prop="accountId" width="120"></el-table-column> 
    
    <el-table-column label="模块id" prop="functionId" width="120"></el-table-column> 
    
    <el-table-column label="图片类型:image&card" prop="imageType" width="120"></el-table-column> 
    
    <el-table-column label="描述" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="图片url" prop="url" width="120"></el-table-column> 
    
    <el-table-column label="图片宽" prop="width" width="120"></el-table-column> 
    
    <el-table-column label="图片高" prop="height" width="120"></el-table-column> 
    
    <el-table-column label="图片展示x轴长度" prop="positionX" width="120"></el-table-column> 
    
    <el-table-column label="图片展示y轴长度" prop="positionY" width="120"></el-table-column> 
    
    <el-table-column label="图片唯一标识" prop="uniqueId" width="120"></el-table-column> 
    
    <el-table-column label="内置图片" prop="group" width="120"></el-table-column> 
    
    <el-table-column label="组封面图" prop="cover" width="120"></el-table-column> 
    
    <el-table-column label="图片来源哪个任务ID" prop="taskId" width="120"></el-table-column> 
    
    <el-table-column label="任务类型" prop="taskType" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateImages(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="中台标识:">
            <el-input v-model="formData.appkey" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="渠道标识:"><el-input v-model.number="formData.channel" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="账户id:">
            <el-input v-model="formData.accountId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="模块id:">
            <el-input v-model="formData.functionId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片类型:image&card:">
            <el-input v-model="formData.imageType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="描述:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片url:">
            <el-input v-model="formData.url" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片宽:">
            <el-input v-model="formData.width" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片高:">
            <el-input v-model="formData.height" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片展示x轴长度:">
            <el-input v-model="formData.positionX" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片展示y轴长度:">
            <el-input v-model="formData.positionY" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片唯一标识:">
            <el-input v-model="formData.uniqueId" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="内置图片:">
            <el-input v-model="formData.group" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="组封面图:">
            <el-input v-model="formData.cover" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="图片来源哪个任务ID:"><el-input v-model.number="formData.taskId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="任务类型:">
            <el-input v-model="formData.taskType" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createImages,
    deleteImages,
    deleteImagesByIds,
    updateImages,
    findImages,
    getImagesList
} from "@/api/images";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Images",
  mixins: [infoList],
  data() {
    return {
      listApi: getImagesList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            appkey:"",
            channel:0,
            accountId:"",
            functionId:"",
            imageType:"",
            name:"",
            url:"",
            width:"",
            height:"",
            positionX:"",
            positionY:"",
            uniqueId:"",
            group:"",
            cover:"",
            taskId:0,
            taskType:"",
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10                    
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteImages(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteImagesByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateImages(row) {
      const res = await findImages({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reimages;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          appkey:"",
          channel:0,
          accountId:"",
          functionId:"",
          imageType:"",
          name:"",
          url:"",
          width:"",
          height:"",
          positionX:"",
          positionY:"",
          uniqueId:"",
          group:"",
          cover:"",
          taskId:0,
          taskType:"",
          
      };
    },
    async deleteImages(row) {
      const res = await deleteImages({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createImages(this.formData);
          break;
        case "update":
          res = await updateImages(this.formData);
          break;
        default:
          res = await createImages(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
