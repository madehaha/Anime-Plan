<template>
    <div class="add-head">
      <h3>加入收藏</h3>
      <button class="close-addview" @click="AddToCollectBox">X</button>
    </div>
    <div class="add-choose">
      <el-radio-group v-model="radio">
        <el-radio :label="1" size="large">想看</el-radio>
        <el-radio :label="2" size="large">看过</el-radio>
        <el-radio :label="3" size="large">在看</el-radio>
        <el-radio :label="4" size="large">搁置</el-radio>
        <el-radio :label="5" size="large">抛弃</el-radio>
      </el-radio-group>
    </div>
    <hr />
    <div class="my-evaluate" v-if="radio != 1">
      <h5>我的评价</h5>
      <div class="my-evaluate-rate-block">
        <el-rate
          v-model="value"
          allow-half
          size="small"
          text-color="red"
          font-size="small"
          show-score
        />
      </div>
    </div>
    <div class="add-complain-to-collector">
      <h5>当前正在看的章节</h5>
      <div class="add-complain-content">
        <input class="post-input-1" type="text" />
      </div>
    </div>
    <div class="add-complain-to-collector">
      <h5>吐槽,简评</h5>
      <div class="add-complain-content">
        <input class="post-input" type="text" />
      </div>
    </div>
    <button @click="AddSubmit">提交</button>
  </template>
  
  <script setup>
  import { getCurrentInstance, ref } from "vue";
  import { service } from "@/router/demoAxios";
  const isadd = getCurrentInstance()?.appContext.config.globalProperties.$isadd;
  const AddToCollectBox = () => {
    isadd.value = !isadd.value;
  };
  const radio = getCurrentInstance()?.appContext.config.globalProperties.$radio;

  const postTo = async (x1,x2,x3,x4,x5) => {
    await service
      .post(`/collection/${localStorage.current_item_id}`,{
        type:x1,
        has_comment:x2,
        comment:x3,
        score:x4,
        ep_status:x5
      })
      .then((res) => {
         alert("上传成功");
      })
      .catch((err) => {
        console.log(err);
      });
  };
 const label=ref();
  const AddSubmit = () => {
    alert(typeof(label.value));
    isadd.value = !isadd.value;
    var search_content=document.getElementsByClassName("post-input");
    var now_charpter=document.getElementsByClassName("post-input-1");
     postTo(label,true,search_content[0].value,value.value,parseInt(now_charpter));
    }
  
  const value = ref();

  
  </script>
  
  <style>
  .add-choose {
    width: 100%;
  }
  .add-head {
    background-color: rgb(161, 114, 145);
    margin: 0px;
    text-align: left;
  }
  .add-head h3 {
    margin: 0px;
    padding: 10px 0px 10px 10px;
    text-align: left;
    display: inline-block;
  }
  .close-addview {
    border-radius: 100%;
    width: 20px;
    height: 20px;
    float: right;
    margin: 10px 10px 10px 0px;
    border: 2px solid white;
    text-align: center;
    padding: 1px 0px 0px 0px;
  }
  .close-addview:hover {
    cursor: pointer;
  }
  
  .my-evaluate {
    width: 100%;
    text-align: left;
  }
  .my-evaluate h5 {
    margin: 0px 0px 0px 10px;
    text-align: left;
    font-size: small;
  }
  .post-input{
    height: 110px;
    width: 500px;
    margin: 0px 10px 10px 10px;
    border-radius: 6px;
    border: 1px solid rgb(129, 116, 116);
  }

  .post-input-1{
    height: 20px;
    width: 500px;
    margin: 0px 10px 10px 10px;
    border-radius: 6px;
    border: 1px solid rgb(129, 116, 116);
  }
  .add-complain-to-collector {
    width: 100%;
    height: fit-content;
  }
  .add-complain-to-collector h5 {
    margin: 0px 0px 0px 10px;
    text-align: left;
    font-size: small;
  }
  .add-complain-to-collector p {
    text-align: left;
    font-size: smaller;
    margin: 0px;
    padding: 5px 10px 5px 10px;
  }
  
  .add-complain-content {
    height: fit-content;
    margin: 0px 10px 10px 10px;
    border-radius: 6px;
 
  }
  .my-evaluate-rate-block {
    margin: 0px 0px 10px 10px;
  }
  </style>