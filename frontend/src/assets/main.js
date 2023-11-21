// 输入DOM
let dom_ip = document.getElementById("ip");
let dom_path = document.getElementById("path");
dom_ip.focus();

// 默认读取
window.onload = function () {
  window.runtime.WindowSetTitle("百灵快传 - 主电脑配置");
  
  window.go.main.App.LoadConfig().then((res)=>{
    dom_ip.value=res.Ip;
    dom_path.value=res.Path;
    //alert(JSON.stringify(res));
  });
}

// 提交配置
window.submitConfig = function () {
  let ip = dom_ip.value;
  let path = dom_path.value;
  window.go.main.App.SubmitConfig(ip,path).then((res) => {
    if(res=="OK"){
      window.runtime.WindowSetTitle("百灵快传 - 主电脑");
      document.getElementById("result").innerText = "启动成功";
      if(ip.charAt(0)==":"){ip="127.0.0.1"+ip;}
      if(ip.indexOf(":")<=0){ip="127.0.0.1:"+ip;}
      window.location.href="http://"+ip;
    }else{
      alert("启动失败，请检查当前目录下的配置文件");
    } 
  });
};

// 点击提交
window.onkeydown = function (e) {
  console.log(e)
  if (e.keyCode == 13) {
    window.submitConfig();
    WindowSetAlwaysOnTop(true);
  }
}