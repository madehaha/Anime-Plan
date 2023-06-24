import axios from 'axios'
//export将service传出去
export const service = axios.create({
	baseURL: 'https://aivwe.top/'
	//baseURL: import.meta.env.VITE_BASE_URL, //这里也可以使用变量
	//timeout: 30000,//超时设置
	//withCredentials: true, //异步请求携带cookie
	//headers: {
	// 		//设置后端需要的传参类型
	// 		'Content-Type': 'application/json;charset=UTF-8;',
	// 		//'token': 'your token',
	// 		'X-Requested-With': 'XMLHttpRequest'
	// 	}
})
 