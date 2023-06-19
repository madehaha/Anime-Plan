<template>
<div class="main-content">

  <!--左侧主要内容-->
 <div class="content-left">
    <form>
      <div class="li-style">
        <div style="float: left;">国漫</div><br/>
          <hr/>
          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>
          
          
          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>
          
          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>
        <!--小图标链接-->
          <div class="response-block" style="width: 150px;height: 104px;">
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
          </div> 
       </div><br/>


       <div class="li-style">
          <div style="float: left;">日漫</div><br/>
          <hr/>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114223关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              
            </div>
            <p class="response-p">114514关注</p>
          </div>
           <!--小图标链接-->
          <div class="response-block" style="width: 150px;height: 104px;">
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg"  @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
          </div>

       </div><br/>


       <div class="li-style">
          <div style="float: left;">欧美</div><br/>
          <hr/>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114512关注</p>
          </div>
          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>
          
          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>
          
           <!--小图标链接-->
          <div class="response-block" style="width: 150px;height: 104px;">
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
          </div>
       </div><br/>

       <div class="li-style">
          <div style="float: left;">阿三</div><br/>
          <hr/>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114512关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>
          
          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
            </div>
            <p class="response-p">114514关注</p>
          </div>

          <div class="response-block">
            <div class="response-photo" @click="ResponsePhoto">
              <img class="response-img" src="../../assets/example.jpg"/>
              <label class="response-label">【你的名字】</label>
            </div>
            <p class="response-p">114514关注</p>
          </div>

           <!--小图标链接-->
          <div class="response-block" style="width: 150px;height: 104px;">
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
              <div class="response-block-li">
                <div class="response-block-li-photo">
                  <img class="response-img" src="../../assets/example.jpg" @click="ResponsePhoto"/>
                </div>
                <div class="response-block-li-content">
                  <a style="height: 26px;font-size: small;" href="https://www.educoder.net/classrooms/15396/students/">你的名字</a>
                  <p class="response-p">114514关注</p>
                </div> 
              </div>
          </div>
       </div><br/>
    </form>
         
 </div>

 <!--右侧界面内容-->
 <div class="content-right">
    <div class="day-suggest">
          cons
    </div>
 </div>
    
</div>

</template>

<script setup>
const ResponsePhoto=()=>{
  alert("图片响应");
}


</script>

<style>
.main-content{
    margin-left: 8%;
    margin-right: 8%;
    padding-top:40px;
    width: 84%;
    background-color:white;
}

.content-left{
    vertical-align: top;
    display: inline-block;
    width: 60%;
    border-radius: 0%;
    border:0px;
}
.content-right{
    margin-left: 20px;
    vertical-align: top;
    display: inline-block;
    width: 24%;
    height: auto;
    background-color: rgb(209, 200, 200);
    border-radius: 3%;
    border:0px;
}

.li-style{
    width: 100%;
    height: auto;
    text-align: left;
}

.response-block{
   width: 100px;
   height: 120px;
   display: inline-block;
   margin-right: 5px;
   vertical-align: top;
}


.response-block-li{
  width: 150px;
  height: 52px;
 overflow: hidden;
   
}

.response-block-li-photo{
  position: relative;
  height: 49px;
  width: 50px;
  display: inline-block;
  border-radius: 5%;
  border: 1px solid rgb(25, 194, 197);
  overflow: hidden;
  vertical-align: top;
}

.response-block-li-content{
  position: relative;
  width: 95px;
  margin-left: 0px;
  display: inline-block;
  vertical-align: top;
  word-break: break-all;
  
}

.response-photo{
   position: relative;
   width: 100px;
   height: 100px;
   border: 1px solid rgb(25, 194, 197);
   border-radius: 5%;
   overflow: hidden;
   display: inline-block;
   margin-right: 0px;
}

.response-p{
  margin: 0px;
  height: 20px;
  font-size: x-small;
}

.response-img{
    width: 100%;
    height: 100%;
    object-fit: cover;
    z-index: -1;
}

.response-img:hover{
  filter: brightness(80%);
  z-index: -1;
}

.response-label{
  position:absolute;
  left:0px;
  top:76px;
  color: white;
  font-size: x-small;
}

</style>