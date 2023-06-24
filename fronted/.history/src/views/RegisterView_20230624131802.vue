<template>
    <div id="background">
      <div id="contain">
        <div class="rg_left">
          <p style="color: goldenrod">新用户注册</p>
          <p>USER REGISTER</p>
        </div>
        <form action="#" class="form_style">
          <div class="form">
            <label>昵称：</label>
            <input
              id="Telephone"
              placeholder="Nickname"
              type="nickname"
              v-model.trim="new_user.nickname"
              @focus="Reset($event)"
            />
            <div id="dTelephone" style="visibility: hidden; height: 20px">
              <p class="err"></p>
            </div>
          </div>
  
          <div class="form">
            <label>用户名：</label>
            <input
              id="Name"
              placeholder="UserName"
              type="text"
              v-model.trim="new_user.username"
              @focus="Reset($event)"
            />
            <div id="dName" style="visibility: hidden; height: 20px">
              <p class="err"></p>
            </div>
          </div>
  
          <div class="form">
            <label>邮箱：</label>
            <input
              id="Email"
              placeholder="Email"
              type="email"
              v-model.trim="new_user.email"
              @focus="Reset($event)"
            />
            <div id="dEmail" style="visibility: hidden; height: 20px">
              <p class="err"></p>
            </div>
          </div>
  
          <div class="form">
            <label>密码：</label>
            <input
              id="Password"
              placeholder="Password"
              type="password"
              v-model.trim="new_user.password"
              @focus="Reset($event)"
            />
            <div id="dPassword" style="visibility: hidden; height: 20px">
              <p id="err1" class="err"></p>
            </div>
          </div>
  
          <div class="form">
            <label>确认密码：</label>
            <input
              id="NewPassword"
              placeholder="Comfirm the Password"
              type="password"
              v-model.trim="newpassword"
              @focus="Reset($event)"
            />
            <div id="dNewPassword" style="visibility: hidden; height: 20px">
              <p class="err"></p>
            </div>
          </div>
  
          <button @click.prevent="Register">提交</button>
        </form>
        <div class="rg_right">
          <p style="color: black">
            已有账号？<a @click.prevent="ReturnToLogin" style="color: darkcyan"
              >登录</a
            >
          </p>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  //import { service } from "@/router/axios";
  //import { useToken } from "ant-design-vue/es/theme/internal";
  import { service } from "@/router/axios";
  import { ref } from "vue";
  import { useRouter } from "vue-router";
  //const token = ref('');
  const new_user = ref({
    email: "",
    nickname: "",
    password: "",
    username: "",
  });
  const newpassword = ref("");
  
  const TelEmpty = () => {
    var tel = document.getElementById("dTelephone");
    tel.style.visibility = "visible";
    tel.firstChild.innerHTML = "请输入昵称";
    var Input = document.getElementById("Telephone");
    Input.style.borderColor = "red";
  };
  
  const NameEmpty = () => {
    var User = document.getElementById("dName");
    User.style.visibility = "visible";
    User.firstChild.innerHTML = "用户名不能为空";
    var Input = document.getElementById("Name");
    Input.style.borderColor = "red";
  };
  const NameExited = () => {
    var User = document.getElementById("dName");
    User.style.visibility = "visible";
    User.firstChild.innerHTML = "用户名已存在";
    var Input = document.getElementById("Name");
    Input.style.borderColor = "red";
  };
  const EmailEmpty = () => {
    var email = document.getElementById("dEmail");
    email.style.visibility = "visible";
    email.firstChild.innerHTML = "邮箱不能为空";
    var Input = document.getElementById("Email");
    Input.style.borderColor = "red";
  };
  const PasswordEmpty = () => {
    var password = document.getElementById("dPassword");
    password.style.visibility = "visible";
    password.firstChild.innerHTML = "请输入密码";
    var newpassword = document.getElementById("dNewPassword");
    newpassword.style.visibility = "visible";
    newpassword.firstChild.innerHTML = "请确认密码";
    var Input1 = document.getElementById("Password");
    Input1.style.borderColor = "red";
    var Input2 = document.getElementById("NewPassword");
    Input2.style.borderColor = "red";
  };
  const CheckPasswork = () => {
    var newpassword = document.getElementById("dNewPassword");
    newpassword.style.visibility = "visible";
    newpassword.firstChild.innerHTML = "密码不一致";
    var Input = document.getElementById("NewPassword");
    Input.style.borderColor = "red";
  };
  
  const router = new useRouter();
  // const Test=()=>{
  //   console.log(new_user.value.name);
  //   console.log(new_user.value.nickname);
  //   console.log(new_user.value.password);
  //   console.log(new_user.value.mail);
  //   console.log(newpassword.value);
  // }
  const Register = async () => {
    await service
      .post("/register", new_user.value)
      .then((res) => {
        if (res.status == 200) {
          alert("注册成功！");
          router.push({ path: "/" });
        }
      })
      .catch((err) => {
        alert(err);
        if (new_user.value.username === "") {
          NameEmpty();
        } else {
          NameExited();
        }
        if (new_user.value.email === "") {
          EmailEmpty();
        }
        if (new_user.value.password === "") {
          PasswordEmpty();
        }
        if (newpassword.value !== new_user.value.password) {
          CheckPasswork();
        }
        if (new_user.value.nickname === "") {
          TelEmpty();
        }
      });
  };
  
  const Reset = (e) => {
    e.currentTarget.style.borderColor = "#4e5ef3";
    var User = document.getElementById("d" + e.currentTarget.id);
    User.style.visibility = "hidden";
  };
  </script>
  
  //css
  <style scoped>
  #background {
    width: 100%;
    height: 100%;
    background: url("../assets/RegisterView.png");
    background-size: 100% 100%;
    position: fixed;
    top: 0;
    left: 0;
  }
  #contain {
    width: 1080px;
    height: 670px;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: white;
    opacity: 0.7;
    text-align: center;
    border-radius: 20px;
  }
  
  .form {
    color: black;
    margin-left: 10%;
    margin-top: 43px;
    font-size: 20px;
    text-align: left;
  }
  label {
    float: left;
    width: 5em;
    margin-right: 1em;
    text-align: right;
  }
  input {
    margin-left: 10px;
    padding: 4px;
    outline: 0;
    border: solid 1px #4e5ef3;
    font: normal 13px/100% Verdana, Tahoma, sans-serif;
    width: 300px;
    height: 20px;
    background: #f1f1f190;
    box-shadow: rgba(202, 179, 179, 0.1) 0px 0px 8px;
  }
  
  input:hover,
  input:focus {
    border: solid 1px #0d0aa1;
  }
  
  button {
    position: relative;
    height: 33px;
    width: 150px;
    background: rgba(57, 149, 182, 0.98);
    border-radius: 10px;
    margin-top: 40px;
    border: 1px solid;
    text-transform: uppercase;
    font-size: 14px;
    color: white;
    margin-left: 40px;
  }
  
  button:hover {
    box-shadow: 0 12px 16px 0 rgba(0, 0, 0, 0.24),
      0 17px 50px 0 rgba(0, 0, 0, 0.19);
  }
  .rg_right {
    float: right;
    margin: 15px;
  }
  .rg_left {
    float: left;
    margin: 15px;
    width: 20%;
  }
  
  a,
  button:hover {
    cursor: pointer;
  }
  .form_style {
    text-align: center;
    float: left;
    margin-top: 60px;
    width: 600px;
  }
  
  .err {
    color: red;
    margin-top: 0px;
    margin-bottom: 0px;
    width: 300px;
    position: relative;
    left: 130px;
    font-size: small;
  }
  
  .rightImg {
    float: right;
    margin: 17px 12px;
    cursor: pointer;
  }
  </style>
  