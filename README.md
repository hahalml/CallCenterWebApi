CallCenterWebApi （语音电话Web接口程序）
=======================================

###1、功能说明

		实现了语音电话的接口，然后转换成Web接口，方便第三方程序调用。

###2、配置文件
		[程序设置]
		端口=:8082
		静态文件=/static/
		模板文件=/view/*.html
		DSN名称=YinZi_Call
		数据库用户名=sa
		数据库密码=123456

		注意！数据库采用的是MSSSQL，通过配置DSN连接数据库
###3、运行环境
		系统之在window 2008 32位、Window 7 32位测试过
###4、使用步骤
		(1)、运行CallCenterWebApi.exe
		(2)、在浏览器输入http://localhost:8082/
		(3)、输入测试的数据，语音设备能拨打电话说明程序运行成功
###5、源码编译
		 go install vistech/CallCenterWebApi