<template>
  <div class="body">
    <div class="shell">
      <div id="img-box">
        <img src="../assets/62f52b3a9a496d5499de328dcc04d13e.jpeg" alt="" />
      </div>
      <form >
        <div id="form-body">
          <div id="welcome-lines">
            <div id="w-line-1">Welcome to</div>
            <div id="w-line-2">Anime-Plan</div>
          </div>
          <div id="input-area">
            <div class="f-inp">
              <input type="text" placeholder="Email Address" v-model="user.email" />
            </div>
            <div class="f-inp">
              <input type="password" placeholder="Password" v-model="user.password"/>
            </div>
          </div>
          <div id="submit-button-cvr">
            <button  id="submit-button" @click.prevent='loginIn'>LOGIN</button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>
 
<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { service } from "@/router/demoAxios";
const user = ref({ email: "", password: "" });
const router = useRouter();
const token = ref('');
const loginIn = async () => {
		await service.post('/login',user.value).then(
			res => {
				if (res.status == 200) {
          token.value=res.data.data.token;
          // console.log(res.data)
          console.log(token)
				} else {
					alert('邮箱或者密码错误！')
				}
			})
	}
  // router.push({path: '/center'});
</script>
 
<style scoped>
* {
  padding: 0;
  margin: 0;
  outline: none;
}

.body {
  background: linear-gradient(45deg, #fbda61, #ff5acd);
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  /* background-image: url('../assets/RegisterView.png');
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: 100% 100%; */
}

.star {
  height: 100px;
  width: 100px;
  background-size: 100px 100px;
}

.shell,
form {
  position: relative;
}

.shell {
  display: flex;
  justify-content: center;
}

form {
  width: 562px;
  height: 520px;
  background-color: #fff;
  box-shadow: 5px 0px 10px 10px rgba(182,53,78,0.2);
  border-radius: 15px;
  display: flex;
  justify-content: center;
  align-items: center;
}

#img-box {
  width: 330px;
  height: 520px;
}

#img-box img {
  height: 100%;
  margin-left: -175px;
  border-radius: 20px 0 0 20px;
  box-shadow: -5px 0px 10px 10px rgba(182,53,78,0.2);
}

#form-body {
  width: 320px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}

#welcome-lines {
  width: 100%;
  text-align: center;
  line-height: 1;
}

#w-line-1 {
  color: #ea1b1b;
  font-size: 40px;
  font-family: flower;
}

#w-line-2 {
  color: #ea0e0eb4;
  font-size: 50px;
  margin-top: 17px;
  font-family: flower;
}

#input-area {
  width: 100%;
  margin-top: 40px;
}

.f-inp {
  padding: 13px 25px;
  border: 2px solid #6e6d6d;
  line-height: 1;
  border-radius: 20px;
  margin-bottom: 15px;
}

.f-inp input {
  width: 100%;
  font-size: 18px;
  padding: 0;
  margin: 0;
  border: 0;
}

.f-inp input::placeholder {
  color: #083a28;
  font-family: flower;
}

#submit-button-cvr {
  margin-top: 20px;
}

#submit-button {
  display: block;
  width: 100%;
  color: #fff;
  font-size: 14px;
  margin: 0;
  padding: 14px 40px;
  border: 0;
  background-color: #f5506e;
  border-radius: 25px;
  line-height: 1;
  cursor: pointer;
}


</style>