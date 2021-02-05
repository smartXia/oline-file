<template>
<div>
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
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createImages,
    updateImages,
    findImages
} from "@/api/images";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Images",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
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
  methods: {
    async save() {
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
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findImages({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.reimages
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>